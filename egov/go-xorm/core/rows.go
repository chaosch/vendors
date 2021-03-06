package core

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/binlaniua/SqlParser"
	"reflect"
	"strconv"
	"sync"
	"time"
)

type Rows struct {
	*sql.Rows
	Mapper      IMapper
	ColumnTypes map[string]reflect.Kind
	SQLPR       *sqlparse.SQLParserResult
}

func (rs *Rows) ToMapString() ([]map[string]string, error) {
	cols, err := rs.Columns()
	if err != nil {
		return nil, err
	}

	var results = make([]map[string]string, 0, 10)
	for rs.Next() {
		var record = make(map[string]string, len(cols))
		err = rs.ScanMap(&record)
		if err != nil {
			return nil, err
		}
		results = append(results, record)
	}
	return results, nil
}

// scan data to a struct's pointer according field index
func (rs *Rows) ScanStructByIndex(dest ...interface{}) error {
	if len(dest) == 0 {
		return errors.New("at least one struct")
	}

	vvvs := make([]reflect.Value, len(dest))
	for i, s := range dest {
		vv := reflect.ValueOf(s)
		if vv.Kind() != reflect.Ptr || vv.Elem().Kind() != reflect.Struct {
			return errors.New("dest should be a struct's pointer")
		}

		vvvs[i] = vv.Elem()
	}

	cols, err := rs.Columns()
	if err != nil {
		return err
	}
	newDest := make([]interface{}, len(cols))

	var i = 0
	for _, vvv := range vvvs {
		for j := 0; j < vvv.NumField(); j++ {
			newDest[i] = vvv.Field(j).Addr().Interface()
			i = i + 1
		}
	}

	return rs.Rows.Scan(newDest...)
}

var (
	fieldCache      = make(map[reflect.Type]map[string]int)
	fieldCacheMutex sync.RWMutex
)

func fieldByName(v reflect.Value, name string) reflect.Value {
	t := v.Type()
	fieldCacheMutex.RLock()
	cache, ok := fieldCache[t]
	fieldCacheMutex.RUnlock()
	if !ok {
		cache = make(map[string]int)
		for i := 0; i < v.NumField(); i++ {
			cache[t.Field(i).Name] = i
		}
		fieldCacheMutex.Lock()
		fieldCache[t] = cache
		fieldCacheMutex.Unlock()
	}

	if i, ok := cache[name]; ok {
		return v.Field(i)
	}

	return reflect.Zero(t)
}

// scan data to a struct's pointer according field name
func (rs *Rows) ScanStructByName(dest interface{}) error {
	vv := reflect.ValueOf(dest)
	if vv.Kind() != reflect.Ptr || vv.Elem().Kind() != reflect.Struct {
		return errors.New("dest should be a struct's pointer")
	}

	cols, err := rs.Columns()
	if err != nil {
		return err
	}

	newDest := make([]interface{}, len(cols))
	var v EmptyScanner
	for j, name := range cols {
		f := fieldByName(vv.Elem(), rs.Mapper.Table2Obj(name))
		if f.IsValid() {
			newDest[j] = f.Addr().Interface()
		} else {
			newDest[j] = &v
		}
	}

	return rs.Rows.Scan(newDest...)
}

type cacheStruct struct {
	value reflect.Value
	idx   int
}

var (
	reflectCache      = make(map[reflect.Type]*cacheStruct)
	reflectCacheMutex sync.RWMutex
)

func ReflectNew(typ reflect.Type) reflect.Value {
	reflectCacheMutex.RLock()
	cs, ok := reflectCache[typ]
	reflectCacheMutex.RUnlock()

	const newSize = 200

	if !ok || cs.idx+1 > newSize-1 {
		cs = &cacheStruct{reflect.MakeSlice(reflect.SliceOf(typ), newSize, newSize), 0}
		reflectCacheMutex.Lock()
		reflectCache[typ] = cs
		reflectCacheMutex.Unlock()
	} else {
		reflectCacheMutex.Lock()
		cs.idx = cs.idx + 1
		reflectCacheMutex.Unlock()
	}

	if cs.idx >= cs.value.Len() {
		fmt.Println(cs.idx, newSize-1, cs.value.Len())
	}
	return cs.value.Index(cs.idx).Addr()
}

