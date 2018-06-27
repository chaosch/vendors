// Copyright 2015 The Xorm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xorm

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"egov/go-xorm/core"
)

var (
	oracleReservedWords = map[string]bool{
		"ACCESS":                    true,
		"ACCOUNT":                   true,
		"ACTIVATE":                  true,
		"ADD":                       true,
		"ADMIN":                     true,
		"ADVISE":                    true,
		"AFTER":                     true,
		"ALL":                       true,
		"ALL_ROWS":                  true,
		"ALLOCATE":                  true,
		"ALTER":                     true,
		"ANALYZE":                   true,
		"AND":                       true,
		"ANY":                       true,
		"ARCHIVE":                   true,
		"ARCHIVELOG":                true,
		"ARRAY":                     true,
		"AS":                        true,
		"ASC":                       true,
		"AT":                        true,
		"AUDIT":                     true,
		"AUTHENTICATED":             true,
		"AUTHORIZATION":             true,
		"AUTOEXTEND":                true,
		"AUTOMATIC":                 true,
		"BACKUP":                    true,
		"BECOME":                    true,
		"BEFORE":                    true,
		"BEGIN":                     true,
		"BETWEEN":                   true,
		"BFILE":                     true,
		"BITMAP":                    true,
		"BLOB":                      true,
		"BLOCK":                     true,
		"BODY":                      true,
		"BY":                        true,
		"CACHE":                     true,
		"CACHE_INSTANCES":           true,
		"CANCEL":                    true,
		"CASCADE":                   true,
		"CAST":                      true,
		"CFILE":                     true,
		"CHAINED":                   true,
		"CHANGE":                    true,
		"CHAR":                      true,
		"CHAR_CS":                   true,
		"CHARACTER":                 true,
		"CHECK":                     true,
		"CHECKPOINT":                true,
		"CHOOSE":                    true,
		"CHUNK":                     true,
		"CLEAR":                     true,
		"CLOB":                      true,
		"CLONE":                     true,
		"CLOSE":                     true,
		"CLOSE_CACHED_OPEN_CURSORS": true,
		"CLUSTER":                   true,
		"COALESCE":                  true,
		"COLUMN":                    true,
		"COLUMNS":                   true,
		"COMMENT":                   true,
		"COMMIT":                    true,
		"COMMITTED":                 true,
		"COMPATIBILITY":             true,
		"COMPILE":                   true,
		"COMPLETE":                  true,
		"COMPOSITE_LIMIT":           true,
		"COMPRESS":                  true,
		"COMPUTE":                   true,
		"CONNECT":                   true,
		"CONNECT_TIME":              true,
		"CONSTRAINT":                true,
		"CONSTRAINTS":               true,
		"CONTENTS":                  true,
		"CONTINUE":                  true,
		"CONTROLFILE":               true,
		"CONVERT":                   true,
		"COST":                      true,
		"CPU_PER_CALL":              true,
		"CPU_PER_SESSION":           true,
		"CREATE":                    true,
		"CURRENT":                   true,
		"CURRENT_SCHEMA":            true,
		"CURREN_USER":               true,
		"CURSOR":                    true,
		"CYCLE":                     true,
		"DANGLING":                  true,
		"DATABASE":                  true,
		"DATAFILE":                  true,
		"DATAFILES":                 true,
		"DATAOBJNO":                 true,
		"DATE":                      true,
		"DBA":                       true,
		"DBHIGH":                    true,
		"DBLOW":                     true,
		"DBMAC":                     true,
		"DEALLOCATE":                true,
		"DEBUG":                     true,
		"DEC":                       true,
		"DECIMAL":                   true,
		"DECLARE":                   true,
		"DEFAULT":                   true,
		"DEFERRABLE":                true,
		"DEFERRED":                  true,
		"DEGREE":                    true,
		"DELETE":                    true,
		"DEREF":                     true,
		"DESC":                      true,
		"DIRECTORY":                 true,
		"DISABLE":                   true,
		"DISCONNECT":                true,
		"DISMOUNT":                  true,
		"DISTINCT":                  true,
		"DISTRIBUTED":               true,
		"DML":                       true,
		"DOUBLE":                    true,
		"DROP":                      true,
		"DUMP":                      true,
		"EACH":                      true,
		"ELSE":                      true,
		"ENABLE":                    true,
		"END":                       true,
		"ENFORCE":                   true,
		"ENTRY":                     true,
		"ESCAPE":                    true,
		"EXCEPT":                    true,
		"EXCEPTIONS":                true,
		"EXCHANGE":                  true,
		"EXCLUDING":                 true,
		"EXCLUSIVE":                 true,
		"EXECUTE":                   true,
		"EXISTS":                    true,
		"EXPIRE":                    true,
		"EXPLAIN":                   true,
		"EXTENT":                    true,
		"EXTENTS":                   true,
		"EXTERNALLY":                true,
		"FAILED_LOGIN_ATTEMPTS":     true,
		"FALSE":                     true,
		"FAST":                      true,
		"FILE":                      true,
		"FIRST_ROWS":                true,
		"FLAGGER":                   true,
		"FLOAT":                     true,
		"FLOB":                      true,
		"FLUSH":                     true,
		"FOR":                       true,
		"FORCE":                     true,
		"FOREIGN":                   true,
		"FREELIST":                  true,
		"FREELISTS":                 true,
		"FROM":                      true,
		"FULL":                      true,
		"FUNCTION":                  true,
		"GLOBAL":                    true,
		"GLOBALLY":                  true,
		"GLOBAL_NAME":               true,
		"GRANT":                     true,
		"GROUP":                     true,
		"GROUPS":                    true,
		"HASH":                      true,
		"HASHKEYS":                  true,
		"HAVING":                    true,
		"HEADER":                    true,
		"HEAP":                      true,
		"IDENTIFIED":                true,
		"IDGENERATORS":              true,
		"IDLE_TIME":                 true,
		"IF":                        true,
		"IMMEDIATE":                 true,
		"IN":                        true,
		"INCLUDING":                 true,
		"INCREMENT":                 true,
		"INDEX":                     true,
		"INDEXED":                   true,
		"INDEXES":                   true,
		"INDICATOR":                 true,
		"IND_PARTITION":             true,
		"INITIAL":                   true,
		"INITIALLY":                 true,
		"INITRANS":                  true,
		"INSERT":                    true,
		"INSTANCE":                  true,
		"INSTANCES":                 true,
		"INSTEAD":                   true,
		"INT":                       true,
		"INTEGER":                   true,
		"INTERMEDIATE":              true,
		"INTERSECT":                 true,
		"INTO":                      true,
		"IS":                        true,
		"ISOLATION":                 true,
		"ISOLATION_LEVEL":           true,
		"KEEP":                      true,
		"KEY":                       true,
		"KILL":                      true,
		"LABEL":                     true,
		"LAYER":                     true,
		"LESS":                      true,
		"LEVEL":                     true,
		"LIBRARY":                   true,
		"LIKE":                      true,
		"LIMIT":                     true,
		"LINK":                      true,
		"LIST":                      true,
		"LOB":                       true,
		"LOCAL":                     true,
		"LOCK":                      true,
		"LOCKED":                    true,
		"LOG":                       true,
		"LOGFILE":                   true,
		"LOGGING":                   true,
		"LOGICAL_READS_PER_CALL":    true,
		"LOGICAL_READS_PER_SESSION": true,
		"LONG":                      true,
		"MANAGE":                    true,
		"MASTER":                    true,
		"MAX":                       true,
		"MAXARCHLOGS":               true,
		"MAXDATAFILES":              true,
		"MAXEXTENTS":                true,
		"MAXINSTANCES":              true,
		"MAXLOGFILES":               true,
		"MAXLOGHISTORY":             true,
		"MAXLOGMEMBERS":             true,
		"MAXSIZE":                   true,
		"MAXTRANS":                  true,
		"MAXVALUE":                  true,
		"MIN":                       true,
		"MEMBER":                    true,
		"MINIMUM":                   true,
		"MINEXTENTS":                true,
		"MINUS":                     true,
		"MINVALUE":                  true,
		"MLSLABEL":                  true,
		"MLS_LABEL_FORMAT":          true,
		"MODE":                      true,
		"MODIFY":                    true,
		"MOUNT":                     true,
		"MOVE":                      true,
		"MTS_DISPATCHERS":           true,
		"MULTISET":                  true,
		"NATIONAL":                  true,
		"NCHAR":                     true,
		"NCHAR_CS":                  true,
		"NCLOB":                     true,
		"NEEDED":                    true,
		"NESTED":                    true,
		"NETWORK":                   true,
		"NEW":                       true,
		"NEXT":                      true,
		"NOARCHIVELOG":              true,
		"NOAUDIT":                   true,
		"NOCACHE":                   true,
		"NOCOMPRESS":                true,
		"NOCYCLE":                   true,
		"NOFORCE":                   true,
		"NOLOGGING":                 true,
		"NOMAXVALUE":                true,
		"NOMINVALUE":                true,
		"NONE":                      true,
		"NOORDER":                   true,
		"NOOVERRIDE":                true,
		"NOPARALLEL":                true,
		"NOREVERSE":                 true,
		"NORMAL":                    true,
		"NOSORT":                    true,
		"NOT":                       true,
		"NOTHING":                   true,
		"NOWAIT":                    true,
		"NULL":                      true,
		"NUMBER":                    true,
		"NUMERIC":                   true,
		"NVARCHAR2":                 true,
		"OBJECT":                    true,
		"OBJNO":                     true,
		"OBJNO_REUSE":               true,
		"OF":                        true,
		"OFF":                       true,
		"OFFLINE":                   true,
		"OID":                       true,
		"OIDINDEX":                  true,
		"OLD":                       true,
		"ON":                        true,
		"ONLINE":                    true,
		"ONLY":                      true,
		"OPCODE":                    true,
		"OPEN":                      true,
		"OPTIMAL":                   true,
		"OPTIMIZER_GOAL":            true,
		"OPTION":                    true,
		"OR":                        true,
		"ORDER":                     true,
		"ORGANIZATION":              true,
		"OSLABEL":                   true,
		"OVERFLOW":                  true,
		"OWN":                       true,
		"PACKAGE":                   true,
		"PARALLEL":                  true,
		"PARTITION":                 true,
		"PASSWORD":                  true,
		"PASSWORD_GRACE_TIME":       true,
		"PASSWORD_LIFE_TIME":        true,
		"PASSWORD_LOCK_TIME":        true,
		"PASSWORD_REUSE_MAX":        true,
		"PASSWORD_REUSE_TIME":       true,
		"PASSWORD_VERIFY_FUNCTION":  true,
		"PCTFREE":                   true,
		"PCTINCREASE":               true,
		"PCTTHRESHOLD":              true,
		"PCTUSED":                   true,
		"PCTVERSION":                true,
		"PERCENT":                   true,
		"PERMANENT":                 true,
		"PLAN":                      true,
		"PLSQL_DEBUG":               true,
		"POST_TRANSACTION":          true,
		"PRECISION":                 true,
		"PRESERVE":                  true,
		"PRIMARY":                   true,
		"PRIOR":                     true,
		"PRIVATE":                   true,
		"PRIVATE_SGA":               true,
		"PRIVILEGE":                 true,
		"PRIVILEGES":                true,
		"PROCEDURE":                 true,
		"PROFILE":                   true,
		"PUBLIC":                    true,
		"PURGE":                     true,
		"QUEUE":                     true,
		"QUOTA":                     true,
		"RANGE":                     true,
		"RAW":                       true,
		"RBA":                       true,
		"READ":                      true,
		"READUP":                    true,
		"REAL":                      true,
		"REBUILD":                   true,
		"RECOVER":                   true,
		"RECOVERABLE":               true,
		"RECOVERY":                  true,
		"REF":                       true,
		"REFERENCES":                true,
		"REFERENCING":               true,
		"REFRESH":                   true,
		"RENAME":                    true,
		"REPLACE":                   true,
		"RESET":                     true,
		"RESETLOGS":                 true,
		"RESIZE":                    true,
		"RESOURCE":                  true,
		"RESTRICTED":                true,
		"RETURN":                    true,
		"RETURNING":                 true,
		"REUSE":                     true,
		"REVERSE":                   true,
		"REVOKE":                    true,
		"ROLE":                      true,
		"ROLES":                     true,
		"ROLLBACK":                  true,
		"ROW":                       true,
		"ROWID":                     true,
		"ROWNUM":                    true,
		"ROWS":                      true,
		"RULE":                      true,
		"SAMPLE":                    true,
		"SAVEPOINT":                 true,
		"SB4":                       true,
		"SCAN_INSTANCES":            true,
		"SCHEMA":                    true,
		"SCN":                       true,
		"SCOPE":                     true,
		"SD_ALL":                    true,
		"SD_INHIBIT":                true,
		"SD_SHOW":                   true,
		"SEGMENT":                   true,
		"SEG_BLOCK":                 true,
		"SEG_FILE":                  true,
		"SELECT":                    true,
		"SEQUENCE":                  true,
		"SERIALIZABLE":              true,
		"SESSION":                   true,
		"SESSION_CACHED_CURSORS":    true,
		"SESSIONS_PER_USER":         true,
		"SET":                       true,
		"SHARE":                     true,
		"SHARED":                    true,
		"SHARED_POOL":               true,
		"SHRINK":                    true,
		"SIZE":                      true,
		"SKIP":                      true,
		"SKIP_UNUSABLE_INDEXES":     true,
		"SMALLINT":                  true,
		"SNAPSHOT":                  true,
		"SOME":                      true,
		"SORT":                      true,
		"SPECIFICATION":             true,
		"SPLIT":                     true,
		"SQL_TRACE":                 true,
		"STANDBY":                   true,
		"START":                     true,
		"STATEMENT_ID":              true,
		"STATISTICS":                true,
		"STOP":                      true,
		"STORAGE":                   true,
		"STORE":                     true,
		"STRUCTURE":                 true,
		"SUCCESSFUL":                true,
		"SWITCH":                    true,
		"SYS_OP_ENFORCE_NOT_NULL$":  true,
		"SYS_OP_NTCIMG$":            true,
		"SYNONYM":                   true,
		"SYSDATE":                   true,
		"SYSDBA":                    true,
		"SYSOPER":                   true,
		"SYSTEM":                    true,
		"TABLE":                     true,
		"TABLES":                    true,
		"TABLESPACE":                true,
		"TABLESPACE_NO":             true,
		"TABNO":                     true,
		"TEMPORARY":                 true,
		"THAN":                      true,
		"THE":                       true,
		"THEN":                      true,
		"THREAD":                    true,
		"TIMESTAMP":                 true,
		"TIME":                      true,
		"TO":                        true,
		"TOPLEVEL":                  true,
		"TRACE":                     true,
		"TRACING":                   true,
		"TRANSACTION":               true,
		"TRANSITIONAL":              true,
		"TRIGGER":                   true,
		"TRIGGERS":                  true,
		"TRUE":                      true,
		"TRUNCATE":                  true,
		"TX":                        true,
		"TYPE":                      true,
		"UB2":                       true,
		"UBA":                       true,
		"UID":                       true,
		"UNARCHIVED":                true,
		"UNDO":                      true,
		"UNION":                     true,
		"UNIQUE":                    true,
		"UNLIMITED":                 true,
		"UNLOCK":                    true,
		"UNRECOVERABLE":             true,
		"UNTIL":                     true,
		"UNUSABLE":                  true,
		"UNUSED":                    true,
		"UPDATABLE":                 true,
		"UPDATE":                    true,
		"USAGE":                     true,
		"USE":                       true,
		"USER":                      true,
		"USING":                     true,
		"VALIDATE":                  true,
		"VALIDATION":                true,
		"VALUE":                     true,
		"VALUES":                    true,
		"VARCHAR":                   true,
		"VARCHAR2":                  true,
		"VARYING":                   true,
		"VIEW":                      true,
		"WHEN":                      true,
		"WHENEVER":                  true,
		"WHERE":                     true,
		"WITH":                      true,
		"WITHOUT":                   true,
		"WORK":                      true,
		"WRITE":                     true,
		"WRITEDOWN":                 true,
		"WRITEUP":                   true,
		"XID":                       true,
		"YEAR":                      true,
		"ZONE":                      true,
	}
)

