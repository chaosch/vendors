package common

import (
	"egov/yaml.v2"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var Es = map[int]string{
	//0:    "normal error",
	401:  "令牌错误",
	1001: "不支持的操作",
	1002: "未定义表",
	1003: "记录存在",
	1004: "记录不存在",
	1005: "操作返回错误",
	1006: "Where条件非v3格式",
	8005: "不能找到令牌",
	8006: "令牌超时!",
	8007: "错误的令牌类型,需要身份令牌!",
	8008: "账号未找到",
	7001: "接口错误",
	9000: "数据库错误",
	6000: "验证码错误",
	6001: "网络错误",
}

var ErrorMap = make(map[string]map[int]string)

func init() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err == nil {
		files, _ := ioutil.ReadDir(dir + "/localize/")
		for _, ymlFile := range files {
			if strings.HasSuffix(ymlFile.Name(), "yaml") {
				buf, _ := ioutil.ReadFile(dir + "/localize/" + ymlFile.Name())
				key := strings.TrimRight(ymlFile.Name(), ".yaml")
				ErrorMap[key] = make(map[int]string)
				vMap := make(map[int]string)
				yaml.Unmarshal(buf, &vMap)
				ErrorMap[key] = vMap
			}
		}
	}
}



//swagger:model
type ResultTemplate struct {
	//成功
	//example: false
	Ok bool `json:"ok" msgpack:"ok"`
	//错误对象
	Err ErrContext `json:"err" msgpack:"err"`
	//影响记录数
	//example: 0
	Changes int64 `json:"changes" msgpack:"changes"`
	//耗时
	//example: 1000000
	Duration int64 `json:"duration" msgpack:"duration"`
	//sql语句耗时
	//example: 1000000
	SqlDuration int64 `json:"sql_duration" msgpack:"sql_duration"`
	//返回数据
	Data interface{} `json:"data" msgpack:"data"`
	//成功提示:
	//example:
	OKMessage string `json:"ok_message"`

	TransList interface{} `json:"trans_list"`

	PrivateCols interface{} `json:"private_cols"`
}

type ErrContext interface {
	Err() *ErrType
}

type ErrType struct {
	//错误代码
	//example: 1002
	ErrCode int `json:"code" msgpack:"code"`
	//错误消息
	//example: table null
	ErrMsg string `json:"msg" msgpack:"msg"`
	//出错行号
	//example: 35
	ErrLine int `json:"line" msgpack:"line"`
	//出错文件
	//example: /usercenter/src/micros/handler_cors_maint.go
	ErrFile  string `json:"file" msgpack:"file"`
	ErrParas []interface{}
}

func (e *ErrType) Err() *ErrType {
	return e
}

func NewError(errorCode int, errorMsg string) ErrContext {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "???"
		line = 0
	}
	if idx := strings.Index(file, "/src/"); idx >= 0 {
		file = file[idx+5:]
	}

	return &ErrType{ErrCode: errorCode, ErrMsg: errorMsg, ErrLine: line, ErrFile: file}
}

func NewError0(errorCode int, errorMsg string) ErrContext {
	_, file, line, ok := runtime.Caller(0)
	if !ok {
		file = "???"
		line = 0
	}
	if idx := strings.Index(file, "/src/"); idx >= 0 {
		file = file[idx+5:]
	}

	return &ErrType{ErrCode: errorCode, ErrMsg: errorMsg, ErrLine: line, ErrFile: file}
}

func NewErrorWithParam(errorCode int, param ...interface{}) ErrContext {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "???"
		line = 0
	}
	if idx := strings.Index(file, "/src/"); idx >= 0 {
		file = file[idx+5:]
	}

	return &ErrType{ErrCode: errorCode, ErrMsg: "", ErrLine: line, ErrFile: file, ErrParas: param}
}

func NewErrorCode(errorCode int, paras ...interface{}) ErrContext {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "???"
		line = 0
	}
	if idx := strings.Index(file, "/src/"); idx >= 0 {
		file = file[idx+5:]
	}
	return &ErrType{ErrCode: errorCode, ErrMsg: "", ErrLine: line, ErrFile: file, ErrParas: paras}
}

func (e *ErrType) ConfirmErr(lanKey string) {
	if lanKey != "" {
		if e.ErrCode != 0 {
			if v, ok := ErrorMap[lanKey][e.ErrCode]; ok {
				valueCount := strings.Count(v, "%v")
				if valueCount > 0 {
					paraArrary := e.ErrParas[0:valueCount]
					e.ErrMsg = fmt.Sprintf(v, paraArrary...)
				} else {
					e.ErrMsg = fmt.Sprintf(v)
				}
			} else {
				e.ErrMsg = ErrorMap[lanKey][9998]
			}
		} else {
			if e.ErrMsg == "" {
				e.ErrMsg = fmt.Sprintf("%+v", e.ErrParas)
			}
		}
	} else {
		e.ErrMsg = fmt.Sprint(e.ErrParas...)
	}
}

func RetChanges(changes int64) *ResultTemplate {
	res := &ResultTemplate{Ok: true}
	res.Changes = changes
	return res
}

func RetChangesStr(changes int64) string {
	res := &ResultTemplate{Ok: true}
	res.Changes = changes
	str, _ := json.Marshal(res)
	return string(str)
}