// scan data to a slice's pointer, slice's length should equal to columns' number
func (rs *Rows) ScanSlice(dest interface{}) error {
	vv := reflect.ValueOf(dest)
	if vv.Kind() != reflect.Ptr || vv.Elem().Kind() != reflect.Slice {
		return errors.New("dest should be a slice's pointer")
	}

	vvv := vv.Elem()
	cols, err := rs.Columns()
	if err != nil {
		return err
	}

	newDest := make([]interface{}, len(cols))

	for j := 0; j < len(cols); j++ {
		if j >= vvv.Len() {
			newDest[j] = reflect.New(vvv.Type().Elem()).Interface()
		} else {
			newDest[j] = vvv.Index(j).Addr().Interface()
		}
	}

	err = rs.Rows.Scan(newDest...)
	if err != nil {
		return err
	}

	srcLen := vvv.Len()
	for i := srcLen; i < len(cols); i++ {
		vvv = reflect.Append(vvv, reflect.ValueOf(newDest[i]).Elem())
	}
	return nil
}

// scan data to a map's pointer
func (rs *Rows) ScanMap(dest interface{}) error {
	vv := reflect.ValueOf(dest)
	if vv.Kind() != reflect.Ptr || vv.Elem().Kind() != reflect.Map {
		return errors.New("dest should be a map's pointer")
	}

	cols, err := rs.Columns()
	if err != nil {
		return err
	}

	newDest := make([]interface{}, len(cols))
	vvv := vv.Elem()

	for i := range cols {
		//newDest[i] = ReflectNew(vvv.Type().Elem()).Interface()
		v := reflect.New(vvv.Type().Elem())
		newDest[i] = v.Interface()
	}

	err = rs.Rows.Scan(newDest...)
	if err != nil {
		return err
	}

	for i, name := range cols {
		vname := reflect.ValueOf(name)
		Name:=name
		for _, v := range rs.SQLPR.GetDBUser("*").TableMap {
			for _,c:=range v.ColumnMap{
				if c.GetTopAlias()==name{
					Name=c.Name
				}
			}
		}
		if vv.Elem().Kind() == reflect.Map {
			v := reflect.ValueOf(newDest[i])
			if v.Kind() == reflect.Ptr {
				v = v.Elem()
				if v.Kind() == reflect.Interface {
					v = v.Elem()
				}
			}
			//mssql
			switch v.Kind() {
			case reflect.Slice:
				//mysql
				str := string(v.Interface().([]byte))

				k := rs.ColumnTypes[Name]
				switch k {
				case reflect.Int64, reflect.Int, reflect.Int32:
					flt, _ := strconv.ParseInt(str, 10, 64)
					vvv.SetMapIndex(vname, reflect.ValueOf(flt))
				case reflect.Float64, reflect.Float32:
					it, _ := strconv.ParseFloat(str, 64)
					vvv.SetMapIndex(vname, reflect.ValueOf(it))
					//case reflect.Int:
					//	flt, _ := strconv.ParseInt(str, 10, 8)
					//	vvv.SetMapIndex(vname, reflect.ValueOf(flt >= 1))
				case reflect.Struct:
					if t, ok := v.Interface().(time.Time); ok {
						vvv.SetMapIndex(vname, reflect.ValueOf(t.Format(time.RFC3339)))
					}
				case reflect.Bool:
					it, _ := strconv.ParseInt(str, 10, 64)
					vvv.SetMapIndex(vname, reflect.ValueOf(it > 0))
				case reflect.Invalid:
					if !v.IsValid() {
						//vvv.SetMapIndex(vname, reflect.ValueOf(""))
					} else {
						vvv.SetMapIndex(vname, reflect.ValueOf(str))
					}
				default:
					vvv.SetMapIndex(vname, reflect.ValueOf(str))
				}
			case reflect.Int64, reflect.Int, reflect.Int32:
				if ck, _ := rs.ColumnTypes[name]; ck == reflect.Bool {
					vvv.SetMapIndex(vname, reflect.ValueOf(v.Interface().(int64) > 0))
				} else {
					vvv.SetMapIndex(vname, reflect.ValueOf(v.Interface().(int64)))
				}
			case reflect.String:
				vvv.SetMapIndex(vname, reflect.ValueOf(v.Interface().(string)))
			case reflect.Float64, reflect.Float32:
				vvv.SetMapIndex(vname, reflect.ValueOf(v.Interface().(float64)))
			case reflect.Bool:
				vvv.SetMapIndex(vname, reflect.ValueOf(v.Interface().(int) > 0))
			case reflect.Struct:
				if t, ok := v.Interface().(time.Time); ok {
					vvv.SetMapIndex(vname, reflect.ValueOf(t.Format(time.RFC3339)))
				}
			case reflect.Invalid:
				//var i *string = nil
				if !v.IsValid() {
					ck := rs.ColumnTypes[name]
					if ck == reflect.String {
						vvv.SetMapIndex(vname, reflect.ValueOf(""))
					}
				} else {
					ck := rs.ColumnTypes[name]
					if ck == reflect.String {
						vvv.SetMapIndex(vname, reflect.ValueOf(v.Interface().(string)))
					}
				}
			default:
				continue
			}
		} else {
			vvv.SetMapIndex(vname, reflect.ValueOf(newDest[i]).Elem())
		}
	}

	return nil
}

