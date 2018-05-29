package xorm

import (
	"fmt"
	"errors"
	"gopkg.in/mgo.v2/bson"
	"strings"
	"reflect"
	"strconv"
)

func (e *Engine)parseWhereMapToBson(source interface{}) (*bson.M, error) {
	if source == nil {
		return &bson.M{}, nil
	}

	if sourceMap, ok := source.(map[string]interface{}); ok {
		if res, err := e.parseMapToBson(sourceMap); err == nil {
			return res, nil
		} else {
			return nil, err
		}
	} else {
		return nil, errors.New("where not is Object")
	}
}

func (e *Engine)parseMapArrayToString(source []interface{}) (string, error) {
	var sourceString []string
	for _, temp := range source {
		if reflect.TypeOf(temp).Kind() == reflect.Float64 {
			sourceString = append(sourceString, fmt.Sprint(temp))
		} else {
			sourceString = append(sourceString, "'"+fmt.Sprint(temp)+"'")
		}
	}
	resStr := strings.Join(sourceString, ",")
	return resStr, nil
}

func (e *Engine)parseMapToBson(source map[string]interface{}) (*bson.M, error) {
	for key, val := range source {
		valType := reflect.TypeOf(val)
		switch key {
		case "and":
			fallthrough
		case "or":
			for valType.Kind() == reflect.Ptr {
				valType = valType.Elem()
			}
			if valType.Kind() == reflect.Array || valType.Kind() == reflect.Slice {
				key = "$" + key
				if bsonVal, err :=e.parseArrayToBson(val.([]interface{})); err == nil {
					return &bson.M{key: bsonVal}, nil
				} else {
					return nil, err
				}
			} else {
				return nil, errors.New(`the key 'and' or 'or' must be array.`)
			}
		case "in":
			if valType.Kind() == reflect.Map {
				valMap := val.(map[string]interface{})
				for valKey, valTemp := range valMap {
					valTempType := reflect.TypeOf(valTemp)
					if valTempType.Kind() != reflect.Array && valTempType.Kind() != reflect.Slice {
						return nil, errors.New("the key in value must array")
					}
					key = "$" + key
					var valArray []interface{}
					for _, temp := range valTemp.([]interface{}) {
						if value, err := ParseValue(temp); err == nil {
							valArray = append(valArray, value)
						} else {
							return nil, errors.New("can't support cors type in kye 'in'")
						}
					}
					return &bson.M{valKey: bson.M{key: valArray}}, nil
				}
			}
			break
		case "like":
			if valType.Kind() == reflect.Map {
				valMap := val.(map[string]interface{})
				for valKey, valTemp := range valMap {
					if reflect.TypeOf(valTemp).Kind() == reflect.String {
						regex := e.parseSqlToRegex(valTemp.(string))
						return &bson.M{valKey: bson.RegEx{Pattern: "/" + fmt.Sprint(regex) + "/", Options: "i"}}, nil
					}
					return nil, errors.New("can't support cors type in key 'like'")

				}
			}
			break
		case "lt":
			fallthrough // <  操作符
		case "lte":
			fallthrough // <=
		case "gt":
			fallthrough // >
		case "gte":
			fallthrough // >=
		case "ne": // !=
			if valType.Kind() == reflect.Map {
				key = "$" + key
				valMap := val.(map[string]interface{})
				for valKey, valTemp := range valMap {
					if value, err := ParseValue(valTemp); err == nil {
						return &bson.M{valKey: bson.M{key: value}}, nil
					} else {
						return nil, err
					}
				}
			}
			break
		case "eq": // =
			if valType.Kind() == reflect.Map {
				valMap := val.(map[string]interface{})
				for valKey, valTemp := range valMap {
					if value, err := ParseValue(valTemp); err == nil {
						return &bson.M{valKey: value}, nil
					} else {
						return nil, err
					}
				}
			} else {
				valSlice := val.([]interface{})
				if selectString, err := e.ParseSelectMapToSql(valSlice); err == nil {
					fmt.Println(selectString)
				}
			}
			break
		case "ex": // exists
			if valType.Kind() == reflect.Map {
				valMap := val.(map[string]interface{})
				for valKey, valTemp := range valMap {
					if value, err := ParseValue(valTemp); err == nil {
						return &bson.M{valKey: bson.M{"$exists": value}}, nil
					} else {
						return nil, err
					}
				}
			}
			break
		default:
			return nil, errors.New("not support this key in json where")
		}
	}
	return nil, errors.New("not support this json where")
}

