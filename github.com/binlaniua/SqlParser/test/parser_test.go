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
		`SELECT  tap.proposer_name,tap1.proposer_name1 from tab_affairs_proposers  as tap force index (ind_name)  
      join  tab_affairs_proposers1  as tap1 force index (ind_name1) 
       where tap1.proposer_name like binnary '%abc%'
	`)
	r, err := p.DoParser()

	//fmt.Println(r.GetDBUser("*").TableMap["tab_flow_node_info"].Alias.Name,r.GetDBUser("*").TableMap["tab_flow_node_info"].GetTopAlias())

	if err != nil {
		t.Error(err)
	} else {
		t.Log(r.String())
	}

	//fmt.Println(p.GetResult().GetDBUser("*").TableMap["b"].GetTopAlias())
	//fmt.Println(p.GetResult().GetDBUser("*").TableMap["b"].ColumnMap["a"].GetTopAlias())

}


func TestSelect2(t *testing.T) {
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
		`SELECT m.city_id,m.city_name FROM dic_city AS m use index(FK_5c874b534eb9b20001d3b6c4,UQE_dic_city_idx_dic_city_city_code)'
	`)
	r, err := p.DoParser()

	//fmt.Println(r.GetDBUser("*").TableMap["tab_flow_node_info"].Alias.Name,r.GetDBUser("*").TableMap["tab_flow_node_info"].GetTopAlias())

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
