package core

import (
	"egov/common"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type DbType string

type Uri struct {
	DbType     DbType
	Proto      string
	Host       string
	Port       string
	DbName     string
	User       string
	Passwd     string
	Charset    string
	Laddr      string
	Raddr      string
	Timeout    time.Duration
	Schema     string
	EngineName string
}

// a dialect is a driver's wrapper
type Dialect interface {
	SetLogger(logger common.ILogger)
	Init(*DB, *Uri, string, string) error
	URI() *Uri
	DB() *DB
	DBType() DbType
	SqlType(*Column) string
	FormatBytes(b []byte) string

	DriverName() string
	DataSourceName() string

	QuoteStr() string
	IsReserved(string) bool
	Quote(string) string
	AndStr() string
	OrStr() string
	EqStr() string
	RollBackStr() string
	AutoIncrStr() string

	SupportInsertMany() bool
	SupportEngine() bool
	SupportCharset() bool
	SupportDropIfExists() bool
	IndexOnTable() bool
	ShowCreateNull() bool
	GetPhysicalColumn(table *Table, column *Column) *Column
	IndexCheckSql(tableName, idxName string) (string, []interface{})
	TableCheckSql(tableName string) (string, []interface{})

	IsColumnExist(table *Table, col *Column) (bool, error, *Column)

	CreateTableSql(table *Table, tableName, storeEngine, charset string) string
	AlterIncrementSql(table *Table, tableName, storeEngine, charset string) string

	DropTableSql(tableName string) string
	CreateIndexSql(tableName string, index *Index) string
	DropIndexSql(tableName string, index *Index) string

	ModifyColumnSql(tableName string, col *Column) string

	ForUpdateSql(query string) string

	//CreateTableIfNotExists(table *Table, tableName, storeEngine, charset string) error
	//MustDropTable(tableName string) error

	GetColumns(tableName string) ([]string, map[string]*Column, error)
	GetTables() ([]*Table, error)
	GetTablesSingle(tableName string) ([]*Table, error)

	GetIndexes(tableName string) (map[string]*Index, error)

	GetAllTableColumns() (map[string]map[string]*Column, error)

	//IsColumnDifferent(tableName string,colName string,column Column)

	Filters() []Filter
	SetTableComment(d map[string]string, t map[string]string)
}

func OpenDialect(dialect Dialect) (*DB, error) {
	return Open(dialect.DriverName(), dialect.DataSourceName())
}

type Base struct {
	db             *DB
	dialect        Dialect
	driverName     string
	dataSourceName string
	logger         common.ILogger
	*Uri
	DataTable    map[string]string
	Dictionaries map[string]string
}

func (b *Base) SetTableComment(d map[string]string, t map[string]string) {
	b.DataTable = t
	b.Dictionaries = d
}

func (b *Base) DB() *DB {
	return b.db
}

func (b *Base) SetLogger(logger common.ILogger) {
	b.logger = logger
}

func (b *Base) Init(db *DB, dialect Dialect, uri *Uri, drivername, dataSourceName string) error {
	b.db, b.dialect, b.Uri = db, dialect, uri
	b.driverName, b.dataSourceName = drivername, dataSourceName
	return nil
}

func (b *Base) URI() *Uri {
	return b.Uri
}

func (b *Base) DBType() DbType {
	return b.Uri.DbType
}

func (b *Base) FormatBytes(bs []byte) string {
	return fmt.Sprintf("0x%x", bs)
}

func (b *Base) DriverName() string {
	return b.driverName
}

func (b *Base) ShowCreateNull() bool {
	return true
}

func (b *Base) DataSourceName() string {
	return b.dataSourceName
}

func (b *Base) AndStr() string {
	return "AND"
}

func (b *Base) OrStr() string {
	return "OR"
}

func (b *Base) EqStr() string {
	return "="
}

func (db *Base) RollBackStr() string {
	return "ROLL BACK"
}

func (db *Base) SupportDropIfExists() bool {
	return true
}

func (db *Base) DropTableSql(tableName string) string {
	return fmt.Sprintf("DROP TABLE IF EXISTS `%s`", tableName)
}

func (db *Base) HasRecords(query string, args ...interface{}) (bool, error) {
	db.LogSQL(query, args)
	rows, err := db.DB().Query(query, args...)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	if rows.Next() {
		return true, nil
	}
	return false, nil
}

func (db *Base) IsColumnExist(table *Table, column *Column) (bool, error, *Column) {
	query := "SELECT `COLUMN_NAME` FROM `INFORMATION_SCHEMA`.`COLUMNS` WHERE `TABLE_SCHEMA` = ? AND `TABLE_NAME` = ? AND `COLUMN_NAME` = ?"
	query = strings.Replace(query, "`", db.dialect.QuoteStr(), -1)
	//col:=db.GetPhysicalColumn(table,column)
	ifHasRecords, err := db.HasRecords(query, db.DbName, table.Name, column.Name)
	return ifHasRecords, err, nil
}

/*
func (db *Base) CreateTableIfNotExists(table *Table, tableName, storeEngine, charset string) error {
	sql, args := db.dialect.TableCheckSql(tableName)
	rows, err := db.DB().Query(sql, args...)
	if db.Logger != nil {
		db.Logger.Info("[sql]", sql, args)
	}
	if err != nil {
		return err
	}
	defer rows.Close()

	if rows.Next() {
		return nil
	}

	sql = db.dialect.CreateTableSql(table, tableName, storeEngine, charset)
	_, err = db.DB().Exec(sql)
	if db.Logger != nil {
		db.Logger.Info("[sql]", sql)
	}
	return err
}*/

func (db *Base) CreateIndexSql(tableName string, index *Index) string {
	quote := db.dialect.Quote
	var unique string
	var idxName string
	if index.Type == UniqueType {
		unique = " UNIQUE"
	}
	idxName = index.XName(tableName)
	if db.DbType == MSSQL {
		//notnull := quote(strings.Join(index.Cols, quote(" is not null and ")) + " is not null")
		return fmt.Sprintf("CREATE %s INDEX %v ON %v (%v) ", unique,
			quote(idxName), quote(tableName),
			quote(strings.Join(index.Cols, quote(","))))
	} else {
		return fmt.Sprintf("CREATE %s INDEX %v ON %v (%v)", unique,
			quote(idxName), quote(tableName),
			quote(strings.Join(index.Cols, quote(","))))
	}
}

func (db *Base) DropIndexSql(tableName string, index *Index) string {
	quote := db.dialect.Quote
	var name string
	if index.IsRegular {
		name = index.XName(tableName)
	} else {
		name = index.Name
	}
	return fmt.Sprintf("DROP INDEX %v ON %s", quote(name), quote(tableName))
}

func (db *Base) ModifyColumnSql(tableName string, col *Column) string {
	if col.IsPrimaryKey {
		sql := fmt.Sprintf("alter table %s MODIFY COLUMN %s Comment '%s'", tableName, col.ModifyString(db.dialect), col.Comment)
		//		fmt.Println(sql)
		return sql
	} else {
		return fmt.Sprintf("alter table %s MODIFY COLUMN %s Comment '%s'", tableName, col.StringNoPk(db.dialect), col.Comment)
	}
}

func (b *Base) CreateTableSql(table *Table, tableName, storeEngine, charset string) string {
	var sql string
	var startwith int64
	sql = "CREATE TABLE IF NOT EXISTS "
	if tableName == "" {
		tableName = table.Name
	}

	sql += b.dialect.Quote(tableName)
	sql += " ("

	if len(table.ColumnsSeq()) > 0 {
		pkList := table.PrimaryKeys

		for _, colName := range table.ColumnsSeq() {
			col := table.GetColumn(colName)
			if col.IsPrimaryKey && len(pkList) == 1 {
				sql += col.String(b.dialect)
			} else {
				sql += col.StringNoPk(b.dialect)
			}
			sql = strings.TrimSpace(sql)
			if len(col.Comment) > 0 {
				sql += fmt.Sprintf(" Comment '%s'", col.Comment)
			}
			sql += ", "

			if col.IsAutoIncrement {
				startwith = col.StartWith
			}
		}

		if len(pkList) > 1 {
			sql += "PRIMARY KEY ( "
			sql += b.dialect.Quote(strings.Join(pkList, b.dialect.Quote(",")))
			sql += " ), "
		}

		sql = sql[:len(sql)-2]

	}
	sql += ")"
	if strings.HasPrefix(tableName, "dic_") {
		if c, ok := b.Dictionaries[tableName]; ok {
			sql += " Comment='" + c + "'"
		}
	} else {
		if c, ok := b.DataTable[tableName]; ok {
			sql += " Comment='" + c + "'"
		}
	}

	//fmt.Println(table.AutoIncrement)
	///去除自增长字段
	if len(table.AutoIncrement) > 0 {
		sql += fmt.Sprintf("AUTO_INCREMENT=%d", startwith)
	}

	if b.dialect.SupportEngine() && storeEngine != "" {
		sql += " ENGINE=" + storeEngine
	}
	if b.dialect.SupportCharset() {
		if len(charset) == 0 {
			charset = b.dialect.URI().Charset
		}
		if len(charset) > 0 {
			sql += " DEFAULT CHARSET " + charset
		}
	}
	sql = strings.Replace(sql, "default null not null", "not null", -1)
	return sql
}

func (b *Base) AlterIncrementSql(table *Table, tableName, storeEngine, charset string) string {
	var sql string
	var startwith int64
	sql = ""
	if tableName == "" {
		tableName = table.Name
	}
	if len(table.ColumnsSeq()) > 0 {
		for _, colName := range table.ColumnsSeq() {
			col := table.GetColumn(colName)
			if col.IsAutoIncrement && col.IsPrimaryKey {
				startwith = col.StartWith
				sql = fmt.Sprintf("alter table %s AUTO_INCREMENT=%d", tableName, startwith)
			}
		}
	}
	return sql
}

func (b *Base) ForUpdateSql(query string) string {
	return query + " FOR UPDATE"
}

func (b *Base) LogSQL(sql string, args []interface{}) {
	if b.logger != nil && b.logger.IsShowSQL() {
		if len(args) > 0 {
			b.logger.Debugf("[SQL] %v %v", sql, args)
		} else {
			b.logger.Debugf("[SQL] %v", sql)
		}
	}
}

var (
	dialects = map[string]func() Dialect{}
)

// RegisterDialect register database dialect
func RegisterDialect(dbName DbType, dialectFunc func() Dialect) {
	if dialectFunc == nil {
		panic("core: Register dialect is nil")
	}
	dialects[strings.ToLower(string(dbName))] = dialectFunc // !nashtsai! allow override dialect
}

// QueryDialect query if registed database dialect
func QueryDialect(dbName DbType) Dialect {
	if d, ok := dialects[strings.ToLower(string(dbName))]; ok {
		return d()
	}
	return nil
}

func (db *Base) GetPhysicalColumn(table *Table, column *Column) *Column {
	return &Column{}
}

func (db *Base) GetAllTableColumns() (map[string]map[string]*Column, error) {
	sql := `select  	
  1 object_id,
	table_name tName  ,
	'' tComment,
	column_name NAME,
	ORDINAL_POSITION  sortcode,
	0 autoincr,
  case COLUMN_KEY when 'PRI' then 1 else 0 end pk,
  0 fk,
  data_type type,
  0 byte,
  case DATA_TYPE when 'varchar' then CHARACTER_MAXIMUM_LENGTH when 'binary' then 0 else 1 end  length,
  0 %[1]sPRECISION%[1]s,
  case IS_NULLABLE when  'YES' then 1 else 0 end nullable,
  COLUMN_COMMENT comments,
  case DATA_TYPE when 'varchar' then CONCAT('(''',COLUMN_DEFAULT,''')') else CONCAT('((',COLUMN_DEFAULT,'))') end	defaultvalue,
  '' indexes,
  0 indexnum
from information_schema.COLUMNS where TABLE_SCHEMA=?
and table_name in (select table_name from information_schema.TABLES where TABLE_SCHEMA=? and table_type='BASE TABLE' )
`
	//
	sql = fmt.Sprintf(sql, "`")
	rows, err := db.DB().Query(sql, db.DbName,db.DbName)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	return GetStringColumnFormRows(rows), nil
}




func GetStringColumnFormRows(rows *Rows) map[string]map[string]*Column {
	result := make(map[string]map[string]*Column)
	columnProperties := make(map[string]map[string]map[string]string)
	cols, err := rows.Columns()

	if err != nil {
		panic(err)
	}
	for rows.Next() {
		slice := make([]interface{}, len(cols))
		err = rows.ScanSlice(&slice)
		if err != nil {
			fmt.Println(err.Error())
		}
		//var colType reflect.Type
		//switch dbType {
		//case MYSQL:
		//	colType=reflect.TypeOf([]byte)
		//case MSSQL:
		//	colType=reflect.TypeOf("")
		//default:
		//	colType=reflect.TypeOf([]byte)
		//}
		tName := ""
		cName := ""
		if x, ok := slice[1].([]byte); !ok {
			tName = slice[1].(string)
		} else {
			tName = string(x)
		}

		//tName := string(slice[1].([]byte)) //表名
		//cName := string(slice[3].([]byte)) //列明
		if x, ok := slice[3].([]byte); !ok {
			cName = slice[3].(string)
		} else {
			cName = string(x)
		}

		//if cName=="data_interface"{
		//	fmt.Println(cName)
		//	x=true
		//}

		if columnProperties[tName] == nil {
			columnProperties[tName] = make(map[string]map[string]string)
			result[tName] = make(map[string]*Column)
		}
		if columnProperties[tName][cName] == nil {
			columnProperties[tName][cName] = make(map[string]string)
			columnProperties[tName][cName]["name"] = cName
			result[tName][cName] = &Column{}

		}
		for idx, colName := range cols {
			if idx <= 3 { //object_id tName
				continue
			}

			if slice[idx] == nil {
				columnProperties[tName][cName][colName] = ""
				continue
			}
			colKind := reflect.TypeOf(slice[idx]).Kind()
			switch colKind {
			case reflect.Int64:
				columnProperties[tName][cName][colName] = strconv.FormatInt(slice[idx].(int64), 10)
			case reflect.String:
				columnProperties[tName][cName][colName] = slice[idx].(string)
			default:
				columnProperties[tName][cName][colName] = string(slice[idx].([]byte))
			}

		}
		_, col := TransMapStringColumn(100, columnProperties[tName][cName])
		result[tName][cName] = col
	}
	//if x {
	//	fmt.Println(fmt.Sprintf("%+v",result))
	//}
	return result
}

func TransMapStringColumn(maxColLen int, column map[string]string) (string, *Column) {
	content := ""
	col := &Column{}
	col.Name = column["name"]
	//if col.Name=="data_interface"{
	//	fmt.Println(col.Name)
	//}
	col.Default = strings.Trim(column["defaultvalue"], "(")
	col.Default = strings.Trim(col.Default, ")")
	if strings.ToLower(col.Default) == "null" {
		col.Default = "null"
	}
	if col.Default == "" {
		col.Default = "null"
	}
	col.FieldName = column["name"]
	col.Comment = column["comments"]
	col.IsAutoIncrement = column["autoincr"] == "1"
	col.IsPrimaryKey = column["pk"] == "1"
	temp, _ := strconv.ParseInt(column["length"], 10, 64)
	col.Length = int(temp)
	temp, _ = strconv.ParseInt(column["precision"], 10, 64)
	col.Length2 = int(temp)
	col.StartWith = 1
	col.TableName = column["t_name"]
	col.Nullable = column["nullable"] == "1"
	col.SQLType.Name = strings.ToUpper(column["type"])
	col.SQLType.DefaultLength = col.Length
	col.SQLType.DefaultLength2 = col.Length2

	//if col.Name=="description"{
	//	fmt.Println(col.Name)
	//}
	typeString := ""
	if !col.IsPrimaryKey {
		if col.SQLType.Name != Image {
			if col.SQLType.Name == "TINYINT" {
				typeString = "*Boolean"
			} else {
				typeString = "*" + SQLType2Type(col.SQLType).String()
			}
		} else {
			typeString = SQLType2Type(col.SQLType).String()
		}
	} else {
		if col.SQLType.Name != Image {
			typeString = "*" + SQLType2Type(col.SQLType).String()
		} else {
			typeString = SQLType2Type(col.SQLType).String()
		}
	}

	if col.FieldName == "" {
		fmt.Println("col.FieldName is empty", col.Name, col.FieldName)
	}
	//fmt.Println(col.TableName,col.FieldName)
	content += fmt.Sprintf("	%s", strFirstToUpper(col.FieldName)+strings.Repeat(" ", maxColLen-len(col.FieldName)+1))

	content += fmt.Sprintf("%s", typeString+strings.Repeat(" ", 10-len(typeString)+1))

	content += "`xorm:\""
	col.XormTag = ""
	pkString := ""
	if col.IsPrimaryKey {
		pkString = "pk"
	}
	if pkString != "" {
		content += fmt.Sprintf(" %s", pkString)
		col.XormTag += fmt.Sprintf(" %s", pkString)
	}
	autoincrString := ""
	startwith := 1
	if col.IsAutoIncrement {
		autoincrString = "autoincr"
		if strings.Index(col.TableName, "dic_") != 0 {
			startwith = 0
		}
	}
	if autoincrString != "" {
		//		content += fmt.Sprintf(" %s", autoincrString)
		//		col.XormTag += fmt.Sprintf(" %s", autoincrString)
	}
	factString := ""

	coltype := SQLType2Type(col.SQLType).String()

	switch coltype {
	case "string":
		if col.SQLType.Name == "TEXT" || col.SQLType.Name == "CLOB" {
			factString = "text"
		} else {
			factString = fmt.Sprintf("varchar(%d)", col.Length)
		}
	default:
		if col.SQLType.Name == strings.ToUpper("binary") {
			factString = fmt.Sprintf("binary(%d)", col.Length)
		} else {
			factString = ""
		}
	}

	if factString != "" {
		content += fmt.Sprintf(" %s", factString)
		col.XormTag += fmt.Sprintf(" %s", factString)
	}
	nullString := "null"
	if !col.Nullable {
		nullString = "not null"
	}

	if nullString != "" {
		content += fmt.Sprintf(" %s", nullString)
		col.XormTag += fmt.Sprintf(" %s", nullString)
	}

	defaultString := ""
	if len(col.Default) > 0 {
		if col.Default == "null" {
			if col.IsPrimaryKey {
				defaultString = ""
			} else {
				//defaultString = "default null"
				defaultString = ""
			}
		} else {
			switch SQLType2Type(col.SQLType).String() {
			case "string":
				defaultString = fmt.Sprintf("default '%s'", strings.Replace(col.Default, "'", "", -1))
			case "time.Time":
				defaultString = fmt.Sprintf("default '%s'", col.Default)
			case "tinyint":
				defaultString = fmt.Sprintf("default %v", col.Default == "1")
			default:
				defaultString = fmt.Sprintf("default %s", col.Default)
			}
		}
	} else {
		defaultString = ""
	}

	if defaultString != "" {
		content += " " + defaultString
		col.XormTag += " " + defaultString
		col.XormTag = strings.Replace(col.XormTag, "\n", "", -1)
	} //todo 为何会有回车换行？
	//if col.FieldName == "data_interface" || col.FieldName == "need_dist" {
	//	fmt.Print(col.XormTag)
	//	fmt.Println("new line?")
	//}

	if col.FieldName == "create_date" {
		content += " created"
		col.XormTag += " created"
	}

	if col.FieldName == "modify_date" {
		content += " updated"
		col.XormTag += " updated"
	}

	content += " '" + col.FieldName + "'"
	col.XormTag += " '" + col.FieldName + "'"

	content += "\""

	content += " indexes:\""

	indexes := column["indexes"]
	indexNum, _ := strconv.ParseInt(column["indexnum"], 10, 64)

	indexString := ""

	if indexNum >= 1 {
		indexString = " " + indexes
	} else if indexNum == 0 {
		indexString = ""
	}

	content += indexString
	content += "\""

	content += fmt.Sprintf(" comment:\"%s\"", col.Comment)

	content += fmt.Sprintf(" json:\"%s,omitempty\"", col.FieldName)

	if startwith == 0 {
		content += fmt.Sprintf(" startwith:\"%d\"", startwith)
	}

	content += fmt.Sprintf(" bson:\",omitempty\"")

	content += fmt.Sprintf(" msgpack:\"%s,omitempty\"", col.FieldName)

	if string(column["fk"]) != "" {
		content += fmt.Sprintf(" fk:\"%s\"", column["fk"])
	}

	content += "`\n"

	//if col.FieldName=="data_interface"||col.FieldName=="need_dist"{
	//fmt.Println(col.XormTag)
	//}
	return content, col

}

/**
 * 字符串首字母转化为大写
 */
func strFirstToUpper(str string) string {
	//fmt.Println("str:", str)
	//if str == "bithday" {
	//	fmt.Println("str:", str)
	//}
	chars := []rune(str)
	chars[0] = []rune(strings.ToUpper(string(chars[0])))[0]
	return string(chars)
}
