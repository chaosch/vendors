package main

import (
	"egov/go-xorm/core"
	"egov/go-xorm/xorm"
	"errors"
	"flag"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"io/ioutil"
	"os"
	"strings"
)

//根据mssql数据库生成数据结构定义及相关初始化

var dependency = make(map[string][]string)
var addedTable = make(map[string]bool)
var allTablesComment = make(map[string]string)
var allTablesTableName = make(map[string]string)
var schemaTableComment = make(map[string]map[string]string)
var schemaTableName = make(map[string]map[string]string)
var exactSchemaTableComment = make(map[string]map[string]string)
var exactSchemaTableName = make(map[string]map[string]string)
var pks = make(map[string]string)
var dsn = "server=192.168.4.119;port=1433;user id=sa;password=Qwer1234;database=tables;encrypt=disable;"

//var dsn = "server=yizheng.f3322.net;port=1433;user id=sas;password=Qwer1234;database=ttee44;encrypt=disable;"

func funcGetTabsPk(db_name string) string {

	content := "\n "
	content += "	//所有表的主键 \n"
	//content += "	Databases[\"" + db_name + "\"].TabsPk=make(map[string]string)  \n"
	if db_name != "all" {
		for _, v := range exactSchemaTableName[db_name] {
			content += fmt.Sprintf(`	Databases["`+db_name+`"].TabsPk["%s"]="%s"`, v, pks[v])
			content += "\n"
		}
	} else {
		for k, v := range pks {
			content += fmt.Sprintf(`	Databases["`+db_name+`"].TabsPk["%s"]="%s"`, k, v)
			content += "\n"
		}
	}
	//	content += "}\n"
	return content
}

func getAllTables(engine *xorm.Engine, dbName *string) {
	var allTables []map[string]string
	var err error
	if *dbName != "all" {
		allTables, err = engine.QueryString(`select   [name] = c.Name, [comment] = isnull(f.[value], '')
from     sys.objects c left join sys.extended_properties f on f.major_id = c.object_id and f.minor_id = 0 and f.class = 1
left join sys.extended_properties g on g.major_id=c.[object_id] and g.minor_id=0 and g.class=1 and g.name='Application'
where    c.Type = 'U' and f.name='MS_Description' 
and g.[value]='` + *dbName + `'`)
	} else {
		allTables, err = engine.QueryString(`select   [name] = c.Name, [comment] = isnull(f.[value], '')
from     sys.objects c left join sys.extended_properties f on f.major_id = c.object_id and f.minor_id = 0 and f.class = 1
where    c.Type = 'U' and f.name='MS_Description'
order by c.name`)
	}
	if err != nil {
		panic(err)
	}
	allTablesComment = make(map[string]string)
	allTablesTableName = make(map[string]string)
	for _, row := range allTables {
		allTablesComment[row["name"]] = row["comment"]
		allTablesTableName[row["comment"]] = row["name"]
	}
	if len(allTablesTableName) != len(allTablesComment) {
		panic(errors.New("表中文说明有重复"))
	} else {
		fmt.Println(fmt.Sprintf("数据库中共有表%v个", len(allTablesComment)))
	}
}

func getSchemaTables(engine *xorm.Engine, dbName string) {
	schemaTables, err := engine.QueryString(`select   [name] = c.Name, [comment] = isnull(f.[value], ''),g.[value] dbName
from     sys.objects c left join sys.extended_properties f on f.major_id = c.object_id and f.minor_id = 0 and f.class = 1
left join sys.extended_properties g on g.major_id = c.object_id and g.minor_id = 0 and g.class = 1 and g.name='Application'
where    c.Type = 'U' and f.name='MS_Description'
and c.[name] in (select object_name(major_id) from sys.extended_properties where name = 'Application' and (value='` + dbName + `'))
order by c.name`)
	if err != nil {
		panic(err)
	}
	schemaTableComment[dbName] = make(map[string]string)
	schemaTableName[dbName] = make(map[string]string)
	for _, row := range schemaTables {
		schemaTableComment[dbName][row["name"]] = row["comment"]
		schemaTableName[dbName][row["comment"]] = row["name"]
	}
	if len(schemaTableComment[dbName]) != len(schemaTableName[dbName]) {
		panic(errors.New("表中文说明有重复"))
	} else {
		fmt.Println(fmt.Sprintf("数据库%s共有表%v个", dbName, len(schemaTableComment[dbName])))
	}
}

