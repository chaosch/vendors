package xorm

import (
	"strings"
	"time"
	"errors"
	"encoding/json"
	"strconv"
)

// 当前操作字典
var currentSqlMap map[string]string

/**
 * 应该在各个字典中存放和默认字典不一致的内容
 * 参考"$year"指令
 */
var defaultOperatorMap = map[string]string{
	"$year":    "year(%s)",
	"$count":   "count(%s)",
	"$as":      " %s as %s ",
	"$if":      "if(%s,%s,%s)",
	"$isnull":  "%s is null",
	"$notnull": "%s is not null",
	"$sum":     "sum(%s)",
	"$max":     "max(%s)",
	"$min":     "min(%s)",
	"$avg":     "avg(%s)",
	"$round":     "round(%s,%s)",
	"$strconcat":"GROUP_CONCAT(%s)",
	"$now": "now()",
	"$strgetdatetime":"DATE_FORMAT(%s,'%%Y-%%m-%%d %%H:%%i:%%S')",
	"$strgethour":"DATE_FORMAT(%s,'%%Y-%%m-%%d %%H')",
	"$strgetdate":"DATE_FORMAT(%s,'%%Y-%%m-%%d')",
	"$strgetmonth":"DATE_FORMAT(%s,'%%Y-%%m')",
	"$datediffmin":" timestampdiff(minute,%s,%s)",
	"$nvl":"(case  when %s is null then %s else %s end)",
	"$concat":"concat(%s,%s,%s,%s,%s)",
}

var MysqlOperatorMap = map[string]string{
}

var MssqlOperatorMap = map[string]string{
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



func CheckPs(wd string) bool {
	ps := []string{"select", "insert", "update", "delete", "'", "\\", "/*", "\\.", "union", "and", "order", "or", "into", "load_file", "outfile"}
	for _, v := range ps {
		if strings.Contains(wd, v){
			return true
		}
	}
	return true
}


func ParseValue(val interface{}) (interface{}, error) {
	switch val.(type) {
	case string:
		if valTemp, err := time.Parse(time.RFC3339, val.(string)); err == nil {
			return valTemp.Local().Format("2006-01-02 15:04:05"), nil
		}
		return val, nil
	case float64:
		floatString := strconv.FormatFloat(val.(float64), 'f', -1, 64)
		if valInt, err := strconv.Atoi(floatString); err == nil {
			return valInt, nil
		}
		return val, nil
	case int64, uint64:
		return val, nil
	case bool:
		return val, nil
	case int:
		return float64(val.(int)), nil
	case json.Number:
		ret,err:=strconv.Atoi(val.(json.Number).String())
		return int64(ret), err
	default:
		return nil, errors.New("it's not singer type")
	}
}