type oracle struct {
	core.Base
}


func (db *oracle) SetTableComment(d map[string]string, t map[string]string) {
	db.DataTable = t
	db.Dictionaries = d
}

func (db *oracle) Init(d *core.DB, uri *core.Uri, drivername, dataSourceName string) error {
	return db.Base.Init(d, db, uri, drivername, dataSourceName)
}

func (db *oracle) SqlType(c *core.Column) string {
	var res string
	switch t := c.SQLType.Name; t {

	case core.Bool:
		res = "NUMBER"
		//if strings.EqualFold(c.Default, "1") { //bool型统一数字设置初始值
		//	c.Default = "1"
		//} else {
		//	c.Default = "0"
		//}
	case core.Bit, core.TinyInt, core.SmallInt, core.MediumInt, core.Int, core.Integer, core.BigInt, core.Bool, core.Serial, core.BigSerial:
		res = "NUMBER"
	case core.Binary, core.VarBinary, core.Blob, core.TinyBlob, core.MediumBlob, core.LongBlob, core.Bytea:
		return core.Blob
	case core.Time, core.DateTime, core.TimeStamp:
		res = core.TimeStamp
	case core.TimeStampz:
		res = "TIMESTAMP WITH TIME ZONE"
	case core.Float, core.Double, core.Numeric, core.Decimal:
		res = "NUMBER"
	case core.Text, core.MediumText, core.LongText, core.Json:
		res = "CLOB"
	case core.Char, core.Varchar, core.TinyText:
		if c.Length > 1000 {
			c.SQLType.Name = core.Text
			res = core.Text
			c.Length = 0
		} else {
			res = "VARCHAR2"
		}
	default:
		res = t
	}

	hasLen1 := c.Length > 0
	hasLen2 := c.Length2 > 0

	if hasLen2 {
		res += "(" + strconv.Itoa(c.Length) + "," + strconv.Itoa(c.Length2) + ")"
	} else if hasLen1 {
		res += "(" + strconv.Itoa(c.Length) + ")"
	}
	return res
}

