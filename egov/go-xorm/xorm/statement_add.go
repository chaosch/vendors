package xorm

import (
	"bytes"
	"fmt"
	"reflect"
)

// Join The joinOP should be one of INNER, LEFT OUTER, CROSS etc - this will be prepended to JOIN
func (statement *Statement) Union(unionOP string, subQuery string) *Statement {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, " UNION %v %v ", unionOP, subQuery)
	statement.UnionStr += buf.String()
	return statement
}

func (engine *Engine) TbName(v reflect.Value) string {
	if tb, ok := v.Interface().(TableName); ok {
		return tb.TableName()
	}

	if v.Type().Kind() == reflect.Ptr {
		if tb, ok := reflect.Indirect(v).Interface().(TableName); ok {
			return tb.TableName()
		}
	} else if v.CanAddr() {
		if tb, ok := v.Addr().Interface().(TableName); ok {
			return tb.TableName()
		}
	}
	return engine.TableMapper.Obj2Table(reflect.Indirect(v).Type().Name())
}
