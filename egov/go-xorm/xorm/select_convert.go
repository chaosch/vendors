package xorm

import (
	"gopkg.in/mgo.v2/bson"
	"reflect"
	"errors"
	"strings"
	"fmt"
)

func (e *Engine) ParseSelectMapToBson(source interface{}) (bson.M, error) {
	sourceValue := reflect.ValueOf(source)
	if sourceValue.Kind() == reflect.Slice || sourceValue.Kind() == reflect.Array {
		return nil, nil
	} else {
		return nil, errors.New("can't support type,you must use 'slice' or 'array'")
	}
}

func (e *Engine) ParseWhereMapToSql(source interface{}) ([]string, error) {
	sourceValue := reflect.ValueOf(source)
	if sourceValue.Kind() == reflect.Slice || sourceValue.Kind() == reflect.Array {
		resStringSlice := []string{}
		for i := 0; i < sourceValue.Len(); i++ {
			tempValue := sourceValue.Index(i)
			if tempValue.Kind() == reflect.Interface {
				tempValue = tempValue.Elem()
			}

			switch tempValue.Kind() {
			case reflect.String:
				wd := tempValue.Interface().(string)
				rt := CheckPs(wd)
				if !rt{
					return nil, errors.New("has sp words")
				}
				resStringSlice = append(resStringSlice, wd)
			case reflect.Map:
				if v, err := e.parseTempMap(tempValue.Interface().(map[string]interface{})); err == nil {
					resStringSlice = append(resStringSlice, v)
				} else {
					return nil, err
				}
			case reflect.Float64:
				resStringSlice = append(resStringSlice, fmt.Sprint(tempValue.Interface().(float64)))
			default:
				return nil, errors.New("can't support type,you must use 'string' or 'map'")
			}
		}
		return resStringSlice, nil
	} else {
		return nil, errors.New("can't support type,you must use 'slice' or 'array'")
	}
}

func (e *Engine) ParseSelectMapToSql(source interface{}) (string, error) {
	sourceValue := reflect.ValueOf(source)
	if sourceValue.Kind() == reflect.Slice || sourceValue.Kind() == reflect.Array {
		resStringSlice := []string{}
		for i := 0; i < sourceValue.Len(); i++ {
			tempValue := sourceValue.Index(i)
			if tempValue.Kind() == reflect.Interface {
				tempValue = tempValue.Elem()
			}

			switch tempValue.Kind() {
			case reflect.String:
				wd := tempValue.Interface().(string)
				rt := CheckPs(wd)
				if !rt{
					return "", errors.New("has sp words")
				}
				resStringSlice = append(resStringSlice, wd)
			case reflect.Map:
				if v, err := e.parseTempMap(tempValue.Interface().(map[string]interface{})); err == nil {
					resStringSlice = append(resStringSlice, v)
				} else {
					return "", err
				}
			default:
				return "", errors.New("can't support type,you must use 'string' or 'map'")
			}
		}
		return strings.Join(resStringSlice, ","), nil
	} else {
		return "", errors.New("can't support type,you must use 'slice' or 'array'")
	}
}

func(e *Engine)  parseTempArray(operator interface{}, argsIn ...interface{}) ([]interface{}, error) {
	var argsOut []interface{}

	switch reflect.ValueOf(operator).Kind() {
	case reflect.String:
		wd := operator.(string)
		rt := CheckPs(wd)
		if !rt{
			return nil, errors.New("has sp words")
		}
		argsOut = append(argsOut, wd)
	case reflect.Map:
		if v, err := e.parseTempMap(operator.(map[string]interface{})); err == nil {
			argsOut = append(argsOut, v)
		} else {
			return nil, err
		}
	default:
		return nil, errors.New("can't support cors '" + reflect.ValueOf(operator).Kind().String() + "'")
	}
	for _, c := range argsIn {
		switch reflect.ValueOf(c).Kind() {
		case reflect.String:
			argsOut = append(argsOut, c)
		case reflect.Map:
			if v, err := e.parseTempMap(c.(map[string]interface{})); err == nil {
				argsOut = append(argsOut, v)
			} else {
				return nil, err
			}
		default:
			argsOut = append(argsOut, fmt.Sprintf("%v", c))
		}
	}
	return argsOut, nil
}

func (e *Engine) parseTempValue(key string, value interface{}) (string, error) {
	switch reflect.ValueOf(value).Kind() {
	case reflect.String:
		rt := CheckPs(value.(string))
		if !rt{
			return "", errors.New("has sp words")
		}

		return fmt.Sprintf(key, value), nil
	case reflect.Map:
		if val, err := e.parseTempMap(value.(map[string]interface{})); err == nil {
			return fmt.Sprintf(key, val), nil
		} else {
			return "", err
		}
	case reflect.Slice, reflect.Array:
		argsIn := value.([]interface{})
		if val, err := e.parseTempArray(argsIn[0], argsIn[1:]...); err == nil {
			return fmt.Sprintf(key, val...), nil
		} else {
			return "", err
		}
	default:
		return "", errors.New("arg error")
	}
}

func (e *Engine) parseTempMap(temp map[string]interface{}) (string, error) {
	for key, value := range temp {
		if v, ok := e.CurrentSqlMap[key]; ok {
			if data, err := e.parseTempValue(v, value); err == nil {
				return data, nil
			} else {
				return "", err
			}
		} else {
			if v, ok = defaultOperatorMap[key]; ok {
				if data, err := e.parseTempValue(v, value); err == nil {
					return data, nil
				} else {
					return "", err
				}
			} else {
				return "", errors.New("can't support cors function")
			}
		}
	}
	return "", errors.New("cors not function")
}