func (e *Engine)parseSqlToRegex(sql string) string {
	if sql[0] != '%' {
		sql = "^" + sql
	}
	if sql[len(sql)-1] != '%' {
		sql = sql + "$"
	}
	return strings.Replace(sql, "%", ".*", -1)
}

func (e *Engine)parseArrayToBson(anArray []interface{}) ([]bson.M, error) {
	var baconAmory []bson.M
	for _, item := range anArray {
		switch item.(type) {
		case map[string]interface{}:
			if valTemp, err := e.parseMapToBson(item.(map[string]interface{})); err == nil {
				baconAmory = append(baconAmory, *valTemp)
			} else {
				return nil, err
			}
		default:
			return nil, errors.New(`array can't other type, must object!`)
		}
	}
	return baconAmory, nil
}

func (e *Engine)ParseMapToSql(source interface{}) (*string, error) {
	if source == nil {
		sql := ""
		return &sql, nil
	}

	if sourceMap, ok := source.(map[string]interface{}); ok {
		var operate string
		var resultSql string

		for key, val := range sourceMap {
			valType := reflect.TypeOf(val)
			switch key {
			case "and":
				operate = "and"
				fallthrough
			case "or":
				if operate == "" {
					operate = "or"
				}

				for valType.Kind() == reflect.Ptr {
					valType = valType.Elem()
				}
				if valType.Kind() == reflect.Array || valType.Kind() == reflect.Slice {
					if bsonVal, err := e.parseArrayToSql(operate, val.([]interface{})); err == nil {
						resultSql = " ( " + *bsonVal + " ) "
						return &resultSql, nil
					} else {
						return nil, err
					}
				} else {
					return nil, errors.New(`the key 'and' or 'or' must be array.`)
				}
			case "like":
				operate = "like"
				if valType.Kind() == reflect.Map {
					valMap := val.(map[string]interface{})
					for valKey, valTemp := range valMap {
						if reflect.TypeOf(valTemp).Kind() == reflect.String {
							resStr := valKey + " " + operate + " '" + valTemp.(string) + "'"
							return &resStr, nil
						}
						return nil, errors.New("can't support this type in key 'like'")

					}
				}
				break
			case "in":
				operate = "in"
				if valType.Kind() == reflect.Map {
					valMap := val.(map[string]interface{})
					for valKey, valTemp := range valMap {
						valTempType := reflect.TypeOf(valTemp)
						if valTempType.Kind() != reflect.Array && valTempType.Kind() != reflect.Slice {
							return nil, errors.New("the key in value must array")
						}
						if tempStr, err := e.parseMapArrayToString(valTemp.([]interface{})); err == nil {
							resStr := valKey + " " + operate + " ( " + tempStr + " ) "
							return &resStr, nil
						}

					}
				}
				break
			case "lt": // <  操作符
				operate = "<"
				fallthrough
			case "lte": // <=
				if operate == "" {
					operate = "<="
				}
				fallthrough
			case "gt": // >
				if operate == "" {
					operate = ">"
				}
				fallthrough
			case "gte": // >=
				if operate == "" {
					operate = ">="
				}
				fallthrough
			case "ne": // !=
				if operate == "" {
					operate = "!="
				}
				fallthrough
			case "eq": // =
				if operate == "" {
					operate = "="
				}
				if valType.Kind() == reflect.Map {
					valMap := val.(map[string]interface{})
					for valKey, valTemp := range valMap {
						if value, err := ParseValue(valTemp); err == nil {
							var valueString string
							valueType := reflect.TypeOf(value)
							if valueType.Kind() == reflect.Float64 || valueType.Kind() == reflect.Int || valueType.String() == "json.Number" {
								valueString = fmt.Sprint(value)
							} else {
								valueString = "'" + fmt.Sprint(value) + "'"
							}
							resultSql = valKey + " " + operate + " " + valueString
							return &resultSql, nil
						} else {
							return nil, err
						}
					}
				} else if valType.Kind() == reflect.Slice || valType.Kind() == reflect.Array {
					valSlice := val.([]interface{})
					if slicestring, err := e.ParseWhereMapToSql(valSlice); err == nil {
						_, err = strconv.ParseInt(slicestring[1], 10, 64)
						if err == nil {
							resultSql = slicestring[0] + " " + operate + " " + slicestring[1]
						} else {
							resultSql = slicestring[0] + " " + operate + " '" + slicestring[1] + "'"
						}
						return &resultSql, nil
					} else {
						return nil, err
					}
				}
			case "ex": // exists
				if operate == "" {
					operate = "IS"
				}
				if valType.Kind() == reflect.Map {
					valMap := val.(map[string]interface{})
					for valKey, valTemp := range valMap {
						if value, err := ParseValue(valTemp); err == nil {
							var isNull = false
							valueType := reflect.ValueOf(value)
							if valueType.Kind() == reflect.Bool {
								isNull = valueType.Bool()
							} else if valueType.Kind() == reflect.Int {
								isNull = valueType.Int() == 0
							}else {
								return nil, errors.New("can't support this type,please use 'bool' or 'int'")
							}

							if isNull {
								resultSql = valKey + " " + operate + " NULL"
							} else {
								resultSql = valKey + " " + operate + " NOT NULL"
							}

							return &resultSql, nil
						} else {
							return nil, err
						}
					}
				}
			default:
				return nil, errors.New("not support key in this json where")
			}
		}
		return nil, errors.New("where not is Object")
	} else {
		return nil, errors.New("where not is Object")
	}
}