func (db *oracle) AutoIncrStr() string {
	return "AUTO_INCREMENT"
}

func (db *oracle) SupportInsertMany() bool {
	return true
}

func (db *oracle) IsReserved(name string) bool {
	_, ok := oracleReservedWords[name]
	return ok
}

func (db *oracle) Quote(name string) string {
	//数据库对象名不需在双引号里
	return "" + name + ""
}

func (db *oracle) QuoteStr() string {
	//数据库对象名不需在双引号里
	return ""
}

func (db *oracle) SupportEngine() bool {
	return false
}

func (db *oracle) SupportCharset() bool {
	return false
}

func (db *oracle) SupportDropIfExists() bool {
	return false
}

func (db *oracle) IndexOnTable() bool {
	return false
}

func (db *oracle) DropTableSql(tableName string) string {
	return fmt.Sprintf("DROP TABLE `%s`", tableName)
}

func (db *oracle) CreateTableSql(table *core.Table, tableName, storeEngine, charset string) (string) {
	var sql string
	var sequencesql string
	var sqlcomment string
	sql = "CREATE TABLE "
	sequencesql = "CREATE SEQUENCE "

	if tableName == "" {
		tableName = table.Name
	}

	sql += db.Quote(tableName) + " ("

	pkList := table.PrimaryKeys

	for _, colName := range table.ColumnsSeq() {
		col := table.GetColumn(colName)
		/*if col.IsPrimaryKey && len(pkList) == 1 {
			sql += col.String(b.dialect)
		} else {*/

		//core中使用的col.StringNoPk方法null(not null)在default之前，oracle必须default在前 --by chaos
		sql += db.QuoteStr() + col.Name + db.QuoteStr() + " "

		sql += db.SqlType(col) + " "
		if col.Default != "" {
			sql += "DEFAULT " + col.Default + " "
		}

		if db.ShowCreateNull() {
			if col.Nullable {
				sql += "NULL "
			} else {
				sql += "NOT NULL "
			}
		}

		if col.IsAutoIncrement {
			sequencesql = sequencesql + " seq_" + tableName + " start with " + strconv.FormatInt(col.StartWith, 10)
		}
		// by chaos end

		//sql += sql1
		//}
		sql = strings.TrimSpace(sql)
		sql += ", "
		sqlcomment += fmt.Sprintf(" comment on column %s.%s is '%s';", table.Name, col.Name, col.Comment)

	}

	if len(pkList) > 0 {
		sql += "PRIMARY KEY ( "
		sql += db.Quote(strings.Join(pkList, db.Quote(",")))
		sql += " ), "
	}

	sql = sql[:len(sql)-2] + ")"
	if db.SupportEngine() && storeEngine != "" {
		sql += " ENGINE=" + storeEngine
	}
	if db.SupportCharset() {
		if len(charset) == 0 {
			charset = db.URI().Charset
		}
		if len(charset) > 0 {
			sql += " DEFAULT CHARSET " + charset
		}
	}

	sqlcomment+=fmt.Sprintf(" comment on column %s is '%s';", table.Name, )
	//fmt.Println(sql)
	return sql + ";"  + sqlcomment
}

