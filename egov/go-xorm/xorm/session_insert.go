// Copyright 2016 The Xorm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xorm

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"egov/go-xorm/core"
)

// Insert insert one or more beans
func (session *Session) Insert(beans ...interface{}) (int64, error) {
	var affected int64
	var err error

	if session.IsAutoClose {
		defer session.Close()
	}
	defer session.resetStatement()

	for _, bean := range beans {
		sliceValue := reflect.Indirect(reflect.ValueOf(bean))
		if sliceValue.Kind() == reflect.Slice {
			size := sliceValue.Len()
			if size > 0 {
				if session.Engine.SupportInsertMany() {
					cnt, err := session.innerInsertMulti(bean)
					if err != nil {
						return affected, err
					}
					affected += cnt
				} else {
					for i := 0; i < size; i++ {
						cnt, err := session.innerInsert(sliceValue.Index(i).Interface())
						if err != nil {
							return affected, err
						}
						affected += cnt
					}
				}
			}
		} else {
			cnt, err := session.innerInsert(bean)
			if err != nil {
				return affected, err
			}
			affected += cnt
		}
	}

	return affected, err
}

func (session *Session) innerInsertMulti(rowsSlicePtr interface{}) (int64, error) {
	sliceValue := reflect.Indirect(reflect.ValueOf(rowsSlicePtr))
	if sliceValue.Kind() != reflect.Slice {
		return 0, errors.New("needs a pointer to a slice")
	}

	if sliceValue.Len() <= 0 {
		return 0, errors.New("could not insert a empty slice")
	}

	if err := session.Statement.setRefValue(sliceValue.Index(0)); err != nil {
		return 0, err
	}

	if len(session.Statement.TableName()) <= 0 {
		return 0, ErrTableNotFound
	}

	table := session.Statement.RefTable
	size := sliceValue.Len()

	var colNames []string
	var colMultiPlaces []string
	var args []interface{}
	var cols []*core.Column

	colIsNotNil := make(map[string]bool)

	DataType := reflect.Indirect(reflect.ValueOf(rowsSlicePtr)).Type().Elem()
	for j := 0; j < DataType.NumField(); j++ {
		cName := DataType.Field(j).Name
		colDone:=false
		for i := 0; i < size; i++ {
			v := sliceValue.Index(i)
			Data := reflect.Indirect(v)
			if !Data.Field(j).IsNil() {
				colIsNotNil[strings.ToLower(cName)]=false
				colDone=true
				break
			}
		}
		if colDone{
			continue
		}
		colIsNotNil[strings.ToLower(cName)]=true
	}

	for i := 0; i < size; i++ {
		v := sliceValue.Index(i)
		vv := reflect.Indirect(v)
		elemValue := v.Interface()
		var colPlaces []string

		// handle BeforeInsertProcessor
		// !nashtsai! does user expect it's same slice to passed closure when using Before()/After() when insert multi??
		for _, closure := range session.beforeClosures {
			closure(elemValue)
		}

		if processor, ok := interface{}(elemValue).(BeforeInsertProcessor); ok {
			processor.BeforeInsert()
		}
		// --

		if i == 0 {
			for _, col := range table.Columns() {
				ptrFieldValue, err := col.ValueOfV(&vv)
				if err != nil {
					return 0, err
				}
				fieldValue := *ptrFieldValue
				if col.IsAutoIncrement && isZero(fieldValue.Interface()) {
					continue
				}

				if colIsNotNil[col.Name] {
					continue
				}

				if col.MapType == core.ONLYFROMDB {
					continue
				}
				if col.IsDeleted {
					continue
				}
				if session.Statement.ColumnStr != "" {
					if _, ok := getFlagForColumn(session.Statement.columnMap, col); !ok {
						continue
					}
				}
				if session.Statement.OmitStr != "" {
					if _, ok := getFlagForColumn(session.Statement.columnMap, col); ok {
						continue
					}
				}
				if (col.IsCreated && fieldValue.IsNil()) && session.Statement.UseAutoTime {
					val, t := session.Engine.NowTime2(col.SQLType.Name)
					args = append(args, val)

					var colName = col.Name
					session.afterClosures = append(session.afterClosures, func(bean interface{}) {
						col := table.GetColumn(colName)
						setColumnTime(bean, col, t)
					})
				} else if col.IsVersion && session.Statement.checkVersion {
					args = append(args, 1)
					var colName = col.Name
					session.afterClosures = append(session.afterClosures, func(bean interface{}) {
						col := table.GetColumn(colName)
						setColumnInt(bean, col, 1)
					})
				} else {
					arg, err := session.value2Interface(col, fieldValue)
					if err != nil {
						return 0, err
					}
					args = append(args, arg)
				}

				colNames = append(colNames, col.Name)
				cols = append(cols, col)
				colPlaces = append(colPlaces, "?")
			}
		} else {
			for _, col := range cols {
				ptrFieldValue, err := col.ValueOfV(&vv)
				if err != nil {
					return 0, err
				}
				fieldValue := *ptrFieldValue

				if colIsNotNil[col.Name] {
					continue
				}

				if col.IsAutoIncrement && isZero(fieldValue.Interface()) {
					continue
				}
				if col.MapType == core.ONLYFROMDB {
					continue
				}
				if col.IsDeleted {
					continue
				}
				if session.Statement.ColumnStr != "" {
					if _, ok := getFlagForColumn(session.Statement.columnMap, col); !ok {
						continue
					}
				}
				if session.Statement.OmitStr != "" {
					if _, ok := getFlagForColumn(session.Statement.columnMap, col); ok {
						continue
					}
				}
				if (col.IsCreated && fieldValue.IsNil()) && session.Statement.UseAutoTime {
					val, t := session.Engine.NowTime2(col.SQLType.Name)
					args = append(args, val)

					var colName = col.Name
					session.afterClosures = append(session.afterClosures, func(bean interface{}) {
						col := table.GetColumn(colName)
						setColumnTime(bean, col, t)
					})
				} else if col.IsVersion && session.Statement.checkVersion {
					args = append(args, 1)
					var colName = col.Name
					session.afterClosures = append(session.afterClosures, func(bean interface{}) {
						col := table.GetColumn(colName)
						setColumnInt(bean, col, 1)
					})
				} else {
					arg, err := session.value2Interface(col, fieldValue)
					if err != nil {
						return 0, err
					}
					args = append(args, arg)
				}

				colPlaces = append(colPlaces, "?")
			}
		}
		colMultiPlaces = append(colMultiPlaces, strings.Join(colPlaces, ", "))
	}
	cleanupProcessorsClosures(&session.beforeClosures)

	var sql = "INSERT INTO %s (%v%v%v) VALUES (%v)"
	var statement string
	if session.Engine.dialect.DBType() == core.ORACLE {
		sql = "INSERT ALL INTO %s (%v%v%v) VALUES (%v) SELECT 1 FROM DUAL"
		temp := fmt.Sprintf(") INTO %s (%v%v%v) VALUES (",
			session.Engine.Quote(session.Statement.TableName()),
			session.Engine.QuoteStr(),
			strings.Join(colNames, session.Engine.QuoteStr()+", "+session.Engine.QuoteStr()),
			session.Engine.QuoteStr())
		statement = fmt.Sprintf(sql,
			session.Engine.Quote(session.Statement.TableName()),
			session.Engine.QuoteStr(),
			strings.Join(colNames, session.Engine.QuoteStr()+", "+session.Engine.QuoteStr()),
			session.Engine.QuoteStr(),
			strings.Join(colMultiPlaces, temp))
	} else {
		statement = fmt.Sprintf(sql,
			session.Engine.Quote(session.Statement.TableName()),
			session.Engine.QuoteStr(),
			strings.Join(colNames, session.Engine.QuoteStr()+", "+session.Engine.QuoteStr()),
			session.Engine.QuoteStr(),
			strings.Join(colMultiPlaces, "),("))
	}
	res, err := session.exec(statement, args...)
	if err != nil {
		return 0, err
	}

	if cacher := session.Engine.getCacher2(table); cacher != nil && session.Statement.UseCache {
		session.cacheInsert(session.Statement.TableName())
	}

	lenAfterClosures := len(session.afterClosures)
	for i := 0; i < size; i++ {
		elemValue := reflect.Indirect(sliceValue.Index(i)).Addr().Interface()

		// handle AfterInsertProcessor
		if session.IsAutoCommit {
			// !nashtsai! does user expect it's same slice to passed closure when using Before()/After() when insert multi??
			for _, closure := range session.afterClosures {
				closure(elemValue)
			}
			if processor, ok := interface{}(elemValue).(AfterInsertProcessor); ok {
				processor.AfterInsert()
			}
		} else {
			if lenAfterClosures > 0 {
				if value, has := session.afterInsertBeans[elemValue]; has && value != nil {
					*value = append(*value, session.afterClosures...)
				} else {
					afterClosures := make([]func(interface{}), lenAfterClosures)
					copy(afterClosures, session.afterClosures)
					session.afterInsertBeans[elemValue] = &afterClosures
				}
			} else {
				if _, ok := interface{}(elemValue).(AfterInsertProcessor); ok {
					session.afterInsertBeans[elemValue] = nil
				}
			}
		}
	}

	cleanupProcessorsClosures(&session.afterClosures)
	return res.RowsAffected()
}

