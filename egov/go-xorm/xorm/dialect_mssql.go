// Copyright 2015 The Xorm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xorm

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"egov/go-xorm/core"
	"reflect"
)

var (
	mssqlReservedWords = map[string]bool{
		"ADD":                            true,
		"EXTERNAL":                       true,
		"PROCEDURE":                      true,
		"ALL":                            true,
		"FETCH":                          true,
		"PUBLIC":                         true,
		"ALTER":                          true,
		"FILE":                           true,
		"RAISERROR":                      true,
		"AND":                            true,
		"FILLFACTOR":                     true,
		"READ":                           true,
		"ANY":                            true,
		"FOR":                            true,
		"READTEXT":                       true,
		"AS":                             true,
		"FOREIGN":                        true,
		"RECONFIGURE":                    true,
		"ASC":                            true,
		"FREETEXT":                       true,
		"REFERENCES":                     true,
		"AUTHORIZATION":                  true,
		"FREETEXTTABLE":                  true,
		"REPLICATION":                    true,
		"BACKUP":                         true,
		"FROM":                           true,
		"RESTORE":                        true,
		"BEGIN":                          true,
		"FULL":                           true,
		"RESTRICT":                       true,
		"BETWEEN":                        true,
		"FUNCTION":                       true,
		"RETURN":                         true,
		"BREAK":                          true,
		"GOTO":                           true,
		"REVERT":                         true,
		"BROWSE":                         true,
		"GRANT":                          true,
		"REVOKE":                         true,
		"BULK":                           true,
		"GROUP":                          true,
		"RIGHT":                          true,
		"BY":                             true,
		"HAVING":                         true,
		"ROLLBACK":                       true,
		"CASCADE":                        true,
		"HOLDLOCK":                       true,
		"ROWCOUNT":                       true,
		"CASE":                           true,
		"IDENTITY":                       true,
		"ROWGUIDCOL":                     true,
		"CHECK":                          true,
		"IDENTITY_INSERT":                true,
		"RULE":                           true,
		"CHECKPOINT":                     true,
		"IDENTITYCOL":                    true,
		"SAVE":                           true,
		"CLOSE":                          true,
		"IF":                             true,
		"SCHEMA":                         true,
		"CLUSTERED":                      true,
		"IN":                             true,
		"SECURITYAUDIT":                  true,
		"COALESCE":                       true,
		"INDEX":                          true,
		"SELECT":                         true,
		"COLLATE":                        true,
		"INNER":                          true,
		"SEMANTICKEYPHRASETABLE":         true,
		"COLUMN":                         true,
		"INSERT":                         true,
		"SEMANTICSIMILARITYDETAILSTABLE": true,
		"COMMIT":                         true,
		"INTERSECT":                      true,
		"SEMANTICSIMILARITYTABLE":        true,
		"COMPUTE":                        true,
		"INTO":                           true,
		"SESSION_USER":                   true,
		"CONSTRAINT":                     true,
		"IS":                             true,
		"SET":                            true,
		"CONTAINS":                       true,
		"JOIN":                           true,
		"SETUSER":                        true,
		"CONTAINSTABLE":                  true,
		"KEY":                            true,
		"SHUTDOWN":                       true,
		"CONTINUE":                       true,
		"KILL":                           true,
		"SOME":                           true,
		"CONVERT":                        true,
		"LEFT":                           true,
		"STATISTICS":                     true,
		"CREATE":                         true,
		"LIKE":                           true,
		"SYSTEM_USER":                    true,
		"CROSS":                          true,
		"LINENO":                         true,
		"TABLE":                          true,
		"CURRENT":                        true,
		"LOAD":                           true,
		"TABLESAMPLE":                    true,
		"CURRENT_DATE":                   true,
		"MERGE":                          true,
		"TEXTSIZE":                       true,
		"CURRENT_TIME":                   true,
		"NATIONAL":                       true,
		"THEN":                           true,
		"CURRENT_TIMESTAMP":              true,
		"NOCHECK":                        true,
		"TO":                             true,
		"CURRENT_USER":                   true,
		"NONCLUSTERED":                   true,
		"TOP":                            true,
		"CURSOR":                         true,
		"NOT":                            true,
		"TRAN":                           true,
		"DATABASE":                       true,
		"NULL":                           true,
		"TRANSACTION":                    true,
		"DBCC":                           true,
		"NULLIF":                         true,
		"TRIGGER":                        true,
		"DEALLOCATE":                     true,
		"OF":                             true,
		"TRUNCATE":                       true,
		"DECLARE":                        true,
		"OFF":                            true,
		"TRY_CONVERT":                    true,
		"DEFAULT":                        true,
		"OFFSETS":                        true,
		"TSEQUAL":                        true,
		"DELETE":                         true,
		"ON":                             true,
		"UNION":                          true,
		"DENY":                           true,
		"OPEN":                           true,
		"UNIQUE":                         true,
		"DESC":                           true,
		"OPENDATASOURCE":                 true,
		"UNPIVOT":                        true,
		"DISK":                           true,
		"OPENQUERY":                      true,
		"UPDATE":                         true,
		"DISTINCT":                       true,
		"OPENROWSET":                     true,
		"UPDATETEXT":                     true,
		"DISTRIBUTED":                    true,
		"OPENXML":                        true,
		"USE":                            true,
		"DOUBLE":                         true,
		"OPTION":                         true,
		"USER":                           true,
		"DROP":                           true,
		"OR":                             true,
		"VALUES":                         true,
		"DUMP":                           true,
		"ORDER":                          true,
		"VARYING":                        true,
		"ELSE":                           true,
		"OUTER":                          true,
		"VIEW":                           true,
		"END":                            true,
		"OVER":                           true,
		"WAITFOR":                        true,
		"ERRLVL":                         true,
		"PERCENT":                        true,
		"WHEN":                           true,
		"ESCAPE":                         true,
		"PIVOT":                          true,
		"WHERE":                          true,
		"EXCEPT":                         true,
		"PLAN":                           true,
		"WHILE":                          true,
		"EXEC":                           true,
		"PRECISION":                      true,
		"WITH":                           true,
		"EXECUTE":                        true,
		"PRIMARY":                        true,
		"WITHIN":                         true,
		"EXISTS":                         true,
		"PRINT":                          true,
		"WRITETEXT":                      true,
		"EXIT":                           true,
		"PROC":                           true,
	}
)

