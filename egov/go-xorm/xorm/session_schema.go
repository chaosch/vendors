// Copyright 2016 The Xorm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xorm

import (
	"database/sql"
	"errors"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"reflect"
	"strings"

	"egov/go-xorm/core"
)

// Ping test if database is ok
func (session *Session) Ping() error {
	defer session.resetStatement()
	if session.IsAutoClose {
		defer session.Close()
	}

	return session.DB().Ping()
}

// CreateTable create a table according a bean
func (session *Session) CreateTable(bean interface{}) error {
	v := rValue(bean)

	if err := session.Statement.setRefValue(v); err != nil {
		return err
	}
	if !session.Engine.IdentityInsert {
		session.Statement.RefTable.AutoIncrement = session.Statement.RefTable.PKColumns()[0].Name
		session.Statement.RefTable.PKColumns()[0].IsAutoIncrement = true
	}
	defer session.resetStatement()
	if session.IsAutoClose {
		defer session.Close()
	}

	return session.createOneTable()
}

// alter auto increment
func (session *Session) AlterIncrement(bean interface{}) error {
	v := rValue(bean)

	if err := session.Statement.setRefValue(v); err != nil {
		return err
	}

	defer session.resetStatement()
	if session.IsAutoClose {
		defer session.Close()
	}

	return session.alterOneIncrement()
}

// CreateIndexes create indexes
func (session *Session) CreateIndexes(bean interface{}) error {
	v := rValue(bean)
	if err := session.Statement.setRefValue(v); err != nil {
		return err
	}

	defer session.resetStatement()
	if session.IsAutoClose {
		defer session.Close()
	}

	sqls := session.Statement.genIndexSQL()
	for _, sqlStr := range sqls {
		_, err := session.exec(sqlStr)
		if err != nil {
			return err
		}
	}
	return nil
}

// CreateUniques create uniques
func (session *Session) CreateUniques(bean interface{}) error {
	v := rValue(bean)
	if err := session.Statement.setRefValue(v); err != nil {
		return err
	}

	defer session.resetStatement()
	if session.IsAutoClose {
		defer session.Close()
	}

	sqls := session.Statement.genUniqueSQL()
	for _, sqlStr := range sqls {
		_, err := session.exec(sqlStr)
		if err != nil {
			return err
		}
	}
	return nil
}

func (session *Session) createOneTable() error {
	sqlStr := session.Statement.genCreateTableSQL()
	var sqls []string
	//fmt.Println(sqlStr)
	//if (strings.Contains(sqlStr,";")&&session.Engine.dialect.DBType() == core.ORACLE){
	sqls = strings.Split(sqlStr, ";")
	//}else{
	//sqls=sqlStr
	//	sqls=append(sqls,sqlStr)
	//}
	var Sql string
	for _, Sql = range sqls {
		if len(Sql) > 0 {
			//fmt.Println(Sql)
			//session.Tx.Exec()
			_, err := session.exec(Sql)
			if err != nil {
				return err
			}
		}
	}
	return nil

	//_, err2:=session.exec(sqlSeq)
	//return err2
}

func (session *Session) alterOneIncrement() error {
	sqlStr := session.Statement.genAlterIncrementSQL()
	var sqls []string
	//fmt.Println(sqlStr)
	//if (strings.Contains(sqlStr,";")&&session.Engine.dialect.DBType() == core.ORACLE){
	sqls = strings.Split(sqlStr, ";")
	//}else{
	//sqls=sqlStr
	//	sqls=append(sqls,sqlStr)
	//}
	var Sql string
	for _, Sql = range sqls {
		if len(Sql) > 0 {
			//fmt.Println(Sql)
			_, err := session.exec(Sql)
			if err != nil {
				return err
			}
		}
	}
	return nil

	//_, err2:=session.exec(sqlSeq)
	//return err2
}

// to be deleted
func (session *Session) createAll() error {
	if session.IsAutoClose {
		defer session.Close()
	}

	for _, table := range session.Engine.Tables {
		session.Statement.RefTable = table
		session.Statement.tableName = table.Name
		err := session.createOneTable()
		session.resetStatement()
		if err != nil {
			return err
		}
	}
	return nil
}