func (e *Engine)parseStructToBson(inFullStruct interface{}, instruct interface{}, table string) (bson.M, error) {
	var result bson.M

	err := e.setStructValue(inFullStruct, instruct, table)
	if err != nil {
		return nil, err
	}
	data, err := bson.Marshal(inFullStruct)
	if err != nil {
		return nil, err
	}
	err = bson.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}


func (e *Engine)parseArrayToSql(operate string, anArray []interface{}) (*string, error) {
	var baconAmory string

	for key, item := range anArray {
		switch concreteVal := item.(type) {
		case map[string]interface{}:
			if valTemp, err := e.ParseMapToSql(item.(map[string]interface{})); err == nil {
				if baconAmory != "" {
					*valTemp = " " + operate + " " + *valTemp
				}

				baconAmory += *valTemp
			} else {
				return nil, err
			}
		default:
			return nil, errors.New(`array can't other type, must object!`)
			fmt.Println(key, ":::", concreteVal)
		}
	}

	return &baconAmory, nil
}



func (e *Engine)setStructValue(out interface{}, in interface{}, table string) error {
	var outValue = reflect.ValueOf(out)
	if outValue.Kind() == reflect.Ptr {
		outValue = outValue.Elem()
	}
	if outValue.Kind() != reflect.Struct {
		return errors.New("can't set not 'struct' type,please use struct")
	}
	tables := strings.Split(table, ".")[1:]

	filedValue := outValue
	for _, temp := range tables {
		temp = strings.Title(temp)
		filedValue = filedValue.FieldByName(temp)
		if filedValue.Kind() != reflect.Struct {
			return errors.New("can't set not 'struct' type,please use struct")
		}
	}

	inValue := reflect.ValueOf(in)
	if inValue.Kind() == reflect.Ptr {
		inValue = inValue.Elem()
	}

	if filedValue.Type() != inValue.Type() {
		return errors.New("can't set value , because the type not eq")
	}

	filedValue.Set(inValue)
	return nil
}