type mssql struct {
	core.Base
}

func (db *mssql) Init(d *core.DB, uri *core.Uri, drivername, dataSourceName string) error {
	return db.Base.Init(d, db, uri, drivername, dataSourceName)
}

func (db *mssql) SqlType(c *core.Column) string {
	var res string
	switch t := c.SQLType.Name; t {
	case core.Bool:
		res = core.TinyInt
		//if strings.EqualFold(c.Default, "1") { //bool型统一数字设置初始值
		//	c.Default = "1"
		//} else {
		//	c.Default = "0"
		//}
	case core.Serial:
		c.IsAutoIncrement = true
		c.IsPrimaryKey = true
		c.Nullable = false
		res = core.Int
	case core.BigSerial:
		c.IsAutoIncrement = true
		c.IsPrimaryKey = true
		c.Nullable = false
		res = core.BigInt
	case core.Bytea, core.Blob, core.Binary, core.TinyBlob, core.MediumBlob, core.LongBlob, core.Image:
		res = core.Image
		//if c.Length == 0 {
		//	c.Length = 50
		//}
	case core.TimeStamp:
		res = core.DateTime
	case core.TimeStampz:
		res = "DATETIMEOFFSET"
		c.Length = 7
	case core.MediumInt:
		res = core.Int
	case core.Text, core.MediumText, core.TinyText, core.LongText, core.Json:
		res = core.Text
	case core.Double:
		res = core.Real
	case core.Uuid:
		res = core.Varchar
		c.Length = 40
	case core.TinyInt:
		res = core.TinyInt
		c.Length = 4
	case core.Float:
		res = core.Decimal
		c.Length = 38
		c.Length2 = 6
	default:
		res = t
	}

	if res == core.Int {
		return core.Int
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

func (db *mssql) SetTableComment(d map[string]string, t map[string]string) {
	db.DataTable = t
	db.Dictionaries = d
}

func (db *mssql) SupportInsertMany() bool {
	return true
}

func (db *mssql) IsReserved(name string) bool {
	_, ok := mssqlReservedWords[name]
	return ok
}

func (db *mssql) Quote(name string) string {
	//return "\"" + name + "\""
	return name
}

func (db *mssql) QuoteStr() string {
	//return "\""
	return ""
}

func (db *mssql) SupportEngine() bool {
	return false
}

func (db *mssql) AutoIncrStr() string {
	return "IDENTITY"
}

func (db *mssql) DropTableSql(tableName string) string {
	return fmt.Sprintf("IF EXISTS (SELECT * FROM sysobjects WHERE id = "+
		"object_id(N'%s') and OBJECTPROPERTY(id, N'IsUserTable') = 1) "+
		"DROP TABLE \"%s\"", tableName, tableName)
}

func (db *mssql) SupportCharset() bool {
	return false
}

func (db *mssql) IndexOnTable() bool {
	return true
}

func (db *mssql) IndexCheckSql(tableName, idxName string) (string, []interface{}) {
	args := []interface{}{idxName}
	sql := "select name from sysindexes where id=object_id('" + tableName + "') and name=?"
	return sql, args
}

/*func (db *mssql) ColumnCheckSql(tableName, colName string) (string, []interface{}) {
	args := []interface{}{tableName, colName}
	sql := `SELECT "COLUMN_NAME" FROM "INFORMATION_SCHEMA"."COLUMNS" WHERE "TABLE_NAME" = ? AND "COLUMN_NAME" = ?`
	return sql, args
}*/

func (db *mssql) GetPhysicalColumn(table *core.Table, column *core.Column) *core.Column {
	sql := `select   c.object_id,
			tName = c.Name,
			tComment = isnull(f.value, ''),
			name = a.Name,
			sortcode = a.Column_id,
			autoincr = case when is_identity = 1 then 1 else 0 end,
		pk =
		case
		when exists
			(SELECT 1
			FROM   INFORMATION_SCHEMA.KEY_COLUMN_USAGE u
			WHERE  CONSTRAINT_NAME IN (SELECT CONSTRAINT_NAME FROM INFORMATION_SCHEMA.TABLE_CONSTRAINTS WHERE INFORMATION_SCHEMA.TABLE_CONSTRAINTS.CONSTRAINT_TYPE='PRIMARY KEY' ) and c.name = u.TABLE_NAME and a.name = u.COLUMN_NAME)
			then
			1
			else
			0
			end,
			fk = T.f_tab + '(' + T.f_clon + ')',
			type = b.Name,
			byte =
		case
		when a.max_length = -1 and b.Name != 'xml' then 'max/2G'
             when b.Name = 'xml' then ' 2^31-1字节/2G'
             else rtrim(a.max_length)
		end,
		length = ColumnProperty(a.object_id, a.Name, 'Precision'),
		precision = isnull(ColumnProperty(a.object_id, a.Name, 'Scale'), 0),
		nullable = case when a.is_nullable = 1 then 1 else 0 end,
		comments = isnull(e.value, ''),
		defaultvalue = isnull(d.text, ''),
		indexes = g.val,
		indexnum = isnull(g.indexNum, 0)
		from     sys.columns a
		left join sys.types b on a.user_type_id = b.user_type_id
		inner join sys.objects c on a.object_id = c.object_id and c.Type = 'U' and c.name = '%[1]s'
         left join syscomments d on a.default_object_id = d.ID
         left join sys.extended_properties e on e.major_id = c.object_id and e.minor_id = a.Column_id and e.class = 1 and e.name='MS_Description'
         left join sys.extended_properties f on f.major_id = c.object_id and f.minor_id = 0 and f.class = 1 and f.name='MS_Description'
         left join
         (SELECT m_tab, f_tab, m_clon, f_clon
          FROM   (SELECT O3.NAME F_NAME, O1.NAME M_TAB, O2.NAME F_TAB, L1.NAME M_CLON, L2.NAME F_CLON
                  FROM   SYSFOREIGNKEYS A, SYSOBJECTS O1, SYSOBJECTS O2, SYSOBJECTS O3, SYSCOLUMNS L1, SYSCOLUMNS L2
                  WHERE  A.CONSTID = O3.ID AND A.FKEYID = O1.ID AND A.RKEYID = O2.ID AND L1.ID = O1.ID AND L2.ID = O2.ID AND A.FKEY
                         = L1.COLID AND A.RKEY = L2.COLID AND O1.XTYPE = 'U' AND O2.XTYPE = 'U' and o1.name='%[1]s' ) M) T
           on T.m_tab = c.name and T.m_clon = a.name
         left join
         (select   c_name,
                   val =
                     (select i_name + ' '
                      from   ((select case  a1.is_unique when 1  then 'unique('+a1.name+')' else 'index('+a1.name+')' end i_name, c1.name c_name
                               from   sys.indexes a1
                                      inner join sys.index_columns b1 on a1.index_id = b1.index_id
                                      inner join sys.columns c1 on b1.column_id = c1.column_id
                                      inner join sys.tables d1
                                        on c1.object_id = d1.object_id and c1.object_id = object_id('%[1]s')
                               where  d1.object_id = a1.object_id and b1.object_id = d1.object_id and a1.is_primary_key <>
                                      1)) bb
                      where  bb.c_name = aa.c_name
                      for xml path ( '' )),
                   count(*) indexNum
          from     (select a2.name i_name, c2.name c_name
                    from   sys.indexes a2
                           inner join sys.index_columns b2 on a2.index_id = b2.index_id
                           inner join sys.columns c2 on b2.column_id = c2.column_id
                           inner join sys.tables d2 on c2.object_id = d2.object_id and a2.object_id = object_id('%[1]s')
                    where  d2.object_id = a2.object_id and b2.object_id = d2.object_id and a2.is_primary_key <> 1) aa
          group by c_name) g
           on g.c_name = a.name
	where a.name='%[2]s'
`

	//if column.Name=="description"{
	//	fmt.Println("description")
	//}
	sqlQuery := fmt.Sprintf(sql, table.Name, column.Name)
	rows, _ := db.DB().Query(sqlQuery)

	colRow := make(map[string]string)

	//	columns, _ = rows.ToMapString()
	cols, err := rows.Columns()
	if err != nil {
		fmt.Println(err.Error())
	}

	if rows.Next() {
		slice := make([]interface{}, len(cols))
		err = rows.ScanSlice(&slice)
		if err != nil {
			fmt.Println(err.Error())
		}
		//fmt.Println(cols)
		//fmt.Println(slice)
		for idx, colName := range cols {
			if slice[idx] == nil {
				colRow[colName] = ""
				continue
			}
			colKind := reflect.TypeOf(slice[idx]).Kind()
			switch colKind {
			case reflect.String:
				colRow[colName] = slice[idx].(string)
			case reflect.Int64:
				colRow[colName] = strconv.FormatInt(slice[idx].(int64), 10)
			default:
				panic("not except type ")
			}

		}
	}
	//if len(colRow) == 0 {
	//	fmt.Println(sqlQuery)
	//}
	rows.Close()
	_, col := core.TransMapStringColumn(50, colRow)
	return col
}

func (db *mssql) IsColumnExist(table *core.Table, column *core.Column) (bool, error, *core.Column) {
	query := `SELECT "COLUMN_NAME" FROM "INFORMATION_SCHEMA"."COLUMNS" WHERE "TABLE_NAME" = ? AND "COLUMN_NAME" = ?`
	col := &core.Column{}
	ifHasRecords, err := db.HasRecords(query, table.Name, column.Name)
	if ifHasRecords {
		//col = db.GetPhysicalColumn(table, column)
	}
	//fmt.Println(col.XormTag,col.Comment)
	return ifHasRecords, err, col
}

func (db *mssql) TableCheckSql(tableName string) (string, []interface{}) {
	args := []interface{}{}
	sql := "select * from sysobjects where id = object_id(N'" + tableName + "') and OBJECTPROPERTY(id, N'IsUserTable') = 1"
	return sql, args
}

func (db *mssql) GetColumns(tableName string) ([]string, map[string]*core.Column, error) {
	args := []interface{}{}
	s := `SELECT a.COLUMN_NAME name,
       data_type ctype,
       case when character_maximum_length is null then 0 else character_maximum_length end max_length,
       case when a.NUMERIC_PRECISION is null then 0 else a.NUMERIC_PRECISION end precision,
       case when a.NUMERIC_SCALE is null then 0 else a.NUMERIC_SCALE  end scale,
       CASE WHEN a.IS_NULLABLE = 'YES' THEN 1 ELSE 0 END nullable,
       case when a.COLUMN_DEFAULT is null then '' else a.COLUMN_DEFAULT  end vdefault, 
       case when b.CONSTRAINT_TYPE='PRIMARY KEY' THEN 1 ELSE 0 END ISPK
  FROM INFORMATION_SCHEMA.COLUMNS a
  left join INFORMATION_SCHEMA.CONSTRAINT_COLUMN_USAGE c on c.TABLE_NAME = a.TABLE_NAME and c.COLUMN_NAME=a.COLUMN_NAME 
  left join INFORMATION_SCHEMA.TABLE_CONSTRAINTS b on b.CONSTRAINT_NAME=c.CONSTRAINT_NAME`
	s = s + " where a.table_name='" + tableName + "'"
	db.LogSQL(s, args)

	rows, err := db.DB().Query(s, args...)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	cols := make(map[string]*core.Column)
	colSeq := make([]string, 0)
	for rows.Next() {
		var name, ctype, vdefault string
		var maxLen, precision, scale, ispk int
		var nullable int
		err = rows.Scan(&name, &ctype, &maxLen, &precision, &scale, &nullable, &vdefault, &ispk)
		if err != nil {
			return nil, nil, err
		}

		col := new(core.Column)
		col.Indexes = make(map[string]int)
		col.Name = strings.Trim(name, "` ")
		col.Nullable = nullable==1
		col.Default = vdefault
		col.IsPrimaryKey = ispk == 1
		ct := strings.ToUpper(ctype)
		if ct == "DECIMAL" {
			col.Length = precision
			col.Length2 = scale
		} else {
			col.Length = maxLen
		}
		switch ct {
		case "DATETIMEOFFSET":
			col.SQLType = core.SQLType{Name: core.TimeStampz, DefaultLength: 0, DefaultLength2: 0}
		case "NVARCHAR":
			col.SQLType = core.SQLType{Name: core.NVarchar, DefaultLength: 0, DefaultLength2: 0}
		case "IMAGE":
			col.SQLType = core.SQLType{Name: core.VarBinary, DefaultLength: 0, DefaultLength2: 0}
		default:
			if _, ok := core.SqlTypes[ct]; ok {
				col.SQLType = core.SQLType{Name: ct, DefaultLength: 0, DefaultLength2: 0}
			} else {
				return nil, nil, fmt.Errorf("Unknown colType %v for %v - %v", ct, tableName, col.Name)
			}
		}

		if col.SQLType.IsText() || col.SQLType.IsTime() {
			if col.Default != "" {
				col.Default = "'" + col.Default + "'"
			} else {
				if col.DefaultIsEmpty {
					col.Default = "''"
				}
			}
		}
		cols[col.Name] = col

		colSeq = append(colSeq, col.Name)
	}
	return colSeq, cols, nil
}

func (db *mssql) GetTables() ([]*core.Table, error) {
	args := []interface{}{}
	s := `select name from sysobjects where xtype ='U'`
	db.LogSQL(s, args)

	rows, err := db.DB().Query(s, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tables := make([]*core.Table, 0)
	for rows.Next() {
		table := core.NewEmptyTable()
		var name string
		err = rows.Scan(&name)
		if err != nil {
			return nil, err
		}
		table.Name = strings.Trim(name, "` ")
		tables = append(tables, table)
	}
	return tables, nil
}

func (db *mssql) GetIndexes(tableName string) (map[string]*core.Index, error) {
	args := []interface{}{tableName}
	s := `SELECT
IXS.NAME                    AS  [INDEX_NAME],
C.NAME                      AS  [COLUMN_NAME],
IXS.is_unique AS [IS_UNIQUE]
FROM SYS.INDEXES IXS
INNER JOIN SYS.INDEX_COLUMNS   IXCS
ON IXS.OBJECT_ID=IXCS.OBJECT_ID  AND IXS.INDEX_ID = IXCS.INDEX_ID
INNER   JOIN SYS.COLUMNS C  ON IXS.OBJECT_ID=C.OBJECT_ID
AND IXCS.COLUMN_ID=C.COLUMN_ID
WHERE IXS.TYPE_DESC='NONCLUSTERED' and OBJECT_NAME(IXS.OBJECT_ID) =?
`
	db.LogSQL(s, args)

	rows, err := db.DB().Query(s, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	indexes := make(map[string]*core.Index, 0)
	for rows.Next() {
		var indexType int
		var indexName, colName, isUnique string

		err = rows.Scan(&indexName, &colName, &isUnique)
		if err != nil {
			return nil, err
		}

		i, err := strconv.ParseBool(isUnique)
		if err != nil {
			return nil, err
		}

		if i {
			indexType = core.UniqueType
		} else {
			indexType = core.IndexType
		}

		colName = strings.Trim(colName, "` ")
		var isRegular bool
		if strings.HasPrefix(indexName, "IDX_"+tableName) || strings.HasPrefix(indexName, "UQE_"+tableName) {
			indexName = indexName[5+len(tableName):]
			isRegular = true
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

func (db *mssql) CreateTableSql(table *core.Table, tableName, storeEngine, charset string) (string) {
	var sql string
	var sqlcomment string
	if tableName == "" {
		tableName = table.Name
	}
	sql = "IF NOT EXISTS (SELECT [name] FROM sys.tables WHERE [name] = '" + tableName + "' ) CREATE TABLE "

	sql += db.QuoteStr() + tableName + db.QuoteStr() + " ("

	pkList := table.PrimaryKeys

	for _, colName := range table.ColumnsSeq() {
		col := table.GetColumn(colName)
		if col.IsPrimaryKey && len(pkList) == 1 {
			sql += col.String(db)
		} else {
			sql += col.StringNoPk(db)
		}
		sql = strings.TrimSpace(sql)
		sql += ", "
		//fmt.Println(col.Comment)
		sqlcomment += fmt.Sprintf("EXEC sys.sp_addextendedproperty @name = N'MS_Description', @value = N'%s', @level0type = N'SCHEMA', @level0name = N'dbo', @level1type = N'TABLE', @level1name = N'%s', @level2type = N'COLUMN', @level2name = N'%s'", col.Comment, tableName, colName)
		//sqlcomment+=fmt.Sprintf("execute sp_addextendedproperty 'MS_Description',  '%s','user', user_name(), 'table', '%s', 'column', '%s';",col.Comment,tableName,colName)
	}

	if len(pkList) > 1 {
		sql += "PRIMARY KEY ( "
		sql += strings.Join(pkList, ",")
		sql += " ), "
	}
	sqlcomment += fmt.Sprintf("exec sys.sp_addextendedproperty 'MS_Description', '%s', 'user', 'dbo', 'table', '%s'", table.Comment, tableName)
	sql = sql[:len(sql)-2] + ")"
	sql += ";" + sqlcomment
	return sql
}

func (db *mssql) ModifyColumnSql(tableName string, col *core.Column) string {
	return fmt.Sprintf("alter table %s alter COLUMN %s", tableName, col.StringNoPk(db))
}

func (db *mssql) AlterIncrementSql(table *core.Table, tableName, storeEngine, charset string) (string) {
	var sql string
	sql = ""
	for _, colName := range table.ColumnsSeq() {
		col := table.GetColumn(colName)
		if table.AutoIncrement == colName && col.IsAutoIncrement {
			sql = fmt.Sprintf("DBCC CHECKIDENT (%s, RESEED, %d)", tableName, col.StartWith)
			break
		}
	}
	return sql
}

func (db *mssql) ForUpdateSql(query string) string {
	return query
}

func (db *mssql) Filters() []core.Filter {
	return []core.Filter{&core.IdFilter{}, &core.QuoteFilter{}}
}

type odbcDriver struct {
}

func (p *odbcDriver) Parse(driverName, dataSourceName string) (*core.Uri, error) {
	kv := strings.Split(dataSourceName, ";")
	var dbName string

	for _, c := range kv {
		vv := strings.Split(strings.TrimSpace(c), "=")
		if len(vv) == 2 {
			switch strings.ToLower(vv[0]) {
			case "database":
				dbName = vv[1]
			}
		}
	}
	if dbName == "" {
		return nil, errors.New("no db name provided")
	}
	return &core.Uri{DbName: dbName, DbType: core.MSSQL}, nil
}

func (db *mssql) GetAllTableColumns() (map[string]map[string]*core.Column, error) {
	sql := `SELECT
	c.object_id,
	tName = c.NAME,
	tComment = isnull(f. VALUE, ''),
	NAME = a.NAME,
	sortcode = a.Column_id,
	autoincr = CASE
WHEN is_identity = 1 THEN
	1
ELSE
	0
END,
 pk = CASE
WHEN EXISTS (
	SELECT
		1
	FROM
		INFORMATION_SCHEMA.KEY_COLUMN_USAGE u
	WHERE
		CONSTRAINT_NAME IN (
			SELECT
				CONSTRAINT_NAME
			FROM
				INFORMATION_SCHEMA.TABLE_CONSTRAINTS
			WHERE
				INFORMATION_SCHEMA.TABLE_CONSTRAINTS.CONSTRAINT_TYPE = 'PRIMARY KEY'
		)
	AND c.NAME = u.TABLE_NAME
	AND a.NAME = u.COLUMN_NAME
) THEN
	1
ELSE
	0
END,
 fk = T.f_tab + '(' + T.f_clon + ')',
 type = b.NAME,
 byte = CASE
WHEN a.max_length = - 1
AND b.NAME != 'xml' THEN
	'max/2G'
WHEN b.NAME = 'xml' THEN
	' 2^31-1字节/2G'
ELSE
	rtrim(a.max_length)
END,
 length = ColumnProperty(
	a.object_id,
	a.NAME,
	'Precision'
),
 PRECISION = isnull(
	ColumnProperty(a.object_id, a.NAME, 'Scale'),
	0
),
 nullable = CASE
WHEN a.is_nullable = 1 THEN
	1
ELSE
	0
END,
 comments = isnull(e. VALUE, ''),
 defaultvalue = isnull(d. TEXT, ''),
 indexes = g.val,
 indexnum = isnull(g.indexNum, 0)
FROM
	sys.COLUMNS a
LEFT JOIN sys.types b ON a.user_type_id = b.user_type_id
INNER JOIN sys.objects c ON a.object_id = c.object_id
AND c.Type = 'U'
LEFT JOIN syscomments d ON a.default_object_id = d.ID
LEFT JOIN sys.extended_properties e ON e.major_id = c.object_id
AND e.minor_id = a.Column_id
AND e.class = 1
AND e.NAME = 'MS_Description'
LEFT JOIN sys.extended_properties f ON f.major_id = c.object_id
AND f.minor_id = 0
AND f.class = 1
AND f.NAME = 'MS_Description'
LEFT JOIN (
	SELECT
		m_tab,
		f_tab,
		m_clon,
		f_clon
	FROM
		(
			SELECT
				O3.NAME F_NAME,
				O1.NAME M_TAB,
				O2.NAME F_TAB,
				L1.NAME M_CLON,
				L2.NAME F_CLON
			FROM
				SYSFOREIGNKEYS A,
				SYSOBJECTS O1,
				SYSOBJECTS O2,
				SYSOBJECTS O3,
				SYSCOLUMNS L1,
				SYSCOLUMNS L2
			WHERE
				A.CONSTID = O3.ID
			AND A.FKEYID = O1.ID
			AND A.RKEYID = O2.ID
			AND L1.ID = O1.ID
			AND L2.ID = O2.ID
			AND A.FKEY = L1.COLID
			AND A.RKEY = L2.COLID
			AND O1.XTYPE = 'U'
			AND O2.XTYPE = 'U'
		) M
) T ON T.m_tab = c.NAME
AND T.m_clon = a.NAME
LEFT JOIN (
	SELECT
		c_name,
    object_id,
		val = (
			SELECT
				i_name + ' '
			FROM
				(
					(
						SELECT
							CASE a1.is_unique
						WHEN 1 THEN
							'unique(' + a1.NAME + ')'
						ELSE
							'index(' + a1.NAME + ')'
						END i_name,
						c1.NAME c_name,
            c1.object_id object_id
					FROM
						sys.indexes a1
					INNER JOIN sys.index_columns b1 ON a1.index_id = b1.index_id
					INNER JOIN sys.COLUMNS c1 ON b1.column_id = c1.column_id and c1.object_id=a1.object_id
					INNER JOIN sys.TABLES d1 ON c1.object_id = d1.object_id
					WHERE
						d1.object_id = a1.object_id
					AND b1.object_id = d1.object_id
					AND a1.is_primary_key <> 1
					)
				) bb
			WHERE
				bb.c_name = aa.c_name and bb.object_id=aa.object_id FOR XML path ('')
		),
		COUNT (*) indexNum
	FROM
		(
			SELECT
				a2.NAME i_name,
				c2.NAME c_name,
        c2.object_id object_id
			FROM
				sys.indexes a2
			INNER JOIN sys.index_columns b2 ON a2.index_id = b2.index_id
			INNER JOIN sys.COLUMNS c2 ON b2.column_id = c2.column_id and c2.object_id=a2.object_id
			INNER JOIN sys.TABLES d2 ON c2.object_id = d2.object_id
			WHERE
				d2.object_id = a2.object_id
			AND b2.object_id = d2.object_id
			AND a2.is_primary_key <> 1
		) aa
	GROUP BY
		c_name,object_id
) g ON g.c_name = a.NAME and g.object_id=a.object_id
`
	rows, err := db.DB().Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	return core.GetStringColumnFormRows(rows), nil
}