// DropIndexes drop indexes
func (session *Session) DropIndexes(bean interface{}) error {
	v := rValue(bean)
	if err := session.Statement.setRefValue(v); err != nil {
		return err
	}

	defer session.resetStatement()
	if session.IsAutoClose {
		defer session.Close()
	}

	sqls := session.Statement.genDelIndexSQL()
	for _, sqlStr := range sqls {
		_, err := session.exec(sqlStr)
		if err != nil {
			return err
		}
	}
	return nil
}

// DropTable drop table will drop table if exist, if drop failed, it will return error
func (session *Session) DropTable(beanOrTableName interface{}) error {
	tableName, err := session.Engine.tableName(beanOrTableName)
	if err != nil {
		return err
	}

	var needDrop = true
	if !session.Engine.dialect.SupportDropIfExists() {
		sqlStr, args := session.Engine.dialect.TableCheckSql(tableName)
		results, err := session.query(sqlStr, args...)
		if err != nil {
			return err
		}
		needDrop = len(results) > 0
	}

	if needDrop {
		sqlStr := session.Engine.Dialect().DropTableSql(tableName)
		_, err = session.exec(sqlStr)
		return err
	}
	return nil
}

// IsTableExist if a table is exist
func (session *Session) IsTableExist(beanOrTableName interface{}) (bool, error) {
	tableName, err := session.Engine.tableName(beanOrTableName)
	if err != nil {
		return false, err
	}

	return session.isTableExist(tableName)
}

func (session *Session) isTableExist(tableName string) (bool, error) {
	defer session.resetStatement()
	if session.IsAutoClose {
		defer session.Close()
	}
	sqlStr, args := session.Engine.dialect.TableCheckSql(tableName)
	results, err := session.query(sqlStr, args...)
	return len(results) > 0, err
}

// IsTableEmpty if table have any records
func (session *Session) IsTableEmpty(bean interface{}) (bool, error) {
	v := rValue(bean)
	t := v.Type()

	if t.Kind() == reflect.String {
		return session.isTableEmpty(bean.(string))
	} else if t.Kind() == reflect.Struct {
		rows, err := session.Count(bean)
		return rows == 0, err
	}
	return false, errors.New("bean should be a struct or struct's point")
}

func (session *Session) isTableEmpty(tableName string) (bool, error) {
	defer session.resetStatement()
	if session.IsAutoClose {
		defer session.Close()
	}

	var total int64
	sqlStr := fmt.Sprintf("select count(*) from %s", session.Engine.Quote(tableName))
	err := session.DB().QueryRow(sqlStr).Scan(&total)
	session.saveLastSQL(sqlStr)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
		}
		return true, err
	}

	return total == 0, nil
}

func (session *Session) isIndexExist(tableName, idxName string, unique bool) (bool, error) {
	defer session.resetStatement()
	if session.IsAutoClose {
		defer session.Close()
	}
	var idx string
	if unique {
		idx = uniqueName(tableName, idxName)
	} else {
		idx = indexName(tableName, idxName)
	}
	sqlStr, args := session.Engine.dialect.IndexCheckSql(tableName, idx)
	results, err := session.query(sqlStr, args...)
	return len(results) > 0, err
}

// find if index is exist according cols
func (session *Session) isIndexExist2(tableName string, dst *core.Index, unique bool) (bool, error) {
	defer session.resetStatement()
	if session.IsAutoClose {
		defer session.Close()
	}

	indexes, err := session.Engine.dialect.GetIndexes(tableName)
	if err != nil {
		return false, err
	}

	for _, index := range indexes {
		typeEq, nameEq, colsEq := index.Equal(dst)

		if typeEq && nameEq && colsEq {
			//if unique&&index.Type==core.IndexType||!unique&&index.Type==core.UniqueType{
			//	//虽然有索引，但类型不对，返回false之前应drop掉该索引
			//	DropIndexSql:=""
			//	switch session.Engine.dialect.DBType() {
			//	case core.MYSQL:
			//		DropIndexSql="drop index "+index.Name+" on "+tableName
			//	default:
			//		DropIndexSql="drop index "+index.Name
			//	}
			//	session.Exec(DropIndexSql)
			//}
			if unique {
				//返回false之前删掉该
				return index.Type == core.UniqueType, nil
			}
			return index.Type == core.IndexType, nil
		}
	}
	return false, nil
}

func (session *Session) CreateOrRenameIndex() {

}

