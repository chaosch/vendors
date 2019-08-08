package xorm

import (
	"egov/go-xorm/builder"
	"egov/go-xorm/core"
	"errors"
	"github.com/binlaniua/SqlParser"
	"reflect"
	"regexp"
	"strings"
)

// union union_operator should be one of INNER, LEFT OUTER, CROSS etc - this will be prepended to JOIN
func (session *Session) Union(unionType string, subQuery string) *Session {
	session.Statement.Union(unionType, subQuery)
	return session
}

func (session *Session) FindReturnWithSql(rowsSlicePtr interface{}, condiBean ...interface{}) (error, string) {
	defer session.resetStatement()
	if session.IsAutoClose {
		defer session.Close()
	}

	sliceValue := reflect.Indirect(reflect.ValueOf(rowsSlicePtr))
	if sliceValue.Kind() != reflect.Slice && sliceValue.Kind() != reflect.Map {
		return errors.New("needs a pointer to a slice or a map"), ""
	}

	sliceElementType := sliceValue.Type().Elem()

	var tp = tpStruct
	if session.Statement.RefTable == nil {
		if sliceElementType.Kind() == reflect.Ptr {
			if sliceElementType.Elem().Kind() == reflect.Struct {
				pv := reflect.New(sliceElementType.Elem())
				if err := session.Statement.setRefValue(pv.Elem()); err != nil {
					return err, ""
				}
			} else {
				tp = tpNonStruct
			}
		} else if sliceElementType.Kind() == reflect.Struct {
			pv := reflect.New(sliceElementType)
			if err := session.Statement.setRefValue(pv.Elem()); err != nil {
				return err, ""
			}
		} else {
			tp = tpNonStruct
		}
	}

	var table = session.Statement.RefTable

	var addedTableName = len(session.Statement.JoinStr) > 0
	var autoCond builder.Cond
	if tp == tpStruct {
		if !session.Statement.noAutoCondition && len(condiBean) > 0 {
			var err error
			autoCond, err = session.Statement.buildConds(table, condiBean[0], true, true, false, true, addedTableName)
			if err != nil {
				panic(err)
			}
		} else {
			// !oinume! Add "<col> IS NULL" to WHERE whatever condiBean is given.
			// See https://egov/go-xorm/xorm/issues/179
			if col := table.DeletedColumn(); col != nil && !session.Statement.unscoped { // tag "deleted" is enabled
				var colName = session.Engine.Quote(col.Name)
				if addedTableName {
					var nm = session.Statement.TableName()
					if len(session.Statement.TableAlias) > 0 {
						nm = session.Statement.TableAlias
					}
					colName = session.Engine.Quote(nm) + "." + colName
				}
				if session.Engine.dialect.DBType() == core.MSSQL {
					autoCond = builder.IsNull{colName}
				} else {
					autoCond = builder.IsNull{colName}.Or(builder.Eq{colName: "0001-01-01 00:00:00"})
				}
			}
		}
	} else {
		var xtable *core.Table
		if len(condiBean) > 0 {
			xtable, _ = session.Engine.AutoMapType(RValue(condiBean[0]))
			session.Statement.RefTable = xtable
			table = session.Statement.RefTable
		}
		if !session.Statement.noAutoCondition && len(condiBean) > 0 {
			var err error
			autoCond, err = session.Statement.buildConds(xtable, condiBean[0], true, true, false, true, addedTableName)
			if err != nil {
				panic(err)
			}
		}
	}

	var sqlStr string
	var args []interface{}
	if session.Statement.RawSQL == "" {
		if len(session.Statement.TableName()) <= 0 {
			return ErrTableNotFound, ""
		}
		var columnStr = session.Statement.ColumnStr
		if len(session.Statement.selectStr) > 0 {
			columnStr = session.Statement.selectStr
			if session.Statement.JoinStr == "" {
				if columnStr == "*" || columnStr == session.Statement.TableAlias+".*" {
					if session.Statement.GroupByStr != "" {
						columnStr = session.Statement.Engine.Quote(strings.Replace(session.Statement.GroupByStr, ",", session.Engine.Quote(","), -1))
					} else {
						v := RValue(condiBean[0])
						t := v.Type()
						//fmt.Println(v.Type().String())
						//fmt.Println(v.Type().Name())
						//fmt.Println(v.Kind().String())
						//
						//for x,_:=range session.Engine.Tables{
						//	fmt.Println(x.Name(),x.Kind().String(),x.String())
						//}
						session.Statement.RefTable = session.Engine.Tables[t]
						columnStr = session.Statement.genColumnStr()
						if session.Statement.TableAlias != "" {
							columnStr = strings.Replace(columnStr, ",", ","+session.Statement.TableAlias+".", -1)
							columnStr = session.Statement.TableAlias + "." + columnStr
						}
					}
				}
			}
		} else {
			if session.Statement.JoinStr == "" {
				if columnStr == "" {
					if session.Statement.GroupByStr != "" {
						columnStr = session.Statement.Engine.Quote(strings.Replace(session.Statement.GroupByStr, ",", session.Engine.Quote(","), -1))
					} else {
						columnStr = session.Statement.genColumnStr()
						if session.Statement.TableAlias != "" {
							columnStr = strings.Replace(columnStr, ",", ","+session.Statement.TableAlias+".", -1)
							columnStr = session.Statement.TableAlias + "." + columnStr
						}
					}
				}
			} else {
				if columnStr == "" {
					if session.Statement.GroupByStr != "" {
						columnStr = session.Statement.Engine.Quote(strings.Replace(session.Statement.GroupByStr, ",", session.Engine.Quote(","), -1))
					} else {
						columnStr = "*"
						columnStr = ""
						for i, c := range session.Statement.RefTable.ColumnsSeq() {
							if session.Statement.TableAlias != "" {
								if i == 0 {
									columnStr = columnStr + session.Statement.TableAlias + "." + c
								} else {
									columnStr = columnStr + "," + session.Statement.TableAlias + "." + c
								}
							}

						}
					}
				}
			}
			if columnStr == "" {
				columnStr = "*"
				columnStr = ""
				for i, c := range session.Statement.RefTable.ColumnsSeq() {
					if session.Statement.TableAlias != "" {
						if i == 0 {
							columnStr = columnStr + session.Statement.TableAlias + "." + c
						} else {
							columnStr = columnStr + "," + session.Statement.TableAlias + "." + c
						}
					}

				}
			}
		}

		condSQL, condArgs, _ := builder.ToSQL(session.Statement.cond.And(autoCond))

		args = append(session.Statement.joinArgs, condArgs...)
		sqlStr = session.Statement.genSelectSQL(columnStr, condSQL)
		// for mssql and use limit
		qs := strings.Count(sqlStr, "?")
		if len(args)*2 == qs {
			args = append(args, args...)
		}
		sqlStr += session.Statement.UnionStr

	} else {
		sqlStr = session.Statement.RawSQL
		args = session.Statement.RawParams
	}

	var err error
	if session.canCache() {
		if cacher := session.Engine.getCacher2(table); cacher != nil &&
			!session.Statement.IsDistinct &&
			!session.Statement.unscoped {
			err = session.cacheFind(sliceElementType, sqlStr, rowsSlicePtr, args...)
			if err != ErrCacheFailed {
				return err, ""
			}
			err = nil // !nashtsai! reset err to nil for ErrCacheFailed
			session.Engine.logger.Warn("Cache Find Failed")
		}
	}

	return session.noCacheFind(table, sliceValue, sqlStr, args...), sqlStr
}

