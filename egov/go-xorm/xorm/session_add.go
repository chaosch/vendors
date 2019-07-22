package xorm

import (
	"egov/go-xorm/builder"
	"egov/go-xorm/core"
	"errors"
	"reflect"
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