func main() {
	var err error
	dependency, err = getDependency()
	if err != nil {
		panic(err)
	}
	schema := flag.String("schema", "all", "指定schema")
	flag.Parse()
	filename := "src/micros/table_define.go"
	engine, err := xorm.NewEngine("mssql", dsn)
	content := ""
	if err != nil {
		fmt.Println(err)
		return
	}
	getAllTables(engine, schema)
	dbList, err := listDatabase(engine, schema)
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintf("共包含字典库共有%v个数据库", len(dbList)))
	initfunc := `package micros

import (
	"sync"
	"time"
	"reflect"
	"errors"
	"fmt"
)


type Boolean bool

func (bit *Boolean) UnmarshalJSON(data []byte) error {
	asString := string(data)
	//	asString = strings.Replace(asString,"\"","",-1)
	//	var tempValue Boolean
	if asString == "1" || asString == "true" {
		*bit = true
	} else if asString == "0" || asString == "false" {
		*bit = false
	} else {
		return errors.New(fmt.Sprintf("Boolean unmarshal error: invalid input %s", asString))
	}
	return nil
}


type DatabaseInfo struct {
    TabsPk           map[string]string
	Tabs                 map[string]*sync.Pool
	TabsPas              map[string]*sync.Pool
	TabsDependency       map[string][]string                  //表依赖关系
	Dictionaries         map[string]string
	DataTables           map[string]string
	DictionariesRevert   map[string]string
	DataTablesRevert     map[string]string
	ImportedDictionaries map[string]bool
	ColumnTypes          map[string]reflect.Kind
	TabsArray 	map[string]*sync.Pool
}

var Databases map[string]*DatabaseInfo

func init(){
	Databases = make(map[string]*DatabaseInfo)
`
	initfunc += `	Databases["all"] = &DatabaseInfo{}` + "\n"
	initfunc += "	" + "allCreate()\n"

	for _, db := range dbList {
		initfunc += `	Databases["` + string(db["db_name"]) + `"] = &DatabaseInfo{}` + "\n"

		initfunc += "	" + string(db["db_name"]) + "Create()\n"
	}
	initfunc += `
}	
`
	createfunc := ""

	//所有表的集合

	fmt.Println(fmt.Sprintf("生成数据库:%s", "all"))
	createfunc += createTab("all", allTablesComment)
	createfunc += getColumnSimpleInfo("all")
	//content = initfunc + createStruct(allTablesComment, engine)
	content = initfunc + createStruct(exactSchemaTableComment["all"], engine)
	content += createfunc + funcGetTabsPk("all") + "}"
	for _, db := range dbList {
		fmt.Println(fmt.Sprintf("生成数据库:%s", db["db_name"]))
		getSchemaTables(engine, db["db_name"])
		content += createTab(string(db["db_name"]), schemaTableComment[db["db_name"]])
		content += getColumnSimpleInfo(db["db_name"])
		content += funcGetTabsPk(db["db_name"])
		content += "}"
	}

	ioutil.WriteFile(filename, []byte(content), os.ModeAppend)
	fmt.Println("代码生成完毕")

}

func listDatabase(engine *xorm.Engine, dbName *string) ([]map[string]string, error) {
	sql := ""
	if *dbName != "all" {
		sql = "select distinct value db_name from  sys.extended_properties where  name='Application' and value<>'dictionary' and value='" + *dbName + "'"
	} else {
		sql = "select distinct value db_name from  sys.extended_properties where  name='Application' and value<>'dictionary'"
	}
	return engine.QueryString(sql)
}

