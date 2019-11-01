package serializer

import (
	"bytes"
	"errors"
	"gopkg.in/vmihailenco/msgpack.v2"
	"reflect"
)

type msgpackDialect struct {
	TransDialect
	ContentType string
}

func (md *msgpackDialect) GetValue(inBuf []byte, path string) (string, error) {
	var operate string
	dec := msgpack.NewDecoder(bytes.NewBuffer(inBuf))
	ops, err1 := dec.Query(path)
	if err1 == nil && len(ops) == 1 {
		v := reflect.TypeOf(ops[0])
		if v.Kind() == reflect.String {
			operate = ops[0].(string)
		} else {
			return "", errors.New("not string value " + path)
		}
	} else {
		return "", err1
	}
	return operate, nil
}

func (md *msgpackDialect) Unmarshal(inBuf []byte, inObject interface{}) error {
	err := msgpack.Unmarshal(inBuf, inObject)
	if err != nil {
		return err
	}
	return nil
}

func (md *msgpackDialect) Marshal(inObject interface{}) ([]byte, error) {
	outBuf, err := msgpack.Marshal(inObject)
	return outBuf, err
}
