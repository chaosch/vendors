package core

import (
	"fmt"
	"sort"
	"strings"
)

const (
	IndexType = iota + 1
	UniqueType
)

// database index
type Index struct {
	IsRegular bool
	Name      string
	Type      int
	Cols      IndexColumns
}

type IndexColumn struct {
	Name string
	Seq  int64
	Desc bool
}

func (index *Index) XName(tableName string) string {
	if !strings.HasPrefix(index.Name, "UDX_") &&
		!strings.HasPrefix(index.Name, "IDX_") {
		if index.Type == UniqueType {
			return fmt.Sprintf("UQE_%v_%v", tableName, index.Name)
		}
		return fmt.Sprintf("IDX_%v_%v", tableName, index.Name)
	}
	return index.Name
}

// add columns which will be composite index
func (index *Index) AddColumn(cols ...IndexColumn) {
	for _, col := range cols {
		index.Cols = append(index.Cols, col)
	}
}

func (index *Index) Equal(dst *Index) (typeEq bool, nameEq bool, colsEq bool) {
	typeEq = index.Type == dst.Type
	nameEq = index.Name == dst.Name
	//sort.StringSlice(index.Cols).Sort()
	//sort.StringSlice(dst.Cols).Sort()

	//sort.Sort(index.Cols)
	//sort.Sort(dst.Cols)

	//for i := 0; i < len(index.Cols); i++ {
	//	if index.Cols[i].Name != dst.Cols[i].Name {
	//		return false
	//	}
	//	if index.Cols[i].Desc != dst.Cols[i].Desc {
	//		return false
	//	}
	//}
	colsEq = index.Cols.Equal(dst.Cols)
	//return true
	return typeEq, nameEq, colsEq
}

func (cols IndexColumns) Equal(dst IndexColumns) bool {
	if len(cols) != len(dst) {
		return false
	}
	sort.Sort(cols)
	sort.Sort(dst)

	for i := 0; i < len(cols); i++ {
		if cols[i].Name != dst[i].Name {
			return false
		}
		if cols[i].Desc != dst[i].Desc {
			return false
		}
	}
	return true
}

// new an index
func NewIndex(name string, indexType int) *Index {
	return &Index{true, name, indexType, make([]IndexColumn, 0)}
}

type IndexColumns []IndexColumn

func (a IndexColumns) Len() int { // 重写 Len() 方法
	return len(a)
}
func (a IndexColumns) Swap(i, j int) { // 重写 Swap() 方法
	a[i], a[j] = a[j], a[i]
}
func (a IndexColumns) Less(i, j int) bool { // 重写 Less() 方法， 从大到小排序
	return a[i].Seq < a[j].Seq
}

func ColumNameArraryToIndex(colNames []string, indexName string) *Index {
	res := &Index{}
	res.Type = IndexType
	res.Name = indexName
	idc := make(IndexColumns, 0)
	for i, colname := range colNames {
		idx := IndexColumn{Name: colname, Seq: int64(i + 1), Desc: false}
		idc = append(idc, idx)
	}
	res.Cols = idc
	return res
}