func (session *Session) isIndexExist3(tableName string, dstIndex *core.Index) (bool, error, string) {
	defer session.resetStatement()
	alterSql := ""
	if session.IsAutoClose {
		defer session.Close()
	}

	indexes, err := session.Engine.dialect.GetIndexes(tableName)
	if err != nil {
		return false, err, alterSql
	}

	for _, index := range indexes {
		typeEq, nameEq, colsEq := index.Equal(dstIndex)
		if colsEq && typeEq && !nameEq { //索引内容一致只是名称不一致，修改名称
			alterSql := session.Engine.dialect.CreateRenameIndexSql(tableName, index, dstIndex)
			return false, nil, alterSql
		}
		if colsEq && nameEq && typeEq {
			return true, nil, ""
		}
	}
	return false, nil, ""
}

func (session *Session) addColumn(col *core.Column) error {
	//	colName:=col.Name
	defer session.resetStatement()
	if session.IsAutoClose {
		defer session.Close()
	}

	//col := session.Statement.RefTable.GetColumn(colName)
	var res error
	x := simplifiedchinese.GBK.NewEncoder()
	sqls, args := session.Statement.genAddColumnStr(col)
	for _, sql := range sqls {
		if session.Engine.dialect.DBType() == core.ORACLE {
			sql, _ = x.String(sql)
			session = session.Engine.Table(col.TableName)
		}
		_, err := session.exec(sql, args...)
		if err != nil {
			if res != nil {
				res = errors.New(res.Error() + ":" + err.Error())
			} else {
				res = err
			}
		}
	}
	return res
}

func (session *Session) AddColumn(col *core.Column) error {
	//	colName:=col.Name
	return session.addColumn(col)
}

func (session *Session) addIndex(tableName, idxName string) error {
	defer session.resetStatement()
	if session.IsAutoClose {
		defer session.Close()
	}
	index, ok := session.Statement.RefTable.Indexes[idxName]
	if !ok {
		return nil
	}
	sqlStr := session.Engine.dialect.CreateIndexSql(tableName, index)

	//if !strings.HasPrefix(sqlStr,";"){
	//	sqlStr=sqlStr+";"
	//}
	_, err := session.exec(sqlStr)
	return err
}

func (session *Session) addUnique(tableName, uqeName string) error {
	defer session.resetStatement()
	if session.IsAutoClose {
		defer session.Close()
	}
	index := session.Statement.RefTable.Indexes[uqeName]
	sqlStr := session.Engine.dialect.CreateIndexSql(tableName, index)
	_, err := session.exec(sqlStr)
	return err
}

// To be deleted
func (session *Session) dropAll() error {
	defer session.resetStatement()
	if session.IsAutoClose {
		defer session.Close()
	}

	for _, table := range session.Engine.Tables {
		session.Statement.Init()
		session.Statement.RefTable = table
		sqlStr := session.Engine.Dialect().DropTableSql(session.Statement.TableName())
		_, err := session.exec(sqlStr)
		if err != nil {
			return err
		}
	}
	return nil
}

