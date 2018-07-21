package common

import (
	"log"
	"reflect"
	"runtime"
	"encoding/json"
	"strings"
)


var Es = map[int]string{
	//0:    "normal error",
	401:  "error token",
	1001: "operate not support",
	1002: "table null",
	1003: "record ext",
	1004: "record not ext",
	1005: "operation return error",
	1006: "where not v3 format",
	8005: "can't find token",
	8006: "access_token timeout!",
	8007: "invalid token type,need access token!",
	8008: "user not found!",
	7001: "dictionary api error",
	9000: "database error",
	6001: "network error",
}
//swagger:model
type ResultTemplate struct {
	//成功
	//example: false
	Ok          bool        `json:"ok" msgpack:"ok"`
	//错误对象
	Err         ErrContext  `json:"err" msgpack:"err"`
	//影响记录数
	//example: 0
	Changes     int64       `json:"changes" msgpack:"changes"`
	//耗时
	//example: 1000000
	Duration    int64       `json:"duration" msgpack:"duration"`
	//sql语句耗时
	//example: 1000000
	SqlDuration int64       `json:"sql_duration" msgpack:"sql_duration"`
	//返回数据
	Data        interface{} `json:"data" msgpack:"data"`
	//成功提示:
	//example:
	OKMessage   string      `json:"ok_message"`
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
	ErrMsg  string `json:"msg" msgpack:"msg"`
	//出错行号
	//example: 35
	ErrLine int `json:"line" msgpack:"line"`
	//出错文件
	//example: /usercenter/src/micros/handler_cors_maint.go
	ErrFile string `json:"file" msgpack:"file"`
}



func (e *ErrType) Err() *ErrType {
	return e
}

func NewError(errorCode int,errorMsg string) ErrContext {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "???"
		line = 0
	}
	return &ErrType{ErrCode: errorCode, ErrMsg: errorMsg,ErrLine:line,ErrFile:file}
}



func RetChanges(changes int64) *ResultTemplate {
	res := &ResultTemplate{Ok: true}
	res.Changes = changes
	return res
}

func RetChangesStr(changes int64) string {
	res := &ResultTemplate{Ok: true}
	res.Changes = changes
	str,_:=json.Marshal(res)
	return string(str)
}


func RetErr(err ErrContext) *ResultTemplate {
	res := &ResultTemplate{Ok: false}
	if value, ok := Es[err.Err().ErrCode]; ok {
		res.Err = NewError(err.Err().ErrCode, value+":"+err.Err().ErrMsg)
		x, _ := json.Marshal(res.Err)
		log.Println(string(x))
		return res
	} else {
		res.Err = NewError(0, value+":"+err.Err().ErrMsg)
		x, _ := json.Marshal(res.Err)
		log.Println(string(x))
		return res
	}
}



func RetErrStr(err ErrContext) string {
	res := &ResultTemplate{Ok: false}
	if value, ok := Es[err.Err().ErrCode]; ok {
		res.Err = NewError(err.Err().ErrCode, value+":"+err.Err().ErrMsg)
		x, _ := json.Marshal(res)
		log.Println(string(x))
		return string(x)
	} else {
		res.Err = NewError(0, value+":"+err.Err().ErrMsg)
		x, _ := json.Marshal(res)
		log.Println(string(x))
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
	str,_:=json.Marshal(res)
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