func (db *oracle) AlterIncrementSql(table *core.Table, tableName, storeEngine, charset string) (string) {
	var sql string
	sql = ""
	for _, colName := range table.ColumnsSeq() {
		col := table.GetColumn(colName)
		if col.IsPrimaryKey && col.IsAutoIncrement {
			sql = fmt.Sprintf("create sequence seq_%s start with %d", tableName, col.StartWith)
		}
	}
	return sql
}

func (db *oracle) IndexCheckSql(tableName, idxName string) (string, []interface{}) {
	args := []interface{}{tableName, idxName}
	return `SELECT INDEX_NAME FROM USER_INDEXES ` +
		`WHERE TABLE_NAME = :1 AND INDEX_NAME = :2`, args
}

func (db *oracle) TableCheckSql(tableName string) (string, []interface{}) {
	args := []interface{}{strings.ToUpper(tableName)}
	return `SELECT table_name FROM user_tables WHERE table_name = :1`, args
}

func (db *oracle) MustDropTable(tableName string) error {
	sql, args := db.TableCheckSql(tableName)
	db.LogSQL(sql, args)

	rows, err := db.DB().Query(sql, args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil
	}

	sql = "Drop Table \"" + tableName + "\""
	db.LogSQL(sql, args)

	_, err = db.DB().Exec(sql)
	return err
}