func RetErr(err ErrContext) *ResultTemplate {
	res := &ResultTemplate{Ok: false, Err: err}
	return res
	//if value, ok := Es[err.Err().ErrCode]; ok {
	//	if value==""{
	//		res.Err = NewError(err.Err().ErrCode, err.Err().ErrMsg)
	//	}else{
	//		res.Err = NewError(err.Err().ErrCode, value+":"+err.Err().ErrMsg)
	//	}
	//	//res.Err = NewError(err.Err().ErrCode, value+":"+err.Err().ErrMsg)
	//	//x, _ := json.Marshal(res.Err)
	//	//fmt.Println(string(x))
	//	return res
	//} else {
	//	if value==""{
	//		res.Err = NewError(0, err.Err().ErrMsg)
	//	}else{
	//		res.Err = NewError(0, value+":"+err.Err().ErrMsg)
	//	}
	//	//x, _ := json.Marshal(res.Err)
	//	//fmt.Println(string(x))
	//	return res
	//}
}

func RetErrStr(err ErrContext) string {
	res := &ResultTemplate{Ok: false}
	if value, ok := Es[err.Err().ErrCode]; ok {
		if value == "" {
			res.Err = NewError(err.Err().ErrCode, err.Err().ErrMsg)
		} else {
			res.Err = NewError(err.Err().ErrCode, value+":"+err.Err().ErrMsg)
		}
		x, _ := json.Marshal(res)
		//fmt.Println(string(x))
		return string(x)
	} else {
		if value == "" {
			res.Err = NewError(0, err.Err().ErrMsg)
		} else {
			res.Err = NewError(0, value+":"+err.Err().ErrMsg)
		}
		x, _ := json.Marshal(res)
		//fmt.Println(string(x))
		return string(x)
	}
}

func RetOk(result interface{}) *ResultTemplate {
	res := &ResultTemplate{Ok: true}
	if result == nil {
		res.Changes = int64(0)
		res.Err = nil
		res.Data = []interface{}{nil}
	}
	resValue := reflect.ValueOf(result)
	if resValue.Kind() == reflect.Array || resValue.Kind() == reflect.Slice {
		res.Data = result
		res.Changes = int64(resValue.Len())
	} else {
		res.Data = []interface{}{result}
		res.Changes = int64(1)
	}
	if res.Data == nil && res.Changes == 1 {
		fmt.Println(res)
	}
	return res
}

func RetOkStr(result interface{}) string {
	res := &ResultTemplate{Ok: true}
	if result == nil {
		res.Changes = int64(0)
		res.Err = nil
		res.Data = []interface{}{nil}
	}
	resValue := reflect.ValueOf(result)
	if resValue.Kind() == reflect.Array || resValue.Kind() == reflect.Slice {
		res.Data = result
		res.Changes = int64(resValue.Len())
	} else {
		res.Data = []interface{}{result}
		res.Changes = int64(1)
	}
	str, _ := json.Marshal(res)
	return string(str)
}

func MixErrContext(errs ...ErrContext) ErrContext {
	errMessage := ""
	for idx, err := range errs {
		if err != nil {
			if idx == 0 {
				errMessage += err.Err().ErrMsg
			} else {
				errMessage += ":" + err.Err().ErrMsg
			}
		} else {
			if idx == 0 {
				errMessage += ""
			} else {
				errMessage += ":"
			}

		}
	}
	if strings.Trim(errMessage, ":") == "" {
		return nil
	} else {
		return NewError(0, errMessage)
	}
}

func MixError(errs ...error) ErrContext {
	errMessage := ""
	for idx, err := range errs {
		if err != nil {
			if idx == 0 {
				errMessage += err.Error()
			} else {

				errMessage += ":" + err.Error()
			}
		} else {
			if idx == 0 {
				errMessage += ""
			} else {
				errMessage += ":"
			}

		}
	}
	if strings.Trim(errMessage, ":") == "" {
		return nil
	} else {
		return NewError(0, errMessage)
	}
}

func MixErrors(errs ...error) error {
	errMessage := ""
	for idx, err := range errs {
		if err != nil {
			if idx == 0 {
				errMessage += err.Error()
			} else {

				errMessage += ";" + err.Error()
			}
		} else {
			if idx == 0 {
				errMessage += ""
			} else {
				errMessage += ";"
			}

		}
	}
	if strings.Trim(errMessage, ";") == "" {
		return nil
	} else {
		return errors.New(errMessage)
	}
}

type ResultSet struct {
	res       []map[string]interface{}
	sortField string
}

func SetResultSet(res []map[string]interface{}, sortFields string) ResultSet {
	Rs := ResultSet{}
	Rs.res = res
	if sortFields != "" {
		for _, r := range Rs.res {
			r[sortFields] = ""
			addStr := ""
			for _, s := range strings.Split(sortFields, ",") {
				switch reflect.TypeOf(r[s]).Kind() {
				case reflect.Int64, reflect.Int, reflect.Int32:
					val := strconv.FormatInt(r[s].(int64), 10)
					addStr = strings.Repeat("0", 20-len(val)) + val
				case reflect.String:
					val := r[s].(string)
					addStr = val
				case reflect.Float64, reflect.Float32:
					val := strconv.FormatFloat(r[s].(float64), 'f', 6, 64)
					addStr = strings.Repeat("0", 20-len(val)) + val
				case reflect.Bool:
					val := "0"
					if r[s].(bool) {
						val = "1"
					}
					addStr = val
				case reflect.Struct:
					val := r[s].(time.Time)
					t := val.UnixNano()
					tStr := strconv.FormatInt(t, 10)
					addStr = strings.Repeat("0", 20-len(tStr)) + tStr
				default:
				}
				r[sortFields] = r[sortFields].(string) + "|" + addStr
			}
		}
	}
	return Rs
}

func (a *ResultSet) Len() int { // 重写 Len() 方法
	return len(a.res)
}
func (a *ResultSet) Swap(i, j int) { // 重写 Swap() 方法
	a.res[i], a.res[j] = a.res[j], a.res[i]
}
func (a *ResultSet) Less(i, j int) bool { // 重写 Less() 方法， 从大到小排序
	return a.res[i][a.sortField].(string) < a.res[j][a.sortField].(string)
}
