package xorm

import (
	"bytes"
	"fmt"
)

// Join The joinOP should be one of INNER, LEFT OUTER, CROSS etc - this will be prepended to JOIN
func (statement *Statement) Union(unionOP string, subQuery string) *Statement {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, " UNION %v %v ",unionOP, subQuery)
	statement.UnionStr += buf.String()
	return statement
}