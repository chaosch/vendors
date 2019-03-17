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
	p := sqlparse.NewSQLParser(`
SELECT
	t1.city_name a,
	t2.province_name b
FROM
	dic_city t1
LEFT JOIN dic_province t2 ON t1.province_id = t2.province_id
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