func IsZeroOfUnderlyingType(x interface{}) bool {
	return x == reflect.Zero(reflect.TypeOf(x)).Interface()
}

type Row struct {
	rows *Rows
	// One of these two will be non-nil:
	err error // deferred error for easy chaining
}

func (row *Row) Columns() ([]string, error) {
	if row.err != nil {
		return nil, row.err
	}
	return row.rows.Columns()
}

func (row *Row) Scan(dest ...interface{}) error {
	if row.err != nil {
		return row.err
	}
	defer row.rows.Close()

	for _, dp := range dest {
		if _, ok := dp.(*sql.RawBytes); ok {
			return errors.New("sql: RawBytes isn't allowed on Row.Scan")
		}
	}

	if !row.rows.Next() {
		if err := row.rows.Err(); err != nil {
			return err
		}
		return sql.ErrNoRows
	}
	err := row.rows.Scan(dest...)
	if err != nil {
		return err
	}
	// Make sure the query can be processed to completion with no errors.
	return row.rows.Close()
}

func (row *Row) ScanStructByName(dest interface{}) error {
	if row.err != nil {
		return row.err
	}
	defer row.rows.Close()

	if !row.rows.Next() {
		if err := row.rows.Err(); err != nil {
			return err
		}
		return sql.ErrNoRows
	}
	err := row.rows.ScanStructByName(dest)
	if err != nil {
		return err
	}
	// Make sure the query can be processed to completion with no errors.
	return row.rows.Close()
}

func (row *Row) ScanStructByIndex(dest interface{}) error {
	if row.err != nil {
		return row.err
	}
	defer row.rows.Close()

	if !row.rows.Next() {
		if err := row.rows.Err(); err != nil {
			return err
		}
		return sql.ErrNoRows
	}
	err := row.rows.ScanStructByIndex(dest)
	if err != nil {
		return err
	}
	// Make sure the query can be processed to completion with no errors.
	return row.rows.Close()
}

// scan data to a slice's pointer, slice's length should equal to columns' number
func (row *Row) ScanSlice(dest interface{}) error {
	if row.err != nil {
		return row.err
	}
	defer row.rows.Close()

	if !row.rows.Next() {
		if err := row.rows.Err(); err != nil {
			return err
		}
		return sql.ErrNoRows
	}
	err := row.rows.ScanSlice(dest)
	if err != nil {
		return err
	}

	// Make sure the query can be processed to completion with no errors.
	return row.rows.Close()
}

// scan data to a map's pointer
func (row *Row) ScanMap(dest interface{}) error {
	if row.err != nil {
		return row.err
	}
	defer row.rows.Close()

	if !row.rows.Next() {
		if err := row.rows.Err(); err != nil {
			return err
		}
		return sql.ErrNoRows
	}
	err := row.rows.ScanMap(dest)
	if err != nil {
		return err
	}

	// Make sure the query can be processed to completion with no errors.
	return row.rows.Close()
}

func (row *Row) ToMapString() (map[string]string, error) {
	cols, err := row.Columns()
	if err != nil {
		return nil, err
	}

	var record = make(map[string]string, len(cols))
	err = row.ScanMap(&record)
	if err != nil {
		return nil, err
	}

	return record, nil
}
