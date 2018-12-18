package xorm


// union union_operator should be one of INNER, LEFT OUTER, CROSS etc - this will be prepended to JOIN
func (session *Session) Union(unionType string, subQuery string) *Session {
	session.Statement.Union(unionType, subQuery)
	return session
}