/**
 * 字符串首字母转化为大写
 */
func strFirstToUpper(str string) string {
	chars := []rune(str)
	chars[0] = []rune(strings.ToUpper(string(chars[0])))[0]
	return string(chars)
}

func createFields(columns []map[string]string, noSqlMode int) (string, string) {
	content := ""
	pk := ""
	maxColLen := 0
	for y := 0; y < len(columns); y++ {
		if maxColLen < len(columns[y]["name"]) {
			maxColLen = len(columns[y]["name"])
		}
	}
	for y := 0; y < len(columns); y++ {
		addContent, col := core.TransMapStringColumn(maxColLen, columns[y])
		content += addContent
		if col.IsPrimaryKey {
			pk = strings.ToLower(col.FieldName)
		}
	}
	return content, pk
}

func createStruct(tablist map[string]string, engine *xorm.Engine) string {
	fmt.Println("正在生成模型……")
	content := ""

	for tName, tComment := range tablist {
		content += fmt.Sprintf(`
//
// 表名: %s
// 说明: %s
// %s
// `, tName, tComment, tComment)
		content += fmt.Sprintf("\ntype %s struct {\n", tName)
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
                                      inner join sys.columns c1 on b1.column_id = c1.column_id and c1.object_id=a1.object_id
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
                           inner join sys.columns c2 on b2.column_id = c2.column_id and c2.object_id=a2.object_id
                           inner join sys.tables d2 on c2.object_id = d2.object_id and a2.object_id = object_id('%[1]s')
                    where  d2.object_id = a2.object_id and b2.object_id = d2.object_id and a2.is_primary_key <> 1) aa
          group by c_name) g
           on g.c_name = a.name
