package test

import (
	"github.com/binlaniua/SqlParser"
	"testing"
)

//-------------------------------------
//
//
//
//-------------------------------------
func TestSelect(t *testing.T) {
	//	p := sqlparse.NewSQLParser(`
	//SELECT
	//	t1.city_name ,
	//	t2.province_name b
	//FROM
	//	dic_city t1
	//LEFT JOIN dic_province t2 ON t1.province_id = t2.province_id
	//	`)

	//	p := sqlparse.NewSQLParser(`
	//  select
	//      sum(t1.e1) ,
	//      t2.f1,
	//      t3.e ccc,
	//      t3.f ddd
	//  from
	//      (select t2.b1 as e1, t2.d1 as f1 from (select a as b1, c as d1 from xx.table1) t2) t1,
	//      (select t2.b as e, t2.d as f from (select a as b, c as d from yy.table2) t2) t3
	//`)

	p := sqlparse.NewSQLParser(
		`SELECT
	b.affairs_info_id
FROM
	tab_affairs_node_status AS a
INNER JOIN tab_affairs_info AS b ON b.affairs_info_id = a.affairs_info_id
INNER JOIN tab_affairs AS c ON c.affairs_id = b.affairs_id
LEFT JOIN tab_affairs_proposers AS e ON e.affairs_info_id = b.affairs_info_id
AND e.primary_proposer = 1
INNER JOIN (
	SELECT
		au2.node_id
	FROM
		tab_authorization AS au2
	WHERE
		(
			au2.user_id = 22
			AND au2.enable_mark = 1
			AND au2.delete_mark = 0
		)
	UNION
		SELECT
			au1.node_id
		FROM
			tab_authorization AS au1
		INNER JOIN tab_group_user_relation AS urg ON urg.group_id = au1.group_id
		AND urg.user_id = 22
		AND urg.enable_mark = 1
		AND urg.delete_mark = 0
		AND au1.enable_mark = 1
		AND au1.delete_mark = 0
) AS au ON au.node_id = a.node_id
WHERE
	(
		(
			(
				a.enable_mark = 1
				AND a.delete_mark = 0
			)
			AND a. STATUS = '2'
		)
		AND a.delete_mark = 0
	)
GROUP BY
	b.affairs_info_id,
	b.affairs_info_code,
	c.affairs_id,
	c.affairs_name,
	b.affairs_start,
	a. STATUS,
	b.user_name,
	e.proposer_name,
	a.status_id,
	a.node_id
`)
	r, err := p.DoParser()
	if err != nil {
		t.Error(err)
	} else {
		t.Log(r.String())
	}
	//fmt.Println(p.GetResult().GetDBUser("*").TableMap["b"].GetTopAlias())
	//fmt.Println(p.GetResult().GetDBUser("*").TableMap["b"].ColumnMap["a"].GetTopAlias())

}

//-------------------------------------
//
//
//
//-------------------------------------
func TestInsert(t *testing.T) {
	p := sqlparse.NewSQLParser(`
		insert into table1 (a,b,c,d) values (1,2,3,4)
	`)
	r, err := p.DoParser()
	if err != nil {
		t.Error(err)
	} else {
		t.Log(r.String())
	}
}

//-------------------------------------
//
//
//
//-------------------------------------
func TestUpdate(t *testing.T) {
	p := sqlparse.NewSQLParser(`
		update table1 set a = 1, b = 2, c = 3
	`)
	r, err := p.DoParser()
	if err != nil {
		t.Error(err)
	} else {
		t.Log(r.String())
	}
}

//-------------------------------------
//
//
//
//-------------------------------------
func TestDelete(t *testing.T) {
	p := sqlparse.NewSQLParser(`
		delete from table1 where a = 1
	`)
	r, err := p.DoParser()
	if err != nil {
		t.Error(err)
	} else {
		t.Log(r.String())
	}
}