func (session *Session) FindReturnWithSqlParseResult(rowsSlicePtr interface{}, condiBean ...interface{}) (error, *sqlparse.SQLParserResult) {
	defer session.resetStatement()
	if session.IsAutoClose {
		defer session.Close()
	}

	sliceValue := reflect.Indirect(reflect.ValueOf(rowsSlicePtr))
	if sliceValue.Kind() != reflect.Slice && sliceValue.Kind() != reflect.Map {
		return errors.New("needs a pointer to a slice or a map"), nil
	}

	sliceElementType := sliceValue.Type().Elem()

	var tp = tpStruct
	if session.Statement.RefTable == nil {
		if sliceElementType.Kind() == reflect.Ptr {
			if sliceElementType.Elem().Kind() == reflect.Struct {
				pv := reflect.New(sliceElementType.Elem())
				if err := session.Statement.setRefValue(pv.Elem()); err != nil {
					return err, nil
				}
			} else {
				tp = tpNonStruct
			}
		} else if sliceElementType.Kind() == reflect.Struct {
			pv := reflect.New(sliceElementType)
			if err := session.Statement.setRefValue(pv.Elem()); err != nil {
				return err, nil
			}
		} else {
			tp = tpNonStruct
		}
	}

	var table = session.Statement.RefTable

	var addedTableName = len(session.Statement.JoinStr) > 0
	var autoCond builder.Cond
	if tp == tpStruct {
		if !session.Statement.noAutoCondition && len(condiBean) > 0 {
			var err error
			autoCond, err = session.Statement.buildConds(table, condiBean[0], true, true, false, true, addedTableName)
			if err != nil {
				panic(err)
			}
		} else {
			// !oinume! Add "<col> IS NULL" to WHERE whatever condiBean is given.
			// See https://egov/go-xorm/xorm/issues/179
			if col := table.DeletedColumn(); col != nil && !session.Statement.unscoped { // tag "deleted" is enabled
				var colName = session.Engine.Quote(col.Name)
				if addedTableName {
					var nm = session.Statement.TableName()
					if len(session.Statement.TableAlias) > 0 {
						nm = session.Statement.TableAlias
					}
					colName = session.Engine.Quote(nm) + "." + colName
				}
				if session.Engine.dialect.DBType() == core.MSSQL {
					autoCond = builder.IsNull{colName}
				} else {
					autoCond = builder.IsNull{colName}.Or(builder.Eq{colName: "0001-01-01 00:00:00"})
				}
			}
		}
	} else {
		var xtable *core.Table
		if len(condiBean) > 0 {
			xtable, _ = session.Engine.AutoMapType(RValue(condiBean[0]))
			session.Statement.RefTable = xtable
			table = session.Statement.RefTable
		}
		if !session.Statement.noAutoCondition && len(condiBean) > 0 {
			var err error
			autoCond, err = session.Statement.buildConds(xtable, condiBean[0], true, true, false, true, addedTableName)
			if err != nil {
				panic(err)
			}
		}
	}

	var sqlStr string
	var args []interface{}
	if session.Statement.RawSQL == "" {
		if len(session.Statement.TableName()) <= 0 {
			return ErrTableNotFound, nil
		}
		var columnStr = session.Statement.ColumnStr
		if len(session.Statement.selectStr) > 0 {
			columnStr = session.Statement.selectStr
			if session.Statement.JoinStr == "" {
				if columnStr == "*" || columnStr == session.Statement.TableAlias+".*" {
					if session.Statement.GroupByStr != "" {
						columnStr = session.Statement.Engine.Quote(strings.Replace(session.Statement.GroupByStr, ",", session.Engine.Quote(","), -1))
					} else {
						v := RValue(condiBean[0])
						t := v.Type()
						//fmt.Println(v.Type().String())
						//fmt.Println(v.Type().Name())
						//fmt.Println(v.Kind().String())
						//
						//for x,_:=range session.Engine.Tables{
						//	fmt.Println(x.Name(),x.Kind().String(),x.String())
						//}
						session.Statement.RefTable = session.Engine.Tables[t]
						columnStr = session.Statement.genColumnStr()
						if session.Statement.TableAlias != "" {
							columnStr = strings.Replace(columnStr, ",", ","+session.Statement.TableAlias+".", -1)
							columnStr = session.Statement.TableAlias + "." + columnStr
						}
					}
				}
			}
		} else {
			if session.Statement.JoinStr == "" {
				if columnStr == "" {
					if session.Statement.GroupByStr != "" {
						columnStr = session.Statement.Engine.Quote(strings.Replace(session.Statement.GroupByStr, ",", session.Engine.Quote(","), -1))
					} else {
						columnStr = session.Statement.genColumnStr()
						if session.Statement.TableAlias != "" {
							columnStr = strings.Replace(columnStr, ",", ","+session.Statement.TableAlias+".", -1)
							columnStr = session.Statement.TableAlias + "." + columnStr
						}
					}
				}
			} else {
				if columnStr == "" {
					if session.Statement.GroupByStr != "" {
						columnStr = session.Statement.Engine.Quote(strings.Replace(session.Statement.GroupByStr, ",", session.Engine.Quote(","), -1))
					} else {
						columnStr = "*"
						columnStr = ""
						for i, c := range session.Statement.RefTable.ColumnsSeq() {
							if session.Statement.TableAlias != "" {
								if i == 0 {
									columnStr = columnStr + session.Statement.TableAlias + "." + c
								} else {
									columnStr = columnStr + "," + session.Statement.TableAlias + "." + c
								}
							}

						}
					}
				}
			}
			if columnStr == "" {
				columnStr = "*"
				columnStr = ""
				for i, c := range session.Statement.RefTable.ColumnsSeq() {
					if session.Statement.TableAlias != "" {
						if i == 0 {
							columnStr = columnStr + session.Statement.TableAlias + "." + c
						} else {
							columnStr = columnStr + "," + session.Statement.TableAlias + "." + c
						}
					}

				}
			}
		}

		condSQL, condArgs, _ := builder.ToSQL(session.Statement.cond.And(autoCond))

		args = append(session.Statement.joinArgs, condArgs...)
		sqlStr = session.Statement.genSelectSQL(columnStr, condSQL)
		// for mssql and use limit
		qs := strings.Count(sqlStr, "?")
		if len(args)*2 == qs {
			args = append(args, args...)
		}
		sqlStr += session.Statement.UnionStr

	} else {
		sqlStr = session.Statement.RawSQL
		args = session.Statement.RawParams
	}

	var err error
	if session.canCache() {
		if cacher := session.Engine.getCacher2(table); cacher != nil &&
			!session.Statement.IsDistinct &&
			!session.Statement.unscoped {
			err = session.cacheFind(sliceElementType, sqlStr, rowsSlicePtr, args...)
			if err != ErrCacheFailed {
				return err, nil
			}
			err = nil // !nashtsai! reset err to nil for ErrCacheFailed
			session.Engine.logger.Warn("Cache Find Failed")
		}
	}

	err, parserRes := session.NoCacheFind(table, sliceValue, sqlStr, args...)
	return err, parserRes
}