/*func (db *oracle) ColumnCheckSql(tableName, colName string) (string, []interface{}) {
	args := []interface{}{strings.ToUpper(tableName), strings.ToUpper(colName)}
	return "SELECT column_name FROM USER_TAB_COLUMNS WHERE table_name = ?" +
		" AND column_name = ?", args
}*/

func (db *oracle) IsColumnExist(table *core.Table, column *core.Column) (bool, error, *core.Column) {
	query := fmt.Sprintf("SELECT column_name CNT FROM USER_TAB_COLUMNS WHERE table_name = upper('%s') AND column_name = upper('%s')", table.Name, column.Name)
	//result:=db.DB().QueryRow(query)

	args := []interface{}{}
	//s := "SELECT table_name FROM user_tables"
	db.LogSQL(query, args)

	rows, err := db.DB().Query(query, args...)
	if err != nil {
		return false, err, nil
	}
	defer rows.Close()

	i := 0
	for rows.Next() {
		i++
	}
	//col:=db.GetPhysicalColumn(table,column)
	return i > 0, nil, nil
}

func (db *oracle) GetColumns(tableName string) ([]string, map[string]*core.Column, error) {
	args := []interface{}{tableName}
	s := "SELECT lower(a.column_name) column_name,null data_default,data_type,data_length,data_precision,data_scale, " +
		"nullable,b.comments comments FROM USER_TAB_COLUMNS a ,user_col_comments b where a.table_name=b.table_name and a.column_name=b.column_name and a.table_name = upper( :1) order by a.column_id"
	db.LogSQL(s, args)

	//x,err:=db.DB().Exec(s,args)
	//
	//fmt.Println(x.RowsAffected())

	rows, err := db.DB().Query(s, args...)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	cols := make(map[string]*core.Column)
	colSeq := make([]string, 0)
	for rows.Next() {
		col := new(core.Column)
		col.Indexes = make(map[string]int)

		var colName, colDefault, nullable, dataType, dataPrecision, comments, dataScale *string
		var dataLen int

		err = rows.Scan(&colName, &colDefault, &dataType, &dataLen, &dataPrecision,
			&dataScale, &nullable, &comments)
		if err != nil {
			return nil, nil, err
		}
		col.Name = strings.Trim(*colName, `" `)
		if comments != nil {
			col.Comment = *comments
		}
		col.DefaultIsEmpty = true
		if colDefault != nil {
			col.Default = *colDefault
			if *colDefault != "''" {
				col.DefaultIsEmpty = false
			}
		}

		if *nullable == "Y" {
			col.Nullable = true
		} else {
			col.Nullable = false
		}

		var ignore bool

		var dt string
		var len1, len2 int
		dts := strings.Split(*dataType, "(")
		dt = dts[0]
		if len(dts) > 1 {
			lens := strings.Split(dts[1][:len(dts[1])-1], ",")
			if len(lens) > 1 {
				len1, _ = strconv.Atoi(lens[0])
				len2, _ = strconv.Atoi(lens[1])
			} else {
				len1, _ = strconv.Atoi(lens[0])
			}
		}

		switch dt {
		case "VARCHAR2":
			col.SQLType = core.SQLType{Name: core.Varchar, DefaultLength: len1, DefaultLength2: len2}
			col.Length = dataLen
		case "NVARCHAR2":
			col.SQLType = core.SQLType{Name: core.Varchar, DefaultLength: len1, DefaultLength2: len2}
			col.Length = dataLen
		case "TIMESTAMP WITH TIME ZONE":
			col.SQLType = core.SQLType{Name: core.DateTime, DefaultLength: 0, DefaultLength2: 0}
		case "NUMBER":
			col.SQLType = core.SQLType{Name: core.BigInt, DefaultLength: len1, DefaultLength2: len2}
		case "CLOB":
			col.SQLType = core.SQLType{Name: core.Text, DefaultLength: len1, DefaultLength2: len2}
		case "DATE":
			col.SQLType = core.SQLType{Name: core.DateTime, DefaultLength: 0, DefaultLength2: 0}
		case "LONG", "LONG RAW":
			col.SQLType = core.SQLType{Name: core.Text, DefaultLength: 0, DefaultLength2: 0}
		case "RAW":
			col.SQLType = core.SQLType{Name: core.Binary, DefaultLength: 0, DefaultLength2: 0}
		case "ROWID":
			col.SQLType = core.SQLType{Name: core.Varchar, DefaultLength: 18, DefaultLength2: 0}
			col.Length = 18
		case "AQ$_SUBSCRIBERS":
			ignore = true
		default:
			col.SQLType = core.SQLType{Name: strings.ToUpper(dt), DefaultLength: len1, DefaultLength2: len2}
		}

		if ignore {
			continue
		}

		if _, ok := core.SqlTypes[col.SQLType.Name]; !ok {
			return nil, nil, fmt.Errorf("Unknown colType %v %v", *dataType, col.SQLType)
		}

		if col.SQLType.IsText() || col.SQLType.IsTime() {
			if !col.DefaultIsEmpty {
				col.Default = "'" + col.Default + "'"
			}
		}
		cols[col.Name] = col
		colSeq = append(colSeq, col.Name)
	}

	return colSeq, cols, nil
}