//对齐所有数据的插入列,在所有数据中看对应列是否为空
func (session *Session) allIsNil(rowsSlicePtr interface{}) (error, bool) {
	return nil, true
}

// InsertMulti insert multiple records
func (session *Session) InsertMulti(rowsSlicePtr interface{}) (int64, error) {
	defer session.resetStatement()
	if session.IsAutoClose {
		defer session.Close()
	}

	sliceValue := reflect.Indirect(reflect.ValueOf(rowsSlicePtr))
	if sliceValue.Kind() != reflect.Slice {
		return 0, ErrParamsType

	}

	if sliceValue.Len() <= 0 {
		return 0, nil
	}

	return session.innerInsertMulti(rowsSlicePtr)
}

func (session *Session) innerInsert(bean interface{}) (int64, error) {
	if err := session.Statement.setRefValue(rValue(bean)); err != nil {
		return 0, err
	}
	if len(session.Statement.TableName()) <= 0 {
		return 0, ErrTableNotFound
	}

	table := session.Statement.RefTable

	// handle BeforeInsertProcessor
	for _, closure := range session.beforeClosures {
		closure(bean)
	}
	cleanupProcessorsClosures(&session.beforeClosures) // cleanup after used

	if processor, ok := interface{}(bean).(BeforeInsertProcessor); ok {
		processor.BeforeInsert()
	}
	// --
	colNames, args, err := genCols(session.Statement.RefTable, session, bean, false, false, false)
	if err != nil {
		return 0, err
	}

	// insert expr columns, override if exists
	exprColumns := session.Statement.getExpr()
	exprColVals := make([]string, 0, len(exprColumns))
	for _, v := range exprColumns {
		// remove the expr columns
		for i, colName := range colNames {
			if colName == v.colName {
				colNames = append(colNames[:i], colNames[i+1:]...)
				args = append(args[:i], args[i+1:]...)
			}
		}

		// append expr column to the end
		colNames = append(colNames, v.colName)
		exprColVals = append(exprColVals, v.expr)
	}

	colPlaces := strings.Repeat("?, ", len(colNames)-len(exprColumns))
	if len(exprColVals) > 0 {
		colPlaces = colPlaces + strings.Join(exprColVals, ", ")
	} else {
		colPlaces = colPlaces[0 : len(colPlaces)-2]
	}

	sqlStr := fmt.Sprintf("INSERT INTO %s (%v%v%v) VALUES (%v)",
		session.Engine.Quote(session.Statement.TableName()),
		session.Engine.dialect.QuoteStr(),
		strings.Join(colNames, session.Engine.Quote(", ")),
		session.Engine.dialect.QuoteStr(),
		colPlaces)

	handleAfterInsertProcessorFunc := func(bean interface{}) {
		if session.IsAutoCommit {
			for _, closure := range session.afterClosures {
				closure(bean)
			}
			if processor, ok := interface{}(bean).(AfterInsertProcessor); ok {
				processor.AfterInsert()
			}
		} else {
			lenAfterClosures := len(session.afterClosures)
			if lenAfterClosures > 0 {
				if value, has := session.afterInsertBeans[bean]; has && value != nil {
					*value = append(*value, session.afterClosures...)
				} else {
					afterClosures := make([]func(interface{}), lenAfterClosures)
					copy(afterClosures, session.afterClosures)
					session.afterInsertBeans[bean] = &afterClosures
				}

			} else {
				if _, ok := interface{}(bean).(AfterInsertProcessor); ok {
					session.afterInsertBeans[bean] = nil
				}
			}
		}
		cleanupProcessorsClosures(&session.afterClosures) // cleanup after used
	}

	// for postgres, many of them didn't implement lastInsertId, so we should
	// implemented it ourself.
	if session.Engine.dialect.DBType() == core.ORACLE && len(table.AutoIncrement) > 0 {
		//fmt.Println(table.AutoIncrement)
		var id int64
		//fmt.Println("before if",id)

		//fmt.Println("after if",id)
		aiValue, err := table.AutoIncrColumn().ValueOf(bean)
		if err != nil {
			session.Engine.logger.Error(err)
		}

		//fmt.Println(!aiValue.CanSet())
		//
		if aiValue == nil || !aiValue.IsValid() || !aiValue.CanSet() {
			return 0, nil
		}
		//fmt.Println("in parameter",id)\

		getSeq := false
		if aiValue.Kind() == reflect.Ptr {
			if aiValue.Pointer() == 0 {
				getSeq = true
			}
		} else {
			if aiValue.Interface().(int64) == 0 {
				getSeq = true
			}
		}

		if getSeq {
			if !session.IdentityInsert {
				if len(table.AutoIncrement) > 0 {
					sql := fmt.Sprintf("select seq_%s.nextval %s from dual", table.Name, table.AutoIncrement)
					//fmt.Println(sql);
					result, err := session.query(sql)
					if err != nil {
						return 1, err
					}
					idByte := result[0][strings.ToUpper(table.AutoIncrement)]
					//fmt.Println(idByte);
					id, err = strconv.ParseInt(string(idByte), 10, 64)
					if err != nil {
						return 0, err
					}
					//id=yizheng.Idt.Idprefix+id 暂不加前缀

					//fmt.Println("in if",id)
				}
				if aiValue.Kind() == reflect.Ptr {
					aiValue.Set(reflect.ValueOf(&id))
					//aiValue.Set(int64ToIntValue(id, aiValue.Type()))
				} else {
					aiValue.Set(int64ToIntValue(id, aiValue.Type()))
				}

			}
			//aiValue.Set(id)
			if err != nil {
				return 0, err
			}
		}
		//fmt.Println(id)
		//fmt.Println(bean)
		colNames, args, err := genCols(session.Statement.RefTable, session, bean, false, false, false)
		//fmt.Println(colNames)
		colPlaces := strings.Repeat("?, ", len(colNames))
		if len(exprColVals) > 0 {
			colPlaces = colPlaces + strings.Join(exprColVals, ", ")
		} else {
			colPlaces = colPlaces[0 : len(colPlaces)-2]
		}

		sqlStr := fmt.Sprintf("INSERT INTO %s (%v%v%v) VALUES (%v)",
			session.Engine.Quote(session.Statement.TableName()),
			session.Engine.QuoteStr(),
			strings.Join(colNames, session.Engine.Quote(", ")),
			session.Engine.QuoteStr(),
			colPlaces)

		res, err := session.exec(sqlStr, args...)
		//fmt.Println(sqlStr, ":", args)
		if err != nil {
			return 0, err
		}

		defer handleAfterInsertProcessorFunc(bean)

		if cacher := session.Engine.getCacher2(table); cacher != nil && session.Statement.UseCache {
			session.cacheInsert(session.Statement.TableName())
		}

		if table.Version != "" && session.Statement.checkVersion {
			verValue, err := table.VersionColumn().ValueOf(bean)
			if err != nil {
				session.Engine.logger.Error(err)
			} else if verValue.IsValid() && verValue.CanSet() {
				verValue.SetInt(1)
			}
		}

		//if len(res) < 1 {
		//	return 0, errors.New("insert no error but not returned id")
		//}

		//for _, rs := range  res {
		//	for j,s:=range rs {
		//		fmt.Println(j, s)
		//	}
		//}

		//session.Commit();
		//return 1, nil
		return res.RowsAffected()
	} else if session.Engine.dialect.DBType() == core.POSTGRES && len(table.AutoIncrement) > 0 {
		//assert table.AutoIncrement != ""
		sqlStr = sqlStr + " RETURNING " + session.Engine.Quote(table.AutoIncrement)
		res, err := session.query(sqlStr, args...)

		if err != nil {
			return 0, err
		}
		handleAfterInsertProcessorFunc(bean)

		if cacher := session.Engine.getCacher2(table); cacher != nil && session.Statement.UseCache {
			session.cacheInsert(session.Statement.TableName())
		}

		if table.Version != "" && session.Statement.checkVersion {
			verValue, err := table.VersionColumn().ValueOf(bean)
			if err != nil {
				session.Engine.logger.Error(err)
			} else if verValue.IsValid() && verValue.CanSet() {
				verValue.SetInt(1)
			}
		}

		if len(res) < 1 {
			return 0, errors.New("insert no error but not returned id")
		}

		idByte := res[0][table.AutoIncrement]
		id, err := strconv.ParseInt(string(idByte), 10, 64)
		if err != nil || id <= 0 {
			return 1, err
		}

		aiValue, err := table.AutoIncrColumn().ValueOf(bean)
		if err != nil {
			session.Engine.logger.Error(err)
		}

		if aiValue == nil || !aiValue.IsValid() || !aiValue.CanSet() {
			return 1, nil
		}

		aiValue.Set(int64ToIntValue(id, aiValue.Type()))

		return 1, nil
	} else {
		res, err := session.exec(sqlStr, args...)
		if err != nil {
			return 0, err
		}

		defer handleAfterInsertProcessorFunc(bean)

		if cacher := session.Engine.getCacher2(table); cacher != nil && session.Statement.UseCache {
			session.cacheInsert(session.Statement.TableName())
		}

		if table.Version != "" && session.Statement.checkVersion {
			verValue, err := table.VersionColumn().ValueOf(bean)
			if err != nil {
				session.Engine.logger.Error(err)
			} else if verValue.IsValid() && verValue.CanSet() {
				if verValue.Kind() == reflect.Ptr {
					i := int64(1)
					verValue.Set(reflect.ValueOf(&i))
				} else {
					verValue.SetInt(1)
				}

			}
		}

		if table.AutoIncrement == "" {
			return res.RowsAffected()
		}

		var id int64
		id, err = res.LastInsertId()

		if err != nil || id <= 0 {
			return res.RowsAffected()
		}

		aiValue, err := table.AutoIncrColumn().ValueOf(bean)
		if err != nil {
			session.Engine.logger.Error(err)
		}

		if aiValue == nil || !aiValue.IsValid() || !aiValue.CanSet() {
			return res.RowsAffected()
		}

		aiValue.Set(int64ToIntValue(id, aiValue.Type()))
		return res.RowsAffected()
	}
}

// InsertOne insert only one struct into database as a record.
// The in parameter bean must a struct or a point to struct. The return
// parameter is inserted and error
func (session *Session) InsertOne(bean interface{}) (int64, error) {
	defer session.resetStatement()
	if session.IsAutoClose {
		defer session.Close()
	}

	return session.innerInsert(bean)
}

func (session *Session) cacheInsert(tables ...string) error {
	if session.Statement.RefTable == nil {
		return ErrCacheFailed
	}

	table := session.Statement.RefTable
	cacher := session.Engine.getCacher2(table)

	for _, t := range tables {
		session.Engine.logger.Debug("[cache] clear sql:", t)
		cacher.ClearIds(t)
	}

	return nil
}