func (session *Session) NoCacheFind(table *core.Table, containerValue reflect.Value, sqlStr string, args ...interface{}) (error, *sqlparse.SQLParserResult) {
	var rawRows *core.Rows
	var err error

	//if session.Engine.showSQL {
	//	fmt.Println(sqlStr)
	//}
	err=session.ParserSqlAllColumns(&sqlStr)

	if err!=nil{
		return err,nil
	}

	session.queryPreprocess(&sqlStr, args...)

	if session.IsAutoCommit {
		_, rawRows, err = session.innerQuery(sqlStr, args...)
	} else {
		rawRows, err = session.Tx.Query(sqlStr, args...)
	}
	if err != nil {
		return err, nil
	}
	reg := regexp.MustCompile(`(?i: offset )\d*$`)
	sqlStr = reg.ReplaceAllString(sqlStr, "")

	p := sqlparse.NewSQLParser(sqlStr)
	p.DoParser()
	rawRows.SQLPR = p.GetResult()
	defer rawRows.Close()

	fields, err := rawRows.Columns()
	if err != nil {
		return err, nil
	}
	var newElemFunc func(fields []string) reflect.Value
	elemType := containerValue.Type().Elem()
	var isPointer bool
	if elemType.Kind() == reflect.Ptr {
		isPointer = true
		elemType = elemType.Elem()
	}
	if elemType.Kind() == reflect.Ptr {
		return errors.New("pointer to pointer is not supported"), nil
	}

	newElemFunc = func(fields []string) reflect.Value {
		switch elemType.Kind() {
		case reflect.Slice:
			slice := reflect.MakeSlice(elemType, len(fields), len(fields))
			x := reflect.New(slice.Type())
			x.Elem().Set(slice)
			return x
		case reflect.Map:
			mp := reflect.MakeMap(elemType)
			x := reflect.New(mp.Type())
			x.Elem().Set(mp)
			return x
		}
		return reflect.New(elemType)
	}

	var containerValueSetFunc func(*reflect.Value, core.PK) error

	if containerValue.Kind() == reflect.Slice {
		containerValueSetFunc = func(newValue *reflect.Value, pk core.PK) error {
			if isPointer {
				containerValue.Set(reflect.Append(containerValue, newValue.Elem().Addr()))
			} else {
				containerValue.Set(reflect.Append(containerValue, newValue.Elem()))
			}
			return nil
		}
	} else {
		keyType := containerValue.Type().Key()
		if len(table.PrimaryKeys) == 0 {
			return errors.New("don't support multiple primary key's map has non-slice key type"), nil
		}
		if len(table.PrimaryKeys) > 1 && keyType.Kind() != reflect.Slice {
			return errors.New("don't support multiple primary key's map has non-slice key type"), nil
		}

		containerValueSetFunc = func(newValue *reflect.Value, pk core.PK) error {
			keyValue := reflect.New(keyType)
			err := convertPKToValue(table, keyValue.Interface(), pk)
			if err != nil {
				return err
			}
			if isPointer {
				containerValue.SetMapIndex(keyValue.Elem(), newValue.Elem().Addr())
			} else {
				containerValue.SetMapIndex(keyValue.Elem(), newValue.Elem())
			}
			return nil
		}
	}

	if elemType.Kind() == reflect.Struct {
		var newValue = newElemFunc(fields)
		dataStruct := rValue(newValue.Interface())
		tb, err := session.Engine.autoMapType(dataStruct)
		if err != nil {
			return err, nil
		}
		return session.rows2Beans(rawRows, fields, len(fields), tb, newElemFunc, containerValueSetFunc), nil
	}

	rawRows.ColumnTypes = session.Engine.ColumnTypes

	for rawRows.Next() {
		var newValue = newElemFunc(fields)
		bean := newValue.Interface()

		switch elemType.Kind() {
		case reflect.Slice:
			err = rawRows.ScanSlice(bean)
		case reflect.Map:
			err = rawRows.ScanMap(bean)
		default:
			err = rawRows.Scan(bean)
		}

		if err != nil {
			return err, nil
		}

		if err := containerValueSetFunc(&newValue, nil); err != nil {
			return err, nil
		}
	}
	return nil, rawRows.SQLPR
}