func (db *oracle) GetTables() ([]*core.Table, error) {
	args := []interface{}{}
	s := "SELECT lower(table_name) table_name FROM user_tables"
	db.LogSQL(s, args)

	rows, err := db.DB().Query(s, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tables := make([]*core.Table, 0)
	for rows.Next() {
		table := core.NewEmptyTable()
		err = rows.Scan(&table.Name)
		if err != nil {
			return nil, err
		}

		tables = append(tables, table)
	}
	return tables, nil
}

func (db *oracle) GetIndexes(tableName string) (map[string]*core.Index, error) {
	args := []interface{}{tableName}
	s := "SELECT t.column_name,i.uniqueness,i.index_name FROM user_ind_columns t,user_indexes i " +
		"WHERE t.index_name = upper(i.index_name) and t.table_name = upper(i.table_name) and t.table_name =upper(:1)"
	db.LogSQL(s, args)

	rows, err := db.DB().Query(s, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	indexes := make(map[string]*core.Index, 0)
	for rows.Next() {
		var indexType int
		var indexName, colName, uniqueness string

		err = rows.Scan(&colName, &uniqueness, &indexName)
		if err != nil {
			return nil, err
		}

		indexName = strings.Trim(indexName, `" `)

		var isRegular bool
		if strings.HasPrefix(indexName, "IDX_"+tableName) || strings.HasPrefix(indexName, "UQE_"+tableName) {
			indexName = indexName[5+len(tableName):]
			isRegular = true
		}

		if uniqueness == "UNIQUE" {
			indexType = core.UniqueType
		} else {
			indexType = core.IndexType
		}

		var index *core.Index
		var ok bool
		if index, ok = indexes[indexName]; !ok {
			index = new(core.Index)
			index.Type = indexType
			index.Name = indexName
			index.IsRegular = isRegular
			indexes[indexName] = index
		}
		index.AddColumn(colName)
	}
	return indexes, nil
}

func (db *oracle) Filters() []core.Filter {
	return []core.Filter{&core.QuoteFilter{}, &core.SeqFilter{Prefix: ":", Start: 1}, &core.IdFilter{}}
}

type goracleDriver struct {
}

func (cfg *goracleDriver) Parse(driverName, dataSourceName string) (*core.Uri, error) {
	db := &core.Uri{DbType: core.ORACLE}
	dsnPattern := regexp.MustCompile(
		`^(?:(?P<user>.*?)(?::(?P<passwd>.*))?@)?` + // [user[:password]@]
			`(?:(?P<net>[^\(]*)(?:\((?P<addr>[^\)]*)\))?)?` + // [net[(addr)]]
			`\/(?P<dbname>.*?)` + // /dbname
			`(?:\?(?P<params>[^\?]*))?$`) // [?param1=value1&paramN=valueN]
	matches := dsnPattern.FindStringSubmatch(dataSourceName)
	//tlsConfigRegister := make(map[string]*tls.Config)
	names := dsnPattern.SubexpNames()

	for i, match := range matches {
		switch names[i] {
		case "dbname":
			db.DbName = match
		}
	}
	if db.DbName == "" {
		return nil, errors.New("dbname is empty")
	}
	return db, nil
}

type oci8Driver struct {
}

//dataSourceName=user/password@ipv4:port/dbname
//dataSourceName=user/password@[ipv6]:port/dbname
func (p *oci8Driver) Parse(driverName, dataSourceName string) (*core.Uri, error) {
	db := &core.Uri{DbType: core.ORACLE}
	dsnPattern := regexp.MustCompile(
		`^(?P<user>.*)\/(?P<password>.*)@` + // user:password@
			`(?P<net>.*)` + // ip:port
			`\/(?P<dbname>.*)`) // dbname
	matches := dsnPattern.FindStringSubmatch(dataSourceName)
	names := dsnPattern.SubexpNames()
	for i, match := range matches {
		switch names[i] {
		case "dbname":
			db.DbName = match
		}
	}
	if db.DbName == "" {
		return nil, errors.New("dbname is empty")
	}

	return db, nil
}

//func (col * github.com.go-xorm.core.Column) StringNoPk(d Dialect) string {

//	sql := d.QuoteStr() + col.Name + d.QuoteStr() + " "
//
//	sql += d.SqlType(col) + " "
//
//	if d.DBType() != ORACLE {
//		if d.ShowCreateNull() {
//			if col.Nullable {
//				sql += "NULL "
//			} else {
//				sql += "NOT NULL "
//			}
//		}
//
//		if col.Default != "" {
//			sql += "DEFAULT " + col.Default + " "
//		}
//
//	}else {
//		if col.Default != "" {
//			sql += "DEFAULT " + col.Default + " "
//		}
//
//		if d.ShowCreateNull() {
//			if col.Nullable {
//				sql += "NULL "
//			} else {
//				sql += "NOT NULL "
//			}
//		}
//
//	}
//
//
//	return sql
//}

func (db *oracle) GetPhysicalColumn(table *core.Table, column *core.Column) *core.Column {
	return &core.Column{}
}

func (db *oracle) GetAllTableColumns() (map[string]map[string]*core.Column, error) {
	result := make(map[string]map[string]*core.Column)
	return result, nil
}
