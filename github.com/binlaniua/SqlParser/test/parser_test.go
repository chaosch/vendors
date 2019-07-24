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
	p.affairs_proposers_id apid,
	p.split_code psc,
	ai.split_code,
	p.affairs_proposers_id,
	ai.affairs_info_id,
	ai.affairs_info_code,
	af.affairs_name,
	ai.user_name,
	ai.affairs_start,
	ai.next_user_name,
	ai.node_current_name,
	ai.consult_affair,
	af.affairs_id aid,
	af.affairs_local_id,
	p.proposer_name,
	p.phone,
	ai.modify_userid,
	ai.modify_username,
	ai.modify_date,
	ai.reg_attribution_id,
	ai.reg_attribution_name,
	reg.region_id AS modify_region_id,
	reg.region_name AS modify_region_name,
	ai.affairs_deadline,
	cun.standard_code_id,
	cun.standard_code_name,
	cun.standard_code AS node_standard_code,
	cun.node_name AS node_result_name,
	cun.node_id,
	p.credential_type,
	p.credential_code,
	p.proposer_gender,
	pt.credential_type_id,
	pt.credential_type_name,
	pg.gender_name,
	bm.department_id,
	bm.department_name,
	cun.monitor_id,
	cun.red_or_yellow,
	cun.is_hook,
	cun.warnning_deadline,
	cun.warnning_deadline_date,
	cun.deadline,
	cun.deadline_date,
	cun.time_limit,
	cun.red_or_yellow_m,
	cun.warnning_deadline_m,
	cun.deadline_m,
	cun.red_or_yellow_mm,
	cun.warnning_deadline_mm,
	cun.deadline_mm
FROM
	tab_affairs_info AS ai
INNER JOIN (
	SELECT
		cu.affairs_info_id,
		c.standard_code,
		c.standard_code_id,
		c.standard_code_name,
		n.node_id,
		n.node_name,
		m.red_or_yellow,
		m.is_hook,
		m.warnning_deadline,
		m.warnning_deadline_date,
		m.deadline,
		m.deadline_date,
		max(m.monitor_id) AS monitor_id,
		m.time_limit,
		mm.red_or_yellow red_or_yellow_m,
		mm.warnning_deadline warnning_deadline_m,
		mm.deadline deadline_m,
		mmm.red_or_yellow red_or_yellow_mm,
		mmm.warnning_deadline warnning_deadline_mm,
		mmm.deadline deadline_mm
	FROM
		tab_affairs_node_current AS cu
	INNER JOIN tab_affairs_mark_result AS mr ON cu.node_id = mr.node_id
	AND cu.affairs_info_id = mr.affairs_info_id
	AND mr.node_result_id IS NULL
	INNER JOIN tab_affairs_info AS a ON cu.affairs_info_id = a.affairs_info_id
	INNER JOIN tab_flow_node_info AS n ON cu.node_id = n.node_id
	INNER JOIN dic_node_standard_code AS c ON n.node_standard_code = c.standard_code_id
	INNER JOIN (
		SELECT
			au2.node_id
		FROM
			tab_authorization AS au2
		WHERE
			(
				au2.user_id IN (3025855999639557)
				AND au2.enable_mark = 1
				AND au2.delete_mark = 0
			)
		UNION
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
				UNION
					SELECT
						au1.node_id
					FROM
						tab_authorization AS au1
					INNER JOIN tab_group_user_relation AS urg ON urg.group_id = au1.group_id
					AND urg.enable_mark = 1
					AND urg.delete_mark = 0
					AND au1.enable_mark = 1
					AND au1.delete_mark = 0
					WHERE
						urg.user_id IN (3025855999639557)
	) AS au ON au.node_id = n.node_id
	LEFT JOIN tab_monitors AS m ON cu.affairs_info_id = m.affairs_info_id
	AND mr.mark_id = m.mark_id
	AND m.is_finished = 0
	AND m.delete_mark = 0
	LEFT JOIN tab_monitors AS mm ON cu.affairs_info_id = mm.affairs_info_id
	AND mm.is_finished = 0
	AND mm.delete_mark = 0
	AND mm.is_affair = 1
	LEFT JOIN tab_monitors AS mmm ON mmm.parent_monitor_id = m.monitor_id
	AND mmm.affairs_info_id = cu.affairs_info_id
	AND mmm.is_finished = 0
	WHERE
		(
			(
				(
					mr.next_user_id IS NULL
					OR mr.next_user_id = 22
				)
				AND (
					(
						mr.next_attribution_id IS NULL
						AND mr.next_attribution_admin_level IS NULL
					)
					OR mr.next_attribution_id = 3025855999639553
					OR (
						mr.next_attribution_id IS NULL
						AND mr.next_attribution_admin_level = 1
						AND mr.next_attribution_area_id = 1
					)
				)
				AND a.enable_mark = 1
				AND a.delete_mark = 0
			)
			AND (
				n.node_standard_code = 1002
				OR n.node_standard_code = 1004
				OR n.node_standard_code = 1007
				OR n.node_standard_code = 1001
			)
		)
	GROUP BY
		cu.affairs_info_id,
		c.standard_code,
		c.standard_code_id,
		c.standard_code_name,
		n.node_id,
		n.node_name,
		m.red_or_yellow,
		m.is_hook,
		m.warnning_deadline,
		m.warnning_deadline_date,
		m.deadline,
		m.deadline_date,
		m.time_limit,
		mm.red_or_yellow,
		mm.warnning_deadline,
		mm.deadline,
		mmm.red_or_yellow,
		mmm.warnning_deadline,
		mmm.deadline
) AS cun ON ai.affairs_info_id = cun.affairs_info_id
INNER JOIN tab_affairs_proposers AS p ON ai.affairs_info_id = p.affairs_info_id
LEFT JOIN tab_user AS us ON us.user_id = ai.modify_userid
LEFT JOIN tab_region AS reg ON reg.region_id = us.location_id
LEFT JOIN dic_credential_type AS pt ON p.credential_type = pt.credential_type_id
LEFT JOIN dic_gender AS pg ON p.proposer_gender = pg.gender_id
INNER JOIN tab_affairs AS af ON af.affairs_id = ai.affairs_id
LEFT JOIN dic_department AS bm ON af.department_id = bm.department_id
WHERE
	(
		(
			(
				ai.enable_mark = 1
				AND ai.delete_mark = 0
			)
			AND p.primary_proposer = 1
		)
		AND ai.delete_mark = 0
	)
ORDER BY
	ai.create_date DESC
LIMIT 8
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