// Sync2 synchronize structs to database tables
func (session *Session) Sync2(beans ...interface{}) error {
	engine := session.Engine

	tables, err := engine.DBMetas()
	if err != nil {
		return err
	}

	var structTables []*core.Table

	for _, bean := range beans {
		v := rValue(bean)
		table, err := engine.mapType(v)
		if err != nil {
			return err
		}
		structTables = append(structTables, table)
		var tbName = session.tbNameNoSchema(table)

		var oriTable *core.Table
		for _, tb := range tables {
			if strings.EqualFold(tb.Name, tbName) {
				oriTable = tb
				break
			}
		}

		if oriTable == nil {
			err = session.StoreEngine(session.Statement.StoreEngine).CreateTable(bean)
			if err != nil {
				return err
			}

			err = session.CreateUniques(bean)
			if err != nil {
				return err
			}

			err = session.CreateIndexes(bean)
			if err != nil {
				return err
			}
		} else {
			for _, col := range table.Columns() {
				var oriCol *core.Column
				for _, col2 := range oriTable.Columns() {
					if strings.EqualFold(col.Name, col2.Name) {
						oriCol = col2
						break
					}
				}

				if oriCol != nil {
					expectedType := engine.dialect.SqlType(col)
					curType := engine.dialect.SqlType(oriCol)
					if expectedType != curType {
						if expectedType == core.Text &&
							strings.HasPrefix(curType, core.Varchar) {
							// currently only support mysql & postgres
							if engine.dialect.DBType() == core.MYSQL ||
								engine.dialect.DBType() == core.POSTGRES {
								engine.logger.Infof("Table %s column %s change type from %s to %s\n",
									tbName, col.Name, curType, expectedType)
								_, err = engine.Exec(engine.dialect.ModifyColumnSql(table.Name, col))
							} else {
								engine.logger.Warnf("Table %s column %s db type is %s, struct type is %s\n",
									tbName, col.Name, curType, expectedType)
							}
						} else if strings.HasPrefix(curType, core.Varchar) && strings.HasPrefix(expectedType, core.Varchar) {
							if engine.dialect.DBType() == core.MYSQL {
								if oriCol.Length < col.Length {
									engine.logger.Infof("Table %s column %s change type from varchar(%d) to varchar(%d)\n",
										tbName, col.Name, oriCol.Length, col.Length)
									_, err = engine.Exec(engine.dialect.ModifyColumnSql(table.Name, col))
								}
							}
						} else {
							if !(strings.HasPrefix(curType, expectedType) && curType[len(expectedType)] == '(') {
								engine.logger.Warnf("Table %s column %s db type is %s, struct type is %s",
									tbName, col.Name, curType, expectedType)
							}
						}
					} else if expectedType == core.Varchar {
						if engine.dialect.DBType() == core.MYSQL {
							if oriCol.Length < col.Length {
								engine.logger.Infof("Table %s column %s change type from varchar(%d) to varchar(%d)\n",
									tbName, col.Name, oriCol.Length, col.Length)
								_, err = engine.Exec(engine.dialect.ModifyColumnSql(table.Name, col))
							}
						}
					}
					if col.Default != oriCol.Default {
						engine.logger.Warnf("Table %s Column %s db default is %s, struct default is %s",
							tbName, col.Name, oriCol.Default, col.Default)
					}
					if col.Nullable != oriCol.Nullable {
						engine.logger.Warnf("Table %s Column %s db nullable is %v, struct nullable is %v",
							tbName, col.Name, oriCol.Nullable, col.Nullable)
					}
				} else {
					session := engine.NewSession()
					session.Statement.RefTable = table
					session.Statement.tableName = tbName
					defer session.Close()
					err = session.addColumn(col)
				}
				if err != nil {
					return err
				}
			}

			var foundIndexNames = make(map[string]bool)
			var addedNames = make(map[string]*core.Index)

			for name, index := range table.Indexes {
				var oriIndex *core.Index
				for name2, index2 := range oriTable.Indexes {
					typeEq, nameEq, colsEq := index.Equal(index2)
					if typeEq && nameEq && colsEq {
						oriIndex = index2
						foundIndexNames[name2] = true
						break
					}
				}

				if oriIndex != nil {
					if oriIndex.Type != index.Type {
						sql := engine.dialect.DropIndexSql(tbName, oriIndex)
						_, err = engine.Exec(sql)
						if err != nil {
							return err
						}
						oriIndex = nil
					}
				}

				if oriIndex == nil {
					addedNames[name] = index
				}
			}

			for name2, index2 := range oriTable.Indexes {
				if _, ok := foundIndexNames[name2]; !ok {
					sql := engine.dialect.DropIndexSql(tbName, index2)
					_, err = engine.Exec(sql)
					if err != nil {
						return err
					}
				}
			}

			for name, index := range addedNames {
				if index.Type == core.UniqueType {
					session := engine.NewSession()
					session.Statement.RefTable = table
					session.Statement.tableName = tbName
					defer session.Close()
					err = session.addUnique(tbName, name)
				} else if index.Type == core.IndexType {
					session := engine.NewSession()
					session.Statement.RefTable = table
					session.Statement.tableName = tbName
					defer session.Close()
					err = session.addIndex(tbName, name)
				}
				if err != nil {
					return err
				}
			}
		}
	}

	for _, table := range tables {
		var oriTable *core.Table
		for _, structTable := range structTables {
			if strings.EqualFold(table.Name, session.tbNameNoSchema(structTable)) {
				oriTable = structTable
				break
			}
		}

		if oriTable == nil {
			//engine.LogWarnf("Table %s has no struct to mapping it", table.Name)
			continue
		}

		for _, colName := range table.ColumnsSeq() {
			if oriTable.GetColumn(colName) == nil {
				engine.logger.Warnf("Table %s has column %s but struct has not related field", table.Name, colName)
			}
		}
	}
	return nil
}
