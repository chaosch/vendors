package xorm

// 当前操作字典
var currentSqlMap map[string]string

/**
 * 应该在各个字典中存放和默认字典不一致的内容
 * 参考"$year"指令
 */
var defaultOperatorMap = map[string]string{
	"$count":   "count(%s)",
	"$as":      " %s as %s ",
	"$isnull":  "%s is null",
	"$distinct":  "distinct %s",
	"$notnull": "%s is not null",
	"$sum":     "sum(%s)",
	"$max":     "max(%s)",
	"$min":     "min(%s)",
	"$avg":     "avg(%s)",
	"$round":     "round(%s,%s)",
}

var MysqlOperatorMap = map[string]string{
	"$count":   "count(%s)",
	"$as":      " %s as %s ",
	"$isnull":  "%s is null",
	"$distinct":  "distinct %s",
	"$notnull": "%s is not null",
	"$sum":     "sum(%s)",
	"$max":     "max(%s)",
	"$min":     "min(%s)",
	"$avg":     "avg(%s)",
	"$round":     "round(%s,%s)",
	"$year":    "year(%s)",
	"$now": "now()",
	"$if":      "if(%s,%s,%s)",
	"$strconcat":"GROUP_CONCAT(%s)",
	"$strgetdatetime":"DATE_FORMAT(%s,'%%Y-%%m-%%d %%H:%%i:%%S')",
	"$strgethour":"DATE_FORMAT(%s,'%%Y-%%m-%%d %%H')",
	"$strgetdate":"DATE_FORMAT(%s,'%%Y-%%m-%%d')",
	"$strgetmonth":"DATE_FORMAT(%s,'%%Y-%%m')",
	"$datediffmin":" timestampdiff(minute,%s,%s)",
	"$nvl":"(case  when %s is null then %s else %s end)",
	"$concat":"concat(%s,%s,%s,%s,%s)",
}

var MssqlOperatorMap = map[string]string{
	"$count":   "count(%s)",
	"$as":      " %s as %s ",
	"$isnull":  "%s is null",
	"$distinct":  "distinct %s",
	"$notnull": "%s is not null",
	"$sum":     "sum(%s)",
	"$max":     "max(%s)",
	"$min":     "min(%s)",
	"$avg":     "avg(%s)",
	"$round":     "round(%s,%s)",
	"$year": "Datename(year,%s)",
	"$now":  "getdate()",
	"$if":"case when %s then %s else %s end ",
	"$strconcat":"stuff(dbo.strconcat(','+%s),1,1,'') ",//todo 在mssql服务器运行/initscript/mssql/strconcat.bat发布strconcat聚合函数
	"$strgetdatetime":"CONVERT(varchar,%s,120)",
	"$strgethour":"CONVERT(varchar(13),%s,120)",
	"$strgetdate":"CONVERT(varchar(10),%s,120)",
	"$strgetmonth":"CONVERT(varchar(7),%s,120)",
	"$datediffmin":"datediff(minute,%s,%s)",
	"$nvl":"(case  when %s is null then %s else %s end)",
	"$concat":"concat(%s,%s,%s,%s,%s)",
}

var OracleOperatorMap = map[string]string{
	"$count":   "count(%s)",
	"$as":      " %s as %s ",
	"$isnull":  "%s is null",
	"$distinct":  "distinct %s",
	"$notnull": "%s is not null",
	"$sum":     "sum(%s)",
	"$max":     "max(%s)",
	"$min":     "min(%s)",
	"$avg":     "avg(%s)",
	"$round":     "round(%s,%s)",
	"$year": "to_char(%s,'YYYY')",
	"$now":  "sysdate",
	"$if":"case when %s then %s else %s end ",
	"$strconcat":"wm_concat(%s)",
	"$strgetdatetime":"to_char(%s,'YYYY-MM-DD HH24:MI:SS')",
	"$strgethour":"to_char(%s,'YYYY-MM-DD HH24')",
	"$strgetdate":"to_char(%s,'YYYY-MM-DD')",
	"$strgetmonth":"to_char(%s,'YYYY-MM')",
	"$datediffmin":" -ceil((%s-%s)*24*60)",
	"$nvl":"(case  when %s is null then %s else %s end)",
	"$concat":"%s||%s||%s||%s||%s",
}

var PostsqlOperatorMap = map[string]string{
}