order by a.column_id`
		sqlQuery := fmt.Sprintf(sql, tName)
		columns, _ := engine.QueryString(sqlQuery)
		con, pk := createFields(columns, 0)
		content += con
		content += fmt.Sprintf("}\n")
		pks[tName] = pk
	}

	return content
}

func createSubTabMap() map[string]string {
	subTab := make(map[string]string)
	subTab["Tab_affairs_info"] = "tab_affairs_info"             //事项受理
	subTab["Proposers_info"] = "tab_affairs_proposers"          //事项受理-申请人
	subTab["Agent_info"] = "tab_affairs_agent"                  //事项受理-代理人
	subTab["Materials_check_info"] = "tab_materials_check"      //事项受理-材料签入签出
	subTab["Affairs_result"] = "tab_affairs_result"             //事项受理-办结信息
	subTab["Affairs_mark_result"] = "tab_affairs_mark_result"   //事项受理-流程审核结果信息
	subTab["proposers_materials"] = "tab_affair_material"       //事项受理-材料-包含结果物
	subTab["Affairs_info_outcome"] = "tab_affairs_info_outcome" //事项受理-结果物
	subTab["Agent_materials"] = "tab_affair_material_agent"     //事项受理-代理人材料-包含结果物
	subTab["Affairs_attachment"] = "tab_affairs_attachment"     //事项受理-附件表
	subTab["Affairs_node_results"] = "tab_affairs_node_results" //事项受理-已审核结点结果
	subTab["Affairs_node_current"] = "tab_affairs_node_current" // 事项受理-当前待审核结点
	subTab["Affairs_node_status"] = "tab_affairs_node_status"   // 事项受理-节点流转状态
	return subTab
}

type colTransForm struct {
	physicalName string
	logicName    string
}

func createSubObjectMap() map[string]colTransForm {
	subTab := make(map[string]colTransForm)
	//subTab["affairs_id"] = colTransForm{physicalName: "tab_affairs", logicName: "affairs_info"}
	//subTab["user_id"] = colTransForm{physicalName: "tab_user", logicName: "user_info"}
	//subTab["next_user_id"] = colTransForm{physicalName: "tab_user", logicName: "next_user_info"}
	//subTab["reg_attribution_id"] = colTransForm{physicalName: "tab_region", logicName: "reg_attribution_info"}
	//subTab["next_attribution_id"] = colTransForm{physicalName: "tab_region", logicName: "next_attribution_info"}
	//subTab["finish_attribution_id"] = colTransForm{physicalName: "tab_region", logicName: "finish_attribution_info"}
	return subTab
}

func createTab(db_name string, tablist map[string]string) string {
	content := "\n"
	content += "//将数据库对象存于pool\n"
	content += fmt.Sprintf(`func ` + db_name + `Create()  {
	Databases["` + db_name + `"].TabsPk=make(map[string]string)
	Databases["` + db_name + `"].Tabs=make(map[string]*sync.Pool)
	Databases["` + db_name + `"].TabsPas=make(map[string]*sync.Pool)
	Databases["` + db_name + `"].TabsDependency=make(map[string][]string)
	Databases["` + db_name + `"].Dictionaries=make(map[string]string)
	Databases["` + db_name + `"].DictionariesRevert=make(map[string]string)
	Databases["` + db_name + `"].DataTablesRevert=make(map[string]string)
	Databases["` + db_name + `"].DataTables=make(map[string]string)
	Databases["` + db_name + `"].ColumnTypes=make(map[string]reflect.Kind)
	Databases["` + db_name + `"].TabsArray=make(map[string]*sync.Pool)
`)

	addedTable = make(map[string]bool)
	exactSchemaTableComment[db_name] = make(map[string]string)
	exactSchemaTableName[db_name] = make(map[string]string)
	for name, comment := range tablist {
		addScript := "\n"
		currentScript := addTabs(addScript, db_name, name, comment)
		content += currentScript

	}

	fmt.Println(fmt.Sprintf("根据依赖关系实际收集表%v个", len(exactSchemaTableComment[db_name])))

	content_dictionaries := "	//字典表列表\n"
	content_datatables := "	//本地数据表列表\n"

	for tName, tComment := range exactSchemaTableComment[db_name] {
		if strings.HasPrefix(tName, "dic_") {
			content_dictionaries += fmt.Sprintf(`	Databases["`+db_name+`"]`+".DictionariesRevert[\"%s\"]=\"%s\" \n", tComment, tName)
			content_dictionaries += fmt.Sprintf(`	Databases["`+db_name+`"]`+".Dictionaries[\"%s\"]=\"%s\" \n", tName, tComment)
		} else {
			content_datatables += fmt.Sprintf(`	Databases["`+db_name+`"]`+".DataTablesRevert[\"%s\"]=\"%s\" \n", tComment, tName)
			content_datatables += fmt.Sprintf(`	Databases["`+db_name+`"]`+".DataTables[\"%s\"]=\"%s\" \n", tName, tComment)
		}

	}

	if db_name == "accept" {

		subTables := createSubTabMap()

		content += fmt.Sprintf(`    
	//
    //
    // 说明: %s
    //
`, "tab_affairs_info及其子表")
		for _, value := range subTables {
			content += fmt.Sprintf(`	Databases["`+db_name+`"]`+`.TabsPas["%s"] = &sync.Pool{
			New : func() interface{} {
				return &%s{}
			},
		}
`, value, value)
			content += fmt.Sprintf(`	Databases["`+db_name+`"]`+`.TabsArray["%s"] = &sync.Pool{
			New : func() interface{} {
			x:=make([]%s,0)
			return &x
			},
		}
`, value, value)

		}
	}

	content += createtabDependencyFast(db_name)

	content += content_dictionaries

	content += content_datatables

	return content
}

func addTabs(codeString string, dbName string, tableName string, comment string) string {
	//先加入被依赖的表
	for _, t := range dependency[tableName] {
		if _, ok := addedTable[t]; !ok {
			//fmt.Println("		",t,"not added!")
			codeString = addTabs(codeString, dbName, t, allTablesComment[t])
		}
	}
	if _, ok := addedTable[tableName]; !ok {
		addedTable[tableName] = true
		//fmt.Println("		",tableName,"added!")
		//加入自己
		codeString += fmt.Sprintf(`	Databases["`+dbName+`"]`+`.Tabs["%s"] = &sync.Pool{
			New : func() interface{} {
				return &%s{}
			},
		}
`, tableName, tableName)

		codeString += fmt.Sprintf(`	Databases["`+dbName+`"]`+`.TabsArray["%s"] = &sync.Pool{
			New : func() interface{} {
			x:=make([]%s,0)
			return &x
			},
		}
`, tableName, tableName)
		exactSchemaTableComment[dbName][tableName] = comment
		exactSchemaTableName[dbName][comment] = tableName
	}
	return codeString
}

func getDependency() (map[string][]string, error) {
	sql := `SELECT M_TAB,
       F_TAB =
          (stuff (
              (SELECT ',' + '"'+F_TAB+'"'
                 FROM (SELECT O1.NAME M_TAB, O2.NAME F_TAB
                         FROM SYSFOREIGNKEYS A,
                              SYSOBJECTS O1,
                              SYSOBJECTS O2,
                              SYSOBJECTS O3
                        WHERE     A.CONSTID = O3.ID
                              AND A.FKEYID = O1.ID
                              AND A.RKEYID = O2.ID
                              AND O1.XTYPE = 'U'
                              AND O2.XTYPE = 'U') A
                WHERE a.m_TAB = T.M_TAB
               FOR XML PATH ( '' )),
              1,
              1,
              ''))
  FROM (SELECT O1.NAME M_TAB, O2.NAME F_TAB
          FROM SYSFOREIGNKEYS A,
               SYSOBJECTS O1,
               SYSOBJECTS O2,
               SYSOBJECTS O3
         WHERE     A.CONSTID = O3.ID
               AND A.FKEYID = O1.ID
               AND A.RKEYID = O2.ID
               AND O1.XTYPE = 'U'
               AND O2.XTYPE = 'U') T
GROUP BY M_TAB`
	engine, err := xorm.NewEngine("mssql", dsn)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	tabList, err := engine.QueryString(sql)
	if err != nil {
		return nil, err
	}
	res := make(map[string][]string)
	for x := 0; x < len(tabList); x++ {
		res[tabList[x]["M_TAB"]] = strings.Split(strings.Replace(tabList[x]["F_TAB"], "\"", "", -1), ",")
	}
	return res, nil
}

func createtabDependency(dbName string) string {
	sql := `SELECT M_TAB,
       F_TAB =
          (stuff (
              (SELECT ',' + '"'+F_TAB+'"'
                 FROM (SELECT O1.NAME M_TAB, O2.NAME F_TAB
                         FROM SYSFOREIGNKEYS A,
                              SYSOBJECTS O1,
                              SYSOBJECTS O2,
                              SYSOBJECTS O3
                        WHERE     A.CONSTID = O3.ID
                              AND A.FKEYID = O1.ID
                              AND A.RKEYID = O2.ID
                              AND O1.XTYPE = 'U'
                              AND O2.XTYPE = 'U') A
                WHERE a.m_TAB = T.M_TAB
               FOR XML PATH ( '' )),
              1,
              1,
              ''))
  FROM (SELECT O1.NAME M_TAB, O2.NAME F_TAB
          FROM SYSFOREIGNKEYS A,
               SYSOBJECTS O1,
               SYSOBJECTS O2,
               SYSOBJECTS O3
         WHERE     A.CONSTID = O3.ID
               AND A.FKEYID = O1.ID
               AND A.RKEYID = O2.ID
               AND O1.XTYPE = 'U'
               AND O2.XTYPE = 'U') T
GROUP BY M_TAB`
	engine, err := xorm.NewEngine("mssql", dsn)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	tabList, err := engine.QueryString(sql)
	if err != nil {
		return ""
	}
	content := "\n "
	content += "	//表的依赖关系 \n"
	for x := 0; x < len(tabList); x++ {
		content += `	Databases["` + dbName + `"]` + ".TabsDependency[\"" + tabList[x]["M_TAB"] + "\"]=[]string{" + tabList[x]["F_TAB"] + "} \n"
	}
	return content
}

func createtabDependencyFast(dbName string) string {

	content := "\n "
	content += "	//表的依赖关系 \n"
	for tName := range schemaTableComment[dbName] {
		content += `	Databases["` + dbName + `"]` + ".TabsDependency[\"" + tName + "\"]=" + outPutStringArray(dependency[tName]) + "\n"
	}
	//for x := 0; x < len(tabList); x++ {
	//	content += `	Databases["` + dbName + `"]` + ".TabsDependency[\"" + tabList[x]["M_TAB"] + "\"]=[]string{" + tabList[x]["F_TAB"] + "} \n"
	//}
	return content
}

func getColumnSimpleInfo(db_name string) string {
	sql := ""
	if db_name != "all" {
		sql = `SELECT  a.column_name,max(data_type) data_type
  FROM INFORMATION_SCHEMA.columns a
  where table_name in ( 
  (select TABLE_NAME from  INFORMATION_SCHEMA.TABLES a
  join sys.extended_properties g on g.major_id=OBJECT_ID(table_name) and g.minor_id=0 and g.class=1 and g.name='Application' and g.[value]='` + db_name + `')
  union
  (SELECT  O2.NAME F_TAB
          FROM SYSFOREIGNKEYS A,
               SYSOBJECTS O1,
               SYSOBJECTS O2,
               SYSOBJECTS O3,
			   sys.extended_properties g
         WHERE     A.CONSTID = O3.ID
               AND A.FKEYID = O1.ID
               AND A.RKEYID = O2.ID
               AND O1.XTYPE = 'U'
               AND O2.XTYPE = 'U'
			   and (g.major_id=OBJECT_ID(o1.name) and g.minor_id=0 and g.class=1 and g.name='Application' and g.[value]='` + db_name + `'))
  )  
  group by a.COLUMN_NAME
 order by a.column_name
`
	} else {
		sql = `SELECT distinct column_name, data_type
  FROM INFORMATION_SCHEMA.columns
 WHERE COLUMN_NAME IN (SELECT COLUMN_NAME
                         FROM INFORMATION_SCHEMA.columns
                       GROUP BY COLUMN_NAME
                       HAVING count (DISTINCT data_type) = 1)`
	}
	engine, err := xorm.NewEngine("mssql", dsn)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	colList, err := engine.QueryString(sql)
	content := "\n "
	content += "	//所有列的数据类型 \n"
	content += "	Databases[\"" + db_name + "\"].ColumnTypes=map[string]reflect.Kind{\"none_field\":1 "
	for x := 0; x < len(colList); x++ {
		st := core.SQLType{}
		st.Name = colList[x]["data_type"]
		if colList[x]["column_name"] == "telephone" {
			a := 1
			_ = a
		}
		//var ColumnTypes			map[string]reflect.Kind	=map[string]reflect.Kind{"a":1,"b":1}				 //所有数据库列简单结构
		content += fmt.Sprintf(",\"%s\":%d", colList[x]["column_name"], core.SQLType2Type(st).Kind())
		//content +=fmt.Sprintf("f.ColumnTypes[\"%s\"]=%d \n",colList[x]["column_name"],core.SQLType2Type(st).Kind())
	}
	content += "}\n"

	return content
}

func outPutStringArray(strArray []string) string {
	result := "[]string{"
	for _, str := range strArray {
		result += "\"" + str + "\","
	}
	result += "}"
	return result
}

func outPutMapToSqlString(m map[string]string) string {
	result := "'x'"
	for tName := range m {
		result += ",'" + tName + "'"
	}
	return result
}