func (session *Session) ParserSqlAllColumns(sqlStr *string) error {

	sql := *sqlStr
	reg := regexp.MustCompile(`(?i: offset )\d*$`)
	sql = reg.ReplaceAllString(sql, "")

	ps := sqlparse.NewSQLParser(sql)
	x, err := ps.DoParser()

	if err!=nil {
		return err
	}
	//if x == nil {
	//	return
	//}
	p := x.GetDBUser("*")

	//if p == nil {
	//	return
	//}
	//fmt.Println(p.TableMap)
	for _, t := range p.TableMap {
		sqlTab := ""
		sqlCols := ""

		//T, Tok := session.Engine.Tabs[t.Name]
		//if _, ok := t.ColumnMap[T.PrimaryKeys[0]]; !ok && Tok && len(t.ColumnMap) > 1 && t.Alias.Name == t.GetTopAlias() { //t表主键不存在于select
		//	if strings.HasPrefix(*sqlStr, "select") {
		//		*sqlStr = strings.Replace(*sqlStr, "select ", "select "+t.GetTopAlias()+"."+T.PrimaryKeys[0]+" "+t.GetTopAlias()+"_"+T.PrimaryKeys[0]+",", 1)
		//	}
		//	if strings.HasPrefix(*sqlStr, "SELECT") {
		//		*sqlStr = strings.Replace(*sqlStr, "SELECT ", "SELECT "+t.GetTopAlias()+"."+T.PrimaryKeys[0]+" "+t.GetTopAlias()+"_"+T.PrimaryKeys[0]+",", 1)
		//	}
		//}
		//
		//if _, ok := t.ColumnMap["split_code"]; !ok && Tok && len(t.ColumnMap) > 1 && Tok && !strings.HasPrefix(t.Name, "dic_") && t.Alias.Name == t.GetTopAlias() { //t.split_code不存在于select
		//	if strings.HasPrefix(*sqlStr, "select") {
		//		*sqlStr = strings.Replace(*sqlStr, "select ", "select "+t.GetTopAlias()+".split_code"+" "+t.GetTopAlias()+"_split_code,", 1)
		//	}
		//	if strings.HasPrefix(*sqlStr, "SELECT") {
		//		*sqlStr = strings.Replace(*sqlStr, "SELECT ", "SELECT "+t.GetTopAlias()+".split_code"+" "+t.GetTopAlias()+"_split_code,", 1)
		//	}
		//}

		for _, c := range t.ColumnMap {
			if c.Name == "*" {
				//fmt.Println(t.GetTopAlias()+"."+"*", t.Name+"."+"*")
				var xt *core.Table
				xt = session.Engine.Tabs[t.Name]
				sqlTab = t.GetTopAlias()
				for _, col := range xt.Columns() {
					sqlCols += sqlTab + "." + col.Name + ","
				}
				sqlCols += "'x'"
			}
		}
		if sqlCols != "" {
			*sqlStr = strings.Replace(*sqlStr, sqlTab+".*", sqlCols, -1)
		}

	}
	return nil
}
