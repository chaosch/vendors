package serializer

import (
	"github.com/tidwall/gjson"
	"github.com/pquerna/ffjson/ffjson"
	"errors"
	"encoding/json"
	//"strings"
	//"reflect"
	//"egov/json"
	"strings"
	//"strconv"
	//"fmt"
)

type jsonDialect struct {
	TransDialect
	ContentType string
}

func (jd *jsonDialect) GetValue(inBuf []byte, path string) (string, error) {
	var e error
	result := gjson.GetBytes(inBuf, path)
	if result.String() == "" {
		e = errors.New("cant not get " + path)
	}
	return result.String(), e
}

func (jd *jsonDialect) Unmarshal(inBuf []byte, inObject interface{}) (error) {

	decoder := json.NewDecoder(strings.NewReader(string(inBuf)))
	//decoder.UseNumber()
	err := decoder.Decode(&inObject)

	if err != nil {
		return err
	}


	//inObjectAddr:=ReplaceJsonNumber(inObject)
	////
	//inObject=&inObjectAddr
	//
	//fmt.Println(fmt.Sprintf("%+v",inObjectAddr))
	//	err := ffjson.Unmarshal(inBuf, inObject)
	return nil
}

func (jd *jsonDialect) Marshal(inObject interface{}) ([]byte, error) {
	outBuf, err := ffjson.Marshal(inObject)
	return outBuf, err
}


//func (jd *jsonDialect) ConvertIntToBool(obj map[string]interface{}) (int64, map[string]interface{}) {
//	converted := int64(0)
//	for key, kind := range core.ColumnTypes {
//		if kind == reflect.Bool {
//			if x, ok := obj[key]; ok {
//				if reflect.ValueOf(x).Kind() != reflect.Bool {
//					obj[key] = x.(float64) > 0
//					converted++
//				}
//			}
//		}
//		continue
//	}
//	return converted, obj
//}
//
//func (jd *jsonDialect) ConvertIntToBoolInObject(inBuf []byte) ([]byte, error, int64) {
//	converted := int64(0)
//	if jd.IntConvertToBool {
//		res := gjson.GetBytes(inBuf, "object")
//		if res.Raw == "" {
//			return inBuf, nil, 0
//		}
//		var mpobject map[string]interface{}
//		err := ffjson.Unmarshal([]byte(res.Raw), &mpobject)
//		if err != nil {
//			return nil, NewError(0, err.Error()), 0
//		}
//		converted, mpobject = jd.ConvertIntToBool(mpobject)
//		mpobyte, e := ffjson.Marshal(mpobject)
//		if e != nil {
//			return nil, NewError(0, e.Error()), 0
//		}
//		inBuf = bytes.Replace(inBuf, []byte(res.Raw), mpobyte, len([]byte(res.Raw)))
//	}
//	return inBuf, nil, converted
//}
//
//func (jd *jsonDialect) ConvertBooltoIntInResult(inBuf []byte) ([]byte, error) {
//	if jd.BoolValueReturnInt {
//		inBuf = bytes.Replace(inBuf, []byte(":true"), []byte(":1"), -1)
//		inBuf = bytes.Replace(inBuf, []byte(":false"), []byte(":0"), -1)
//	}
//	return inBuf, nil
//}
