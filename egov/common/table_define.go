package common

import (
	"sync"
	"time"
	"reflect"
)

type DatabaseInfo struct {
	Tabs                 map[string]*sync.Pool
	TabsPas              map[string]*sync.Pool
	TabsDependency       map[string][]string                  //表依赖关系
	Dictionaries         map[string]string
	DataTables           map[string]string
	DictionariesRevert   map[string]string
	DataTablesRevert     map[string]string
	ImportedDictionaries map[string]bool
	ColumnTypes          map[string]reflect.Kind
}

var Databases map[string]*DatabaseInfo

func init(){
	Databases = make(map[string]*DatabaseInfo)
	Databases["all"] = &DatabaseInfo{}
	allCreate()
	Databases["accept"] = &DatabaseInfo{}
	acceptCreate()
	Databases["community"] = &DatabaseInfo{}
	communityCreate()
	Databases["outter"] = &DatabaseInfo{}
	outterCreate()
	Databases["synccenter"] = &DatabaseInfo{}
	synccenterCreate()

}	

//将数据库对象存于pool
func acceptCreate()  {
	Databases["accept"].Tabs=make(map[string]*sync.Pool)
	Databases["accept"].TabsPas=make(map[string]*sync.Pool)
	Databases["accept"].TabsDependency=make(map[string][]string)
	Databases["accept"].Dictionaries=make(map[string]string)
	Databases["accept"].DictionariesRevert=make(map[string]string)
	Databases["accept"].DataTablesRevert=make(map[string]string)
	Databases["accept"].DataTables=make(map[string]string)
	Databases["accept"].ColumnTypes=make(map[string]reflect.Kind)

	Databases["accept"].Tabs["tab_virtual_machines"] = &sync.Pool{
			New : func() interface{} {
				return &tab_virtual_machines{}
			},
		}
	Databases["accept"].Tabs["tab_vm_accounts"] = &sync.Pool{
			New : func() interface{} {
				return &tab_vm_accounts{}
			},
		}

	Databases["accept"].Tabs["dic_user_level"] = &sync.Pool{
			New : func() interface{} {
				return &dic_user_level{}
			},
		}
	Databases["accept"].Tabs["dic_gender"] = &sync.Pool{
			New : func() interface{} {
				return &dic_gender{}
			},
		}
	Databases["accept"].Tabs["dic_positions"] = &sync.Pool{
			New : func() interface{} {
				return &dic_positions{}
			},
		}
	Databases["accept"].Tabs["tab_user"] = &sync.Pool{
			New : func() interface{} {
				return &tab_user{}
			},
		}
	Databases["accept"].Tabs["tab_group"] = &sync.Pool{
			New : func() interface{} {
				return &tab_group{}
			},
		}
	Databases["accept"].Tabs["tab_group_user_relation"] = &sync.Pool{
			New : func() interface{} {
				return &tab_group_user_relation{}
			},
		}

	Databases["accept"].Tabs["dic_time_limit_type"] = &sync.Pool{
			New : func() interface{} {
				return &dic_time_limit_type{}
			},
		}
	Databases["accept"].Tabs["dic_node_type"] = &sync.Pool{
			New : func() interface{} {
				return &dic_node_type{}
			},
		}
	Databases["accept"].Tabs["dic_node_standard_code"] = &sync.Pool{
			New : func() interface{} {
				return &dic_node_standard_code{}
			},
		}
	Databases["accept"].Tabs["dic_affairs_type"] = &sync.Pool{
			New : func() interface{} {
				return &dic_affairs_type{}
			},
		}
	Databases["accept"].Tabs["dic_subject_nature"] = &sync.Pool{
			New : func() interface{} {
				return &dic_subject_nature{}
			},
		}
	Databases["accept"].Tabs["dic_proposer_type"] = &sync.Pool{
			New : func() interface{} {
				return &dic_proposer_type{}
			},
		}
	Databases["accept"].Tabs["dic_universal_range"] = &sync.Pool{
			New : func() interface{} {
				return &dic_universal_range{}
			},
		}
	Databases["accept"].Tabs["dic_handling_form"] = &sync.Pool{
			New : func() interface{} {
				return &dic_handling_form{}
			},
		}
	Databases["accept"].Tabs["dic_department"] = &sync.Pool{
			New : func() interface{} {
				return &dic_department{}
			},
		}
	Databases["accept"].Tabs["dic_affairs_subject"] = &sync.Pool{
			New : func() interface{} {
				return &dic_affairs_subject{}
			},
		}
	Databases["accept"].Tabs["dic_corporate_subject"] = &sync.Pool{
			New : func() interface{} {
				return &dic_corporate_subject{}
			},
		}
	Databases["accept"].Tabs["dic_affairs_objects"] = &sync.Pool{
			New : func() interface{} {
				return &dic_affairs_objects{}
			},
		}
	Databases["accept"].Tabs["dic_life_events"] = &sync.Pool{
			New : func() interface{} {
				return &dic_life_events{}
			},
		}
	Databases["accept"].Tabs["dic_net_level"] = &sync.Pool{
			New : func() interface{} {
				return &dic_net_level{}
			},
		}
	Databases["accept"].Tabs["dic_result_receive"] = &sync.Pool{
			New : func() interface{} {
				return &dic_result_receive{}
			},
		}
	Databases["accept"].Tabs["tab_affairs"] = &sync.Pool{
			New : func() interface{} {
				return &tab_affairs{}
			},
		}
	Databases["accept"].Tabs["tab_flow_template"] = &sync.Pool{
			New : func() interface{} {
				return &tab_flow_template{}
			},
		}
	Databases["accept"].Tabs["tab_flow_node_info"] = &sync.Pool{
			New : func() interface{} {
				return &tab_flow_node_info{}
			},
		}
	Databases["accept"].Tabs["tab_flow_node_result_info"] = &sync.Pool{
			New : func() interface{} {
				return &tab_flow_node_result_info{}
			},
		}
	Databases["accept"].Tabs["dic_administrative_level"] = &sync.Pool{
			New : func() interface{} {
				return &dic_administrative_level{}
			},
		}
	Databases["accept"].Tabs["dic_city"] = &sync.Pool{
			New : func() interface{} {
				return &dic_city{}
			},
		}
	Databases["accept"].Tabs["dic_district"] = &sync.Pool{
			New : func() interface{} {
				return &dic_district{}
			},
		}
	Databases["accept"].Tabs["dic_street"] = &sync.Pool{
			New : func() interface{} {
				return &dic_street{}
			},
		}
	Databases["accept"].Tabs["dic_community"] = &sync.Pool{
			New : func() interface{} {
				return &dic_community{}
			},
		}
	Databases["accept"].Tabs["tab_region"] = &sync.Pool{
			New : func() interface{} {
				return &tab_region{}
			},
		}
	Databases["accept"].Tabs["dic_apply_from"] = &sync.Pool{
			New : func() interface{} {
				return &dic_apply_from{}
			},
		}
	Databases["accept"].Tabs["dic_satisfy_level"] = &sync.Pool{
			New : func() interface{} {
				return &dic_satisfy_level{}
			},
		}
	Databases["accept"].Tabs["tab_affairs_info"] = &sync.Pool{
			New : func() interface{} {
				return &tab_affairs_info{}
			},
		}
	Databases["accept"].Tabs["tab_affairs_mark_result"] = &sync.Pool{
			New : func() interface{} {
				return &tab_affairs_mark_result{}
			},
		}
	Databases["accept"].Tabs["tab_monitors"] = &sync.Pool{
			New : func() interface{} {
				return &tab_monitors{}
			},
		}
	Databases["accept"].Tabs["tab_monitor_hook_log"] = &sync.Pool{
			New : func() interface{} {
				return &tab_monitor_hook_log{}
			},
		}


	Databases["accept"].Tabs["tab_region_events"] = &sync.Pool{
			New : func() interface{} {
				return &tab_region_events{}
			},
		}


	Databases["accept"].Tabs["tab_affairs_attachment"] = &sync.Pool{
			New : func() interface{} {
				return &tab_affairs_attachment{}
			},
		}

	Databases["accept"].Tabs["tab_affairs_delicate"] = &sync.Pool{
			New : func() interface{} {
				return &tab_affairs_delicate{}
			},
		}

	Databases["accept"].Tabs["tab_coporation_info"] = &sync.Pool{
			New : func() interface{} {
				return &tab_coporation_info{}
			},
		}

	Databases["accept"].Tabs["tab_system_config"] = &sync.Pool{
			New : func() interface{} {
				return &tab_system_config{}
			},
		}

	Databases["accept"].Tabs["tab_affairs_agent"] = &sync.Pool{
			New : func() interface{} {
				return &tab_affairs_agent{}
			},
		}

	Databases["accept"].Tabs["tab_dal_deploy"] = &sync.Pool{
			New : func() interface{} {
				return &tab_dal_deploy{}
			},
		}

	Databases["accept"].Tabs["tab_holiday_adjusted"] = &sync.Pool{
			New : func() interface{} {
				return &tab_holiday_adjusted{}
			},
		}

	Databases["accept"].Tabs["tab_legal_person_info"] = &sync.Pool{
			New : func() interface{} {
				return &tab_legal_person_info{}
			},
		}

	Databases["accept"].Tabs["tab_material_objects"] = &sync.Pool{
			New : func() interface{} {
				return &tab_material_objects{}
			},
		}
	Databases["accept"].Tabs["tab_affairs_proposers"] = &sync.Pool{
			New : func() interface{} {
				return &tab_affairs_proposers{}
			},
		}
	Databases["accept"].Tabs["dic_material_upload_type"] = &sync.Pool{
			New : func() interface{} {
				return &dic_material_upload_type{}
			},
		}
	Databases["accept"].Tabs["dic_reuse_type"] = &sync.Pool{
			New : func() interface{} {
				return &dic_reuse_type{}
			},
		}
	Databases["accept"].Tabs["dic_material_type"] = &sync.Pool{
			New : func() interface{} {
				return &dic_material_type{}
			},
		}
	Databases["accept"].Tabs["tab_affair_material"] = &sync.Pool{
			New : func() interface{} {
				return &tab_affair_material{}
			},
		}

	Databases["accept"].Tabs["tab_affairs_info_outcome"] = &sync.Pool{
			New : func() interface{} {
				return &tab_affairs_info_outcome{}
			},
		}

	Databases["accept"].Tabs["tab_authorization"] = &sync.Pool{
			New : func() interface{} {
				return &tab_authorization{}
			},
		}


	Databases["accept"].Tabs["tab_credential_field"] = &sync.Pool{
			New : func() interface{} {
				return &tab_credential_field{}
			},
		}
	Databases["accept"].Tabs["tab_credential_detail"] = &sync.Pool{
			New : func() interface{} {
				return &tab_credential_detail{}
			},
		}

	Databases["accept"].Tabs["tab_monitors_extend_log"] = &sync.Pool{
			New : func() interface{} {
				return &tab_monitors_extend_log{}
			},
		}


	Databases["accept"].Tabs["tab_affair_tags"] = &sync.Pool{
			New : func() interface{} {
				return &tab_affair_tags{}
			},
		}



	Databases["accept"].Tabs["tab_credential_catalog"] = &sync.Pool{
			New : func() interface{} {
				return &tab_credential_catalog{}
			},
		}


	Databases["accept"].Tabs["tab_affairs_node_current"] = &sync.Pool{
			New : func() interface{} {
				return &tab_affairs_node_current{}
			},
		}

	Databases["accept"].Tabs["tab_autoinput"] = &sync.Pool{
			New : func() interface{} {
				return &tab_autoinput{}
			},
		}

	Databases["accept"].Tabs["tab_client_version"] = &sync.Pool{
			New : func() interface{} {
				return &tab_client_version{}
			},
		}


	Databases["accept"].Tabs["dic_node_status"] = &sync.Pool{
			New : func() interface{} {
				return &dic_node_status{}
			},
		}
	Databases["accept"].Tabs["tab_affairs_node_status"] = &sync.Pool{
			New : func() interface{} {
				return &tab_affairs_node_status{}
			},
		}
	Databases["accept"].Tabs["tab_materials_check"] = &sync.Pool{
			New : func() interface{} {
				return &tab_materials_check{}
			},
		}

	Databases["accept"].Tabs["tab_affair_material_agent"] = &sync.Pool{
			New : func() interface{} {
				return &tab_affair_material_agent{}
			},
		}


	Databases["accept"].Tabs["tab_affairs_result"] = &sync.Pool{
			New : func() interface{} {
				return &tab_affairs_result{}
			},
		}

	Databases["accept"].Tabs["tab_client_machine"] = &sync.Pool{
			New : func() interface{} {
				return &tab_client_machine{}
			},
		}

	Databases["accept"].Tabs["tab_enterprise_reg_info"] = &sync.Pool{
			New : func() interface{} {
				return &tab_enterprise_reg_info{}
			},
		}


	Databases["accept"].Tabs["tab_coporation_change"] = &sync.Pool{
			New : func() interface{} {
				return &tab_coporation_change{}
			},
		}

	Databases["accept"].Tabs["tab_coporation_off"] = &sync.Pool{
			New : func() interface{} {
				return &tab_coporation_off{}
			},
		}


	Databases["accept"].Tabs["dic_material_info"] = &sync.Pool{
			New : func() interface{} {
				return &dic_material_info{}
			},
		}
	Databases["accept"].Tabs["tab_material_relation"] = &sync.Pool{
			New : func() interface{} {
				return &tab_material_relation{}
			},
		}

	Databases["accept"].Tabs["dic_credential_type"] = &sync.Pool{
			New : func() interface{} {
				return &dic_credential_type{}
			},
		}
	Databases["accept"].Tabs["tab_natural_person_info"] = &sync.Pool{
			New : func() interface{} {
				return &tab_natural_person_info{}
			},
		}

	Databases["accept"].Tabs["tab_policy_interpretation"] = &sync.Pool{
			New : func() interface{} {
				return &tab_policy_interpretation{}
			},
		}


	Databases["accept"].Tabs["tab_affairs_outcome"] = &sync.Pool{
			New : func() interface{} {
				return &tab_affairs_outcome{}
			},
		}

	Databases["accept"].Tabs["tab_dedicate_account"] = &sync.Pool{
			New : func() interface{} {
				return &tab_dedicate_account{}
			},
		}


	Databases["accept"].Tabs["tab_check_bill"] = &sync.Pool{
			New : func() interface{} {
				return &tab_check_bill{}
			},
		}


	Databases["accept"].Tabs["tab_dedicate_limit"] = &sync.Pool{
			New : func() interface{} {
				return &tab_dedicate_limit{}
			},
		}

	Databases["accept"].Tabs["tab_region_honors"] = &sync.Pool{
			New : func() interface{} {
				return &tab_region_honors{}
			},
		}

	Databases["accept"].Tabs["tab_document_info"] = &sync.Pool{
			New : func() interface{} {
				return &tab_document_info{}
			},
		}


	Databases["accept"].Tabs["tab_vpn_connect"] = &sync.Pool{
			New : func() interface{} {
				return &tab_vpn_connect{}
			},
		}

	Databases["accept"].Tabs["tab_affairs_node_results"] = &sync.Pool{
			New : func() interface{} {
				return &tab_affairs_node_results{}
			},
		}

	Databases["accept"].Tabs["tab_region_photos"] = &sync.Pool{
			New : func() interface{} {
				return &tab_region_photos{}
			},
		}

	Databases["accept"].Tabs["tab_scripts"] = &sync.Pool{
			New : func() interface{} {
				return &tab_scripts{}
			},
		}
    
	//
    //
    // 说明: tab_affairs_info及其子表
    //
	Databases["accept"].TabsPas["tab_materials_check"] = &sync.Pool{
			New : func() interface{} {
				return &tab_materials_check{}
			},
		}
	Databases["accept"].TabsPas["tab_affairs_info_outcome"] = &sync.Pool{
			New : func() interface{} {
				return &tab_affairs_info_outcome{}
			},
		}
	Databases["accept"].TabsPas["tab_affair_material_agent"] = &sync.Pool{
			New : func() interface{} {
				return &tab_affair_material_agent{}
			},
		}
	Databases["accept"].TabsPas["tab_affairs_node_results"] = &sync.Pool{
			New : func() interface{} {
				return &tab_affairs_node_results{}
			},
		}
	Databases["accept"].TabsPas["tab_affairs_node_current"] = &sync.Pool{
			New : func() interface{} {
				return &tab_affairs_node_current{}
			},
		}
	Databases["accept"].TabsPas["tab_affairs_proposers"] = &sync.Pool{
			New : func() interface{} {
				return &tab_affairs_proposers{}
			},
		}
	Databases["accept"].TabsPas["tab_affairs_agent"] = &sync.Pool{
			New : func() interface{} {
				return &tab_affairs_agent{}
			},
		}
	Databases["accept"].TabsPas["tab_affairs_result"] = &sync.Pool{
			New : func() interface{} {
				return &tab_affairs_result{}
			},
		}
	Databases["accept"].TabsPas["tab_affairs_mark_result"] = &sync.Pool{
			New : func() interface{} {
				return &tab_affairs_mark_result{}
			},
		}
	Databases["accept"].TabsPas["tab_affair_material"] = &sync.Pool{
			New : func() interface{} {
				return &tab_affair_material{}
			},
		}
	Databases["accept"].TabsPas["tab_affairs_attachment"] = &sync.Pool{
			New : func() interface{} {
				return &tab_affairs_attachment{}
			},
		}
	Databases["accept"].TabsPas["tab_affairs_node_status"] = &sync.Pool{
			New : func() interface{} {
				return &tab_affairs_node_status{}
			},
		}
	Databases["accept"].TabsPas["tab_affairs_info"] = &sync.Pool{
			New : func() interface{} {
				return &tab_affairs_info{}
			},
		}

 	//表的依赖关系 
	Databases["accept"].TabsDependency["tab_holiday_adjusted"]=[]string{}
	Databases["accept"].TabsDependency["tab_legal_person_info"]=[]string{}
	Databases["accept"].TabsDependency["tab_affairs_agent"]=[]string{"tab_affairs_info",}
	Databases["accept"].TabsDependency["tab_dal_deploy"]=[]string{}
	Databases["accept"].TabsDependency["tab_authorization"]=[]string{"tab_user","tab_group","tab_flow_node_info","tab_affairs",}
	Databases["accept"].TabsDependency["tab_group"]=[]string{}
	Databases["accept"].TabsDependency["tab_affair_material"]=[]string{"tab_affairs_proposers","dic_material_upload_type","dic_reuse_type","dic_material_type","tab_affairs_info",}
	Databases["accept"].TabsDependency["tab_affairs_info_outcome"]=[]string{"dic_result_receive","tab_affairs_info",}
	Databases["accept"].TabsDependency["tab_credential_detail"]=[]string{"tab_credential_field",}
	Databases["accept"].TabsDependency["tab_monitors_extend_log"]=[]string{"tab_monitors",}
	Databases["accept"].TabsDependency["tab_affairs_proposers"]=[]string{"dic_proposer_type","tab_material_objects","tab_affairs_info",}
	Databases["accept"].TabsDependency["tab_credential_catalog"]=[]string{}
	Databases["accept"].TabsDependency["tab_region"]=[]string{"dic_city","dic_district","dic_street","dic_community",}
	Databases["accept"].TabsDependency["tab_user"]=[]string{"dic_user_level","dic_user_level","dic_gender","dic_positions",}
	Databases["accept"].TabsDependency["tab_affair_tags"]=[]string{}
	Databases["accept"].TabsDependency["tab_affairs_mark_result"]=[]string{"tab_user","tab_flow_node_result_info","dic_node_type","dic_node_standard_code","tab_user","dic_administrative_level","tab_region","tab_affairs_info",}
	Databases["accept"].TabsDependency["tab_client_version"]=[]string{}
	Databases["accept"].TabsDependency["tab_affairs_node_current"]=[]string{"tab_flow_node_info","tab_user","tab_region","tab_group","tab_affairs_info",}
	Databases["accept"].TabsDependency["tab_autoinput"]=[]string{"tab_affairs",}
	Databases["accept"].TabsDependency["tab_affairs_result"]=[]string{"tab_affairs_info",}
	Databases["accept"].TabsDependency["tab_client_machine"]=[]string{"tab_region",}
	Databases["accept"].TabsDependency["tab_enterprise_reg_info"]=[]string{}
	Databases["accept"].TabsDependency["tab_flow_template"]=[]string{"tab_affairs",}
	Databases["accept"].TabsDependency["tab_materials_check"]=[]string{"tab_flow_node_info","tab_affairs_node_status","dic_node_status","tab_affairs_info",}
	Databases["accept"].TabsDependency["tab_affair_material_agent"]=[]string{"tab_affairs_agent","dic_material_upload_type","dic_material_type","tab_affairs_info",}
	Databases["accept"].TabsDependency["tab_affairs"]=[]string{"dic_affairs_type","dic_subject_nature","dic_proposer_type","dic_time_limit_type","dic_universal_range","dic_handling_form","dic_department","dic_affairs_subject","dic_corporate_subject","dic_affairs_objects","dic_life_events","dic_net_level","dic_result_receive",}
	Databases["accept"].TabsDependency["tab_coporation_off"]=[]string{}
	Databases["accept"].TabsDependency["tab_affairs_info"]=[]string{"tab_affairs","tab_user","tab_region","dic_administrative_level","dic_apply_from","dic_satisfy_level","tab_flow_template",}
	Databases["accept"].TabsDependency["tab_coporation_change"]=[]string{}
	Databases["accept"].TabsDependency["tab_natural_person_info"]=[]string{"dic_credential_type",}
	Databases["accept"].TabsDependency["tab_policy_interpretation"]=[]string{}
	Databases["accept"].TabsDependency["tab_flow_node_result_info"]=[]string{"tab_flow_node_info","tab_flow_node_info",}
	Databases["accept"].TabsDependency["tab_material_relation"]=[]string{"tab_material_objects","tab_affairs","dic_material_info","dic_reuse_type","dic_material_type",}
	Databases["accept"].TabsDependency["tab_dedicate_account"]=[]string{"tab_user",}
	Databases["accept"].TabsDependency["tab_material_objects"]=[]string{"tab_affairs","dic_time_limit_type",}
	Databases["accept"].TabsDependency["tab_affairs_node_status"]=[]string{"tab_flow_node_info","tab_flow_node_info","dic_node_status","tab_region","tab_affairs_info",}
	Databases["accept"].TabsDependency["tab_affairs_outcome"]=[]string{"tab_affairs","dic_material_info","tab_material_objects",}
	Databases["accept"].TabsDependency["tab_dedicate_limit"]=[]string{"tab_virtual_machines",}
	Databases["accept"].TabsDependency["tab_region_honors"]=[]string{"tab_region",}
	Databases["accept"].TabsDependency["tab_check_bill"]=[]string{}
	Databases["accept"].TabsDependency["tab_credential_field"]=[]string{}
	Databases["accept"].TabsDependency["tab_vpn_connect"]=[]string{"tab_virtual_machines",}
	Databases["accept"].TabsDependency["tab_document_info"]=[]string{}
	Databases["accept"].TabsDependency["tab_flow_node_info"]=[]string{"dic_node_type","dic_node_standard_code","tab_flow_template","dic_time_limit_type",}
	Databases["accept"].TabsDependency["tab_scripts"]=[]string{}
	Databases["accept"].TabsDependency["tab_affairs_node_results"]=[]string{"tab_flow_node_result_info","tab_affairs_info",}
	Databases["accept"].TabsDependency["tab_region_photos"]=[]string{"tab_region",}
	Databases["accept"].TabsDependency["tab_monitors"]=[]string{"dic_time_limit_type","tab_affairs_mark_result","tab_flow_node_info","dic_node_standard_code","tab_affairs_info",}
	Databases["accept"].TabsDependency["tab_region_events"]=[]string{"tab_region",}
	Databases["accept"].TabsDependency["tab_virtual_machines"]=[]string{}
	Databases["accept"].TabsDependency["tab_vm_accounts"]=[]string{"tab_virtual_machines",}
	Databases["accept"].TabsDependency["tab_group_user_relation"]=[]string{"tab_user","tab_group",}
	Databases["accept"].TabsDependency["tab_monitor_hook_log"]=[]string{"tab_monitors",}
	Databases["accept"].TabsDependency["tab_coporation_info"]=[]string{}
	Databases["accept"].TabsDependency["tab_system_config"]=[]string{}
	Databases["accept"].TabsDependency["tab_affairs_attachment"]=[]string{"tab_affairs_info",}
	Databases["accept"].TabsDependency["tab_affairs_delicate"]=[]string{"tab_affairs",}
	//字典表列表
	Databases["accept"].DictionariesRevert["字典-流程节点标准编码"]="dic_node_standard_code" 
	Databases["accept"].Dictionaries["dic_node_standard_code"]="字典-流程节点标准编码" 
	Databases["accept"].DictionariesRevert["字典-人生事件"]="dic_life_events" 
	Databases["accept"].Dictionaries["dic_life_events"]="字典-人生事件" 
	Databases["accept"].DictionariesRevert["字典-证照类型"]="dic_credential_type" 
	Databases["accept"].Dictionaries["dic_credential_type"]="字典-证照类型" 
	Databases["accept"].DictionariesRevert["字典-评价结果等级"]="dic_satisfy_level" 
	Databases["accept"].Dictionaries["dic_satisfy_level"]="字典-评价结果等级" 
	Databases["accept"].DictionariesRevert["字典-实施主体性质"]="dic_subject_nature" 
	Databases["accept"].Dictionaries["dic_subject_nature"]="字典-实施主体性质" 
	Databases["accept"].DictionariesRevert["字典-用户级别"]="dic_user_level" 
	Databases["accept"].Dictionaries["dic_user_level"]="字典-用户级别" 
	Databases["accept"].DictionariesRevert["字典-性别"]="dic_gender" 
	Databases["accept"].Dictionaries["dic_gender"]="字典-性别" 
	Databases["accept"].DictionariesRevert["字典-材料类型"]="dic_material_type" 
	Databases["accept"].Dictionaries["dic_material_type"]="字典-材料类型" 
	Databases["accept"].DictionariesRevert["字典-材料定义"]="dic_material_info" 
	Databases["accept"].Dictionaries["dic_material_info"]="字典-材料定义" 
	Databases["accept"].DictionariesRevert["字典-行政区划-社区"]="dic_community" 
	Databases["accept"].Dictionaries["dic_community"]="字典-行政区划-社区" 
	Databases["accept"].DictionariesRevert["字典-事项类型"]="dic_affairs_type" 
	Databases["accept"].Dictionaries["dic_affairs_type"]="字典-事项类型" 
	Databases["accept"].DictionariesRevert["字典-事项对象"]="dic_affairs_objects" 
	Databases["accept"].Dictionaries["dic_affairs_objects"]="字典-事项对象" 
	Databases["accept"].DictionariesRevert["字典-部门"]="dic_department" 
	Databases["accept"].Dictionaries["dic_department"]="字典-部门" 
	Databases["accept"].DictionariesRevert["字典-网上办理深度"]="dic_net_level" 
	Databases["accept"].Dictionaries["dic_net_level"]="字典-网上办理深度" 
	Databases["accept"].DictionariesRevert["字典-大厅岗位设置"]="dic_positions" 
	Databases["accept"].Dictionaries["dic_positions"]="字典-大厅岗位设置" 
	Databases["accept"].DictionariesRevert["字典-时限类型表"]="dic_time_limit_type" 
	Databases["accept"].Dictionaries["dic_time_limit_type"]="字典-时限类型表" 
	Databases["accept"].DictionariesRevert["字典-流程节点类型"]="dic_node_type" 
	Databases["accept"].Dictionaries["dic_node_type"]="字典-流程节点类型" 
	Databases["accept"].DictionariesRevert["字典-申请来源"]="dic_apply_from" 
	Databases["accept"].Dictionaries["dic_apply_from"]="字典-申请来源" 
	Databases["accept"].DictionariesRevert["字典-材料上传类型"]="dic_material_upload_type" 
	Databases["accept"].Dictionaries["dic_material_upload_type"]="字典-材料上传类型" 
	Databases["accept"].DictionariesRevert["字典-申请人类型"]="dic_proposer_type" 
	Databases["accept"].Dictionaries["dic_proposer_type"]="字典-申请人类型" 
	Databases["accept"].DictionariesRevert["字典-行政区划-区县"]="dic_district" 
	Databases["accept"].Dictionaries["dic_district"]="字典-行政区划-区县" 
	Databases["accept"].DictionariesRevert["字典-行政区划-街道"]="dic_street" 
	Databases["accept"].Dictionaries["dic_street"]="字典-行政区划-街道" 
	Databases["accept"].DictionariesRevert["字典-通办范围"]="dic_universal_range" 
	Databases["accept"].Dictionaries["dic_universal_range"]="字典-通办范围" 
	Databases["accept"].DictionariesRevert["字典-事项办理形式"]="dic_handling_form" 
	Databases["accept"].Dictionaries["dic_handling_form"]="字典-事项办理形式" 
	Databases["accept"].DictionariesRevert["字典-事项主题"]="dic_affairs_subject" 
	Databases["accept"].Dictionaries["dic_affairs_subject"]="字典-事项主题" 
	Databases["accept"].DictionariesRevert["字典-复用类型"]="dic_reuse_type" 
	Databases["accept"].Dictionaries["dic_reuse_type"]="字典-复用类型" 
	Databases["accept"].DictionariesRevert["字典-节点流转状态"]="dic_node_status" 
	Databases["accept"].Dictionaries["dic_node_status"]="字典-节点流转状态" 
	Databases["accept"].DictionariesRevert["字典-结果送达方式"]="dic_result_receive" 
	Databases["accept"].Dictionaries["dic_result_receive"]="字典-结果送达方式" 
	Databases["accept"].DictionariesRevert["字典-行政级别"]="dic_administrative_level" 
	Databases["accept"].Dictionaries["dic_administrative_level"]="字典-行政级别" 
	Databases["accept"].DictionariesRevert["字典-法人事项主题"]="dic_corporate_subject" 
	Databases["accept"].Dictionaries["dic_corporate_subject"]="字典-法人事项主题" 
	Databases["accept"].DictionariesRevert["字典-行政区划-地市"]="dic_city" 
	Databases["accept"].Dictionaries["dic_city"]="字典-行政区划-地市" 
	//本地数据表列表
	Databases["accept"].DataTablesRevert["用户表"]="tab_user" 
	Databases["accept"].DataTables["tab_user"]="用户表" 
	Databases["accept"].DataTablesRevert["事项受理-代理人"]="tab_affairs_agent" 
	Databases["accept"].DataTables["tab_affairs_agent"]="事项受理-代理人" 
	Databases["accept"].DataTablesRevert["事项受理-材料签入签出"]="tab_materials_check" 
	Databases["accept"].DataTables["tab_materials_check"]="事项受理-材料签入签出" 
	Databases["accept"].DataTablesRevert["WPF客户机信息"]="tab_client_machine" 
	Databases["accept"].DataTables["tab_client_machine"]="WPF客户机信息" 
	Databases["accept"].DataTablesRevert["自然人信息表"]="tab_natural_person_info" 
	Databases["accept"].DataTables["tab_natural_person_info"]="自然人信息表" 
	Databases["accept"].DataTablesRevert["脚本表"]="tab_scripts" 
	Databases["accept"].DataTables["tab_scripts"]="脚本表" 
	Databases["accept"].DataTablesRevert["流程模板-流程节点表"]="tab_flow_node_info" 
	Databases["accept"].DataTables["tab_flow_node_info"]="流程模板-流程节点表" 
	Databases["accept"].DataTablesRevert["证照详细信息"]="tab_credential_detail" 
	Databases["accept"].DataTables["tab_credential_detail"]="证照详细信息" 
	Databases["accept"].DataTablesRevert["证件信息表"]="tab_document_info" 
	Databases["accept"].DataTables["tab_document_info"]="证件信息表" 
	Databases["accept"].DataTablesRevert["事项受理-已审核结点结果"]="tab_affairs_node_results" 
	Databases["accept"].DataTables["tab_affairs_node_results"]="事项受理-已审核结点结果" 
	Databases["accept"].DataTablesRevert["事项定义表"]="tab_affairs" 
	Databases["accept"].DataTables["tab_affairs"]="事项定义表" 
	Databases["accept"].DataTablesRevert["事项标签"]="tab_affair_tags" 
	Databases["accept"].DataTables["tab_affair_tags"]="事项标签" 
	Databases["accept"].DataTablesRevert["事项定义-事项专线关系"]="tab_affairs_delicate" 
	Databases["accept"].DataTables["tab_affairs_delicate"]="事项定义-事项专线关系" 
	Databases["accept"].DataTablesRevert["法定代表人信息"]="tab_legal_person_info" 
	Databases["accept"].DataTables["tab_legal_person_info"]="法定代表人信息" 
	Databases["accept"].DataTablesRevert["事项受理-节点流转状态"]="tab_affairs_node_status" 
	Databases["accept"].DataTables["tab_affairs_node_status"]="事项受理-节点流转状态" 
	Databases["accept"].DataTablesRevert["DAL部署表"]="tab_dal_deploy" 
	Databases["accept"].DataTables["tab_dal_deploy"]="DAL部署表" 
	Databases["accept"].DataTablesRevert["事项受理-当前待审核结点"]="tab_affairs_node_current" 
	Databases["accept"].DataTables["tab_affairs_node_current"]="事项受理-当前待审核结点" 
	Databases["accept"].DataTablesRevert["事项对象结果物定义"]="tab_affairs_outcome" 
	Databases["accept"].DataTables["tab_affairs_outcome"]="事项对象结果物定义" 
	Databases["accept"].DataTablesRevert["查询机-所获荣誉"]="tab_region_honors" 
	Databases["accept"].DataTables["tab_region_honors"]="查询机-所获荣誉" 
	Databases["accept"].DataTablesRevert["除正常工作日（周一到周五）外，国家规定有变化的工作日和休息日的调整记录"]="tab_holiday_adjusted" 
	Databases["accept"].DataTables["tab_holiday_adjusted"]="除正常工作日（周一到周五）外，国家规定有变化的工作日和休息日的调整记录" 
	Databases["accept"].DataTablesRevert["WPF客户端版本"]="tab_client_version" 
	Databases["accept"].DataTables["tab_client_version"]="WPF客户端版本" 
	Databases["accept"].DataTablesRevert["材料迁入迁出批次表"]="tab_check_bill" 
	Databases["accept"].DataTables["tab_check_bill"]="材料迁入迁出批次表" 
	Databases["accept"].DataTablesRevert["专线设置-虚拟机登录帐号"]="tab_vm_accounts" 
	Databases["accept"].DataTables["tab_vm_accounts"]="专线设置-虚拟机登录帐号" 
	Databases["accept"].DataTablesRevert["事项受理-附件表"]="tab_affairs_attachment" 
	Databases["accept"].DataTables["tab_affairs_attachment"]="事项受理-附件表" 
	Databases["accept"].DataTablesRevert["材料关系表"]="tab_material_relation" 
	Databases["accept"].DataTables["tab_material_relation"]="材料关系表" 
	Databases["accept"].DataTablesRevert["专线设置-专线虚拟机限定"]="tab_dedicate_limit" 
	Databases["accept"].DataTables["tab_dedicate_limit"]="专线设置-专线虚拟机限定" 
	Databases["accept"].DataTablesRevert["行政监察挂起记录表"]="tab_monitor_hook_log" 
	Databases["accept"].DataTables["tab_monitor_hook_log"]="行政监察挂起记录表" 
	Databases["accept"].DataTablesRevert["查询机-大事记"]="tab_region_events" 
	Databases["accept"].DataTables["tab_region_events"]="查询机-大事记" 
	Databases["accept"].DataTablesRevert["证件照面字段定义"]="tab_credential_field" 
	Databases["accept"].DataTables["tab_credential_field"]="证件照面字段定义" 
	Databases["accept"].DataTablesRevert["企业登记信息"]="tab_enterprise_reg_info" 
	Databases["accept"].DataTables["tab_enterprise_reg_info"]="企业登记信息" 
	Databases["accept"].DataTablesRevert["查询机-机构图片"]="tab_region_photos" 
	Databases["accept"].DataTables["tab_region_photos"]="查询机-机构图片" 
	Databases["accept"].DataTablesRevert["行政监察补正延长记录"]="tab_monitors_extend_log" 
	Databases["accept"].DataTables["tab_monitors_extend_log"]="行政监察补正延长记录" 
	Databases["accept"].DataTablesRevert["法人信息表"]="tab_coporation_info" 
	Databases["accept"].DataTables["tab_coporation_info"]="法人信息表" 
	Databases["accept"].DataTablesRevert["事项受理-材料-包含结果物"]="tab_affair_material" 
	Databases["accept"].DataTables["tab_affair_material"]="事项受理-材料-包含结果物" 
	Databases["accept"].DataTablesRevert["组别用户关系表"]="tab_group_user_relation" 
	Databases["accept"].DataTables["tab_group_user_relation"]="组别用户关系表" 
	Databases["accept"].DataTablesRevert["流程模板-流程模版表"]="tab_flow_template" 
	Databases["accept"].DataTables["tab_flow_template"]="流程模板-流程模版表" 
	Databases["accept"].DataTablesRevert["流程模板-流程节点结果信息表"]="tab_flow_node_result_info" 
	Databases["accept"].DataTables["tab_flow_node_result_info"]="流程模板-流程节点结果信息表" 
	Databases["accept"].DataTablesRevert["事项受理-结果物"]="tab_affairs_info_outcome" 
	Databases["accept"].DataTables["tab_affairs_info_outcome"]="事项受理-结果物" 
	Databases["accept"].DataTablesRevert["法人注销信息"]="tab_coporation_off" 
	Databases["accept"].DataTables["tab_coporation_off"]="法人注销信息" 
	Databases["accept"].DataTablesRevert["专线设置-虚拟机列表"]="tab_virtual_machines" 
	Databases["accept"].DataTables["tab_virtual_machines"]="专线设置-虚拟机列表" 
	Databases["accept"].DataTablesRevert["事项受理"]="tab_affairs_info" 
	Databases["accept"].DataTables["tab_affairs_info"]="事项受理" 
	Databases["accept"].DataTablesRevert["材料对象表"]="tab_material_objects" 
	Databases["accept"].DataTables["tab_material_objects"]="材料对象表" 
	Databases["accept"].DataTablesRevert["自动填充表"]="tab_autoinput" 
	Databases["accept"].DataTables["tab_autoinput"]="自动填充表" 
	Databases["accept"].DataTablesRevert["事项受理-办结信息"]="tab_affairs_result" 
	Databases["accept"].DataTables["tab_affairs_result"]="事项受理-办结信息" 
	Databases["accept"].DataTablesRevert["归属地"]="tab_region" 
	Databases["accept"].DataTables["tab_region"]="归属地" 
	Databases["accept"].DataTablesRevert["事项受理-申请人"]="tab_affairs_proposers" 
	Databases["accept"].DataTables["tab_affairs_proposers"]="事项受理-申请人" 
	Databases["accept"].DataTablesRevert["证照目录信息"]="tab_credential_catalog" 
	Databases["accept"].DataTables["tab_credential_catalog"]="证照目录信息" 
	Databases["accept"].DataTablesRevert["政策解读"]="tab_policy_interpretation" 
	Databases["accept"].DataTables["tab_policy_interpretation"]="政策解读" 
	Databases["accept"].DataTablesRevert["专线设置-专线vpn连接"]="tab_vpn_connect" 
	Databases["accept"].DataTables["tab_vpn_connect"]="专线设置-专线vpn连接" 
	Databases["accept"].DataTablesRevert["事项受理-流程审核结果信息"]="tab_affairs_mark_result" 
	Databases["accept"].DataTables["tab_affairs_mark_result"]="事项受理-流程审核结果信息" 
	Databases["accept"].DataTablesRevert["分组表"]="tab_group" 
	Databases["accept"].DataTables["tab_group"]="分组表" 
	Databases["accept"].DataTablesRevert["系统配置"]="tab_system_config" 
	Databases["accept"].DataTables["tab_system_config"]="系统配置" 
	Databases["accept"].DataTablesRevert["权限表"]="tab_authorization" 
	Databases["accept"].DataTables["tab_authorization"]="权限表" 
	Databases["accept"].DataTablesRevert["事项受理-代理人材料-包含结果物"]="tab_affair_material_agent" 
	Databases["accept"].DataTables["tab_affair_material_agent"]="事项受理-代理人材料-包含结果物" 
	Databases["accept"].DataTablesRevert["法人变更信息"]="tab_coporation_change" 
	Databases["accept"].DataTables["tab_coporation_change"]="法人变更信息" 
	Databases["accept"].DataTablesRevert["专线设置-专线登录帐号"]="tab_dedicate_account" 
	Databases["accept"].DataTables["tab_dedicate_account"]="专线设置-专线登录帐号" 
	Databases["accept"].DataTablesRevert["行政监察表"]="tab_monitors" 
	Databases["accept"].DataTables["tab_monitors"]="行政监察表" 

 	//所有列的数据类型 
	Databases["accept"].ColumnTypes=map[string]reflect.Kind{"none_field":1 ,"accept_condition":24,"accept_department":24,"accept_file_code":24,"access_token":24,"account":24,"account_id":6,"account_name":24,"account_pwd":24,"address":24,"administrative_level":6,"affair_promise_type":1,"affair_type":1,"affairs_agent_id":6,"affairs_count":6,"affairs_deadline":25,"affairs_delicate_id":6,"affairs_end":25,"affairs_finish":25,"affairs_global_id":24,"affairs_guide":24,"affairs_id":6,"affairs_info_code":24,"affairs_info_id":6,"affairs_local_id":24,"affairs_material_id":6,"affairs_name":24,"affairs_object_id":6,"affairs_outcome_id":6,"affairs_proposers_id":6,"affairs_real_finish":25,"affairs_reset_time":25,"affairs_result_description":24,"affairs_result_file_url":24,"affairs_result_id":6,"affairs_simple_name":24,"affairs_start":25,"affairs_start_time":25,"affairs_subject_id":6,"affairs_type":6,"after_change":24,"age":6,"age_id":6,"agent_gender":6,"agent_id":6,"agent_name":24,"agent_type":6,"apply_from":6,"apply_from_no":24,"approval_date":25,"attachment_id":6,"attachment_name":24,"attachment_url":24,"attribution_id":6,"attribution_name":24,"authority_division":24,"authority_type":6,"authorization_id":6,"auto_notify":1,"autoinput_id":6,"before_change":24,"begin_date":25,"birtyday":25,"bithday":25,"bithday_id":25,"bitmap_content":23,"booking_perday":6,"booking_support":1,"browser_mode":1,"change_date":25,"change_issue":24,"change_record":24,"changes":24,"check_attribution_id":6,"check_attribution_name":24,"check_bill_id":6,"check_bill_no":24,"check_date":25,"check_id":6,"check_point":24,"check_point_url":24,"check_standard":24,"check_type":1,"check_user_id":6,"check_user_name":24,"chipid":6,"cid":6,"city_id":6,"client_version_no":13,"commit_count":6,"community_id":6,"confer_date":25,"confer_sector":24,"config_description":24,"config_id":6,"config_name":24,"config_value":24,"confirm_department":24,"consult_affair":1,"consult_phone":24,"consultation_path":24,"consultation_phone":24,"content":24,"coporation_change_id":6,"coporation_id":6,"coporation_off_id":6,"corporate_subject_id":6,"correct_extend":1,"correct_timing":1,"country_id":6,"create_date":25,"create_userid":6,"create_username":24,"credential_catalog_code":24,"credential_catalog_id":6,"credential_code":24,"credential_code_id":24,"credential_name":24,"credential_type":6,"credential_type_id":6,"credit_code":24,"current_id":6,"current_red_or_yellow":6,"current_red_status":6,"current_version":1,"current_yellow_status":6,"dal_district_name":24,"dal_id":6,"dal_key":24,"dal_name":24,"dal_position":24,"date_id":6,"dbip":24,"dbname":24,"dbport":6,"dbpwd":24,"dbtype":24,"dbuser":24,"deadline":6,"deadline_date":25,"debug":6,"dedicate_account":24,"dedicate_account_id":6,"dedicate_protect":1,"dedicate_pwd":24,"default_start":24,"delete_mark":1,"department_id":6,"detail_id":6,"did":6,"district_id":6,"docs":6,"document_code":24,"document_info_id":6,"document_type":6,"document_url":24,"documentl_code":24,"documentl_name":24,"enable_mark":1,"end_date":25,"enterprise_reg_id":6,"error_date":25,"error_times":6,"event_content":24,"event_date":25,"event_id":6,"exercise_content":24,"exercise_level":6,"express_bill":24,"express_bill_no":24,"express_company":24,"express_company_id":6,"express_support":1,"extend_id":6,"extend_time":6,"extended_limit_time":6,"fact_capital":6,"faqs":24,"fee_basis":24,"fee_standard":13,"fee_standard_text":24,"fee_tag":1,"field_id":6,"field_name":24,"field_type":6,"field_value":24,"finger_print":24,"finish_attribution_id":6,"finish_attribution_name":24,"first_login_ip":24,"first_login_time":25,"fixed_phone":24,"fixed_phone_id":24,"flow_chart":24,"flow_chart_url":24,"flow_description":24,"founded_date":25,"gathered":1,"gathered_mode":6,"gathered_number":6,"gathered_time":25,"gender":6,"gender_id":6,"grant_object":6,"group_id":6,"group_name":24,"handling_form":6,"history_red_or_yellow":6,"history_red_status":6,"history_yellow_status":6,"holder":24,"home_url":24,"honor_desc":24,"honor_id":6,"hook_comment":24,"hook_id":6,"hook_type":6,"idcard":24,"image_url":24,"implement_code":24,"implement_sector":24,"include_received_day":1,"inter_service":24,"interpretation_id":6,"is_affair":1,"is_consultation":1,"is_default_result":1,"is_department":1,"is_enabled":1,"is_finished":1,"is_flow_check":1,"is_hook":6,"is_online":1,"is_outcome":1,"is_public":1,"issue_date":25,"issuer":24,"last_compute_time":25,"last_hook_time":25,"last_login_ip":24,"last_login_time":25,"latitude":24,"law_accept_condistion":24,"law_finish_limit":6,"legal_person":24,"legal_person_card_no":24,"legal_person_id":6,"legal_person_name":24,"life_event_id":6,"limit_date":25,"limit_id":6,"limit_time":6,"live_address":24,"local_msg_port":6,"local_msg_server":24,"location_id":6,"log_level":6,"log_server":24,"longitude":24,"machine_id":6,"manual_node":1,"mark_commit":24,"mark_id":6,"mashine_identy":24,"material_doc_url":24,"material_file_id":24,"material_file_url":24,"material_group":6,"material_id":6,"material_info_id":6,"material_name":24,"material_needed":24,"material_provider_id":6,"material_region_id":6,"material_region_name":24,"material_relation_id":6,"material_relative_field":24,"material_retention":1,"material_reuse_type_id":6,"material_sample_url":24,"material_type_changable":1,"material_type_id":6,"material_upload_type_id":6,"material_valid_term":25,"materials_count":6,"message_template":24,"mobile":24,"modify_date":25,"modify_userid":6,"modify_username":24,"monitor_id":6,"msg_port":6,"msg_server":24,"must_collect":6,"name":24,"name_id":24,"natural_man_status":6,"natural_person_id":6,"nature_deadline":6,"nature_warnning_deadline":6,"need_collected":1,"need_scaned":1,"net_level":6,"next_attribution_admin_level":6,"next_attribution_area_id":6,"next_attribution_area_name":24,"next_attribution_id":6,"next_attribution_name":24,"next_node_id":6,"next_user_id":6,"next_user_name":24,"node_current":24,"node_current_name":24,"node_end_time":25,"node_id":6,"node_name":24,"node_result_id":6,"node_result_name":24,"node_results":24,"node_standard_code":6,"node_start_time":25,"node_type":6,"node_type_id":6,"notice_list_id":24,"number_limit":6,"object_id":6,"object_name":24,"off_date":25,"off_reason":24,"off_sector":24,"offset":6,"online_status":1,"online_support":1,"online_user_count":6,"op_user_id":6,"op_user_name":24,"operator_id":6,"operator_name":24,"operator_time":25,"org_code":24,"org_name":24,"org_status":6,"org_type":6,"original_check":1,"outcome_id":6,"outcome_name":24,"outcome_tag":1,"passwd":24,"password":24,"pay_online_support":1,"person_rights":24,"phone":24,"phone_id":24,"photo_url":24,"pickup_date":25,"pickup_man":24,"pid":6,"plugin_folder":24,"policy_basis":24,"port":6,"position":6,"position_id":6,"position_x":6,"position_y":6,"post_address_id":24,"post_date":25,"postcode":24,"postcode_id":24,"potrait":24,"potrait_url":24,"power_status":6,"power_update_type":6,"prepare_perday":6,"primary_proposer":1,"print_bill":24,"progress_result_query":24,"prohibitive_requirement":24,"project_code":24,"promise_finish_limit":6,"proposer_gender":6,"proposer_id":6,"proposer_name":24,"proposer_type":6,"queue_call_time":25,"queue_time":25,"red_card_status":6,"red_or_yellow":6,"reg_address":24,"reg_address_id":24,"reg_attribution_id":6,"reg_attribution_name":24,"reg_capital":6,"reg_date":24,"reg_sector":24,"region_brief":24,"region_history":24,"region_id":6,"region_name":24,"region_photo_id":6,"relation_file":24,"relation_id":6,"replication_error":24,"replication_status":1,"result_name":24,"result_outcome_code":24,"result_receive":6,"result_receive_id":6,"result_sample_url":24,"result_valid_time":6,"results_id":6,"return_failed":1,"return_success":1,"reuse_type_id":6,"running_system":6,"satisfy_level_id":6,"satisfy_result":24,"satisfy_time":25,"script_content":24,"script_id":6,"script_name":24,"separate_dbip":24,"separate_dbname":24,"separate_dbport":6,"separate_dbpwd":24,"separate_dbtype":24,"separate_dbuser":24,"service_time":24,"setup_basis":24,"sex":6,"signer_date":25,"signer_name":24,"signer_photo":24,"sort_code":6,"special_group_id":6,"special_region":6,"special_user_id":6,"specified_version_no":13,"standard_code_id":6,"status":6,"status_id":6,"storage_url":24,"street_id":6,"subject_nature":6,"supervise_phone":24,"supervision_path":24,"supervision_phone":24,"sync":6,"sync_data":6,"sync_error":24,"sync_online":1,"sync_status":6,"tag_id":6,"tag_name":24,"tag_type":24,"tel":24,"telephone":24,"template_description":24,"template_id":6,"template_name":24,"term_end":25,"term_start":25,"the_date":25,"time_limit":6,"time_limit_type":6,"time_limit_type_id":6,"time_limited":6,"time_limited_type":6,"timing_reset":1,"title":24,"tm":24,"transact_place":24,"transact_time":24,"transfer_days":6,"type_id":6,"unflow_check_status":6,"union_sector":24,"universal_range":6,"use_material_pool":1,"user_dept":24,"user_id":6,"user_level_id":6,"user_name":24,"user_role":6,"valid_period_end":25,"valid_period_start":25,"valid_term":6,"valid_term_measure":6,"valid_time_measure":6,"valided_date":25,"version":6,"version_code":24,"version_id":6,"version_no":13,"vm_address":24,"vm_code":24,"vm_description":24,"vm_id":6,"vm_name":24,"vm_special_address":24,"vpn_account":24,"vpn_excutable_path":24,"vpn_id":6,"vpn_pwd":24,"warnning_deadline":6,"warnning_deadline_date":25,"window_name":24,"window_no":6,"work_tag":1,"worker_no":24,"yellow_card_status":6}
}
//将数据库对象存于pool
func communityCreate()  {
	Databases["community"].Tabs=make(map[string]*sync.Pool)
	Databases["community"].TabsPas=make(map[string]*sync.Pool)
	Databases["community"].TabsDependency=make(map[string][]string)
	Databases["community"].Dictionaries=make(map[string]string)
	Databases["community"].DictionariesRevert=make(map[string]string)
	Databases["community"].DataTablesRevert=make(map[string]string)
	Databases["community"].DataTables=make(map[string]string)
	Databases["community"].ColumnTypes=make(map[string]reflect.Kind)

	Databases["community"].Tabs["dic_speciality"] = &sync.Pool{
			New : func() interface{} {
				return &dic_speciality{}
			},
		}

	Databases["community"].Tabs["dic_technology_duty"] = &sync.Pool{
			New : func() interface{} {
				return &dic_technology_duty{}
			},
		}

	Databases["community"].Tabs["dic_unit_property"] = &sync.Pool{
			New : func() interface{} {
				return &dic_unit_property{}
			},
		}

	Databases["community"].Tabs["dic_unit_trade"] = &sync.Pool{
			New : func() interface{} {
				return &dic_unit_trade{}
			},
		}

	Databases["community"].Tabs["dic_work"] = &sync.Pool{
			New : func() interface{} {
				return &dic_work{}
			},
		}

	Databases["community"].Tabs["dic_work_status"] = &sync.Pool{
			New : func() interface{} {
				return &dic_work_status{}
			},
		}

	Databases["community"].Tabs["tab_operation_log"] = &sync.Pool{
			New : func() interface{} {
				return &tab_operation_log{}
			},
		}

	Databases["community"].Tabs["dic_married_type"] = &sync.Pool{
			New : func() interface{} {
				return &dic_married_type{}
			},
		}

	Databases["community"].Tabs["dic_nation"] = &sync.Pool{
			New : func() interface{} {
				return &dic_nation{}
			},
		}

	Databases["community"].Tabs["dic_military_service_status"] = &sync.Pool{
			New : func() interface{} {
				return &dic_military_service_status{}
			},
		}

	Databases["community"].Tabs["dic_card_type"] = &sync.Pool{
			New : func() interface{} {
				return &dic_card_type{}
			},
		}

	Databases["community"].Tabs["dic_census_register_type"] = &sync.Pool{
			New : func() interface{} {
				return &dic_census_register_type{}
			},
		}

	Databases["community"].Tabs["dic_country"] = &sync.Pool{
			New : func() interface{} {
				return &dic_country{}
			},
		}

	Databases["community"].Tabs["dic_family_master_relation"] = &sync.Pool{
			New : func() interface{} {
				return &dic_family_master_relation{}
			},
		}

	Databases["community"].Tabs["dic_health_status"] = &sync.Pool{
			New : func() interface{} {
				return &dic_health_status{}
			},
		}

	Databases["community"].Tabs["dic_part_time_job"] = &sync.Pool{
			New : func() interface{} {
				return &dic_part_time_job{}
			},
		}

	Databases["community"].Tabs["dic_politics_status"] = &sync.Pool{
			New : func() interface{} {
				return &dic_politics_status{}
			},
		}

	Databases["community"].Tabs["dic_blood_type"] = &sync.Pool{
			New : func() interface{} {
				return &dic_blood_type{}
			},
		}

	Databases["community"].Tabs["tab_files"] = &sync.Pool{
			New : func() interface{} {
				return &tab_files{}
			},
		}

	Databases["community"].Tabs["dic_duty_level"] = &sync.Pool{
			New : func() interface{} {
				return &dic_duty_level{}
			},
		}

	Databases["community"].Tabs["dic_education"] = &sync.Pool{
			New : func() interface{} {
				return &dic_education{}
			},
		}

	Databases["community"].Tabs["dic_marriage_status"] = &sync.Pool{
			New : func() interface{} {
				return &dic_marriage_status{}
			},
		}

	Databases["community"].Tabs["dic_not_transfer_registered_cause"] = &sync.Pool{
			New : func() interface{} {
				return &dic_not_transfer_registered_cause{}
			},
		}

	Databases["community"].Tabs["dic_religion"] = &sync.Pool{
			New : func() interface{} {
				return &dic_religion{}
			},
		}

	Databases["community"].Tabs["tab_person"] = &sync.Pool{
			New : func() interface{} {
				return &tab_person{}
			},
		}

	Databases["community"].Tabs["dic_census_register_status"] = &sync.Pool{
			New : func() interface{} {
				return &dic_census_register_status{}
			},
		}

 	//表的依赖关系 
	Databases["community"].TabsDependency["dic_census_register_status"]=[]string{}
	Databases["community"].TabsDependency["dic_duty_level"]=[]string{}
	Databases["community"].TabsDependency["dic_education"]=[]string{}
	Databases["community"].TabsDependency["dic_marriage_status"]=[]string{}
	Databases["community"].TabsDependency["dic_not_transfer_registered_cause"]=[]string{}
	Databases["community"].TabsDependency["dic_religion"]=[]string{}
	Databases["community"].TabsDependency["tab_person"]=[]string{}
	Databases["community"].TabsDependency["tab_operation_log"]=[]string{}
	Databases["community"].TabsDependency["dic_married_type"]=[]string{}
	Databases["community"].TabsDependency["dic_speciality"]=[]string{}
	Databases["community"].TabsDependency["dic_technology_duty"]=[]string{}
	Databases["community"].TabsDependency["dic_unit_property"]=[]string{}
	Databases["community"].TabsDependency["dic_unit_trade"]=[]string{}
	Databases["community"].TabsDependency["dic_work"]=[]string{}
	Databases["community"].TabsDependency["dic_work_status"]=[]string{}
	Databases["community"].TabsDependency["dic_military_service_status"]=[]string{}
	Databases["community"].TabsDependency["dic_nation"]=[]string{}
	Databases["community"].TabsDependency["dic_politics_status"]=[]string{}
	Databases["community"].TabsDependency["dic_blood_type"]=[]string{}
	Databases["community"].TabsDependency["dic_card_type"]=[]string{}
	Databases["community"].TabsDependency["dic_census_register_type"]=[]string{}
	Databases["community"].TabsDependency["dic_country"]=[]string{}
	Databases["community"].TabsDependency["dic_family_master_relation"]=[]string{}
	Databases["community"].TabsDependency["dic_health_status"]=[]string{}
	Databases["community"].TabsDependency["dic_part_time_job"]=[]string{}
	Databases["community"].TabsDependency["tab_files"]=[]string{}
	//字典表列表
	Databases["community"].DictionariesRevert["自治家园字典-学历"]="dic_education" 
	Databases["community"].Dictionaries["dic_education"]="自治家园字典-学历" 
	Databases["community"].DictionariesRevert["自治家园字典-民族"]="dic_nation" 
	Databases["community"].Dictionaries["dic_nation"]="自治家园字典-民族" 
	Databases["community"].DictionariesRevert["自治家园字典-户口类别"]="dic_census_register_type" 
	Databases["community"].Dictionaries["dic_census_register_type"]="自治家园字典-户口类别" 
	Databases["community"].DictionariesRevert["自治家园字典-未迁户口原因"]="dic_not_transfer_registered_cause" 
	Databases["community"].Dictionaries["dic_not_transfer_registered_cause"]="自治家园字典-未迁户口原因" 
	Databases["community"].DictionariesRevert["自治家园字典-户籍状态"]="dic_census_register_status" 
	Databases["community"].Dictionaries["dic_census_register_status"]="自治家园字典-户籍状态" 
	Databases["community"].DictionariesRevert["自治家园字典-单位性质"]="dic_unit_property" 
	Databases["community"].Dictionaries["dic_unit_property"]="自治家园字典-单位性质" 
	Databases["community"].DictionariesRevert["自治家园字典-单位所属行业"]="dic_unit_trade" 
	Databases["community"].Dictionaries["dic_unit_trade"]="自治家园字典-单位所属行业" 
	Databases["community"].DictionariesRevert["自治家园字典-已婚分类"]="dic_married_type" 
	Databases["community"].Dictionaries["dic_married_type"]="自治家园字典-已婚分类" 
	Databases["community"].DictionariesRevert["自治家园字典-兵役状况"]="dic_military_service_status" 
	Databases["community"].Dictionaries["dic_military_service_status"]="自治家园字典-兵役状况" 
	Databases["community"].DictionariesRevert["自治家园字典-国家"]="dic_country" 
	Databases["community"].Dictionaries["dic_country"]="自治家园字典-国家" 
	Databases["community"].DictionariesRevert["自治家园字典-与户主关系"]="dic_family_master_relation" 
	Databases["community"].Dictionaries["dic_family_master_relation"]="自治家园字典-与户主关系" 
	Databases["community"].DictionariesRevert["自治家园字典-社会兼职"]="dic_part_time_job" 
	Databases["community"].Dictionaries["dic_part_time_job"]="自治家园字典-社会兼职" 
	Databases["community"].DictionariesRevert["自治家园字典-专业技术职务"]="dic_technology_duty" 
	Databases["community"].Dictionaries["dic_technology_duty"]="自治家园字典-专业技术职务" 
	Databases["community"].DictionariesRevert["自治家园字典-从业状况"]="dic_work_status" 
	Databases["community"].Dictionaries["dic_work_status"]="自治家园字典-从业状况" 
	Databases["community"].DictionariesRevert["自治家园字典-干部职务级别"]="dic_duty_level" 
	Databases["community"].Dictionaries["dic_duty_level"]="自治家园字典-干部职务级别" 
	Databases["community"].DictionariesRevert["自治家园字典-宗教信仰"]="dic_religion" 
	Databases["community"].Dictionaries["dic_religion"]="自治家园字典-宗教信仰" 
	Databases["community"].DictionariesRevert["自治家园字典-政治面貌"]="dic_politics_status" 
	Databases["community"].Dictionaries["dic_politics_status"]="自治家园字典-政治面貌" 
	Databases["community"].DictionariesRevert["自治家园字典-血型"]="dic_blood_type" 
	Databases["community"].Dictionaries["dic_blood_type"]="自治家园字典-血型" 
	Databases["community"].DictionariesRevert["自治家园字典-证件类型"]="dic_card_type" 
	Databases["community"].Dictionaries["dic_card_type"]="自治家园字典-证件类型" 
	Databases["community"].DictionariesRevert["自治家园字典-健康状况"]="dic_health_status" 
	Databases["community"].Dictionaries["dic_health_status"]="自治家园字典-健康状况" 
	Databases["community"].DictionariesRevert["自治家园字典-婚姻状况"]="dic_marriage_status" 
	Databases["community"].Dictionaries["dic_marriage_status"]="自治家园字典-婚姻状况" 
	Databases["community"].DictionariesRevert["自治家园字典-特长"]="dic_speciality" 
	Databases["community"].Dictionaries["dic_speciality"]="自治家园字典-特长" 
	Databases["community"].DictionariesRevert["自治家园字典-职业"]="dic_work" 
	Databases["community"].Dictionaries["dic_work"]="自治家园字典-职业" 
	//本地数据表列表
	Databases["community"].DataTablesRevert["自治家园-文件表"]="tab_files" 
	Databases["community"].DataTables["tab_files"]="自治家园-文件表" 
	Databases["community"].DataTablesRevert["自治家园-居民信息操作记录表"]="tab_operation_log" 
	Databases["community"].DataTables["tab_operation_log"]="自治家园-居民信息操作记录表" 
	Databases["community"].DataTablesRevert["自治家园-居民信表"]="tab_person" 
	Databases["community"].DataTables["tab_person"]="自治家园-居民信表" 

 	//所有列的数据类型 
	Databases["community"].ColumnTypes=map[string]reflect.Kind{"none_field":1 ,"age":6,"birthday":25,"birthplace":24,"blood_type_code":24,"blood_type_id":6,"blood_type_name":24,"blood_type_prompt":24,"card_code":24,"card_type_code":24,"card_type_id":6,"card_type_name":24,"card_type_prompt":24,"census_register_status_code":24,"census_register_status_id":6,"census_register_status_name":24,"census_register_status_prompt":24,"census_register_type_code":24,"census_register_type_id":6,"census_register_type_name":24,"census_register_type_prompt":24,"country_code":24,"country_id":6,"country_name":24,"country_prompt":24,"create_date":25,"create_userid":6,"create_username":24,"delete_mark":1,"duty_level_code":24,"duty_level_id":6,"duty_level_name":24,"duty_level_prompt":24,"dwell_region_id":6,"education_code":24,"education_id":6,"education_name":24,"education_prompt":24,"email":24,"enable_mark":1,"extend_id":6,"family_master_relation_code":24,"family_master_relation_id":6,"family_master_relation_name":24,"family_master_relation_prompt":24,"file_ext":24,"file_id":6,"file_kind":24,"file_path":24,"file_show_name":24,"file_type":6,"former_name":24,"health_status_code":24,"health_status_id":6,"health_status_name":24,"health_status_prompt":24,"join_work_date":25,"marriage_status_code":24,"marriage_status_id":6,"marriage_status_name":24,"marriage_status_prompt":24,"married_type_code":24,"married_type_id":6,"married_type_name":24,"married_type_prompt":24,"military_service_status_code":24,"military_service_status_id":6,"military_service_status_name":24,"military_service_status_prompt":24,"mobile_phone":24,"modify_date":25,"modify_userid":6,"modify_username":24,"name":24,"nation_code":24,"nation_id":6,"nation_name":24,"nation_prompt":24,"not_transfer_registered_cause_code":24,"not_transfer_registered_cause_id":6,"not_transfer_registered_cause_name":24,"not_transfer_registered_cause_prompt":24,"operation_log_id":6,"operation_mark":6,"part_time_job_code":24,"part_time_job_id":6,"part_time_job_name":24,"part_time_job_prompt":24,"person_id":6,"phone":24,"photo":24,"politics_status_code":24,"politics_status_id":6,"politics_status_name":24,"politics_status_prompt":24,"register_region_id":6,"religion_code":24,"religion_id":6,"religion_name":24,"religion_prompt":24,"reside_address":24,"sort_code":6,"speciality_code":24,"speciality_id":6,"speciality_name":24,"speciality_prompt":24,"stature":6,"tags":6,"technology_duty_code":24,"technology_duty_id":6,"technology_duty_name":24,"technology_duty_prompt":24,"unit_property_code":24,"unit_property_id":6,"unit_property_name":24,"unit_property_prompt":24,"unit_trade_code":24,"unit_trade_id":6,"unit_trade_name":24,"unit_trade_prompt":24,"version":6,"work_code":24,"work_id":6,"work_name":24,"work_prompt":24,"work_status_code":24,"work_status_id":6,"work_status_name":24,"work_status_prompt":24}
}
//将数据库对象存于pool
func outterCreate()  {
	Databases["outter"].Tabs=make(map[string]*sync.Pool)
	Databases["outter"].TabsPas=make(map[string]*sync.Pool)
	Databases["outter"].TabsDependency=make(map[string][]string)
	Databases["outter"].Dictionaries=make(map[string]string)
	Databases["outter"].DictionariesRevert=make(map[string]string)
	Databases["outter"].DataTablesRevert=make(map[string]string)
	Databases["outter"].DataTables=make(map[string]string)
	Databases["outter"].ColumnTypes=make(map[string]reflect.Kind)

	Databases["outter"].Tabs["tab_outer_user"] = &sync.Pool{
			New : func() interface{} {
				return &tab_outer_user{}
			},
		}
	Databases["outter"].Tabs["tab_prepare_affairs"] = &sync.Pool{
			New : func() interface{} {
				return &tab_prepare_affairs{}
			},
		}
	Databases["outter"].Tabs["tab_prepare_agent"] = &sync.Pool{
			New : func() interface{} {
				return &tab_prepare_agent{}
			},
		}

	Databases["outter"].Tabs["tab_prepare_materials"] = &sync.Pool{
			New : func() interface{} {
				return &tab_prepare_materials{}
			},
		}

	Databases["outter"].Tabs["tab_real_name_materials"] = &sync.Pool{
			New : func() interface{} {
				return &tab_real_name_materials{}
			},
		}

	Databases["outter"].Tabs["tab_region_schedule"] = &sync.Pool{
			New : func() interface{} {
				return &tab_region_schedule{}
			},
		}
	Databases["outter"].Tabs["tab_booking"] = &sync.Pool{
			New : func() interface{} {
				return &tab_booking{}
			},
		}

	Databases["outter"].Tabs["tab_outer_consultation"] = &sync.Pool{
			New : func() interface{} {
				return &tab_outer_consultation{}
			},
		}
	Databases["outter"].Tabs["tab_consultation_replies"] = &sync.Pool{
			New : func() interface{} {
				return &tab_consultation_replies{}
			},
		}


	Databases["outter"].Tabs["tab_prepare_proposer"] = &sync.Pool{
			New : func() interface{} {
				return &tab_prepare_proposer{}
			},
		}




 	//表的依赖关系 
	Databases["outter"].TabsDependency["tab_booking"]=[]string{"tab_outer_user","tab_region_schedule",}
	Databases["outter"].TabsDependency["tab_consultation_replies"]=[]string{"tab_outer_consultation",}
	Databases["outter"].TabsDependency["tab_prepare_agent"]=[]string{"tab_prepare_affairs",}
	Databases["outter"].TabsDependency["tab_prepare_materials"]=[]string{}
	Databases["outter"].TabsDependency["tab_real_name_materials"]=[]string{}
	Databases["outter"].TabsDependency["tab_outer_consultation"]=[]string{"tab_outer_user",}
	Databases["outter"].TabsDependency["tab_outer_user"]=[]string{}
	Databases["outter"].TabsDependency["tab_prepare_affairs"]=[]string{"tab_outer_user",}
	Databases["outter"].TabsDependency["tab_prepare_proposer"]=[]string{"tab_prepare_affairs",}
	Databases["outter"].TabsDependency["tab_region_schedule"]=[]string{}
	//字典表列表
	//本地数据表列表
	Databases["outter"].DataTablesRevert["外网-咨询回复表"]="tab_consultation_replies" 
	Databases["outter"].DataTables["tab_consultation_replies"]="外网-咨询回复表" 
	Databases["outter"].DataTablesRevert["外网-预审申请表"]="tab_prepare_affairs" 
	Databases["outter"].DataTables["tab_prepare_affairs"]="外网-预审申请表" 
	Databases["outter"].DataTablesRevert["外网-归属地预约时段表"]="tab_region_schedule" 
	Databases["outter"].DataTables["tab_region_schedule"]="外网-归属地预约时段表" 
	Databases["outter"].DataTablesRevert["外网-网上预约表"]="tab_booking" 
	Databases["outter"].DataTables["tab_booking"]="外网-网上预约表" 
	Databases["outter"].DataTablesRevert["外网-咨询表"]="tab_outer_consultation" 
	Databases["outter"].DataTables["tab_outer_consultation"]="外网-咨询表" 
	Databases["outter"].DataTablesRevert["外网-申请人表"]="tab_prepare_proposer" 
	Databases["outter"].DataTables["tab_prepare_proposer"]="外网-申请人表" 
	Databases["outter"].DataTablesRevert["外网-用户表"]="tab_outer_user" 
	Databases["outter"].DataTables["tab_outer_user"]="外网-用户表" 
	Databases["outter"].DataTablesRevert["外网-代理人表"]="tab_prepare_agent" 
	Databases["outter"].DataTables["tab_prepare_agent"]="外网-代理人表" 
	Databases["outter"].DataTablesRevert["外网-附件及材料表"]="tab_prepare_materials" 
	Databases["outter"].DataTables["tab_prepare_materials"]="外网-附件及材料表" 
	Databases["outter"].DataTablesRevert["外网-实名认证资料"]="tab_real_name_materials" 
	Databases["outter"].DataTables["tab_real_name_materials"]="外网-实名认证资料" 

 	//所有列的数据类型 
	Databases["outter"].ColumnTypes=map[string]reflect.Kind{"none_field":1 ,"accept_department":24,"affairs_id":6,"affairs_info_code":24,"affairs_info_id":6,"affairs_name":24,"affairs_object":24,"affairs_subject":24,"agent_id":6,"agent_name":24,"booking_date":24,"booking_id":6,"booking_seq":24,"check_mess":24,"check_result":1,"check_standard":24,"client_id":24,"consultation_id":6,"consultation_type":24,"content":24,"create_date":25,"create_userid":6,"create_username":24,"credit_code":24,"delete_mark":1,"email":24,"enable_mark":1,"finished_operator":6,"finished_tag":1,"finished_time":25,"finshed_window_no":6,"id_code":24,"if_pauzed":1,"is_admin_user":1,"is_nature_person":1,"is_public":1,"is_real_name":1,"last_login_time":25,"life_event":24,"location_id":6,"location_name":24,"material_info_id":6,"material_name":24,"material_needed":24,"materials_name":24,"materials_url":24,"mobile":24,"modify_date":25,"modify_userid":6,"modify_username":24,"name":24,"open_id_qq":24,"open_id_wx":24,"orgnize_name":24,"out_user_id":6,"passwd":24,"photo_url":24,"policy_basis":24,"prepare_affairs_id":6,"prepare_consultation_id":6,"prepare_materials_id":6,"proposer_id":6,"proposer_name":24,"pubnumbers":6,"real_materials_id":6,"region_id":6,"region_name":24,"reply_content":24,"reply_id":6,"reply_user_id":6,"reply_user_name":24,"schedule_duration":24,"schedule_id":6,"sex":6,"sort_code":6,"title":24,"user_id":6,"user_name":24,"version":6}
}
//将数据库对象存于pool
func synccenterCreate()  {
	Databases["synccenter"].Tabs=make(map[string]*sync.Pool)
	Databases["synccenter"].TabsPas=make(map[string]*sync.Pool)
	Databases["synccenter"].TabsDependency=make(map[string][]string)
	Databases["synccenter"].Dictionaries=make(map[string]string)
	Databases["synccenter"].DictionariesRevert=make(map[string]string)
	Databases["synccenter"].DataTablesRevert=make(map[string]string)
	Databases["synccenter"].DataTables=make(map[string]string)
	Databases["synccenter"].ColumnTypes=make(map[string]reflect.Kind)

	Databases["synccenter"].Tabs["tab_system_list"] = &sync.Pool{
			New : func() interface{} {
				return &tab_system_list{}
			},
		}
	Databases["synccenter"].Tabs["tab_application_user_list"] = &sync.Pool{
			New : func() interface{} {
				return &tab_application_user_list{}
			},
		}


	Databases["synccenter"].Tabs["dic_user_level"] = &sync.Pool{
			New : func() interface{} {
				return &dic_user_level{}
			},
		}
	Databases["synccenter"].Tabs["dic_gender"] = &sync.Pool{
			New : func() interface{} {
				return &dic_gender{}
			},
		}
	Databases["synccenter"].Tabs["dic_positions"] = &sync.Pool{
			New : func() interface{} {
				return &dic_positions{}
			},
		}
	Databases["synccenter"].Tabs["tab_user"] = &sync.Pool{
			New : func() interface{} {
				return &tab_user{}
			},
		}
	Databases["synccenter"].Tabs["tab_system_user_map"] = &sync.Pool{
			New : func() interface{} {
				return &tab_system_user_map{}
			},
		}

 	//表的依赖关系 
	Databases["synccenter"].TabsDependency["tab_application_user_list"]=[]string{"tab_system_list",}
	Databases["synccenter"].TabsDependency["tab_system_list"]=[]string{}
	Databases["synccenter"].TabsDependency["tab_system_user_map"]=[]string{"tab_system_list","tab_user",}
	//字典表列表
	Databases["synccenter"].DictionariesRevert["字典-用户级别"]="dic_user_level" 
	Databases["synccenter"].Dictionaries["dic_user_level"]="字典-用户级别" 
	Databases["synccenter"].DictionariesRevert["字典-性别"]="dic_gender" 
	Databases["synccenter"].Dictionaries["dic_gender"]="字典-性别" 
	Databases["synccenter"].DictionariesRevert["字典-大厅岗位设置"]="dic_positions" 
	Databases["synccenter"].Dictionaries["dic_positions"]="字典-大厅岗位设置" 
	//本地数据表列表
	Databases["synccenter"].DataTablesRevert["用户中心-系统列表"]="tab_system_list" 
	Databases["synccenter"].DataTables["tab_system_list"]="用户中心-系统列表" 
	Databases["synccenter"].DataTablesRevert["用户中心-应用系统用户列表"]="tab_application_user_list" 
	Databases["synccenter"].DataTables["tab_application_user_list"]="用户中心-应用系统用户列表" 
	Databases["synccenter"].DataTablesRevert["用户表"]="tab_user" 
	Databases["synccenter"].DataTables["tab_user"]="用户表" 
	Databases["synccenter"].DataTablesRevert["用户中心-系统用户对应表"]="tab_system_user_map" 
	Databases["synccenter"].DataTables["tab_system_user_map"]="用户中心-系统用户对应表" 

 	//所有列的数据类型 
	Databases["synccenter"].ColumnTypes=map[string]reflect.Kind{"none_field":1 ,"bind_tag":1,"create_date":25,"create_userid":6,"create_username":24,"delete_mark":1,"enable_mark":1,"list_id":6,"map_id":6,"modify_date":25,"modify_userid":6,"modify_username":24,"sort_code":6,"system_account":24,"system_code":24,"system_id":6,"system_name":24,"system_user_id":24,"user_id":6,"version":6}
}
//将数据库对象存于pool
func allCreate()  {
	Databases["all"].Tabs=make(map[string]*sync.Pool)
	Databases["all"].TabsPas=make(map[string]*sync.Pool)
	Databases["all"].TabsDependency=make(map[string][]string)
	Databases["all"].Dictionaries=make(map[string]string)
	Databases["all"].DictionariesRevert=make(map[string]string)
	Databases["all"].DataTablesRevert=make(map[string]string)
	Databases["all"].DataTables=make(map[string]string)
	Databases["all"].ColumnTypes=make(map[string]reflect.Kind)

	Databases["all"].Tabs["dic_affairs_type"] = &sync.Pool{
			New : func() interface{} {
				return &dic_affairs_type{}
			},
		}

	Databases["all"].Tabs["dic_time_limit_type"] = &sync.Pool{
			New : func() interface{} {
				return &dic_time_limit_type{}
			},
		}

	Databases["all"].Tabs["dic_user_level"] = &sync.Pool{
			New : func() interface{} {
				return &dic_user_level{}
			},
		}
	Databases["all"].Tabs["dic_gender"] = &sync.Pool{
			New : func() interface{} {
				return &dic_gender{}
			},
		}
	Databases["all"].Tabs["dic_positions"] = &sync.Pool{
			New : func() interface{} {
				return &dic_positions{}
			},
		}
	Databases["all"].Tabs["tab_user"] = &sync.Pool{
			New : func() interface{} {
				return &tab_user{}
			},
		}
	Databases["all"].Tabs["tab_group"] = &sync.Pool{
			New : func() interface{} {
				return &tab_group{}
			},
		}
	Databases["all"].Tabs["dic_node_type"] = &sync.Pool{
			New : func() interface{} {
				return &dic_node_type{}
			},
		}
	Databases["all"].Tabs["dic_node_standard_code"] = &sync.Pool{
			New : func() interface{} {
				return &dic_node_standard_code{}
			},
		}
	Databases["all"].Tabs["dic_subject_nature"] = &sync.Pool{
			New : func() interface{} {
				return &dic_subject_nature{}
			},
		}
	Databases["all"].Tabs["dic_proposer_type"] = &sync.Pool{
			New : func() interface{} {
				return &dic_proposer_type{}
			},
		}
	Databases["all"].Tabs["dic_universal_range"] = &sync.Pool{
			New : func() interface{} {
				return &dic_universal_range{}
			},
		}
	Databases["all"].Tabs["dic_handling_form"] = &sync.Pool{
			New : func() interface{} {
				return &dic_handling_form{}
			},
		}
	Databases["all"].Tabs["dic_department"] = &sync.Pool{
			New : func() interface{} {
				return &dic_department{}
			},
		}
	Databases["all"].Tabs["dic_affairs_subject"] = &sync.Pool{
			New : func() interface{} {
				return &dic_affairs_subject{}
			},
		}
	Databases["all"].Tabs["dic_corporate_subject"] = &sync.Pool{
			New : func() interface{} {
				return &dic_corporate_subject{}
			},
		}
	Databases["all"].Tabs["dic_affairs_objects"] = &sync.Pool{
			New : func() interface{} {
				return &dic_affairs_objects{}
			},
		}
	Databases["all"].Tabs["dic_life_events"] = &sync.Pool{
			New : func() interface{} {
				return &dic_life_events{}
			},
		}
	Databases["all"].Tabs["dic_net_level"] = &sync.Pool{
			New : func() interface{} {
				return &dic_net_level{}
			},
		}
	Databases["all"].Tabs["dic_result_receive"] = &sync.Pool{
			New : func() interface{} {
				return &dic_result_receive{}
			},
		}
	Databases["all"].Tabs["tab_affairs"] = &sync.Pool{
			New : func() interface{} {
				return &tab_affairs{}
			},
		}
	Databases["all"].Tabs["tab_flow_template"] = &sync.Pool{
			New : func() interface{} {
				return &tab_flow_template{}
			},
		}
	Databases["all"].Tabs["tab_flow_node_info"] = &sync.Pool{
			New : func() interface{} {
				return &tab_flow_node_info{}
			},
		}
	Databases["all"].Tabs["tab_authorization"] = &sync.Pool{
			New : func() interface{} {
				return &tab_authorization{}
			},
		}

	Databases["all"].Tabs["tab_outer_user"] = &sync.Pool{
			New : func() interface{} {
				return &tab_outer_user{}
			},
		}
	Databases["all"].Tabs["tab_prepare_affairs"] = &sync.Pool{
			New : func() interface{} {
				return &tab_prepare_affairs{}
			},
		}
	Databases["all"].Tabs["tab_prepare_agent"] = &sync.Pool{
			New : func() interface{} {
				return &tab_prepare_agent{}
			},
		}

	Databases["all"].Tabs["dic_node_status"] = &sync.Pool{
			New : func() interface{} {
				return &dic_node_status{}
			},
		}

	Databases["all"].Tabs["tab_affairs_delicate"] = &sync.Pool{
			New : func() interface{} {
				return &tab_affairs_delicate{}
			},
		}

	Databases["all"].Tabs["tab_autoinput"] = &sync.Pool{
			New : func() interface{} {
				return &tab_autoinput{}
			},
		}

	Databases["all"].Tabs["tab_document_info"] = &sync.Pool{
			New : func() interface{} {
				return &tab_document_info{}
			},
		}

	Databases["all"].Tabs["tab_real_name_materials"] = &sync.Pool{
			New : func() interface{} {
				return &tab_real_name_materials{}
			},
		}

	Databases["all"].Tabs["tab_material_objects"] = &sync.Pool{
			New : func() interface{} {
				return &tab_material_objects{}
			},
		}
	Databases["all"].Tabs["dic_city"] = &sync.Pool{
			New : func() interface{} {
				return &dic_city{}
			},
		}
	Databases["all"].Tabs["dic_district"] = &sync.Pool{
			New : func() interface{} {
				return &dic_district{}
			},
		}
	Databases["all"].Tabs["dic_street"] = &sync.Pool{
			New : func() interface{} {
				return &dic_street{}
			},
		}
	Databases["all"].Tabs["dic_community"] = &sync.Pool{
			New : func() interface{} {
				return &dic_community{}
			},
		}
	Databases["all"].Tabs["tab_region"] = &sync.Pool{
			New : func() interface{} {
				return &tab_region{}
			},
		}
	Databases["all"].Tabs["dic_administrative_level"] = &sync.Pool{
			New : func() interface{} {
				return &dic_administrative_level{}
			},
		}
	Databases["all"].Tabs["dic_apply_from"] = &sync.Pool{
			New : func() interface{} {
				return &dic_apply_from{}
			},
		}
	Databases["all"].Tabs["dic_satisfy_level"] = &sync.Pool{
			New : func() interface{} {
				return &dic_satisfy_level{}
			},
		}
	Databases["all"].Tabs["tab_affairs_info"] = &sync.Pool{
			New : func() interface{} {
				return &tab_affairs_info{}
			},
		}
	Databases["all"].Tabs["tab_affairs_proposers"] = &sync.Pool{
			New : func() interface{} {
				return &tab_affairs_proposers{}
			},
		}
	Databases["all"].Tabs["dic_material_upload_type"] = &sync.Pool{
			New : func() interface{} {
				return &dic_material_upload_type{}
			},
		}
	Databases["all"].Tabs["dic_reuse_type"] = &sync.Pool{
			New : func() interface{} {
				return &dic_reuse_type{}
			},
		}
	Databases["all"].Tabs["dic_material_type"] = &sync.Pool{
			New : func() interface{} {
				return &dic_material_type{}
			},
		}
	Databases["all"].Tabs["tab_affair_material"] = &sync.Pool{
			New : func() interface{} {
				return &tab_affair_material{}
			},
		}

	Databases["all"].Tabs["tab_affairs_node_status"] = &sync.Pool{
			New : func() interface{} {
				return &tab_affairs_node_status{}
			},
		}




	Databases["all"].Tabs["tab_affairs_attachment"] = &sync.Pool{
			New : func() interface{} {
				return &tab_affairs_attachment{}
			},
		}




	Databases["all"].Tabs["dic_marriage_status"] = &sync.Pool{
			New : func() interface{} {
				return &dic_marriage_status{}
			},
		}

	Databases["all"].Tabs["dic_military_service_status"] = &sync.Pool{
			New : func() interface{} {
				return &dic_military_service_status{}
			},
		}

	Databases["all"].Tabs["dic_part_time_job"] = &sync.Pool{
			New : func() interface{} {
				return &dic_part_time_job{}
			},
		}


	Databases["all"].Tabs["dic_family_master_relation"] = &sync.Pool{
			New : func() interface{} {
				return &dic_family_master_relation{}
			},
		}


	Databases["all"].Tabs["dic_card_type"] = &sync.Pool{
			New : func() interface{} {
				return &dic_card_type{}
			},
		}

	Databases["all"].Tabs["dic_duty_level"] = &sync.Pool{
			New : func() interface{} {
				return &dic_duty_level{}
			},
		}

	Databases["all"].Tabs["dic_unit_property"] = &sync.Pool{
			New : func() interface{} {
				return &dic_unit_property{}
			},
		}

	Databases["all"].Tabs["tab_virtual_machines"] = &sync.Pool{
			New : func() interface{} {
				return &tab_virtual_machines{}
			},
		}
	Databases["all"].Tabs["tab_vm_accounts"] = &sync.Pool{
			New : func() interface{} {
				return &tab_vm_accounts{}
			},
		}

	Databases["all"].Tabs["tab_vpn_connect"] = &sync.Pool{
			New : func() interface{} {
				return &tab_vpn_connect{}
			},
		}

	Databases["all"].Tabs["dic_health_status"] = &sync.Pool{
			New : func() interface{} {
				return &dic_health_status{}
			},
		}

	Databases["all"].Tabs["tab_coporation_change"] = &sync.Pool{
			New : func() interface{} {
				return &tab_coporation_change{}
			},
		}

	Databases["all"].Tabs["tab_flow_node_result_info"] = &sync.Pool{
			New : func() interface{} {
				return &tab_flow_node_result_info{}
			},
		}
	Databases["all"].Tabs["tab_affairs_mark_result"] = &sync.Pool{
			New : func() interface{} {
				return &tab_affairs_mark_result{}
			},
		}
	Databases["all"].Tabs["tab_monitors"] = &sync.Pool{
			New : func() interface{} {
				return &tab_monitors{}
			},
		}
	Databases["all"].Tabs["tab_monitors_extend_log"] = &sync.Pool{
			New : func() interface{} {
				return &tab_monitors_extend_log{}
			},
		}


	Databases["all"].Tabs["dic_politics_status"] = &sync.Pool{
			New : func() interface{} {
				return &dic_politics_status{}
			},
		}

	Databases["all"].Tabs["dic_unit_trade"] = &sync.Pool{
			New : func() interface{} {
				return &dic_unit_trade{}
			},
		}

	Databases["all"].Tabs["tab_credential_field"] = &sync.Pool{
			New : func() interface{} {
				return &tab_credential_field{}
			},
		}
	Databases["all"].Tabs["tab_credential_detail"] = &sync.Pool{
			New : func() interface{} {
				return &tab_credential_detail{}
			},
		}

	Databases["all"].Tabs["tab_holiday_adjusted"] = &sync.Pool{
			New : func() interface{} {
				return &tab_holiday_adjusted{}
			},
		}

	Databases["all"].Tabs["dic_blood_type"] = &sync.Pool{
			New : func() interface{} {
				return &dic_blood_type{}
			},
		}

	Databases["all"].Tabs["tab_scripts"] = &sync.Pool{
			New : func() interface{} {
				return &tab_scripts{}
			},
		}

	Databases["all"].Tabs["tab_system_config"] = &sync.Pool{
			New : func() interface{} {
				return &tab_system_config{}
			},
		}



	Databases["all"].Tabs["dic_technology_duty"] = &sync.Pool{
			New : func() interface{} {
				return &dic_technology_duty{}
			},
		}

	Databases["all"].Tabs["tab_system_list"] = &sync.Pool{
			New : func() interface{} {
				return &tab_system_list{}
			},
		}
	Databases["all"].Tabs["tab_application_user_list"] = &sync.Pool{
			New : func() interface{} {
				return &tab_application_user_list{}
			},
		}

	Databases["all"].Tabs["tab_outer_consultation"] = &sync.Pool{
			New : func() interface{} {
				return &tab_outer_consultation{}
			},
		}
	Databases["all"].Tabs["tab_consultation_replies"] = &sync.Pool{
			New : func() interface{} {
				return &tab_consultation_replies{}
			},
		}

	Databases["all"].Tabs["dic_credential_type"] = &sync.Pool{
			New : func() interface{} {
				return &dic_credential_type{}
			},
		}
	Databases["all"].Tabs["tab_natural_person_info"] = &sync.Pool{
			New : func() interface{} {
				return &tab_natural_person_info{}
			},
		}


	Databases["all"].Tabs["tab_affairs_agent"] = &sync.Pool{
			New : func() interface{} {
				return &tab_affairs_agent{}
			},
		}

	Databases["all"].Tabs["tab_dedicate_limit"] = &sync.Pool{
			New : func() interface{} {
				return &tab_dedicate_limit{}
			},
		}

	Databases["all"].Tabs["tab_enterprise_reg_info"] = &sync.Pool{
			New : func() interface{} {
				return &tab_enterprise_reg_info{}
			},
		}


	Databases["all"].Tabs["tab_system_user_map"] = &sync.Pool{
			New : func() interface{} {
				return &tab_system_user_map{}
			},
		}

	Databases["all"].Tabs["dic_nation"] = &sync.Pool{
			New : func() interface{} {
				return &dic_nation{}
			},
		}

	Databases["all"].Tabs["dic_not_transfer_registered_cause"] = &sync.Pool{
			New : func() interface{} {
				return &dic_not_transfer_registered_cause{}
			},
		}

	Databases["all"].Tabs["dic_material_info"] = &sync.Pool{
			New : func() interface{} {
				return &dic_material_info{}
			},
		}
	Databases["all"].Tabs["tab_affairs_outcome"] = &sync.Pool{
			New : func() interface{} {
				return &tab_affairs_outcome{}
			},
		}

	Databases["all"].Tabs["tab_affairs_result"] = &sync.Pool{
			New : func() interface{} {
				return &tab_affairs_result{}
			},
		}

	Databases["all"].Tabs["tab_operation_log"] = &sync.Pool{
			New : func() interface{} {
				return &tab_operation_log{}
			},
		}


	Databases["all"].Tabs["tab_monitor_hook_log"] = &sync.Pool{
			New : func() interface{} {
				return &tab_monitor_hook_log{}
			},
		}


	Databases["all"].Tabs["tab_affair_material_agent"] = &sync.Pool{
			New : func() interface{} {
				return &tab_affair_material_agent{}
			},
		}


	Databases["all"].Tabs["tab_client_version"] = &sync.Pool{
			New : func() interface{} {
				return &tab_client_version{}
			},
		}

	Databases["all"].Tabs["tab_coporation_info"] = &sync.Pool{
			New : func() interface{} {
				return &tab_coporation_info{}
			},
		}


	Databases["all"].Tabs["tab_materials_check"] = &sync.Pool{
			New : func() interface{} {
				return &tab_materials_check{}
			},
		}

	Databases["all"].Tabs["tab_region_events"] = &sync.Pool{
			New : func() interface{} {
				return &tab_region_events{}
			},
		}


	Databases["all"].Tabs["tab_affairs_node_current"] = &sync.Pool{
			New : func() interface{} {
				return &tab_affairs_node_current{}
			},
		}

	Databases["all"].Tabs["dic_country"] = &sync.Pool{
			New : func() interface{} {
				return &dic_country{}
			},
		}





	Databases["all"].Tabs["tab_affairs_node_results"] = &sync.Pool{
			New : func() interface{} {
				return &tab_affairs_node_results{}
			},
		}

	Databases["all"].Tabs["tab_material_relation"] = &sync.Pool{
			New : func() interface{} {
				return &tab_material_relation{}
			},
		}

	Databases["all"].Tabs["tab_prepare_materials"] = &sync.Pool{
			New : func() interface{} {
				return &tab_prepare_materials{}
			},
		}

	Databases["all"].Tabs["tab_check_bill"] = &sync.Pool{
			New : func() interface{} {
				return &tab_check_bill{}
			},
		}


	Databases["all"].Tabs["dic_connection_type"] = &sync.Pool{
			New : func() interface{} {
				return &dic_connection_type{}
			},
		}

	Databases["all"].Tabs["dic_education"] = &sync.Pool{
			New : func() interface{} {
				return &dic_education{}
			},
		}




	Databases["all"].Tabs["tab_affairs_info_outcome"] = &sync.Pool{
			New : func() interface{} {
				return &tab_affairs_info_outcome{}
			},
		}

	Databases["all"].Tabs["dic_married_type"] = &sync.Pool{
			New : func() interface{} {
				return &dic_married_type{}
			},
		}



	Databases["all"].Tabs["dic_religion"] = &sync.Pool{
			New : func() interface{} {
				return &dic_religion{}
			},
		}

	Databases["all"].Tabs["dic_work"] = &sync.Pool{
			New : func() interface{} {
				return &dic_work{}
			},
		}

	Databases["all"].Tabs["tab_dedicate_account"] = &sync.Pool{
			New : func() interface{} {
				return &tab_dedicate_account{}
			},
		}


	Databases["all"].Tabs["tab_prepare_proposer"] = &sync.Pool{
			New : func() interface{} {
				return &tab_prepare_proposer{}
			},
		}


	Databases["all"].Tabs["dic_census_register_status"] = &sync.Pool{
			New : func() interface{} {
				return &dic_census_register_status{}
			},
		}



	Databases["all"].Tabs["dic_speciality"] = &sync.Pool{
			New : func() interface{} {
				return &dic_speciality{}
			},
		}

	Databases["all"].Tabs["tab_coporation_off"] = &sync.Pool{
			New : func() interface{} {
				return &tab_coporation_off{}
			},
		}


	Databases["all"].Tabs["dic_census_register_type"] = &sync.Pool{
			New : func() interface{} {
				return &dic_census_register_type{}
			},
		}


	Databases["all"].Tabs["tab_person"] = &sync.Pool{
			New : func() interface{} {
				return &tab_person{}
			},
		}




	Databases["all"].Tabs["dic_work_status"] = &sync.Pool{
			New : func() interface{} {
				return &dic_work_status{}
			},
		}

	Databases["all"].Tabs["tab_group_user_relation"] = &sync.Pool{
			New : func() interface{} {
				return &tab_group_user_relation{}
			},
		}

	Databases["all"].Tabs["tab_policy_interpretation"] = &sync.Pool{
			New : func() interface{} {
				return &tab_policy_interpretation{}
			},
		}

	Databases["all"].Tabs["tab_region_schedule"] = &sync.Pool{
			New : func() interface{} {
				return &tab_region_schedule{}
			},
		}





	Databases["all"].Tabs["tab_client_machine"] = &sync.Pool{
			New : func() interface{} {
				return &tab_client_machine{}
			},
		}



	Databases["all"].Tabs["tab_affair_tags"] = &sync.Pool{
			New : func() interface{} {
				return &tab_affair_tags{}
			},
		}

	Databases["all"].Tabs["tab_booking"] = &sync.Pool{
			New : func() interface{} {
				return &tab_booking{}
			},
		}



	Databases["all"].Tabs["tab_region_photos"] = &sync.Pool{
			New : func() interface{} {
				return &tab_region_photos{}
			},
		}

	Databases["all"].Tabs["tab_credential_catalog"] = &sync.Pool{
			New : func() interface{} {
				return &tab_credential_catalog{}
			},
		}

	Databases["all"].Tabs["tab_files"] = &sync.Pool{
			New : func() interface{} {
				return &tab_files{}
			},
		}

	Databases["all"].Tabs["tab_region_honors"] = &sync.Pool{
			New : func() interface{} {
				return &tab_region_honors{}
			},
		}

	Databases["all"].Tabs["tab_dal_deploy"] = &sync.Pool{
			New : func() interface{} {
				return &tab_dal_deploy{}
			},
		}
	Databases["all"].Tabs["tab_application_db_config"] = &sync.Pool{
			New : func() interface{} {
				return &tab_application_db_config{}
			},
		}


	Databases["all"].Tabs["tab_legal_person_info"] = &sync.Pool{
			New : func() interface{} {
				return &tab_legal_person_info{}
			},
		}


 	//表的依赖关系 
	//字典表列表
	Databases["all"].DictionariesRevert["字典-行政区划-区县"]="dic_district" 
	Databases["all"].Dictionaries["dic_district"]="字典-行政区划-区县" 
	Databases["all"].DictionariesRevert["字典-法人事项主题"]="dic_corporate_subject" 
	Databases["all"].Dictionaries["dic_corporate_subject"]="字典-法人事项主题" 
	Databases["all"].DictionariesRevert["字典-行政区划-社区"]="dic_community" 
	Databases["all"].Dictionaries["dic_community"]="字典-行政区划-社区" 
	Databases["all"].DictionariesRevert["字典-评价结果等级"]="dic_satisfy_level" 
	Databases["all"].Dictionaries["dic_satisfy_level"]="字典-评价结果等级" 
	Databases["all"].DictionariesRevert["自治家园字典-兵役状况"]="dic_military_service_status" 
	Databases["all"].Dictionaries["dic_military_service_status"]="自治家园字典-兵役状况" 
	Databases["all"].DictionariesRevert["字典-事项主题"]="dic_affairs_subject" 
	Databases["all"].Dictionaries["dic_affairs_subject"]="字典-事项主题" 
	Databases["all"].DictionariesRevert["字典-大厅岗位设置"]="dic_positions" 
	Databases["all"].Dictionaries["dic_positions"]="字典-大厅岗位设置" 
	Databases["all"].DictionariesRevert["字典-性别"]="dic_gender" 
	Databases["all"].Dictionaries["dic_gender"]="字典-性别" 
	Databases["all"].DictionariesRevert["自治家园字典-证件类型"]="dic_card_type" 
	Databases["all"].Dictionaries["dic_card_type"]="自治家园字典-证件类型" 
	Databases["all"].DictionariesRevert["自治家园字典-健康状况"]="dic_health_status" 
	Databases["all"].Dictionaries["dic_health_status"]="自治家园字典-健康状况" 
	Databases["all"].DictionariesRevert["自治家园字典-未迁户口原因"]="dic_not_transfer_registered_cause" 
	Databases["all"].Dictionaries["dic_not_transfer_registered_cause"]="自治家园字典-未迁户口原因" 
	Databases["all"].DictionariesRevert["字典-事项类型"]="dic_affairs_type" 
	Databases["all"].Dictionaries["dic_affairs_type"]="字典-事项类型" 
	Databases["all"].DictionariesRevert["字典-通办范围"]="dic_universal_range" 
	Databases["all"].Dictionaries["dic_universal_range"]="字典-通办范围" 
	Databases["all"].DictionariesRevert["自治家园字典-从业状况"]="dic_work_status" 
	Databases["all"].Dictionaries["dic_work_status"]="自治家园字典-从业状况" 
	Databases["all"].DictionariesRevert["字典-实施主体性质"]="dic_subject_nature" 
	Databases["all"].Dictionaries["dic_subject_nature"]="字典-实施主体性质" 
	Databases["all"].DictionariesRevert["字典-行政区划-街道"]="dic_street" 
	Databases["all"].Dictionaries["dic_street"]="字典-行政区划-街道" 
	Databases["all"].DictionariesRevert["字典-复用类型"]="dic_reuse_type" 
	Databases["all"].Dictionaries["dic_reuse_type"]="字典-复用类型" 
	Databases["all"].DictionariesRevert["自治家园字典-专业技术职务"]="dic_technology_duty" 
	Databases["all"].Dictionaries["dic_technology_duty"]="自治家园字典-专业技术职务" 
	Databases["all"].DictionariesRevert["自治家园字典-婚姻状况"]="dic_marriage_status" 
	Databases["all"].Dictionaries["dic_marriage_status"]="自治家园字典-婚姻状况" 
	Databases["all"].DictionariesRevert["自治家园字典-宗教信仰"]="dic_religion" 
	Databases["all"].Dictionaries["dic_religion"]="自治家园字典-宗教信仰" 
	Databases["all"].DictionariesRevert["字典-节点流转状态"]="dic_node_status" 
	Databases["all"].Dictionaries["dic_node_status"]="字典-节点流转状态" 
	Databases["all"].DictionariesRevert["字典-流程节点标准编码"]="dic_node_standard_code" 
	Databases["all"].Dictionaries["dic_node_standard_code"]="字典-流程节点标准编码" 
	Databases["all"].DictionariesRevert["自治家园字典-学历"]="dic_education" 
	Databases["all"].Dictionaries["dic_education"]="自治家园字典-学历" 
	Databases["all"].DictionariesRevert["自治家园字典-已婚分类"]="dic_married_type" 
	Databases["all"].Dictionaries["dic_married_type"]="自治家园字典-已婚分类" 
	Databases["all"].DictionariesRevert["字典-流程节点类型"]="dic_node_type" 
	Databases["all"].Dictionaries["dic_node_type"]="字典-流程节点类型" 
	Databases["all"].DictionariesRevert["字典-事项对象"]="dic_affairs_objects" 
	Databases["all"].Dictionaries["dic_affairs_objects"]="字典-事项对象" 
	Databases["all"].DictionariesRevert["字典-证照类型"]="dic_credential_type" 
	Databases["all"].Dictionaries["dic_credential_type"]="字典-证照类型" 
	Databases["all"].DictionariesRevert["字典-用户级别"]="dic_user_level" 
	Databases["all"].Dictionaries["dic_user_level"]="字典-用户级别" 
	Databases["all"].DictionariesRevert["字典-行政区划-地市"]="dic_city" 
	Databases["all"].Dictionaries["dic_city"]="字典-行政区划-地市" 
	Databases["all"].DictionariesRevert["字典-申请来源"]="dic_apply_from" 
	Databases["all"].Dictionaries["dic_apply_from"]="字典-申请来源" 
	Databases["all"].DictionariesRevert["自治家园字典-血型"]="dic_blood_type" 
	Databases["all"].Dictionaries["dic_blood_type"]="自治家园字典-血型" 
	Databases["all"].DictionariesRevert["字典-结果送达方式"]="dic_result_receive" 
	Databases["all"].Dictionaries["dic_result_receive"]="字典-结果送达方式" 
	Databases["all"].DictionariesRevert["字典-部门"]="dic_department" 
	Databases["all"].Dictionaries["dic_department"]="字典-部门" 
	Databases["all"].DictionariesRevert["字典-材料类型"]="dic_material_type" 
	Databases["all"].Dictionaries["dic_material_type"]="字典-材料类型" 
	Databases["all"].DictionariesRevert["自治家园字典-国家"]="dic_country" 
	Databases["all"].Dictionaries["dic_country"]="自治家园字典-国家" 
	Databases["all"].DictionariesRevert["字典-人生事件"]="dic_life_events" 
	Databases["all"].Dictionaries["dic_life_events"]="字典-人生事件" 
	Databases["all"].DictionariesRevert["字典-网上办理深度"]="dic_net_level" 
	Databases["all"].Dictionaries["dic_net_level"]="字典-网上办理深度" 
	Databases["all"].DictionariesRevert["字典-行政级别"]="dic_administrative_level" 
	Databases["all"].Dictionaries["dic_administrative_level"]="字典-行政级别" 
	Databases["all"].DictionariesRevert["自治家园字典-户口类别"]="dic_census_register_type" 
	Databases["all"].Dictionaries["dic_census_register_type"]="自治家园字典-户口类别" 
	Databases["all"].DictionariesRevert["字典-事项办理形式"]="dic_handling_form" 
	Databases["all"].Dictionaries["dic_handling_form"]="字典-事项办理形式" 
	Databases["all"].DictionariesRevert["字典-时限类型表"]="dic_time_limit_type" 
	Databases["all"].Dictionaries["dic_time_limit_type"]="字典-时限类型表" 
	Databases["all"].DictionariesRevert["自治家园字典-单位所属行业"]="dic_unit_trade" 
	Databases["all"].Dictionaries["dic_unit_trade"]="自治家园字典-单位所属行业" 
	Databases["all"].DictionariesRevert["字典-专线连接方式"]="dic_connection_type" 
	Databases["all"].Dictionaries["dic_connection_type"]="字典-专线连接方式" 
	Databases["all"].DictionariesRevert["字典-材料上传类型"]="dic_material_upload_type" 
	Databases["all"].Dictionaries["dic_material_upload_type"]="字典-材料上传类型" 
	Databases["all"].DictionariesRevert["自治家园字典-民族"]="dic_nation" 
	Databases["all"].Dictionaries["dic_nation"]="自治家园字典-民族" 
	Databases["all"].DictionariesRevert["字典-材料定义"]="dic_material_info" 
	Databases["all"].Dictionaries["dic_material_info"]="字典-材料定义" 
	Databases["all"].DictionariesRevert["自治家园字典-户籍状态"]="dic_census_register_status" 
	Databases["all"].Dictionaries["dic_census_register_status"]="自治家园字典-户籍状态" 
	Databases["all"].DictionariesRevert["自治家园字典-政治面貌"]="dic_politics_status" 
	Databases["all"].Dictionaries["dic_politics_status"]="自治家园字典-政治面貌" 
	Databases["all"].DictionariesRevert["自治家园字典-职业"]="dic_work" 
	Databases["all"].Dictionaries["dic_work"]="自治家园字典-职业" 
	Databases["all"].DictionariesRevert["自治家园字典-干部职务级别"]="dic_duty_level" 
	Databases["all"].Dictionaries["dic_duty_level"]="自治家园字典-干部职务级别" 
	Databases["all"].DictionariesRevert["自治家园字典-社会兼职"]="dic_part_time_job" 
	Databases["all"].Dictionaries["dic_part_time_job"]="自治家园字典-社会兼职" 
	Databases["all"].DictionariesRevert["自治家园字典-与户主关系"]="dic_family_master_relation" 
	Databases["all"].Dictionaries["dic_family_master_relation"]="自治家园字典-与户主关系" 
	Databases["all"].DictionariesRevert["字典-申请人类型"]="dic_proposer_type" 
	Databases["all"].Dictionaries["dic_proposer_type"]="字典-申请人类型" 
	Databases["all"].DictionariesRevert["自治家园字典-特长"]="dic_speciality" 
	Databases["all"].Dictionaries["dic_speciality"]="自治家园字典-特长" 
	Databases["all"].DictionariesRevert["自治家园字典-单位性质"]="dic_unit_property" 
	Databases["all"].Dictionaries["dic_unit_property"]="自治家园字典-单位性质" 
	//本地数据表列表
	Databases["all"].DataTablesRevert["事项受理"]="tab_affairs_info" 
	Databases["all"].DataTables["tab_affairs_info"]="事项受理" 
	Databases["all"].DataTablesRevert["材料关系表"]="tab_material_relation" 
	Databases["all"].DataTables["tab_material_relation"]="材料关系表" 
	Databases["all"].DataTablesRevert["事项标签"]="tab_affair_tags" 
	Databases["all"].DataTables["tab_affair_tags"]="事项标签" 
	Databases["all"].DataTablesRevert["外网-实名认证资料"]="tab_real_name_materials" 
	Databases["all"].DataTables["tab_real_name_materials"]="外网-实名认证资料" 
	Databases["all"].DataTablesRevert["自然人信息表"]="tab_natural_person_info" 
	Databases["all"].DataTables["tab_natural_person_info"]="自然人信息表" 
	Databases["all"].DataTablesRevert["证件照面字段定义"]="tab_credential_field" 
	Databases["all"].DataTables["tab_credential_field"]="证件照面字段定义" 
	Databases["all"].DataTablesRevert["自动填充表"]="tab_autoinput" 
	Databases["all"].DataTables["tab_autoinput"]="自动填充表" 
	Databases["all"].DataTablesRevert["外网-预审申请表"]="tab_prepare_affairs" 
	Databases["all"].DataTables["tab_prepare_affairs"]="外网-预审申请表" 
	Databases["all"].DataTablesRevert["外网-咨询回复表"]="tab_consultation_replies" 
	Databases["all"].DataTables["tab_consultation_replies"]="外网-咨询回复表" 
	Databases["all"].DataTablesRevert["事项受理-材料签入签出"]="tab_materials_check" 
	Databases["all"].DataTables["tab_materials_check"]="事项受理-材料签入签出" 
	Databases["all"].DataTablesRevert["专线设置-虚拟机列表"]="tab_virtual_machines" 
	Databases["all"].DataTables["tab_virtual_machines"]="专线设置-虚拟机列表" 
	Databases["all"].DataTablesRevert["脚本表"]="tab_scripts" 
	Databases["all"].DataTables["tab_scripts"]="脚本表" 
	Databases["all"].DataTablesRevert["事项受理-当前待审核结点"]="tab_affairs_node_current" 
	Databases["all"].DataTables["tab_affairs_node_current"]="事项受理-当前待审核结点" 
	Databases["all"].DataTablesRevert["流程模板-流程节点结果信息表"]="tab_flow_node_result_info" 
	Databases["all"].DataTables["tab_flow_node_result_info"]="流程模板-流程节点结果信息表" 
	Databases["all"].DataTablesRevert["外网-用户表"]="tab_outer_user" 
	Databases["all"].DataTables["tab_outer_user"]="外网-用户表" 
	Databases["all"].DataTablesRevert["事项受理-代理人"]="tab_affairs_agent" 
	Databases["all"].DataTables["tab_affairs_agent"]="事项受理-代理人" 
	Databases["all"].DataTablesRevert["法人信息表"]="tab_coporation_info" 
	Databases["all"].DataTables["tab_coporation_info"]="法人信息表" 
	Databases["all"].DataTablesRevert["查询机-机构图片"]="tab_region_photos" 
	Databases["all"].DataTables["tab_region_photos"]="查询机-机构图片" 
	Databases["all"].DataTablesRevert["法定代表人信息"]="tab_legal_person_info" 
	Databases["all"].DataTables["tab_legal_person_info"]="法定代表人信息" 
	Databases["all"].DataTablesRevert["材料对象表"]="tab_material_objects" 
	Databases["all"].DataTables["tab_material_objects"]="材料对象表" 
	Databases["all"].DataTablesRevert["用户中心-系统用户对应表"]="tab_system_user_map" 
	Databases["all"].DataTables["tab_system_user_map"]="用户中心-系统用户对应表" 
	Databases["all"].DataTablesRevert["组别用户关系表"]="tab_group_user_relation" 
	Databases["all"].DataTables["tab_group_user_relation"]="组别用户关系表" 
	Databases["all"].DataTablesRevert["外网-归属地预约时段表"]="tab_region_schedule" 
	Databases["all"].DataTables["tab_region_schedule"]="外网-归属地预约时段表" 
	Databases["all"].DataTablesRevert["事项受理-申请人"]="tab_affairs_proposers" 
	Databases["all"].DataTables["tab_affairs_proposers"]="事项受理-申请人" 
	Databases["all"].DataTablesRevert["归属地"]="tab_region" 
	Databases["all"].DataTables["tab_region"]="归属地" 
	Databases["all"].DataTablesRevert["事项受理-结果物"]="tab_affairs_info_outcome" 
	Databases["all"].DataTables["tab_affairs_info_outcome"]="事项受理-结果物" 
	Databases["all"].DataTablesRevert["自治家园-文件表"]="tab_files" 
	Databases["all"].DataTables["tab_files"]="自治家园-文件表" 
	Databases["all"].DataTablesRevert["外网-代理人表"]="tab_prepare_agent" 
	Databases["all"].DataTables["tab_prepare_agent"]="外网-代理人表" 
	Databases["all"].DataTablesRevert["行政监察挂起记录表"]="tab_monitor_hook_log" 
	Databases["all"].DataTables["tab_monitor_hook_log"]="行政监察挂起记录表" 
	Databases["all"].DataTablesRevert["WPF客户端版本"]="tab_client_version" 
	Databases["all"].DataTables["tab_client_version"]="WPF客户端版本" 
	Databases["all"].DataTablesRevert["查询机-所获荣誉"]="tab_region_honors" 
	Databases["all"].DataTables["tab_region_honors"]="查询机-所获荣誉" 
	Databases["all"].DataTablesRevert["权限表"]="tab_authorization" 
	Databases["all"].DataTables["tab_authorization"]="权限表" 
	Databases["all"].DataTablesRevert["事项受理-节点流转状态"]="tab_affairs_node_status" 
	Databases["all"].DataTables["tab_affairs_node_status"]="事项受理-节点流转状态" 
	Databases["all"].DataTablesRevert["法人变更信息"]="tab_coporation_change" 
	Databases["all"].DataTables["tab_coporation_change"]="法人变更信息" 
	Databases["all"].DataTablesRevert["事项受理-代理人材料-包含结果物"]="tab_affair_material_agent" 
	Databases["all"].DataTables["tab_affair_material_agent"]="事项受理-代理人材料-包含结果物" 
	Databases["all"].DataTablesRevert["事项定义表"]="tab_affairs" 
	Databases["all"].DataTables["tab_affairs"]="事项定义表" 
	Databases["all"].DataTablesRevert["专线设置-虚拟机登录帐号"]="tab_vm_accounts" 
	Databases["all"].DataTables["tab_vm_accounts"]="专线设置-虚拟机登录帐号" 
	Databases["all"].DataTablesRevert["专线设置-专线vpn连接"]="tab_vpn_connect" 
	Databases["all"].DataTables["tab_vpn_connect"]="专线设置-专线vpn连接" 
	Databases["all"].DataTablesRevert["系统配置"]="tab_system_config" 
	Databases["all"].DataTables["tab_system_config"]="系统配置" 
	Databases["all"].DataTablesRevert["查询机-大事记"]="tab_region_events" 
	Databases["all"].DataTables["tab_region_events"]="查询机-大事记" 
	Databases["all"].DataTablesRevert["事项定义-事项专线关系"]="tab_affairs_delicate" 
	Databases["all"].DataTables["tab_affairs_delicate"]="事项定义-事项专线关系" 
	Databases["all"].DataTablesRevert["事项受理-附件表"]="tab_affairs_attachment" 
	Databases["all"].DataTables["tab_affairs_attachment"]="事项受理-附件表" 
	Databases["all"].DataTablesRevert["事项对象结果物定义"]="tab_affairs_outcome" 
	Databases["all"].DataTables["tab_affairs_outcome"]="事项对象结果物定义" 
	Databases["all"].DataTablesRevert["事项受理-已审核结点结果"]="tab_affairs_node_results" 
	Databases["all"].DataTables["tab_affairs_node_results"]="事项受理-已审核结点结果" 
	Databases["all"].DataTablesRevert["用户中心-系统列表"]="tab_system_list" 
	Databases["all"].DataTables["tab_system_list"]="用户中心-系统列表" 
	Databases["all"].DataTablesRevert["外网-咨询表"]="tab_outer_consultation" 
	Databases["all"].DataTables["tab_outer_consultation"]="外网-咨询表" 
	Databases["all"].DataTablesRevert["自治家园-居民信息操作记录表"]="tab_operation_log" 
	Databases["all"].DataTables["tab_operation_log"]="自治家园-居民信息操作记录表" 
	Databases["all"].DataTablesRevert["外网-附件及材料表"]="tab_prepare_materials" 
	Databases["all"].DataTables["tab_prepare_materials"]="外网-附件及材料表" 
	Databases["all"].DataTablesRevert["除正常工作日（周一到周五）外，国家规定有变化的工作日和休息日的调整记录"]="tab_holiday_adjusted" 
	Databases["all"].DataTables["tab_holiday_adjusted"]="除正常工作日（周一到周五）外，国家规定有变化的工作日和休息日的调整记录" 
	Databases["all"].DataTablesRevert["事项受理-办结信息"]="tab_affairs_result" 
	Databases["all"].DataTables["tab_affairs_result"]="事项受理-办结信息" 
	Databases["all"].DataTablesRevert["专线设置-专线登录帐号"]="tab_dedicate_account" 
	Databases["all"].DataTables["tab_dedicate_account"]="专线设置-专线登录帐号" 
	Databases["all"].DataTablesRevert["分组表"]="tab_group" 
	Databases["all"].DataTables["tab_group"]="分组表" 
	Databases["all"].DataTablesRevert["行政监察补正延长记录"]="tab_monitors_extend_log" 
	Databases["all"].DataTables["tab_monitors_extend_log"]="行政监察补正延长记录" 
	Databases["all"].DataTablesRevert["外网-申请人表"]="tab_prepare_proposer" 
	Databases["all"].DataTables["tab_prepare_proposer"]="外网-申请人表" 
	Databases["all"].DataTablesRevert["自治家园-居民信表"]="tab_person" 
	Databases["all"].DataTables["tab_person"]="自治家园-居民信表" 
	Databases["all"].DataTablesRevert["WPF客户机信息"]="tab_client_machine" 
	Databases["all"].DataTables["tab_client_machine"]="WPF客户机信息" 
	Databases["all"].DataTablesRevert["专线设置-专线虚拟机限定"]="tab_dedicate_limit" 
	Databases["all"].DataTables["tab_dedicate_limit"]="专线设置-专线虚拟机限定" 
	Databases["all"].DataTablesRevert["证照详细信息"]="tab_credential_detail" 
	Databases["all"].DataTables["tab_credential_detail"]="证照详细信息" 
	Databases["all"].DataTablesRevert["事项受理-流程审核结果信息"]="tab_affairs_mark_result" 
	Databases["all"].DataTables["tab_affairs_mark_result"]="事项受理-流程审核结果信息" 
	Databases["all"].DataTablesRevert["用户中心-应用系统用户列表"]="tab_application_user_list" 
	Databases["all"].DataTables["tab_application_user_list"]="用户中心-应用系统用户列表" 
	Databases["all"].DataTablesRevert["材料迁入迁出批次表"]="tab_check_bill" 
	Databases["all"].DataTables["tab_check_bill"]="材料迁入迁出批次表" 
	Databases["all"].DataTablesRevert["行政监察表"]="tab_monitors" 
	Databases["all"].DataTables["tab_monitors"]="行政监察表" 
	Databases["all"].DataTablesRevert["应用数据配置"]="tab_application_db_config" 
	Databases["all"].DataTables["tab_application_db_config"]="应用数据配置" 
	Databases["all"].DataTablesRevert["证件信息表"]="tab_document_info" 
	Databases["all"].DataTables["tab_document_info"]="证件信息表" 
	Databases["all"].DataTablesRevert["DAL部署表"]="tab_dal_deploy" 
	Databases["all"].DataTables["tab_dal_deploy"]="DAL部署表" 
	Databases["all"].DataTablesRevert["用户表"]="tab_user" 
	Databases["all"].DataTables["tab_user"]="用户表" 
	Databases["all"].DataTablesRevert["证照目录信息"]="tab_credential_catalog" 
	Databases["all"].DataTables["tab_credential_catalog"]="证照目录信息" 
	Databases["all"].DataTablesRevert["法人注销信息"]="tab_coporation_off" 
	Databases["all"].DataTables["tab_coporation_off"]="法人注销信息" 
	Databases["all"].DataTablesRevert["流程模板-流程节点表"]="tab_flow_node_info" 
	Databases["all"].DataTables["tab_flow_node_info"]="流程模板-流程节点表" 
	Databases["all"].DataTablesRevert["政策解读"]="tab_policy_interpretation" 
	Databases["all"].DataTables["tab_policy_interpretation"]="政策解读" 
	Databases["all"].DataTablesRevert["外网-网上预约表"]="tab_booking" 
	Databases["all"].DataTables["tab_booking"]="外网-网上预约表" 
	Databases["all"].DataTablesRevert["流程模板-流程模版表"]="tab_flow_template" 
	Databases["all"].DataTables["tab_flow_template"]="流程模板-流程模版表" 
	Databases["all"].DataTablesRevert["事项受理-材料-包含结果物"]="tab_affair_material" 
	Databases["all"].DataTables["tab_affair_material"]="事项受理-材料-包含结果物" 
	Databases["all"].DataTablesRevert["企业登记信息"]="tab_enterprise_reg_info" 
	Databases["all"].DataTables["tab_enterprise_reg_info"]="企业登记信息" 

 	//所有列的数据类型 
	Databases["all"].ColumnTypes=map[string]reflect.Kind{"none_field":1 ,"accept_condition":24,"accept_department":24,"accept_file_code":24,"access_token":24,"account":24,"account_id":6,"account_name":24,"account_pwd":24,"address":24,"administrative_level":6,"affair_promise_type":1,"affair_type":1,"affairs_agent_id":6,"affairs_count":6,"affairs_deadline":25,"affairs_delicate_id":6,"affairs_end":25,"affairs_finish":25,"affairs_global_id":24,"affairs_guide":24,"affairs_id":6,"affairs_info_code":24,"affairs_info_id":6,"affairs_local_id":24,"affairs_material_id":6,"affairs_name":24,"affairs_object":24,"affairs_object_code":24,"affairs_object_id":6,"affairs_object_name":24,"affairs_object_prompt":24,"affairs_outcome_id":6,"affairs_proposers_id":6,"affairs_real_finish":25,"affairs_reset_time":25,"affairs_result_description":24,"affairs_result_file_url":24,"affairs_result_id":6,"affairs_simple_name":24,"affairs_start":25,"affairs_start_time":25,"affairs_subject":24,"affairs_subject_code":24,"affairs_subject_id":6,"affairs_subject_name":24,"affairs_subject_prompt":24,"affairs_type":6,"affairs_type_code":24,"affairs_type_id":6,"affairs_type_name":24,"affairs_type_prompt":24,"after_change":24,"age":6,"age_id":6,"agent_gender":6,"agent_id":6,"agent_name":24,"agent_type":6,"app_code":24,"apply_from":6,"apply_from_code":24,"apply_from_id":6,"apply_from_name":24,"apply_from_no":24,"apply_from_prompt":24,"approval_date":25,"attachment_id":6,"attachment_name":24,"attachment_url":24,"attribution_id":6,"attribution_name":24,"authority_division":24,"authority_type":6,"authorization_id":6,"auto_notify":1,"autoincr":1,"autoinput_id":6,"before_change":24,"begin_date":25,"bind_tag":1,"birthday":25,"birthplace":24,"birtyday":25,"bithday":25,"bithday_id":25,"bitmap_content":23,"blood_type_code":24,"blood_type_id":6,"blood_type_name":24,"blood_type_prompt":24,"booking_date":24,"booking_id":6,"booking_perday":6,"booking_seq":24,"booking_support":1,"browser_mode":1,"card_code":24,"card_type_code":24,"card_type_id":6,"card_type_name":24,"card_type_prompt":24,"census_register_status_code":24,"census_register_status_id":6,"census_register_status_name":24,"census_register_status_prompt":24,"census_register_type_code":24,"census_register_type_id":6,"census_register_type_name":24,"census_register_type_prompt":24,"change_date":25,"change_issue":24,"change_record":24,"changes":24,"check_attribution_id":6,"check_attribution_name":24,"check_bill_id":6,"check_bill_no":24,"check_date":25,"check_id":6,"check_mess":24,"check_point":24,"check_point_url":24,"check_result":1,"check_standard":24,"check_type":1,"check_user_id":6,"check_user_name":24,"chipid":6,"cid":6,"city_code":24,"city_id":6,"city_name":24,"city_prompt":24,"client_id":24,"client_version_no":13,"commit_count":6,"community_code":24,"community_id":6,"community_name":24,"community_prompt":24,"community_type":6,"confer_date":25,"confer_sector":24,"config_description":24,"config_id":6,"config_name":24,"config_value":24,"confirm_department":24,"connection_type_code":24,"connection_type_id":6,"connection_type_name":24,"connection_type_prompt":24,"consult_affair":1,"consult_phone":24,"consultation_id":6,"consultation_path":24,"consultation_phone":24,"consultation_type":24,"content":24,"coporation_change_id":6,"coporation_id":6,"coporation_off_id":6,"corporate_subject_id":6,"correct_extend":1,"correct_timing":1,"country_code":24,"country_id":6,"country_name":24,"country_prompt":24,"create_date":25,"create_userid":6,"create_username":24,"credential_catalog_code":24,"credential_catalog_id":6,"credential_code":24,"credential_code_id":24,"credential_name":24,"credential_type":6,"credential_type_code":24,"credential_type_id":6,"credential_type_name":24,"credential_type_prompt":24,"credit_code":24,"current_id":6,"current_red_or_yellow":6,"current_red_status":6,"current_version":1,"current_yellow_status":6,"dal_district_name":24,"dal_id":6,"dal_key":24,"dal_name":24,"dal_position":24,"date_id":6,"dbip":24,"dbname":24,"dbport":6,"dbpwd":24,"dbtype":24,"dbuser":24,"deadline":6,"deadline_date":25,"debug":6,"dedicate_account":24,"dedicate_account_id":6,"dedicate_protect":1,"dedicate_pwd":24,"default_start":24,"delete_mark":1,"department_id":6,"department_name":24,"department_prompt":24,"detail_id":6,"did":6,"district_code":24,"district_id":6,"district_name":24,"district_prompt":24,"docs":6,"document_code":24,"document_info_id":6,"document_type":6,"document_url":24,"documentl_code":24,"documentl_name":24,"duty_level_code":24,"duty_level_id":6,"duty_level_name":24,"duty_level_prompt":24,"dwell_region_id":6,"education_code":24,"education_id":6,"education_name":24,"education_prompt":24,"email":24,"enable_mark":1,"end_date":25,"enterprise_reg_id":6,"error_date":25,"error_times":6,"event_content":24,"event_date":25,"event_id":6,"exercise_content":24,"exercise_level":6,"express_bill":24,"express_bill_no":24,"express_company":24,"express_company_id":6,"express_support":1,"extend_id":6,"extend_time":6,"extended_limit_time":6,"fact_capital":6,"family_master_relation_code":24,"family_master_relation_id":6,"family_master_relation_name":24,"family_master_relation_prompt":24,"faqs":24,"fee_basis":24,"fee_standard":13,"fee_standard_text":24,"fee_tag":1,"field_id":6,"field_name":24,"field_type":6,"field_value":24,"file_ext":24,"file_id":6,"file_kind":24,"file_path":24,"file_show_name":24,"file_type":6,"finger_print":24,"finish_attribution_id":6,"finish_attribution_name":24,"finished_operator":6,"finished_tag":1,"finished_time":25,"finshed_window_no":6,"first_login_ip":24,"first_login_time":25,"fixed_phone":24,"fixed_phone_id":24,"flow_chart":24,"flow_chart_url":24,"flow_description":24,"former_name":24,"founded_date":25,"gathered":1,"gathered_mode":6,"gathered_number":6,"gathered_time":25,"gender":6,"gender_id":6,"gender_name":24,"gender_prompt":24,"grant_object":6,"group_id":6,"group_name":24,"handling_form":6,"handling_form_code":24,"handling_form_id":6,"handling_form_name":24,"handling_form_promt":24,"health_status_code":24,"health_status_id":6,"health_status_name":24,"health_status_prompt":24,"history_red_or_yellow":6,"history_red_status":6,"history_yellow_status":6,"holder":24,"home_url":24,"honor_desc":24,"honor_id":6,"hook_comment":24,"hook_id":6,"hook_type":6,"icon_url":24,"id_code":24,"idcard":24,"if_pauzed":1,"image_url":24,"implement_code":24,"implement_sector":24,"include_received_day":1,"initialized":1,"inter_service":24,"interpretation_id":6,"is_admin_user":1,"is_affair":1,"is_consultation":1,"is_default_result":1,"is_department":1,"is_enabled":1,"is_finished":1,"is_flow_check":1,"is_hook":6,"is_nature_person":1,"is_online":1,"is_outcome":1,"is_public":1,"is_real_name":1,"issue_date":25,"issuer":24,"join_work_date":25,"last_compute_time":25,"last_hook_time":25,"last_login_ip":24,"last_login_time":25,"latitude":24,"law_accept_condistion":24,"law_finish_limit":6,"legal_person":24,"legal_person_card_no":24,"legal_person_id":6,"legal_person_name":24,"level_id":6,"level_name":24,"level_prompt":24,"life_event":24,"life_event_code":24,"life_event_id":6,"life_event_name":24,"life_event_prompt":24,"limit_date":25,"limit_id":6,"limit_time":6,"list_id":6,"live_address":24,"local_msg_port":6,"local_msg_server":24,"location_id":6,"location_name":24,"log_level":6,"log_server":24,"longitude":24,"machine_id":6,"manual_node":1,"map_id":6,"mark_commit":24,"mark_id":6,"marriage_status_code":24,"marriage_status_id":6,"marriage_status_name":24,"marriage_status_prompt":24,"married_type_code":24,"married_type_id":6,"married_type_name":24,"married_type_prompt":24,"mashine_identy":24,"material_alias":24,"material_code":24,"material_doc_url":24,"material_file_id":24,"material_file_url":24,"material_group":6,"material_id":6,"material_info_id":6,"material_name":24,"material_needed":24,"material_prompt":24,"material_provider_id":6,"material_region_id":6,"material_region_name":24,"material_relation_id":6,"material_relative_field":24,"material_retention":1,"material_reuse_type_id":6,"material_sample_url":24,"material_type_changable":1,"material_type_id":6,"material_type_name":24,"material_type_prompt":24,"material_upload_id":6,"material_upload_type_id":6,"material_upload_type_name":24,"material_upload_type_prompt":24,"material_valid_term":25,"materials_count":6,"materials_name":24,"materials_url":24,"message_template":24,"military_service_status_code":24,"military_service_status_id":6,"military_service_status_name":24,"military_service_status_prompt":24,"mobile":24,"mobile_phone":24,"modify_date":25,"modify_userid":6,"modify_username":24,"monitor_id":6,"msg_port":6,"msg_server":24,"must_collect":6,"name":24,"name_id":24,"nation_code":24,"nation_id":6,"nation_name":24,"nation_prompt":24,"natural_man_status":6,"natural_person_id":6,"nature_deadline":6,"nature_warnning_deadline":6,"need_collected":1,"need_scaned":1,"net_level":6,"net_level_code":24,"net_level_id":6,"net_level_name":24,"net_level_prompt":24,"next_attribution_admin_level":6,"next_attribution_area_id":6,"next_attribution_area_name":24,"next_attribution_id":6,"next_attribution_name":24,"next_node_id":6,"next_user_id":6,"next_user_name":24,"node_current":24,"node_current_name":24,"node_end_time":25,"node_id":6,"node_name":24,"node_result_id":6,"node_result_name":24,"node_results":24,"node_standard_code":6,"node_start_time":25,"node_status_code":24,"node_status_id":6,"node_status_name":24,"node_status_prompt":24,"node_type":6,"node_type_id":6,"node_type_name":24,"node_type_prompt":24,"not_transfer_registered_cause_code":24,"not_transfer_registered_cause_id":6,"not_transfer_registered_cause_name":24,"not_transfer_registered_cause_prompt":24,"notice_list_id":24,"number_limit":6,"object_id":6,"object_name":24,"off_date":25,"off_reason":24,"off_sector":24,"offset":6,"online_status":1,"online_support":1,"online_user_count":6,"op_user_id":6,"op_user_name":24,"open_id_qq":24,"open_id_wx":24,"operation_log_id":6,"operation_mark":6,"operator_id":6,"operator_name":24,"operator_time":25,"org_code":24,"org_name":24,"org_status":6,"org_type":6,"orgnize_name":24,"original_check":1,"out_user_id":6,"outcome_id":6,"outcome_name":24,"outcome_tag":1,"part_time_job_code":24,"part_time_job_id":6,"part_time_job_name":24,"part_time_job_prompt":24,"passwd":24,"password":24,"pay_online_support":1,"person_id":6,"person_rights":24,"phone":24,"phone_id":24,"photo":24,"photo_url":24,"pickup_date":25,"pickup_man":24,"pid":6,"plugin_folder":24,"policy_basis":24,"politics_status_code":24,"politics_status_id":6,"politics_status_name":24,"politics_status_prompt":24,"port":6,"position":6,"position_id":6,"position_name":24,"position_prompt":24,"position_x":6,"position_y":6,"post_address_id":24,"post_date":25,"postcode":24,"postcode_id":24,"potrait":24,"potrait_url":24,"power_status":6,"power_update_type":6,"prepare_affairs_id":6,"prepare_consultation_id":6,"prepare_materials_id":6,"prepare_perday":6,"primary_proposer":1,"print_bill":24,"progress_result_query":24,"prohibitive_requirement":24,"project_code":24,"promise_finish_limit":6,"proposer_gender":6,"proposer_id":6,"proposer_name":24,"proposer_type":6,"proposer_type_id":6,"proposer_type_prompt":24,"proposer_type_value":24,"pubnumbers":6,"queue_call_time":25,"queue_time":25,"real_materials_id":6,"red_card_status":6,"red_or_yellow":6,"reg_address":24,"reg_address_id":24,"reg_attribution_id":6,"reg_attribution_name":24,"reg_capital":6,"reg_date":24,"reg_sector":24,"region_brief":24,"region_history":24,"region_id":6,"region_name":24,"region_photo_id":6,"register_region_id":6,"relation_file":24,"relation_id":6,"religion_code":24,"religion_id":6,"religion_name":24,"religion_prompt":24,"replication_error":24,"replication_status":1,"reply_content":24,"reply_id":6,"reply_user_id":6,"reply_user_name":24,"reside_address":24,"result_name":24,"result_outcome_code":24,"result_receive":6,"result_receive_code":24,"result_receive_id":6,"result_receive_name":24,"result_receive_prompt":24,"result_sample_url":24,"result_valid_time":6,"results_id":6,"return_failed":1,"return_success":1,"reuse_type_id":6,"reuse_type_name":24,"reuse_type_promt":24,"running_system":6,"satisfy_level_code":24,"satisfy_level_id":6,"satisfy_level_name":24,"satisfy_level_prompt":24,"satisfy_result":24,"satisfy_time":25,"schedule_duration":24,"schedule_id":6,"script_content":24,"script_id":6,"script_name":24,"separate_dbip":24,"separate_dbname":24,"separate_dbport":6,"separate_dbpwd":24,"separate_dbtype":24,"separate_dbuser":24,"service_time":24,"setup_basis":24,"sex":6,"signer_date":25,"signer_name":24,"signer_photo":24,"sort_code":6,"special_group_id":6,"special_region":6,"special_user_id":6,"speciality_code":24,"speciality_id":6,"speciality_name":24,"speciality_prompt":24,"specified_version_no":13,"standard_code":24,"standard_code_id":6,"standard_code_name":24,"standard_code_prompt":24,"stature":6,"status":6,"status_id":6,"storage_url":24,"street_code":24,"street_id":6,"street_name":24,"street_prompt":24,"subject_nature":6,"subject_nature_code":24,"subject_nature_id":6,"subject_nature_name":24,"subject_nature_prompt":24,"supervise_phone":24,"supervision_path":24,"supervision_phone":24,"sync":6,"sync_data":6,"sync_error":24,"sync_online":1,"sync_status":6,"system_account":24,"system_code":24,"system_id":6,"system_name":24,"system_user_id":24,"tag_id":6,"tag_name":24,"tag_type":24,"tags":6,"technology_duty_code":24,"technology_duty_id":6,"technology_duty_name":24,"technology_duty_prompt":24,"tel":24,"telephone":24,"template_description":24,"template_id":6,"template_name":24,"term_end":25,"term_start":25,"the_date":25,"time_limit":6,"time_limit_type":6,"time_limit_type_id":6,"time_limit_type_name":24,"time_limit_type_prompt":24,"time_limited":6,"time_limited_type":6,"timing_reset":1,"title":24,"tm":24,"transact_place":24,"transact_time":24,"transfer_days":6,"type_id":6,"unflow_check_status":6,"union_sector":24,"unit_property_code":24,"unit_property_id":6,"unit_property_name":24,"unit_property_prompt":24,"unit_trade_code":24,"unit_trade_id":6,"unit_trade_name":24,"unit_trade_prompt":24,"universal_range":6,"universal_range_code":24,"universal_range_id":6,"universal_range_name":24,"universal_range_prompt":24,"use_material_pool":1,"user_dept":24,"user_id":6,"user_level_id":6,"user_level_name":24,"user_level_prompt":24,"user_name":24,"user_role":6,"valid_period_end":25,"valid_period_start":25,"valid_term":6,"valid_term_measure":6,"valid_time_measure":6,"valided_date":25,"value_code":24,"version":6,"version_code":24,"version_id":6,"version_no":13,"vm_address":24,"vm_code":24,"vm_description":24,"vm_id":6,"vm_name":24,"vm_special_address":24,"vpn_account":24,"vpn_excutable_path":24,"vpn_id":6,"vpn_pwd":24,"warnning_deadline":6,"warnning_deadline_date":25,"window_name":24,"window_no":6,"work_code":24,"work_id":6,"work_name":24,"work_prompt":24,"work_status_code":24,"work_status_id":6,"work_status_name":24,"work_status_prompt":24,"work_tag":1,"worker_no":24,"yellow_card_status":6}
}
//
// 表名: dic_street
// 说明: 字典-行政区划-街道
// 字典-行政区划-街道
// 
type dic_street struct {
	Street_id       *int64     `xorm:" pk not null 'street_id'" indexes:"" comment:"街道主键" json:"street_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"street_id,omitempty"`
	Street_name     *string    `xorm:" varchar(255) null 'street_name'" indexes:"" comment:"街道名称" json:"street_name,omitempty" bson:",omitempty" msgpack:"street_name,omitempty"`
	Street_prompt   *string    `xorm:" varchar(255) null 'street_prompt'" indexes:"" comment:"街道简称" json:"street_prompt,omitempty" bson:",omitempty" msgpack:"street_prompt,omitempty"`
	Street_code     *string    `xorm:" varchar(255) null 'street_code'" indexes:" unique(Index_1) " comment:"街道代码" json:"street_code,omitempty" bson:",omitempty" msgpack:"street_code,omitempty"`
	District_id     *int64     `xorm:" null 'district_id'" indexes:"" comment:"区县主键" json:"district_id,omitempty" bson:",omitempty" msgpack:"district_id,omitempty" fk:"dic_district(district_id)"`
	Enable_mark     *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_2) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark     *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_3) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code       *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description     *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date     *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid   *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date     *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid   *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version         *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: tab_affairs_proposers
// 说明: 事项受理-申请人
// 事项受理-申请人
// 
type tab_affairs_proposers struct {
	Affairs_proposers_id *int64     `xorm:" pk not null 'affairs_proposers_id'" indexes:"" comment:"事项申请人主键" json:"affairs_proposers_id,omitempty" bson:",omitempty" msgpack:"affairs_proposers_id,omitempty"`
	Proposer_id          *int64     `xorm:" null 'proposer_id'" indexes:"" comment:"申请人主键" json:"proposer_id,omitempty" bson:",omitempty" msgpack:"proposer_id,omitempty"`
	Proposer_name        *string    `xorm:" varchar(255) null 'proposer_name'" indexes:"" comment:"申请人名称" json:"proposer_name,omitempty" bson:",omitempty" msgpack:"proposer_name,omitempty"`
	Proposer_gender      *int64     `xorm:" null 'proposer_gender'" indexes:"" comment:"申请人性别" json:"proposer_gender,omitempty" bson:",omitempty" msgpack:"proposer_gender,omitempty"`
	Proposer_type        *int64     `xorm:" not null 'proposer_type'" indexes:"" comment:"申请人类型" json:"proposer_type,omitempty" bson:",omitempty" msgpack:"proposer_type,omitempty" fk:"dic_proposer_type(proposer_type_id)"`
	Object_id            *int64     `xorm:" null 'object_id'" indexes:"" comment:"对象主键" json:"object_id,omitempty" bson:",omitempty" msgpack:"object_id,omitempty" fk:"tab_material_objects(object_id)"`
	Affairs_info_id      *int64     `xorm:" null 'affairs_info_id'" indexes:"" comment:"受理主键" json:"affairs_info_id,omitempty" bson:",omitempty" msgpack:"affairs_info_id,omitempty" fk:"tab_affairs_info(affairs_info_id)"`
	Credential_type      *int64     `xorm:" null 'credential_type'" indexes:"" comment:"申请人证件类型" json:"credential_type,omitempty" bson:",omitempty" msgpack:"credential_type,omitempty"`
	Credential_code      *string    `xorm:" varchar(50) null 'credential_code'" indexes:" index(Index_1) " comment:"申请人证件号码" json:"credential_code,omitempty" bson:",omitempty" msgpack:"credential_code,omitempty"`
	Primary_proposer     *Boolean   `xorm:" null default 0 'primary_proposer'" indexes:"" comment:"是否主要申请人" json:"primary_proposer,omitempty" bson:",omitempty" msgpack:"primary_proposer,omitempty"`
	Legal_person_id      *int64     `xorm:" null 'legal_person_id'" indexes:"" comment:"法定代表人主键" json:"legal_person_id,omitempty" bson:",omitempty" msgpack:"legal_person_id,omitempty"`
	Legal_person_name    *string    `xorm:" varchar(50) null 'legal_person_name'" indexes:"" comment:"法定代表人姓名" json:"legal_person_name,omitempty" bson:",omitempty" msgpack:"legal_person_name,omitempty"`
	Phone                *string    `xorm:" varchar(50) null 'phone'" indexes:"" comment:"联系电话" json:"phone,omitempty" bson:",omitempty" msgpack:"phone,omitempty"`
	Post_address         *string    `xorm:" varchar(255) null 'post_address'" indexes:"" comment:"通信地址" json:"post_address,omitempty" bson:",omitempty" msgpack:"post_address,omitempty"`
	Fixed_phone          *string    `xorm:" varchar(50) null 'fixed_phone'" indexes:"" comment:"固定电话" json:"fixed_phone,omitempty" bson:",omitempty" msgpack:"fixed_phone,omitempty"`
	Reg_address          *string    `xorm:" varchar(255) null 'reg_address'" indexes:"" comment:"户籍地址" json:"reg_address,omitempty" bson:",omitempty" msgpack:"reg_address,omitempty"`
	Bithday              *time.Time `xorm:" null 'bithday'" indexes:"" comment:"出生日期" json:"bithday,omitempty" bson:",omitempty" msgpack:"bithday,omitempty"`
	Age                  *int64     `xorm:" null 'age'" indexes:"" comment:"年龄" json:"age,omitempty" bson:",omitempty" msgpack:"age,omitempty"`
	Postcode             *string    `xorm:" varchar(50) null 'postcode'" indexes:"" comment:"邮编" json:"postcode,omitempty" bson:",omitempty" msgpack:"postcode,omitempty"`
	Potrait              *string    `xorm:" varchar(50) null 'potrait'" indexes:"" comment:"申请人头像" json:"potrait,omitempty" bson:",omitempty" msgpack:"potrait,omitempty"`
	Enable_mark          *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_2) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark          *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_3) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code            *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description          *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date          *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid        *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username      *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date          *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid        *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username      *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version              *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: dic_marriage_status
// 说明: 自治家园字典-婚姻状况
// 自治家园字典-婚姻状况
// 
type dic_marriage_status struct {
	Marriage_status_id     *int64     `xorm:" pk not null 'marriage_status_id'" indexes:"" comment:"婚姻状况主键" json:"marriage_status_id,omitempty" bson:",omitempty" msgpack:"marriage_status_id,omitempty"`
	Marriage_status_name   *string    `xorm:" varchar(255) null 'marriage_status_name'" indexes:"" comment:"婚姻状况名称" json:"marriage_status_name,omitempty" bson:",omitempty" msgpack:"marriage_status_name,omitempty"`
	Marriage_status_prompt *string    `xorm:" varchar(255) null 'marriage_status_prompt'" indexes:"" comment:"婚姻状况别称" json:"marriage_status_prompt,omitempty" bson:",omitempty" msgpack:"marriage_status_prompt,omitempty"`
	Marriage_status_code   *string    `xorm:" varchar(255) null 'marriage_status_code'" indexes:"" comment:"婚姻状况编码" json:"marriage_status_code,omitempty" bson:",omitempty" msgpack:"marriage_status_code,omitempty"`
	Enable_mark            *Boolean   `xorm:" null default 1 'enable_mark'" indexes:"" comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark            *Boolean   `xorm:" null default 0 'delete_mark'" indexes:"" comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code              *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description            *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date            *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid          *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username        *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date            *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid          *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username        *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version                *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: dic_military_service_status
// 说明: 自治家园字典-兵役状况
// 自治家园字典-兵役状况
// 
type dic_military_service_status struct {
	Military_service_status_id     *int64     `xorm:" pk not null 'military_service_status_id'" indexes:"" comment:"兵役状况主键" json:"military_service_status_id,omitempty" bson:",omitempty" msgpack:"military_service_status_id,omitempty"`
	Military_service_status_name   *string    `xorm:" varchar(255) null 'military_service_status_name'" indexes:"" comment:"兵役状况名称" json:"military_service_status_name,omitempty" bson:",omitempty" msgpack:"military_service_status_name,omitempty"`
	Military_service_status_prompt *string    `xorm:" varchar(255) null 'military_service_status_prompt'" indexes:"" comment:"兵役状况别称" json:"military_service_status_prompt,omitempty" bson:",omitempty" msgpack:"military_service_status_prompt,omitempty"`
	Military_service_status_code   *string    `xorm:" varchar(255) null 'military_service_status_code'" indexes:"" comment:"兵役状况编码" json:"military_service_status_code,omitempty" bson:",omitempty" msgpack:"military_service_status_code,omitempty"`
	Enable_mark                    *Boolean   `xorm:" null default 1 'enable_mark'" indexes:"" comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark                    *Boolean   `xorm:" null default 0 'delete_mark'" indexes:"" comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code                      *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description                    *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date                    *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid                  *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username                *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date                    *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid                  *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username                *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version                        *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: dic_part_time_job
// 说明: 自治家园字典-社会兼职
// 自治家园字典-社会兼职
// 
type dic_part_time_job struct {
	Part_time_job_id     *int64     `xorm:" pk not null 'part_time_job_id'" indexes:"" comment:"社会兼职主键" json:"part_time_job_id,omitempty" bson:",omitempty" msgpack:"part_time_job_id,omitempty"`
	Part_time_job_name   *string    `xorm:" varchar(255) null 'part_time_job_name'" indexes:"" comment:"社会兼职名称" json:"part_time_job_name,omitempty" bson:",omitempty" msgpack:"part_time_job_name,omitempty"`
	Part_time_job_prompt *string    `xorm:" varchar(255) null 'part_time_job_prompt'" indexes:"" comment:"社会兼职别称" json:"part_time_job_prompt,omitempty" bson:",omitempty" msgpack:"part_time_job_prompt,omitempty"`
	Part_time_job_code   *string    `xorm:" varchar(255) null 'part_time_job_code'" indexes:"" comment:"社会兼职编码" json:"part_time_job_code,omitempty" bson:",omitempty" msgpack:"part_time_job_code,omitempty"`
	Enable_mark          *Boolean   `xorm:" null default 1 'enable_mark'" indexes:"" comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark          *Boolean   `xorm:" null default 0 'delete_mark'" indexes:"" comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code            *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description          *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date          *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid        *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username      *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date          *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid        *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username      *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version              *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: dic_department
// 说明: 字典-部门
// 字典-部门
// 
type dic_department struct {
	Department_id     *int64     `xorm:" pk not null 'department_id'" indexes:"" comment:"部门主键" json:"department_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"department_id,omitempty"`
	Department_name   *string    `xorm:" varchar(50) null 'department_name'" indexes:"" comment:"部门名称" json:"department_name,omitempty" bson:",omitempty" msgpack:"department_name,omitempty"`
	Department_prompt *string    `xorm:" varchar(50) null 'department_prompt'" indexes:"" comment:"显示名称" json:"department_prompt,omitempty" bson:",omitempty" msgpack:"department_prompt,omitempty"`
	Value_code        *string    `xorm:" varchar(50) null 'value_code'" indexes:" unique(Index_1) " comment:"编码" json:"value_code,omitempty" bson:",omitempty" msgpack:"value_code,omitempty"`
	Icon_url          *string    `xorm:" varchar(255) null 'icon_url'" indexes:"" comment:"显示图标" json:"icon_url,omitempty" bson:",omitempty" msgpack:"icon_url,omitempty"`
	Enable_mark       *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_2) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark       *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_3) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code         *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description       *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date       *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid     *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username   *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date       *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid     *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username   *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version           *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: dic_family_master_relation
// 说明: 自治家园字典-与户主关系
// 自治家园字典-与户主关系
// 
type dic_family_master_relation struct {
	Family_master_relation_id     *int64     `xorm:" pk not null 'family_master_relation_id'" indexes:"" comment:"与户主关系主键" json:"family_master_relation_id,omitempty" bson:",omitempty" msgpack:"family_master_relation_id,omitempty"`
	Family_master_relation_name   *string    `xorm:" varchar(255) null 'family_master_relation_name'" indexes:"" comment:"与户主关系名称" json:"family_master_relation_name,omitempty" bson:",omitempty" msgpack:"family_master_relation_name,omitempty"`
	Family_master_relation_prompt *string    `xorm:" varchar(255) null 'family_master_relation_prompt'" indexes:"" comment:"与户主关系别称" json:"family_master_relation_prompt,omitempty" bson:",omitempty" msgpack:"family_master_relation_prompt,omitempty"`
	Family_master_relation_code   *string    `xorm:" varchar(255) null 'family_master_relation_code'" indexes:"" comment:"与户主关系编码" json:"family_master_relation_code,omitempty" bson:",omitempty" msgpack:"family_master_relation_code,omitempty"`
	Enable_mark                   *Boolean   `xorm:" null default 1 'enable_mark'" indexes:"" comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark                   *Boolean   `xorm:" null default 0 'delete_mark'" indexes:"" comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code                     *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description                   *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date                   *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid                 *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username               *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date                   *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid                 *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username               *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version                       *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: dic_apply_from
// 说明: 字典-申请来源
// 字典-申请来源
// 
type dic_apply_from struct {
	Apply_from_id     *int64     `xorm:" pk not null 'apply_from_id'" indexes:"" comment:"申请来源主键" json:"apply_from_id,omitempty" bson:",omitempty" msgpack:"apply_from_id,omitempty"`
	Apply_from_name   *string    `xorm:" varchar(255) null 'apply_from_name'" indexes:" unique(Index_1) " comment:"申请来源名称" json:"apply_from_name,omitempty" bson:",omitempty" msgpack:"apply_from_name,omitempty"`
	Apply_from_prompt *string    `xorm:" varchar(255) null 'apply_from_prompt'" indexes:"" comment:"申请来源简称" json:"apply_from_prompt,omitempty" bson:",omitempty" msgpack:"apply_from_prompt,omitempty"`
	Apply_from_code   *string    `xorm:" varchar(255) null 'apply_from_code'" indexes:" unique(Index_2) " comment:"申请来源编码" json:"apply_from_code,omitempty" bson:",omitempty" msgpack:"apply_from_code,omitempty"`
	Enable_mark       *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_3) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark       *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_4) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code         *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description       *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date       *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid     *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username   *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date       *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid     *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username   *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version           *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: dic_card_type
// 说明: 自治家园字典-证件类型
// 自治家园字典-证件类型
// 
type dic_card_type struct {
	Card_type_id     *int64     `xorm:" pk not null 'card_type_id'" indexes:"" comment:"证件类型主键" json:"card_type_id,omitempty" bson:",omitempty" msgpack:"card_type_id,omitempty"`
	Card_type_name   *string    `xorm:" varchar(255) null 'card_type_name'" indexes:"" comment:"证件类型名称" json:"card_type_name,omitempty" bson:",omitempty" msgpack:"card_type_name,omitempty"`
	Card_type_prompt *string    `xorm:" varchar(255) null 'card_type_prompt'" indexes:"" comment:"证件类型别称" json:"card_type_prompt,omitempty" bson:",omitempty" msgpack:"card_type_prompt,omitempty"`
	Card_type_code   *string    `xorm:" varchar(255) null 'card_type_code'" indexes:"" comment:"证件类型编码" json:"card_type_code,omitempty" bson:",omitempty" msgpack:"card_type_code,omitempty"`
	Enable_mark      *Boolean   `xorm:" null default 1 'enable_mark'" indexes:"" comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark      *Boolean   `xorm:" null default 0 'delete_mark'" indexes:"" comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code        *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description      *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date      *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid    *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username  *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date      *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid    *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username  *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version          *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: dic_duty_level
// 说明: 自治家园字典-干部职务级别
// 自治家园字典-干部职务级别
// 
type dic_duty_level struct {
	Duty_level_id     *int64     `xorm:" pk not null 'duty_level_id'" indexes:"" comment:"干部职务主键" json:"duty_level_id,omitempty" bson:",omitempty" msgpack:"duty_level_id,omitempty"`
	Duty_level_name   *string    `xorm:" varchar(255) null 'duty_level_name'" indexes:"" comment:"干部职务名称" json:"duty_level_name,omitempty" bson:",omitempty" msgpack:"duty_level_name,omitempty"`
	Duty_level_prompt *string    `xorm:" varchar(255) null 'duty_level_prompt'" indexes:"" comment:"干部职务别称" json:"duty_level_prompt,omitempty" bson:",omitempty" msgpack:"duty_level_prompt,omitempty"`
	Duty_level_code   *string    `xorm:" varchar(255) null 'duty_level_code'" indexes:"" comment:"干部职务编码" json:"duty_level_code,omitempty" bson:",omitempty" msgpack:"duty_level_code,omitempty"`
	Enable_mark       *Boolean   `xorm:" null default 1 'enable_mark'" indexes:"" comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark       *Boolean   `xorm:" null default 0 'delete_mark'" indexes:"" comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code         *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description       *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date       *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid     *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username   *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date       *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid     *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username   *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version           *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: dic_unit_property
// 说明: 自治家园字典-单位性质
// 自治家园字典-单位性质
// 
type dic_unit_property struct {
	Unit_property_id     *int64     `xorm:" pk not null 'unit_property_id'" indexes:"" comment:"单位性质主键" json:"unit_property_id,omitempty" bson:",omitempty" msgpack:"unit_property_id,omitempty"`
	Unit_property_name   *string    `xorm:" varchar(255) null 'unit_property_name'" indexes:"" comment:"单位性质名称" json:"unit_property_name,omitempty" bson:",omitempty" msgpack:"unit_property_name,omitempty"`
	Unit_property_prompt *string    `xorm:" varchar(255) null 'unit_property_prompt'" indexes:"" comment:"单位性质别称" json:"unit_property_prompt,omitempty" bson:",omitempty" msgpack:"unit_property_prompt,omitempty"`
	Unit_property_code   *string    `xorm:" varchar(255) null 'unit_property_code'" indexes:"" comment:"单位性质编码" json:"unit_property_code,omitempty" bson:",omitempty" msgpack:"unit_property_code,omitempty"`
	Enable_mark          *Boolean   `xorm:" null default 1 'enable_mark'" indexes:"" comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark          *Boolean   `xorm:" null default 0 'delete_mark'" indexes:"" comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code            *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description          *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date          *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid        *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username      *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date          *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid        *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username      *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version              *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: tab_vm_accounts
// 说明: 专线设置-虚拟机登录帐号
// 专线设置-虚拟机登录帐号
// 
type tab_vm_accounts struct {
	Account_id      *int64     `xorm:" pk not null 'account_id'" indexes:"" comment:"帐号主键" json:"account_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"account_id,omitempty"`
	Vm_id           *int64     `xorm:" null 'vm_id'" indexes:"" comment:"虚拟机主键" json:"vm_id,omitempty" bson:",omitempty" msgpack:"vm_id,omitempty" fk:"tab_virtual_machines(vm_id)"`
	Account_name    *string    `xorm:" varchar(255) null 'account_name'" indexes:"" comment:"帐号名" json:"account_name,omitempty" bson:",omitempty" msgpack:"account_name,omitempty"`
	Account_pwd     *string    `xorm:" varchar(255) null 'account_pwd'" indexes:"" comment:"登录密码" json:"account_pwd,omitempty" bson:",omitempty" msgpack:"account_pwd,omitempty"`
	Enable_mark     *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_1) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark     *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_2) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code       *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description     *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date     *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid   *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date     *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid   *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version         *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: tab_vpn_connect
// 说明: 专线设置-专线vpn连接
// 专线设置-专线vpn连接
// 
type tab_vpn_connect struct {
	Vpn_id             *int64     `xorm:" pk not null 'vpn_id'" indexes:"" comment:"vpn连接主键" json:"vpn_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"vpn_id,omitempty"`
	Vm_id              *int64     `xorm:" null 'vm_id'" indexes:"" comment:"虚拟机主键" json:"vm_id,omitempty" bson:",omitempty" msgpack:"vm_id,omitempty" fk:"tab_virtual_machines(vm_id)"`
	Vpn_excutable_path *string    `xorm:" varchar(500) null 'vpn_excutable_path'" indexes:"" comment:"vpn客户端" json:"vpn_excutable_path,omitempty" bson:",omitempty" msgpack:"vpn_excutable_path,omitempty"`
	Vpn_account        *string    `xorm:" varchar(255) null 'vpn_account'" indexes:"" comment:"vpn帐号" json:"vpn_account,omitempty" bson:",omitempty" msgpack:"vpn_account,omitempty"`
	Vpn_pwd            *string    `xorm:" varchar(255) null 'vpn_pwd'" indexes:"" comment:"vpn密码" json:"vpn_pwd,omitempty" bson:",omitempty" msgpack:"vpn_pwd,omitempty"`
	Enable_mark        *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_1) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark        *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_2) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code          *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description        *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date        *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid      *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username    *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date        *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid      *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username    *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version            *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: dic_health_status
// 说明: 自治家园字典-健康状况
// 自治家园字典-健康状况
// 
type dic_health_status struct {
	Health_status_id     *int64     `xorm:" pk not null 'health_status_id'" indexes:"" comment:"健康状况主键" json:"health_status_id,omitempty" bson:",omitempty" msgpack:"health_status_id,omitempty"`
	Health_status_name   *string    `xorm:" varchar(255) null 'health_status_name'" indexes:"" comment:"健康状况名称" json:"health_status_name,omitempty" bson:",omitempty" msgpack:"health_status_name,omitempty"`
	Health_status_prompt *string    `xorm:" varchar(255) null 'health_status_prompt'" indexes:"" comment:"健康状况别称" json:"health_status_prompt,omitempty" bson:",omitempty" msgpack:"health_status_prompt,omitempty"`
	Health_status_code   *string    `xorm:" varchar(255) null 'health_status_code'" indexes:"" comment:"健康状况编码" json:"health_status_code,omitempty" bson:",omitempty" msgpack:"health_status_code,omitempty"`
	Enable_mark          *Boolean   `xorm:" null default 1 'enable_mark'" indexes:"" comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark          *Boolean   `xorm:" null default 0 'delete_mark'" indexes:"" comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code            *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description          *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date          *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid        *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username      *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date          *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid        *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username      *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version              *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: tab_coporation_change
// 说明: 法人变更信息
// 法人变更信息
// 
type tab_coporation_change struct {
	Coporation_change_id *int64     `xorm:" pk not null 'coporation_change_id'" indexes:"" comment:"法人变更主键" json:"coporation_change_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"coporation_change_id,omitempty"`
	Credit_code          *string    `xorm:" varchar(255) null 'credit_code'" indexes:"" comment:"社会统一信用代码" json:"credit_code,omitempty" bson:",omitempty" msgpack:"credit_code,omitempty"`
	Org_code             *string    `xorm:" varchar(50) null 'org_code'" indexes:"" comment:"组织机构代码" json:"org_code,omitempty" bson:",omitempty" msgpack:"org_code,omitempty"`
	Change_issue         *string    `xorm:" varchar(50) null 'change_issue'" indexes:"" comment:"变更事项" json:"change_issue,omitempty" bson:",omitempty" msgpack:"change_issue,omitempty"`
	Before_change        *string    `xorm:" text null 'before_change'" indexes:"" comment:"变更前" json:"before_change,omitempty" bson:",omitempty" msgpack:"before_change,omitempty"`
	After_change         *string    `xorm:" text null 'after_change'" indexes:"" comment:"变更后" json:"after_change,omitempty" bson:",omitempty" msgpack:"after_change,omitempty"`
	Change_date          *time.Time `xorm:" null 'change_date'" indexes:"" comment:"变更日期" json:"change_date,omitempty" bson:",omitempty" msgpack:"change_date,omitempty"`
	Reg_sector           *string    `xorm:" varchar(255) null 'reg_sector'" indexes:"" comment:"登记注册机关" json:"reg_sector,omitempty" bson:",omitempty" msgpack:"reg_sector,omitempty"`
	Enable_mark          *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_1) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark          *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_2) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code            *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description          *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date          *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid        *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username      *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date          *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid        *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username      *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version              *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: tab_monitors_extend_log
// 说明: 行政监察补正延长记录
// 行政监察补正延长记录
// 
type tab_monitors_extend_log struct {
	Extend_id           *int64     `xorm:" pk not null 'extend_id'" indexes:"" comment:"补正延长记录主键" json:"extend_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"extend_id,omitempty"`
	Monitor_id          *int64     `xorm:" null 'monitor_id'" indexes:"" comment:"监察主键" json:"monitor_id,omitempty" bson:",omitempty" msgpack:"monitor_id,omitempty" fk:"tab_monitors(monitor_id)"`
	Limit_time          *int64     `xorm:" null 'limit_time'" indexes:"" comment:"延长前时限" json:"limit_time,omitempty" bson:",omitempty" msgpack:"limit_time,omitempty"`
	Extended_limit_time *int64     `xorm:" null 'extended_limit_time'" indexes:"" comment:"延长后时限" json:"extended_limit_time,omitempty" bson:",omitempty" msgpack:"extended_limit_time,omitempty"`
	Extend_time         *int64     `xorm:" null 'extend_time'" indexes:"" comment:"延长时限" json:"extend_time,omitempty" bson:",omitempty" msgpack:"extend_time,omitempty"`
	Affairs_start_time  *time.Time `xorm:" null 'affairs_start_time'" indexes:"" comment:"事项开始时间" json:"affairs_start_time,omitempty" bson:",omitempty" msgpack:"affairs_start_time,omitempty"`
	Affairs_reset_time  *time.Time `xorm:" null 'affairs_reset_time'" indexes:"" comment:"事项重置时间" json:"affairs_reset_time,omitempty" bson:",omitempty" msgpack:"affairs_reset_time,omitempty"`
	Enable_mark         *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_1) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark         *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_2) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code           *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description         *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date         *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid       *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username     *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date         *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid       *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username     *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version             *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: tab_outer_user
// 说明: 外网-用户表
// 外网-用户表
// 
type tab_outer_user struct {
	Out_user_id      *int64     `xorm:" pk not null 'out_user_id'" indexes:"" comment:"外网用户主键" json:"out_user_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"out_user_id,omitempty"`
	User_name        *string    `xorm:" varchar(255) null 'user_name'" indexes:"" comment:"用户名" json:"user_name,omitempty" bson:",omitempty" msgpack:"user_name,omitempty"`
	Passwd           *string    `xorm:" varchar(255) not null 'passwd'" indexes:"" comment:"登录密码" json:"passwd,omitempty" bson:",omitempty" msgpack:"passwd,omitempty"`
	Name             *string    `xorm:" varchar(255) not null 'name'" indexes:"" comment:"姓名" json:"name,omitempty" bson:",omitempty" msgpack:"name,omitempty"`
	Sex              *int64     `xorm:" null 'sex'" indexes:"" comment:"性别" json:"sex,omitempty" bson:",omitempty" msgpack:"sex,omitempty"`
	Id_code          *string    `xorm:" varchar(255) null 'id_code'" indexes:" unique(Index_1) " comment:"身份证号码" json:"id_code,omitempty" bson:",omitempty" msgpack:"id_code,omitempty"`
	Mobile           *string    `xorm:" varchar(255) not null 'mobile'" indexes:" unique(Index_2) " comment:"手机" json:"mobile,omitempty" bson:",omitempty" msgpack:"mobile,omitempty"`
	Email            *string    `xorm:" varchar(255) null 'email'" indexes:" unique(Index_3) " comment:"邮箱" json:"email,omitempty" bson:",omitempty" msgpack:"email,omitempty"`
	Photo_url        *string    `xorm:" varchar(255) null 'photo_url'" indexes:"" comment:"照片" json:"photo_url,omitempty" bson:",omitempty" msgpack:"photo_url,omitempty"`
	Is_nature_person *Boolean   `xorm:" null default 1 'is_nature_person'" indexes:"" comment:"是否自然人" json:"is_nature_person,omitempty" bson:",omitempty" msgpack:"is_nature_person,omitempty"`
	Credit_code      *string    `xorm:" varchar(50) null 'credit_code'" indexes:"" comment:"组织机构社会统一信用代码" json:"credit_code,omitempty" bson:",omitempty" msgpack:"credit_code,omitempty"`
	Orgnize_name     *string    `xorm:" varchar(255) null 'orgnize_name'" indexes:"" comment:"组织机构名称" json:"orgnize_name,omitempty" bson:",omitempty" msgpack:"orgnize_name,omitempty"`
	Is_admin_user    *Boolean   `xorm:" null default 0 'is_admin_user'" indexes:"" comment:"是否管理员" json:"is_admin_user,omitempty" bson:",omitempty" msgpack:"is_admin_user,omitempty"`
	Is_real_name     *Boolean   `xorm:" null default 0 'is_real_name'" indexes:"" comment:"是否已实名认证" json:"is_real_name,omitempty" bson:",omitempty" msgpack:"is_real_name,omitempty"`
	Open_id_qq       *string    `xorm:" varchar(255) null 'open_id_qq'" indexes:"" comment:"QQ号" json:"open_id_qq,omitempty" bson:",omitempty" msgpack:"open_id_qq,omitempty"`
	Open_id_wx       *string    `xorm:" varchar(255) null 'open_id_wx'" indexes:"" comment:"微信号" json:"open_id_wx,omitempty" bson:",omitempty" msgpack:"open_id_wx,omitempty"`
	Location_id      *int64     `xorm:" null 'location_id'" indexes:"" comment:"归属地" json:"location_id,omitempty" bson:",omitempty" msgpack:"location_id,omitempty"`
	Location_name    *string    `xorm:" varchar(255) null 'location_name'" indexes:"" comment:"归属地名称" json:"location_name,omitempty" bson:",omitempty" msgpack:"location_name,omitempty"`
	Client_id        *string    `xorm:" varchar(255) null 'client_id'" indexes:"" comment:"客户端Id" json:"client_id,omitempty" bson:",omitempty" msgpack:"client_id,omitempty"`
	Last_login_time  *time.Time `xorm:" null 'last_login_time'" indexes:"" comment:"最后登录时间" json:"last_login_time,omitempty" bson:",omitempty" msgpack:"last_login_time,omitempty"`
	Enable_mark      *Boolean   `xorm:" null default 1 'enable_mark'" indexes:"" comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark      *Boolean   `xorm:" null default 0 'delete_mark'" indexes:"" comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code        *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description      *string    `xorm:" varchar(255) null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date      *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid    *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username  *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date      *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid    *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username  *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version          *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: dic_politics_status
// 说明: 自治家园字典-政治面貌
// 自治家园字典-政治面貌
// 
type dic_politics_status struct {
	Politics_status_id     *int64     `xorm:" pk not null 'politics_status_id'" indexes:"" comment:"政治面貌主键" json:"politics_status_id,omitempty" bson:",omitempty" msgpack:"politics_status_id,omitempty"`
	Politics_status_name   *string    `xorm:" varchar(255) null 'politics_status_name'" indexes:"" comment:"政治面貌名称" json:"politics_status_name,omitempty" bson:",omitempty" msgpack:"politics_status_name,omitempty"`
	Politics_status_prompt *string    `xorm:" varchar(255) null 'politics_status_prompt'" indexes:"" comment:"政治面貌别称" json:"politics_status_prompt,omitempty" bson:",omitempty" msgpack:"politics_status_prompt,omitempty"`
	Politics_status_code   *string    `xorm:" varchar(255) null 'politics_status_code'" indexes:"" comment:"政治面貌编码" json:"politics_status_code,omitempty" bson:",omitempty" msgpack:"politics_status_code,omitempty"`
	Enable_mark            *Boolean   `xorm:" null default 1 'enable_mark'" indexes:"" comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark            *Boolean   `xorm:" null default 0 'delete_mark'" indexes:"" comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code              *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description            *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date            *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid          *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username        *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date            *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid          *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username        *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version                *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: dic_unit_trade
// 说明: 自治家园字典-单位所属行业
// 自治家园字典-单位所属行业
// 
type dic_unit_trade struct {
	Unit_trade_id     *int64     `xorm:" pk not null 'unit_trade_id'" indexes:"" comment:"单位所属行业主键" json:"unit_trade_id,omitempty" bson:",omitempty" msgpack:"unit_trade_id,omitempty"`
	Unit_trade_name   *string    `xorm:" varchar(255) null 'unit_trade_name'" indexes:"" comment:"单位所属行业名称" json:"unit_trade_name,omitempty" bson:",omitempty" msgpack:"unit_trade_name,omitempty"`
	Unit_trade_prompt *string    `xorm:" varchar(255) null 'unit_trade_prompt'" indexes:"" comment:"单位所属行业别称" json:"unit_trade_prompt,omitempty" bson:",omitempty" msgpack:"unit_trade_prompt,omitempty"`
	Unit_trade_code   *string    `xorm:" varchar(255) null 'unit_trade_code'" indexes:"" comment:"单位所属行业编码" json:"unit_trade_code,omitempty" bson:",omitempty" msgpack:"unit_trade_code,omitempty"`
	Enable_mark       *Boolean   `xorm:" null default 1 'enable_mark'" indexes:"" comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark       *Boolean   `xorm:" null default 0 'delete_mark'" indexes:"" comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code         *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description       *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date       *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid     *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username   *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date       *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid     *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username   *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version           *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: tab_credential_detail
// 说明: 证照详细信息
// 证照详细信息
// 
type tab_credential_detail struct {
	Detail_id       *int64     `xorm:" pk not null 'detail_id'" indexes:"" comment:"证照详细信息主键" json:"detail_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"detail_id,omitempty"`
	Field_id        *int64     `xorm:" null 'field_id'" indexes:"" comment:"证面字段主键" json:"field_id,omitempty" bson:",omitempty" msgpack:"field_id,omitempty" fk:"tab_credential_field(field_id)"`
	Field_name      *string    `xorm:" varchar(255) null 'field_name'" indexes:"" comment:"证面字段名称" json:"field_name,omitempty" bson:",omitempty" msgpack:"field_name,omitempty"`
	Field_value     *string    `xorm:" varchar(255) null 'field_value'" indexes:"" comment:"证面字段值" json:"field_value,omitempty" bson:",omitempty" msgpack:"field_value,omitempty"`
	Enable_mark     *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_1) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark     *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_2) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code       *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description     *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date     *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid   *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date     *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid   *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version         *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: tab_holiday_adjusted
// 说明: 除正常工作日（周一到周五）外，国家规定有变化的工作日和休息日的调整记录
// 除正常工作日（周一到周五）外，国家规定有变化的工作日和休息日的调整记录
// 
type tab_holiday_adjusted struct {
	Date_id         *int64     `xorm:" pk not null 'date_id'" indexes:"" comment:"日期主键" json:"date_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"date_id,omitempty"`
	The_date        *time.Time `xorm:" null 'the_date'" indexes:" unique(Index_1) " comment:"日期" json:"the_date,omitempty" bson:",omitempty" msgpack:"the_date,omitempty"`
	Work_tag        *Boolean   `xorm:" null default 0 'work_tag'" indexes:"" comment:"是否工作" json:"work_tag,omitempty" bson:",omitempty" msgpack:"work_tag,omitempty"`
	Enable_mark     *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_2) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark     *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_3) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code       *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description     *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date     *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid   *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date     *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid   *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version         *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: dic_blood_type
// 说明: 自治家园字典-血型
// 自治家园字典-血型
// 
type dic_blood_type struct {
	Blood_type_id     *int64     `xorm:" pk not null 'blood_type_id'" indexes:"" comment:"血型主键" json:"blood_type_id,omitempty" bson:",omitempty" msgpack:"blood_type_id,omitempty"`
	Blood_type_name   *string    `xorm:" varchar(255) null 'blood_type_name'" indexes:"" comment:"血型名称" json:"blood_type_name,omitempty" bson:",omitempty" msgpack:"blood_type_name,omitempty"`
	Blood_type_prompt *string    `xorm:" varchar(255) null 'blood_type_prompt'" indexes:"" comment:"血型别称" json:"blood_type_prompt,omitempty" bson:",omitempty" msgpack:"blood_type_prompt,omitempty"`
	Blood_type_code   *string    `xorm:" varchar(255) null 'blood_type_code'" indexes:"" comment:"血型编码" json:"blood_type_code,omitempty" bson:",omitempty" msgpack:"blood_type_code,omitempty"`
	Enable_mark       *Boolean   `xorm:" null default 1 'enable_mark'" indexes:"" comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark       *Boolean   `xorm:" null default 0 'delete_mark'" indexes:"" comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code         *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description       *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date       *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid     *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username   *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date       *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid     *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username   *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version           *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: tab_scripts
// 说明: 脚本表
// 脚本表
// 
type tab_scripts struct {
	Script_id       *int64     `xorm:" pk not null 'script_id'" indexes:"" comment:"脚本主键" json:"script_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"script_id,omitempty"`
	Script_name     *string    `xorm:" varchar(255) null 'script_name'" indexes:"" comment:"脚本名称" json:"script_name,omitempty" bson:",omitempty" msgpack:"script_name,omitempty"`
	Script_content  *string    `xorm:" text null 'script_content'" indexes:"" comment:"脚本内容" json:"script_content,omitempty" bson:",omitempty" msgpack:"script_content,omitempty"`
	Enable_mark     *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_1) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark     *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_2) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code       *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description     *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date     *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid   *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date     *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid   *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version         *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: tab_system_config
// 说明: 系统配置
// 系统配置
// 
type tab_system_config struct {
	Config_id          *int64     `xorm:" pk not null 'config_id'" indexes:"" comment:"配置项主键" json:"config_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"config_id,omitempty"`
	Config_name        *string    `xorm:" varchar(255) null 'config_name'" indexes:" unique(Index_1) " comment:"配置项" json:"config_name,omitempty" bson:",omitempty" msgpack:"config_name,omitempty"`
	Config_value       *string    `xorm:" varchar(255) null 'config_value'" indexes:"" comment:"配置项参数" json:"config_value,omitempty" bson:",omitempty" msgpack:"config_value,omitempty"`
	Config_description *string    `xorm:" varchar(500) null 'config_description'" indexes:"" comment:"描述" json:"config_description,omitempty" bson:",omitempty" msgpack:"config_description,omitempty"`
	Enable_mark        *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_2) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark        *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_3) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code          *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description        *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date        *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid      *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username    *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date        *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid      *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username    *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version            *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: dic_net_level
// 说明: 字典-网上办理深度
// 字典-网上办理深度
// 
type dic_net_level struct {
	Net_level_id     *int64     `xorm:" pk not null 'net_level_id'" indexes:"" comment:"办理深度主键" json:"net_level_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"net_level_id,omitempty"`
	Net_level_name   *string    `xorm:" varchar(255) null 'net_level_name'" indexes:" unique(Index_1) " comment:"办理深度名称" json:"net_level_name,omitempty" bson:",omitempty" msgpack:"net_level_name,omitempty"`
	Net_level_prompt *string    `xorm:" varchar(255) null 'net_level_prompt'" indexes:"" comment:"办理深度简称" json:"net_level_prompt,omitempty" bson:",omitempty" msgpack:"net_level_prompt,omitempty"`
	Net_level_code   *string    `xorm:" varchar(50) null 'net_level_code'" indexes:" unique(Index_2) " comment:"办理深度编码" json:"net_level_code,omitempty" bson:",omitempty" msgpack:"net_level_code,omitempty"`
	Enable_mark      *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_3) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark      *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_4) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code        *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description      *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date      *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid    *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username  *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date      *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid    *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username  *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version          *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: dic_positions
// 说明: 字典-大厅岗位设置
// 字典-大厅岗位设置
// 
type dic_positions struct {
	Position_id     *int64     `xorm:" pk not null 'position_id'" indexes:"" comment:"岗位主键" json:"position_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"position_id,omitempty"`
	Position_name   *string    `xorm:" varchar(50) null 'position_name'" indexes:"" comment:"岗位名称" json:"position_name,omitempty" bson:",omitempty" msgpack:"position_name,omitempty"`
	Position_prompt *string    `xorm:" varchar(50) null 'position_prompt'" indexes:"" comment:"岗位简称" json:"position_prompt,omitempty" bson:",omitempty" msgpack:"position_prompt,omitempty"`
	Value_code      *string    `xorm:" varchar(50) null 'value_code'" indexes:" unique(Index_1) " comment:"岗位编码" json:"value_code,omitempty" bson:",omitempty" msgpack:"value_code,omitempty"`
	Enable_mark     *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_2) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark     *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_3) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code       *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description     *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date     *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid   *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date     *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid   *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version         *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: dic_technology_duty
// 说明: 自治家园字典-专业技术职务
// 自治家园字典-专业技术职务
// 
type dic_technology_duty struct {
	Technology_duty_id     *int64     `xorm:" pk not null 'technology_duty_id'" indexes:"" comment:"专业技术职务主键" json:"technology_duty_id,omitempty" bson:",omitempty" msgpack:"technology_duty_id,omitempty"`
	Technology_duty_name   *string    `xorm:" varchar(255) null 'technology_duty_name'" indexes:"" comment:"专业技术职务名称" json:"technology_duty_name,omitempty" bson:",omitempty" msgpack:"technology_duty_name,omitempty"`
	Technology_duty_prompt *string    `xorm:" varchar(255) null 'technology_duty_prompt'" indexes:"" comment:"专业技术职务别称" json:"technology_duty_prompt,omitempty" bson:",omitempty" msgpack:"technology_duty_prompt,omitempty"`
	Technology_duty_code   *string    `xorm:" varchar(255) null 'technology_duty_code'" indexes:"" comment:"专业技术职务编码" json:"technology_duty_code,omitempty" bson:",omitempty" msgpack:"technology_duty_code,omitempty"`
	Enable_mark            *Boolean   `xorm:" null default 1 'enable_mark'" indexes:"" comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark            *Boolean   `xorm:" null default 0 'delete_mark'" indexes:"" comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code              *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description            *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date            *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid          *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username        *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date            *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid          *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username        *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version                *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: tab_application_user_list
// 说明: 用户中心-应用系统用户列表
// 用户中心-应用系统用户列表
// 
type tab_application_user_list struct {
	List_id         *int64     `xorm:" pk not null 'list_id'" indexes:"" comment:"应用系统用户列表主键" json:"list_id,omitempty" bson:",omitempty" msgpack:"list_id,omitempty"`
	System_id       *int64     `xorm:" null 'system_id'" indexes:"" comment:"系统主键" json:"system_id,omitempty" bson:",omitempty" msgpack:"system_id,omitempty" fk:"tab_system_list(system_id)"`
	System_account  *string    `xorm:" varchar(50) null 'system_account'" indexes:"" comment:"应用系统账户" json:"system_account,omitempty" bson:",omitempty" msgpack:"system_account,omitempty"`
	Bind_tag        *Boolean   `xorm:" null default 0 'bind_tag'" indexes:"" comment:"是否绑定" json:"bind_tag,omitempty" bson:",omitempty" msgpack:"bind_tag,omitempty"`
	Enable_mark     *Boolean   `xorm:" null default 1 'enable_mark'" indexes:"" comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark     *Boolean   `xorm:" null default 0 'delete_mark'" indexes:"" comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code       *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description     *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date     *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid   *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date     *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid   *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version         *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: tab_consultation_replies
// 说明: 外网-咨询回复表
// 外网-咨询回复表
// 
type tab_consultation_replies struct {
	Reply_id        *int64     `xorm:" pk not null 'reply_id'" indexes:"" comment:"回复主键" json:"reply_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"reply_id,omitempty"`
	Consultation_id *int64     `xorm:" null 'consultation_id'" indexes:"" comment:"咨询主键" json:"consultation_id,omitempty" bson:",omitempty" msgpack:"consultation_id,omitempty" fk:"tab_outer_consultation(consultation_id)"`
	Reply_content   *string    `xorm:" text null 'reply_content'" indexes:"" comment:"回复内容" json:"reply_content,omitempty" bson:",omitempty" msgpack:"reply_content,omitempty"`
	Reply_user_id   *int64     `xorm:" null 'reply_user_id'" indexes:"" comment:"回复人员" json:"reply_user_id,omitempty" bson:",omitempty" msgpack:"reply_user_id,omitempty"`
	Reply_user_name *string    `xorm:" varchar(255) null 'reply_user_name'" indexes:"" comment:"回复人姓名" json:"reply_user_name,omitempty" bson:",omitempty" msgpack:"reply_user_name,omitempty"`
	Enable_mark     *Boolean   `xorm:" null default 1 'enable_mark'" indexes:"" comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark     *Boolean   `xorm:" null default 0 'delete_mark'" indexes:"" comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code       *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description     *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date     *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid   *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date     *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid   *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version         *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: tab_natural_person_info
// 说明: 自然人信息表
// 自然人信息表
// 
type tab_natural_person_info struct {
	Natural_person_id  *int64     `xorm:" pk not null 'natural_person_id'" indexes:"" comment:"自然人主键" json:"natural_person_id,omitempty" bson:",omitempty" msgpack:"natural_person_id,omitempty"`
	Document_info_id   *int64     `xorm:" null 'document_info_id'" indexes:"" comment:"证件信息" json:"document_info_id,omitempty" bson:",omitempty" msgpack:"document_info_id,omitempty"`
	Account            *string    `xorm:" varchar(255) null 'account'" indexes:"" comment:"登录账号" json:"account,omitempty" bson:",omitempty" msgpack:"account,omitempty"`
	Passwd             *string    `xorm:" varchar(255) null 'passwd'" indexes:"" comment:"登录密码" json:"passwd,omitempty" bson:",omitempty" msgpack:"passwd,omitempty"`
	Name               *string    `xorm:" varchar(255) null 'name'" indexes:"" comment:"姓名" json:"name,omitempty" bson:",omitempty" msgpack:"name,omitempty"`
	Sex                *int64     `xorm:" null 'sex'" indexes:"" comment:"性别" json:"sex,omitempty" bson:",omitempty" msgpack:"sex,omitempty"`
	Credential_type_id *int64     `xorm:" null 'credential_type_id'" indexes:"" comment:"证照类型主键" json:"credential_type_id,omitempty" bson:",omitempty" msgpack:"credential_type_id,omitempty" fk:"dic_credential_type(credential_type_id)"`
	Document_code      *string    `xorm:" varchar(255) null 'document_code'" indexes:"" comment:"证件号码" json:"document_code,omitempty" bson:",omitempty" msgpack:"document_code,omitempty"`
	Birtyday           *time.Time `xorm:" null 'birtyday'" indexes:"" comment:"出生日期" json:"birtyday,omitempty" bson:",omitempty" msgpack:"birtyday,omitempty"`
	Tel                *string    `xorm:" varchar(255) null 'tel'" indexes:"" comment:"固定电话" json:"tel,omitempty" bson:",omitempty" msgpack:"tel,omitempty"`
	Mobile             *string    `xorm:" varchar(255) null 'mobile'" indexes:"" comment:"手机" json:"mobile,omitempty" bson:",omitempty" msgpack:"mobile,omitempty"`
	Reg_address        *string    `xorm:" varchar(255) null 'reg_address'" indexes:"" comment:"户籍所在地址" json:"reg_address,omitempty" bson:",omitempty" msgpack:"reg_address,omitempty"`
	Live_address       *string    `xorm:" varchar(255) null 'live_address'" indexes:"" comment:"当前居住地址" json:"live_address,omitempty" bson:",omitempty" msgpack:"live_address,omitempty"`
	Photo_url          *string    `xorm:" varchar(255) null 'photo_url'" indexes:"" comment:"照片" json:"photo_url,omitempty" bson:",omitempty" msgpack:"photo_url,omitempty"`
	Natural_man_status *int64     `xorm:" null 'natural_man_status'" indexes:"" comment:"自然人状态" json:"natural_man_status,omitempty" bson:",omitempty" msgpack:"natural_man_status,omitempty"`
	Enable_mark        *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_1) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark        *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_2) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code          *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description        *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date        *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid      *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username    *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date        *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid      *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username    *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version            *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: dic_affairs_subject
// 说明: 字典-事项主题
// 字典-事项主题
// 
type dic_affairs_subject struct {
	Affairs_subject_id     *int64     `xorm:" pk not null 'affairs_subject_id'" indexes:"" comment:"事项主题主键" json:"affairs_subject_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"affairs_subject_id,omitempty"`
	Proposer_type_id       *int64     `xorm:" null 'proposer_type_id'" indexes:"" comment:"申请人类型主键" json:"proposer_type_id,omitempty" bson:",omitempty" msgpack:"proposer_type_id,omitempty" fk:"dic_proposer_type(proposer_type_id)"`
	Affairs_subject_name   *string    `xorm:" varchar(255) null 'affairs_subject_name'" indexes:"" comment:"事项主题名称" json:"affairs_subject_name,omitempty" bson:",omitempty" msgpack:"affairs_subject_name,omitempty"`
	Affairs_subject_prompt *string    `xorm:" varchar(255) null 'affairs_subject_prompt'" indexes:"" comment:"事项主题简称" json:"affairs_subject_prompt,omitempty" bson:",omitempty" msgpack:"affairs_subject_prompt,omitempty"`
	Affairs_subject_code   *string    `xorm:" varchar(255) null 'affairs_subject_code'" indexes:" unique(Index_1) " comment:"事项主题编码" json:"affairs_subject_code,omitempty" bson:",omitempty" msgpack:"affairs_subject_code,omitempty"`
	Icon_url               *string    `xorm:" varchar(255) null 'icon_url'" indexes:"" comment:"图标" json:"icon_url,omitempty" bson:",omitempty" msgpack:"icon_url,omitempty"`
	Enable_mark            *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_2) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark            *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_3) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code              *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description            *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date            *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid          *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username        *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date            *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid          *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username        *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version                *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: tab_affairs_agent
// 说明: 事项受理-代理人
// 事项受理-代理人
// 
type tab_affairs_agent struct {
	Affairs_agent_id  *int64     `xorm:" pk not null 'affairs_agent_id'" indexes:"" comment:"事项代理人主键" json:"affairs_agent_id,omitempty" bson:",omitempty" msgpack:"affairs_agent_id,omitempty"`
	Affairs_info_id   *int64     `xorm:" null 'affairs_info_id'" indexes:"" comment:"受理主键" json:"affairs_info_id,omitempty" bson:",omitempty" msgpack:"affairs_info_id,omitempty" fk:"tab_affairs_info(affairs_info_id)"`
	Agent_id          *int64     `xorm:" null 'agent_id'" indexes:"" comment:"代理人主键" json:"agent_id,omitempty" bson:",omitempty" msgpack:"agent_id,omitempty"`
	Agent_name        *string    `xorm:" varchar(255) null 'agent_name'" indexes:"" comment:"代理人名称" json:"agent_name,omitempty" bson:",omitempty" msgpack:"agent_name,omitempty"`
	Agent_type        *int64     `xorm:" null 'agent_type'" indexes:"" comment:"代理人类型" json:"agent_type,omitempty" bson:",omitempty" msgpack:"agent_type,omitempty"`
	Agent_gender      *int64     `xorm:" null 'agent_gender'" indexes:"" comment:"代理人性别" json:"agent_gender,omitempty" bson:",omitempty" msgpack:"agent_gender,omitempty"`
	Credential_type   *int64     `xorm:" null 'credential_type'" indexes:"" comment:"代理人证件类型" json:"credential_type,omitempty" bson:",omitempty" msgpack:"credential_type,omitempty"`
	Credential_code   *string    `xorm:" varchar(50) null 'credential_code'" indexes:" index(Index_1) " comment:"代理人证件号码" json:"credential_code,omitempty" bson:",omitempty" msgpack:"credential_code,omitempty"`
	Legal_person_id   *int64     `xorm:" null 'legal_person_id'" indexes:"" comment:"法定代表人主键" json:"legal_person_id,omitempty" bson:",omitempty" msgpack:"legal_person_id,omitempty"`
	Legal_person_name *string    `xorm:" varchar(50) null 'legal_person_name'" indexes:"" comment:"法定代表人姓名" json:"legal_person_name,omitempty" bson:",omitempty" msgpack:"legal_person_name,omitempty"`
	Phone             *string    `xorm:" varchar(50) null 'phone'" indexes:"" comment:"联系电话" json:"phone,omitempty" bson:",omitempty" msgpack:"phone,omitempty"`
	Post_address      *string    `xorm:" varchar(255) null 'post_address'" indexes:"" comment:"通信地址" json:"post_address,omitempty" bson:",omitempty" msgpack:"post_address,omitempty"`
	Fixed_phone       *string    `xorm:" varchar(50) null 'fixed_phone'" indexes:"" comment:"固定电话" json:"fixed_phone,omitempty" bson:",omitempty" msgpack:"fixed_phone,omitempty"`
	Reg_address       *string    `xorm:" varchar(50) null 'reg_address'" indexes:"" comment:"户籍地址" json:"reg_address,omitempty" bson:",omitempty" msgpack:"reg_address,omitempty"`
	Bithday           *time.Time `xorm:" null 'bithday'" indexes:"" comment:"出生日期" json:"bithday,omitempty" bson:",omitempty" msgpack:"bithday,omitempty"`
	Age               *int64     `xorm:" null 'age'" indexes:"" comment:"年龄" json:"age,omitempty" bson:",omitempty" msgpack:"age,omitempty"`
	Postcode          *string    `xorm:" varchar(50) null 'postcode'" indexes:"" comment:"邮编" json:"postcode,omitempty" bson:",omitempty" msgpack:"postcode,omitempty"`
	Potrait           *string    `xorm:" varchar(50) null 'potrait'" indexes:"" comment:"代理人头像" json:"potrait,omitempty" bson:",omitempty" msgpack:"potrait,omitempty"`
	Enable_mark       *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_2) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark       *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_3) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code         *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description       *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date       *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid     *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username   *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date       *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid     *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username   *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version           *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: tab_dedicate_limit
// 说明: 专线设置-专线虚拟机限定
// 专线设置-专线虚拟机限定
// 
type tab_dedicate_limit struct {
	Limit_id        *int64     `xorm:" pk not null 'limit_id'" indexes:"" comment:"专线虚拟机限定主键" json:"limit_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"limit_id,omitempty"`
	Vm_id           *int64     `xorm:" null 'vm_id'" indexes:"" comment:"虚拟机主键" json:"vm_id,omitempty" bson:",omitempty" msgpack:"vm_id,omitempty" fk:"tab_virtual_machines(vm_id)"`
	Enable_mark     *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_1) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark     *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_2) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code       *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description     *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date     *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid   *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date     *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid   *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version         *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: tab_enterprise_reg_info
// 说明: 企业登记信息
// 企业登记信息
// 
type tab_enterprise_reg_info struct {
	Enterprise_reg_id *int64     `xorm:" pk not null 'enterprise_reg_id'" indexes:"" comment:"企业登记信息主键" json:"enterprise_reg_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"enterprise_reg_id,omitempty"`
	Credit_code       *string    `xorm:" varchar(255) null 'credit_code'" indexes:"" comment:"社会统一信用代码" json:"credit_code,omitempty" bson:",omitempty" msgpack:"credit_code,omitempty"`
	Org_code          *string    `xorm:" varchar(50) null 'org_code'" indexes:"" comment:"企业住所/经营地址" json:"org_code,omitempty" bson:",omitempty" msgpack:"org_code,omitempty"`
	Org_name          *string    `xorm:" varchar(50) null 'org_name'" indexes:"" comment:"经营范围" json:"org_name,omitempty" bson:",omitempty" msgpack:"org_name,omitempty"`
	Document_type     *int64     `xorm:" null 'document_type'" indexes:"" comment:"许可经营项目" json:"document_type,omitempty" bson:",omitempty" msgpack:"document_type,omitempty"`
	Reg_capital       *int64     `xorm:" null 'reg_capital'" indexes:"" comment:"注册资本" json:"reg_capital,omitempty" bson:",omitempty" msgpack:"reg_capital,omitempty"`
	Fact_capital      *int64     `xorm:" null 'fact_capital'" indexes:"" comment:"实收资本" json:"fact_capital,omitempty" bson:",omitempty" msgpack:"fact_capital,omitempty"`
	Reg_address       *string    `xorm:" varchar(255) null 'reg_address'" indexes:"" comment:"企业类型代码" json:"reg_address,omitempty" bson:",omitempty" msgpack:"reg_address,omitempty"`
	Reg_date          *string    `xorm:" varchar(255) null 'reg_date'" indexes:"" comment:"行业类型" json:"reg_date,omitempty" bson:",omitempty" msgpack:"reg_date,omitempty"`
	Org_status        *int64     `xorm:" null 'org_status'" indexes:"" comment:"行业小类" json:"org_status,omitempty" bson:",omitempty" msgpack:"org_status,omitempty"`
	Founded_date      *time.Time `xorm:" null 'founded_date'" indexes:"" comment:"成立日期" json:"founded_date,omitempty" bson:",omitempty" msgpack:"founded_date,omitempty"`
	Term_start        *time.Time `xorm:" null 'term_start'" indexes:"" comment:"营业期限起" json:"term_start,omitempty" bson:",omitempty" msgpack:"term_start,omitempty"`
	Term_end          *time.Time `xorm:" null 'term_end'" indexes:"" comment:"营业期限止" json:"term_end,omitempty" bson:",omitempty" msgpack:"term_end,omitempty"`
	Approval_date     *time.Time `xorm:" null 'approval_date'" indexes:"" comment:"核准日期" json:"approval_date,omitempty" bson:",omitempty" msgpack:"approval_date,omitempty"`
	Reg_sector        *string    `xorm:" varchar(255) null 'reg_sector'" indexes:"" comment:"登记机关" json:"reg_sector,omitempty" bson:",omitempty" msgpack:"reg_sector,omitempty"`
	Phone             *string    `xorm:" varchar(50) null 'phone'" indexes:"" comment:"联系电话" json:"phone,omitempty" bson:",omitempty" msgpack:"phone,omitempty"`
	Enable_mark       *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_1) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark       *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_2) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code         *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description       *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date       *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid     *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username   *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date       *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid     *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username   *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version           *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: tab_flow_node_info
// 说明: 流程模板-流程节点表
// 流程模板-流程节点表
// 
type tab_flow_node_info struct {
	Node_id              *int64     `xorm:" pk not null 'node_id'" indexes:"" comment:"节点编号" json:"node_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"node_id,omitempty"`
	Node_name            *string    `xorm:" varchar(255) null 'node_name'" indexes:"" comment:"节点名称" json:"node_name,omitempty" bson:",omitempty" msgpack:"node_name,omitempty"`
	Node_type            *int64     `xorm:" null 'node_type'" indexes:"" comment:"节点类型" json:"node_type,omitempty" bson:",omitempty" msgpack:"node_type,omitempty" fk:"dic_node_type(node_type_id)"`
	Node_standard_code   *int64     `xorm:" null 'node_standard_code'" indexes:"" comment:"流程标准编码" json:"node_standard_code,omitempty" bson:",omitempty" msgpack:"node_standard_code,omitempty" fk:"dic_node_standard_code(standard_code_id)"`
	Template_id          *int64     `xorm:" null 'template_id'" indexes:"" comment:"模版主键" json:"template_id,omitempty" bson:",omitempty" msgpack:"template_id,omitempty" fk:"tab_flow_template(template_id)"`
	Administrative_level *int64     `xorm:" null 'administrative_level'" indexes:"" comment:"行政级别" json:"administrative_level,omitempty" bson:",omitempty" msgpack:"administrative_level,omitempty"`
	Authority_type       *int64     `xorm:" null 'authority_type'" indexes:"" comment:"权限类型" json:"authority_type,omitempty" bson:",omitempty" msgpack:"authority_type,omitempty"`
	Position_x           *int64     `xorm:" null 'position_x'" indexes:"" comment:"坐标水平位置" json:"position_x,omitempty" bson:",omitempty" msgpack:"position_x,omitempty"`
	Position_y           *int64     `xorm:" null 'position_y'" indexes:"" comment:"坐标垂直位置" json:"position_y,omitempty" bson:",omitempty" msgpack:"position_y,omitempty"`
	Auto_notify          *Boolean   `xorm:" null default 0 'auto_notify'" indexes:"" comment:"是否自动发送短信" json:"auto_notify,omitempty" bson:",omitempty" msgpack:"auto_notify,omitempty"`
	Message_template     *string    `xorm:" varchar(500) null 'message_template'" indexes:"" comment:"短信内容模板" json:"message_template,omitempty" bson:",omitempty" msgpack:"message_template,omitempty"`
	Time_limited         *int64     `xorm:" null 'time_limited'" indexes:"" comment:"操作时限" json:"time_limited,omitempty" bson:",omitempty" msgpack:"time_limited,omitempty"`
	Time_limited_type    *int64     `xorm:" null default 1 'time_limited_type'" indexes:"" comment:"操作时限类型" json:"time_limited_type,omitempty" bson:",omitempty" msgpack:"time_limited_type,omitempty" fk:"dic_time_limit_type(time_limit_type_id)"`
	Timing_reset         *Boolean   `xorm:" null 'timing_reset'" indexes:"" comment:"补正后否重新计时" json:"timing_reset,omitempty" bson:",omitempty" msgpack:"timing_reset,omitempty"`
	Enable_mark          *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_1) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark          *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_2) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code            *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description          *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date          *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid        *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username      *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date          *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid        *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username      *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version              *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: tab_system_user_map
// 说明: 用户中心-系统用户对应表
// 用户中心-系统用户对应表
// 
type tab_system_user_map struct {
	Map_id          *int64     `xorm:" pk not null 'map_id'" indexes:"" comment:"映射主键" json:"map_id,omitempty" bson:",omitempty" msgpack:"map_id,omitempty"`
	System_id       *int64     `xorm:" null 'system_id'" indexes:" unique(Index_1) " comment:"系统主键" json:"system_id,omitempty" bson:",omitempty" msgpack:"system_id,omitempty" fk:"tab_system_list(system_id)"`
	User_id         *int64     `xorm:" null 'user_id'" indexes:"" comment:"用户主键" json:"user_id,omitempty" bson:",omitempty" msgpack:"user_id,omitempty" fk:"tab_user(user_id)"`
	System_user_id  *string    `xorm:" varchar(50) not null 'system_user_id'" indexes:" unique(Index_1) " comment:"系统用户主键" json:"system_user_id,omitempty" bson:",omitempty" msgpack:"system_user_id,omitempty"`
	Enable_mark     *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_2) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark     *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_3) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code       *int64     `xorm:" null 'sort_code'" indexes:" index(Index_4) " comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description     *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date     *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid   *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date     *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid   *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version         *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: dic_nation
// 说明: 自治家园字典-民族
// 自治家园字典-民族
// 
type dic_nation struct {
	Nation_id       *int64     `xorm:" pk not null 'nation_id'" indexes:"" comment:"民族主键" json:"nation_id,omitempty" bson:",omitempty" msgpack:"nation_id,omitempty"`
	Nation_name     *string    `xorm:" varchar(255) null 'nation_name'" indexes:"" comment:"名称" json:"nation_name,omitempty" bson:",omitempty" msgpack:"nation_name,omitempty"`
	Nation_prompt   *string    `xorm:" varchar(255) null 'nation_prompt'" indexes:"" comment:"别称" json:"nation_prompt,omitempty" bson:",omitempty" msgpack:"nation_prompt,omitempty"`
	Nation_code     *string    `xorm:" varchar(255) null 'nation_code'" indexes:"" comment:"民族编码" json:"nation_code,omitempty" bson:",omitempty" msgpack:"nation_code,omitempty"`
	Enable_mark     *Boolean   `xorm:" null default 1 'enable_mark'" indexes:"" comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark     *Boolean   `xorm:" null default 0 'delete_mark'" indexes:"" comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code       *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description     *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date     *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid   *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date     *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid   *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version         *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: dic_not_transfer_registered_cause
// 说明: 自治家园字典-未迁户口原因
// 自治家园字典-未迁户口原因
// 
type dic_not_transfer_registered_cause struct {
	Not_transfer_registered_cause_id     *int64     `xorm:" pk not null 'not_transfer_registered_cause_id'" indexes:"" comment:"未迁户口原因主键" json:"not_transfer_registered_cause_id,omitempty" bson:",omitempty" msgpack:"not_transfer_registered_cause_id,omitempty"`
	Not_transfer_registered_cause_name   *string    `xorm:" varchar(255) null 'not_transfer_registered_cause_name'" indexes:"" comment:"未迁户口原因名称" json:"not_transfer_registered_cause_name,omitempty" bson:",omitempty" msgpack:"not_transfer_registered_cause_name,omitempty"`
	Not_transfer_registered_cause_prompt *string    `xorm:" varchar(255) null 'not_transfer_registered_cause_prompt'" indexes:"" comment:"未迁户口原因别称" json:"not_transfer_registered_cause_prompt,omitempty" bson:",omitempty" msgpack:"not_transfer_registered_cause_prompt,omitempty"`
	Not_transfer_registered_cause_code   *string    `xorm:" varchar(255) null 'not_transfer_registered_cause_code'" indexes:"" comment:"未迁户口原因编码" json:"not_transfer_registered_cause_code,omitempty" bson:",omitempty" msgpack:"not_transfer_registered_cause_code,omitempty"`
	Enable_mark                          *Boolean   `xorm:" null default 1 'enable_mark'" indexes:"" comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark                          *Boolean   `xorm:" null default 0 'delete_mark'" indexes:"" comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code                            *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description                          *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date                          *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid                        *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username                      *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date                          *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid                        *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username                      *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version                              *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: tab_affairs_outcome
// 说明: 事项对象结果物定义
// 事项对象结果物定义
// 
type tab_affairs_outcome struct {
	Outcome_id       *int64     `xorm:" pk not null 'outcome_id'" indexes:"" comment:"事项结果物主键" json:"outcome_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"outcome_id,omitempty"`
	Affairs_id       *int64     `xorm:" null 'affairs_id'" indexes:"" comment:"事项主键" json:"affairs_id,omitempty" bson:",omitempty" msgpack:"affairs_id,omitempty" fk:"tab_affairs(affairs_id)"`
	Material_info_id *int64     `xorm:" null 'material_info_id'" indexes:"" comment:"材料主键" json:"material_info_id,omitempty" bson:",omitempty" msgpack:"material_info_id,omitempty" fk:"dic_material_info(material_info_id)"`
	Object_id        *int64     `xorm:" null 'object_id'" indexes:"" comment:"对象主键" json:"object_id,omitempty" bson:",omitempty" msgpack:"object_id,omitempty" fk:"tab_material_objects(object_id)"`
	Enable_mark      *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_1) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark      *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_2) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code        *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description      *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date      *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid    *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username  *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date      *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid    *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username  *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version          *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: tab_affairs_result
// 说明: 事项受理-办结信息
// 事项受理-办结信息
// 
type tab_affairs_result struct {
	Affairs_result_id          *int64     `xorm:" pk not null 'affairs_result_id'" indexes:"" comment:"事项办结主键" json:"affairs_result_id,omitempty" bson:",omitempty" msgpack:"affairs_result_id,omitempty"`
	Affairs_info_id            *int64     `xorm:" null 'affairs_info_id'" indexes:"" comment:"受理主键" json:"affairs_info_id,omitempty" bson:",omitempty" msgpack:"affairs_info_id,omitempty" fk:"tab_affairs_info(affairs_info_id)"`
	Affairs_info_code          *string    `xorm:" varchar(255) null 'affairs_info_code'" indexes:"" comment:"受理流水号" json:"affairs_info_code,omitempty" bson:",omitempty" msgpack:"affairs_info_code,omitempty"`
	Affairs_name               *string    `xorm:" varchar(255) null 'affairs_name'" indexes:"" comment:"事项名称" json:"affairs_name,omitempty" bson:",omitempty" msgpack:"affairs_name,omitempty"`
	User_id                    *int64     `xorm:" null 'user_id'" indexes:"" comment:"受理人主键" json:"user_id,omitempty" bson:",omitempty" msgpack:"user_id,omitempty"`
	User_name                  *string    `xorm:" varchar(50) null 'user_name'" indexes:"" comment:"受理人姓名" json:"user_name,omitempty" bson:",omitempty" msgpack:"user_name,omitempty"`
	Region_id                  *int64     `xorm:" null 'region_id'" indexes:"" comment:"受理人归属地" json:"region_id,omitempty" bson:",omitempty" msgpack:"region_id,omitempty"`
	Region_name                *string    `xorm:" varchar(255) null 'region_name'" indexes:"" comment:"受理人归属地名称" json:"region_name,omitempty" bson:",omitempty" msgpack:"region_name,omitempty"`
	User_dept                  *string    `xorm:" varchar(255) null 'user_dept'" indexes:"" comment:"受理人所在部门" json:"user_dept,omitempty" bson:",omitempty" msgpack:"user_dept,omitempty"`
	Outcome_id                 *int64     `xorm:" null 'outcome_id'" indexes:"" comment:"结果物主键" json:"outcome_id,omitempty" bson:",omitempty" msgpack:"outcome_id,omitempty"`
	Outcome_name               *string    `xorm:" varchar(50) null 'outcome_name'" indexes:"" comment:"结果物名称" json:"outcome_name,omitempty" bson:",omitempty" msgpack:"outcome_name,omitempty"`
	Result_outcome_code        *string    `xorm:" varchar(50) null 'result_outcome_code'" indexes:"" comment:"结果证照编号" json:"result_outcome_code,omitempty" bson:",omitempty" msgpack:"result_outcome_code,omitempty"`
	Affairs_result_file_url    *string    `xorm:" varchar(255) null 'affairs_result_file_url'" indexes:"" comment:"办件结果电子文书" json:"affairs_result_file_url,omitempty" bson:",omitempty" msgpack:"affairs_result_file_url,omitempty"`
	Affairs_result_description *string    `xorm:" varchar(255) null 'affairs_result_description'" indexes:"" comment:"办件结果描述" json:"affairs_result_description,omitempty" bson:",omitempty" msgpack:"affairs_result_description,omitempty"`
	Result_valid_time          *int64     `xorm:" null 'result_valid_time'" indexes:"" comment:"结果时效" json:"result_valid_time,omitempty" bson:",omitempty" msgpack:"result_valid_time,omitempty"`
	Valid_time_measure         *int64     `xorm:" null 'valid_time_measure'" indexes:"" comment:"结果时效单位" json:"valid_time_measure,omitempty" bson:",omitempty" msgpack:"valid_time_measure,omitempty"`
	Affairs_start              *time.Time `xorm:" null 'affairs_start'" indexes:"" comment:"事项开始时间" json:"affairs_start,omitempty" bson:",omitempty" msgpack:"affairs_start,omitempty"`
	Affairs_finish             *time.Time `xorm:" null 'affairs_finish'" indexes:"" comment:"事项办结时间" json:"affairs_finish,omitempty" bson:",omitempty" msgpack:"affairs_finish,omitempty"`
	Enable_mark                *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_1) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark                *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_2) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code                  *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description                *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date                *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid              *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username            *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date                *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid              *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username            *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version                    *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: tab_operation_log
// 说明: 自治家园-居民信息操作记录表
// 自治家园-居民信息操作记录表
// 
type tab_operation_log struct {
	Operation_log_id                 *int64     `xorm:" pk not null 'operation_log_id'" indexes:"" comment:"居民信息操作主键" json:"operation_log_id,omitempty" bson:",omitempty" msgpack:"operation_log_id,omitempty"`
	Nation_id                        *int64     `xorm:" null 'nation_id'" indexes:"" comment:"民族主键" json:"nation_id,omitempty" bson:",omitempty" msgpack:"nation_id,omitempty"`
	Health_status_id                 *int64     `xorm:" null 'health_status_id'" indexes:"" comment:"健康状况主键" json:"health_status_id,omitempty" bson:",omitempty" msgpack:"health_status_id,omitempty"`
	Blood_type_id                    *int64     `xorm:" null 'blood_type_id'" indexes:"" comment:"血型主键" json:"blood_type_id,omitempty" bson:",omitempty" msgpack:"blood_type_id,omitempty"`
	Marriage_status_id               *int64     `xorm:" null 'marriage_status_id'" indexes:"" comment:"婚姻状况主键" json:"marriage_status_id,omitempty" bson:",omitempty" msgpack:"marriage_status_id,omitempty"`
	Married_type_id                  *int64     `xorm:" null 'married_type_id'" indexes:"" comment:"已婚分类主键" json:"married_type_id,omitempty" bson:",omitempty" msgpack:"married_type_id,omitempty"`
	Military_service_status_id       *int64     `xorm:" null 'military_service_status_id'" indexes:"" comment:"兵役状况主键" json:"military_service_status_id,omitempty" bson:",omitempty" msgpack:"military_service_status_id,omitempty"`
	Politics_status_id               *int64     `xorm:" null 'politics_status_id'" indexes:"" comment:"政治面貌主键" json:"politics_status_id,omitempty" bson:",omitempty" msgpack:"politics_status_id,omitempty"`
	Education_id                     *int64     `xorm:" null 'education_id'" indexes:"" comment:"学历主键" json:"education_id,omitempty" bson:",omitempty" msgpack:"education_id,omitempty"`
	Family_master_relation_id        *int64     `xorm:" null 'family_master_relation_id'" indexes:"" comment:"与户主关系主键" json:"family_master_relation_id,omitempty" bson:",omitempty" msgpack:"family_master_relation_id,omitempty"`
	Card_type_id                     *int64     `xorm:" null 'card_type_id'" indexes:"" comment:"证件类型主键" json:"card_type_id,omitempty" bson:",omitempty" msgpack:"card_type_id,omitempty"`
	Country_id                       *int64     `xorm:" null 'country_id'" indexes:"" comment:"国家主键" json:"country_id,omitempty" bson:",omitempty" msgpack:"country_id,omitempty"`
	Religion_id                      *int64     `xorm:" null 'religion_id'" indexes:"" comment:"宗教信仰" json:"religion_id,omitempty" bson:",omitempty" msgpack:"religion_id,omitempty"`
	Work_status_id                   *int64     `xorm:" null 'work_status_id'" indexes:"" comment:"从业状况主键" json:"work_status_id,omitempty" bson:",omitempty" msgpack:"work_status_id,omitempty"`
	Part_time_job_id                 *int64     `xorm:" null 'part_time_job_id'" indexes:"" comment:"社会兼职主键" json:"part_time_job_id,omitempty" bson:",omitempty" msgpack:"part_time_job_id,omitempty"`
	Speciality_id                    *int64     `xorm:" null 'speciality_id'" indexes:"" comment:"特长主键" json:"speciality_id,omitempty" bson:",omitempty" msgpack:"speciality_id,omitempty"`
	Census_register_status_id        *int64     `xorm:" null 'census_register_status_id'" indexes:"" comment:"户籍状态主键" json:"census_register_status_id,omitempty" bson:",omitempty" msgpack:"census_register_status_id,omitempty"`
	Not_transfer_registered_cause_id *int64     `xorm:" null 'not_transfer_registered_cause_id'" indexes:"" comment:"未迁户口原因主键" json:"not_transfer_registered_cause_id,omitempty" bson:",omitempty" msgpack:"not_transfer_registered_cause_id,omitempty"`
	Census_register_type_id          *int64     `xorm:" null 'census_register_type_id'" indexes:"" comment:"户口类别主键" json:"census_register_type_id,omitempty" bson:",omitempty" msgpack:"census_register_type_id,omitempty"`
	Work_id                          *int64     `xorm:" null 'work_id'" indexes:"" comment:"职业主键" json:"work_id,omitempty" bson:",omitempty" msgpack:"work_id,omitempty"`
	Unit_property_id                 *int64     `xorm:" null 'unit_property_id'" indexes:"" comment:"单位性质主键" json:"unit_property_id,omitempty" bson:",omitempty" msgpack:"unit_property_id,omitempty"`
	Unit_trade_id                    *int64     `xorm:" null 'unit_trade_id'" indexes:"" comment:"单位所属行业主键" json:"unit_trade_id,omitempty" bson:",omitempty" msgpack:"unit_trade_id,omitempty"`
	Duty_level_id                    *int64     `xorm:" null 'duty_level_id'" indexes:"" comment:"干部职务主键" json:"duty_level_id,omitempty" bson:",omitempty" msgpack:"duty_level_id,omitempty"`
	Technology_duty_id               *int64     `xorm:" null 'technology_duty_id'" indexes:"" comment:"专业技术职务主键" json:"technology_duty_id,omitempty" bson:",omitempty" msgpack:"technology_duty_id,omitempty"`
	Register_region_id               *int64     `xorm:" null 'register_region_id'" indexes:"" comment:"户籍地所属归属地" json:"register_region_id,omitempty" bson:",omitempty" msgpack:"register_region_id,omitempty"`
	Dwell_region_id                  *int64     `xorm:" null 'dwell_region_id'" indexes:"" comment:"居住地所属归属地" json:"dwell_region_id,omitempty" bson:",omitempty" msgpack:"dwell_region_id,omitempty"`
	Card_code                        *string    `xorm:" varchar(20) not null 'card_code'" indexes:"" comment:"证件号码" json:"card_code,omitempty" bson:",omitempty" msgpack:"card_code,omitempty"`
	Name                             *string    `xorm:" varchar(50) not null 'name'" indexes:"" comment:"姓名" json:"name,omitempty" bson:",omitempty" msgpack:"name,omitempty"`
	Former_name                      *string    `xorm:" varchar(50) null 'former_name'" indexes:"" comment:"曾用名" json:"former_name,omitempty" bson:",omitempty" msgpack:"former_name,omitempty"`
	Phone                            *string    `xorm:" varchar(20) null 'phone'" indexes:"" comment:"联系电话" json:"phone,omitempty" bson:",omitempty" msgpack:"phone,omitempty"`
	Mobile_phone                     *string    `xorm:" varchar(20) null 'mobile_phone'" indexes:"" comment:"手机号码" json:"mobile_phone,omitempty" bson:",omitempty" msgpack:"mobile_phone,omitempty"`
	Join_work_date                   *time.Time `xorm:" null 'join_work_date'" indexes:"" comment:"参加工作日期" json:"join_work_date,omitempty" bson:",omitempty" msgpack:"join_work_date,omitempty"`
	Birthday                         *time.Time `xorm:" null 'birthday'" indexes:"" comment:"出生日期" json:"birthday,omitempty" bson:",omitempty" msgpack:"birthday,omitempty"`
	Age                              *int64     `xorm:" null 'age'" indexes:"" comment:"年龄" json:"age,omitempty" bson:",omitempty" msgpack:"age,omitempty"`
	Birthplace                       *string    `xorm:" varchar(255) null 'birthplace'" indexes:"" comment:"出生地" json:"birthplace,omitempty" bson:",omitempty" msgpack:"birthplace,omitempty"`
	Reside_address                   *string    `xorm:" varchar(255) null 'reside_address'" indexes:"" comment:"居住地址" json:"reside_address,omitempty" bson:",omitempty" msgpack:"reside_address,omitempty"`
	Stature                          *int64     `xorm:" null 'stature'" indexes:"" comment:"身高,单位cm" json:"stature,omitempty" bson:",omitempty" msgpack:"stature,omitempty"`
	Email                            *string    `xorm:" varchar(255) null 'email'" indexes:"" comment:"电子邮箱" json:"email,omitempty" bson:",omitempty" msgpack:"email,omitempty"`
	Photo                            *string    `xorm:" varchar(255) null 'photo'" indexes:"" comment:"照片" json:"photo,omitempty" bson:",omitempty" msgpack:"photo,omitempty"`
	Tags                             *int64     `xorm:" null 'tags'" indexes:"" comment:"标签(所有标签用一个字段表示)" json:"tags,omitempty" bson:",omitempty" msgpack:"tags,omitempty"`
	Operation_mark                   *int64     `xorm:" not null 'operation_mark'" indexes:"" comment:"操作标记(改删改)" json:"operation_mark,omitempty" bson:",omitempty" msgpack:"operation_mark,omitempty"`
	Enable_mark                      *Boolean   `xorm:" null default 1 'enable_mark'" indexes:"" comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark                      *Boolean   `xorm:" null default 0 'delete_mark'" indexes:"" comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code                        *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description                      *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date                      *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid                    *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username                  *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date                      *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid                    *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username                  *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version                          *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: tab_prepare_affairs
// 说明: 外网-预审申请表
// 外网-预审申请表
// 
type tab_prepare_affairs struct {
	Prepare_affairs_id *int64     `xorm:" pk not null 'prepare_affairs_id'" indexes:"" comment:"预审申请主键" json:"prepare_affairs_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"prepare_affairs_id,omitempty"`
	Out_user_id        *int64     `xorm:" null 'out_user_id'" indexes:"" comment:"外网用户主键" json:"out_user_id,omitempty" bson:",omitempty" msgpack:"out_user_id,omitempty" fk:"tab_outer_user(out_user_id)"`
	Region_id          *int64     `xorm:" null 'region_id'" indexes:"" comment:"归属地" json:"region_id,omitempty" bson:",omitempty" msgpack:"region_id,omitempty"`
	Region_name        *string    `xorm:" varchar(255) null 'region_name'" indexes:"" comment:"归属地名称" json:"region_name,omitempty" bson:",omitempty" msgpack:"region_name,omitempty"`
	Affairs_id         *int64     `xorm:" null 'affairs_id'" indexes:"" comment:"事项主键" json:"affairs_id,omitempty" bson:",omitempty" msgpack:"affairs_id,omitempty"`
	Affairs_info_code  *string    `xorm:" varchar(255) null 'affairs_info_code'" indexes:"" comment:"事项流水号" json:"affairs_info_code,omitempty" bson:",omitempty" msgpack:"affairs_info_code,omitempty"`
	Affairs_name       *string    `xorm:" varchar(255) null 'affairs_name'" indexes:"" comment:"事项名称" json:"affairs_name,omitempty" bson:",omitempty" msgpack:"affairs_name,omitempty"`
	Policy_basis       *string    `xorm:" text null 'policy_basis'" indexes:"" comment:"政策依据" json:"policy_basis,omitempty" bson:",omitempty" msgpack:"policy_basis,omitempty"`
	Material_needed    *string    `xorm:" text null 'material_needed'" indexes:"" comment:"所需材料" json:"material_needed,omitempty" bson:",omitempty" msgpack:"material_needed,omitempty"`
	Accept_department  *string    `xorm:" text null 'accept_department'" indexes:"" comment:"受理部门" json:"accept_department,omitempty" bson:",omitempty" msgpack:"accept_department,omitempty"`
	Check_standard     *string    `xorm:" text null 'check_standard'" indexes:"" comment:"审查标准" json:"check_standard,omitempty" bson:",omitempty" msgpack:"check_standard,omitempty"`
	Affairs_subject    *string    `xorm:" varchar(255) null 'affairs_subject'" indexes:"" comment:"事项主题" json:"affairs_subject,omitempty" bson:",omitempty" msgpack:"affairs_subject,omitempty"`
	Affairs_object     *string    `xorm:" varchar(255) null 'affairs_object'" indexes:"" comment:"事项对象" json:"affairs_object,omitempty" bson:",omitempty" msgpack:"affairs_object,omitempty"`
	Department         *string    `xorm:" varchar(255) null 'department'" indexes:"" comment:"部门" json:"department,omitempty" bson:",omitempty" msgpack:"department,omitempty"`
	Life_event         *string    `xorm:" varchar(255) null 'life_event'" indexes:"" comment:"人生事件" json:"life_event,omitempty" bson:",omitempty" msgpack:"life_event,omitempty"`
	Service_subject    *string    `xorm:" varchar(255) null 'service_subject'" indexes:"" comment:"服务对象" json:"service_subject,omitempty" bson:",omitempty" msgpack:"service_subject,omitempty"`
	Check_result       *Boolean   `xorm:" null 'check_result'" indexes:"" comment:"审核结果" json:"check_result,omitempty" bson:",omitempty" msgpack:"check_result,omitempty"`
	Check_mess         *string    `xorm:" text null 'check_mess'" indexes:"" comment:"反馈信息" json:"check_mess,omitempty" bson:",omitempty" msgpack:"check_mess,omitempty"`
	Affairs_info_id    *int64     `xorm:" null 'affairs_info_id'" indexes:"" comment:"事项受理主键" json:"affairs_info_id,omitempty" bson:",omitempty" msgpack:"affairs_info_id,omitempty"`
	Enable_mark        *Boolean   `xorm:" null default 1 'enable_mark'" indexes:"" comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark        *Boolean   `xorm:" null default 0 'delete_mark'" indexes:"" comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code          *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description        *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date        *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid      *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username    *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date        *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid      *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username    *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version            *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: tab_monitor_hook_log
// 说明: 行政监察挂起记录表
// 行政监察挂起记录表
// 
type tab_monitor_hook_log struct {
	Hook_id         *int64     `xorm:" pk not null 'hook_id'" indexes:"" comment:"挂起操作主键" json:"hook_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"hook_id,omitempty"`
	Monitor_id      *int64     `xorm:" null 'monitor_id'" indexes:"" comment:"监察主键" json:"monitor_id,omitempty" bson:",omitempty" msgpack:"monitor_id,omitempty" fk:"tab_monitors(monitor_id)"`
	Hook_type       *int64     `xorm:" null 'hook_type'" indexes:"" comment:"挂起类型" json:"hook_type,omitempty" bson:",omitempty" msgpack:"hook_type,omitempty"`
	Enable_mark     *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_1) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark     *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_2) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code       *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description     *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date     *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid   *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date     *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid   *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version         *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: tab_outer_consultation
// 说明: 外网-咨询表
// 外网-咨询表
// 
type tab_outer_consultation struct {
	Consultation_id   *int64     `xorm:" pk not null 'consultation_id'" indexes:"" comment:"咨询主键" json:"consultation_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"consultation_id,omitempty"`
	Out_user_id       *int64     `xorm:" null 'out_user_id'" indexes:"" comment:"外网用户主键" json:"out_user_id,omitempty" bson:",omitempty" msgpack:"out_user_id,omitempty" fk:"tab_outer_user(out_user_id)"`
	Title             *string    `xorm:" varchar(255) null 'title'" indexes:"" comment:"标题" json:"title,omitempty" bson:",omitempty" msgpack:"title,omitempty"`
	Consultation_type *string    `xorm:" varchar(255) null 'consultation_type'" indexes:"" comment:"咨询类型" json:"consultation_type,omitempty" bson:",omitempty" msgpack:"consultation_type,omitempty"`
	Is_public         *Boolean   `xorm:" null default 1 'is_public'" indexes:"" comment:"是否公开" json:"is_public,omitempty" bson:",omitempty" msgpack:"is_public,omitempty"`
	Content           *string    `xorm:" text null 'content'" indexes:"" comment:"咨询内容" json:"content,omitempty" bson:",omitempty" msgpack:"content,omitempty"`
	Enable_mark       *Boolean   `xorm:" null default 1 'enable_mark'" indexes:"" comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark       *Boolean   `xorm:" null default 0 'delete_mark'" indexes:"" comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code         *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description       *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date       *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid     *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username   *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date       *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid     *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username   *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version           *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: tab_affair_material_agent
// 说明: 事项受理-代理人材料-包含结果物
// 事项受理-代理人材料-包含结果物
// 
type tab_affair_material_agent struct {
	Affairs_material_id     *int64     `xorm:" pk not null 'affairs_material_id'" indexes:"" comment:"事项受理材料主键" json:"affairs_material_id,omitempty" bson:",omitempty" msgpack:"affairs_material_id,omitempty"`
	Affairs_agent_id        *int64     `xorm:" null 'affairs_agent_id'" indexes:"" comment:"事项申请人主键" json:"affairs_agent_id,omitempty" bson:",omitempty" msgpack:"affairs_agent_id,omitempty" fk:"tab_affairs_agent(affairs_agent_id)"`
	Material_provider_id    *int64     `xorm:" null 'material_provider_id'" indexes:"" comment:"材料提供人主键" json:"material_provider_id,omitempty" bson:",omitempty" msgpack:"material_provider_id,omitempty"`
	Material_info_id        *int64     `xorm:" null 'material_info_id'" indexes:"" comment:"材料主键" json:"material_info_id,omitempty" bson:",omitempty" msgpack:"material_info_id,omitempty"`
	Material_name           *string    `xorm:" varchar(255) null 'material_name'" indexes:"" comment:"材料名称" json:"material_name,omitempty" bson:",omitempty" msgpack:"material_name,omitempty"`
	Material_file_id        *string    `xorm:" varchar(255) null 'material_file_id'" indexes:"" comment:"材料文件ID" json:"material_file_id,omitempty" bson:",omitempty" msgpack:"material_file_id,omitempty"`
	Material_file_url       *string    `xorm:" varchar(255) null 'material_file_url'" indexes:"" comment:"材料文件位置" json:"material_file_url,omitempty" bson:",omitempty" msgpack:"material_file_url,omitempty"`
	Material_upload_type_id *int64     `xorm:" not null 'material_upload_type_id'" indexes:"" comment:"上传操作类型" json:"material_upload_type_id,omitempty" bson:",omitempty" msgpack:"material_upload_type_id,omitempty" fk:"dic_material_upload_type(material_upload_id)"`
	Material_type_id        *int64     `xorm:" null 'material_type_id'" indexes:"" comment:"材料类型" json:"material_type_id,omitempty" bson:",omitempty" msgpack:"material_type_id,omitempty" fk:"dic_material_type(material_type_id)"`
	Affairs_info_id         *int64     `xorm:" null 'affairs_info_id'" indexes:"" comment:"受理主键" json:"affairs_info_id,omitempty" bson:",omitempty" msgpack:"affairs_info_id,omitempty" fk:"tab_affairs_info(affairs_info_id)"`
	Material_region_id      *int64     `xorm:" null 'material_region_id'" indexes:"" comment:"材料提交归属地主键" json:"material_region_id,omitempty" bson:",omitempty" msgpack:"material_region_id,omitempty"`
	Material_region_name    *string    `xorm:" varchar(255) null 'material_region_name'" indexes:"" comment:"材料提交归属地名称" json:"material_region_name,omitempty" bson:",omitempty" msgpack:"material_region_name,omitempty"`
	Material_valid_term     *time.Time `xorm:" null 'material_valid_term'" indexes:"" comment:"材料到期时间" json:"material_valid_term,omitempty" bson:",omitempty" msgpack:"material_valid_term,omitempty"`
	Material_relative_field *string    `xorm:" varchar(255) null 'material_relative_field'" indexes:"" comment:"材料相关字段" json:"material_relative_field,omitempty" bson:",omitempty" msgpack:"material_relative_field,omitempty"`
	Gathered                *Boolean   `xorm:" null 'gathered'" indexes:"" comment:"是否已收取" json:"gathered,omitempty" bson:",omitempty" msgpack:"gathered,omitempty"`
	Gathered_mode           *int64     `xorm:" null 'gathered_mode'" indexes:"" comment:"收取方式" json:"gathered_mode,omitempty" bson:",omitempty" msgpack:"gathered_mode,omitempty"`
	Gathered_number         *int64     `xorm:" null 'gathered_number'" indexes:"" comment:"收取数量" json:"gathered_number,omitempty" bson:",omitempty" msgpack:"gathered_number,omitempty"`
	Gathered_time           *time.Time `xorm:" null 'gathered_time'" indexes:"" comment:"收取时间" json:"gathered_time,omitempty" bson:",omitempty" msgpack:"gathered_time,omitempty"`
	Return_success          *Boolean   `xorm:" null default 0 'return_success'" indexes:"" comment:"成功时退回" json:"return_success,omitempty" bson:",omitempty" msgpack:"return_success,omitempty"`
	Return_failed           *Boolean   `xorm:" null default 0 'return_failed'" indexes:"" comment:"不成功退回" json:"return_failed,omitempty" bson:",omitempty" msgpack:"return_failed,omitempty"`
	Outcome_tag             *Boolean   `xorm:" not null default 0 'outcome_tag'" indexes:"" comment:"是否结果物" json:"outcome_tag,omitempty" bson:",omitempty" msgpack:"outcome_tag,omitempty"`
	Signer_date             *time.Time `xorm:" null 'signer_date'" indexes:"" comment:"取回时间" json:"signer_date,omitempty" bson:",omitempty" msgpack:"signer_date,omitempty"`
	Signer_name             *string    `xorm:" varchar(50) null 'signer_name'" indexes:"" comment:"取回人签字" json:"signer_name,omitempty" bson:",omitempty" msgpack:"signer_name,omitempty"`
	Signer_photo            *string    `xorm:" varchar(255) null 'signer_photo'" indexes:"" comment:"签字人照片" json:"signer_photo,omitempty" bson:",omitempty" msgpack:"signer_photo,omitempty"`
	Enable_mark             *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_1) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark             *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_2) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code               *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description             *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date             *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid           *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username         *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date             *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid           *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username         *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version                 *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: tab_affairs_mark_result
// 说明: 事项受理-流程审核结果信息
// 事项受理-流程审核结果信息
// 
type tab_affairs_mark_result struct {
	Mark_id                      *int64     `xorm:" pk not null 'mark_id'" indexes:"" comment:"审核主键" json:"mark_id,omitempty" bson:",omitempty" msgpack:"mark_id,omitempty"`
	User_id                      *int64     `xorm:" null 'user_id'" indexes:"" comment:"审核人主键" json:"user_id,omitempty" bson:",omitempty" msgpack:"user_id,omitempty" fk:"tab_user(user_id)"`
	User_name                    *string    `xorm:" varchar(50) null 'user_name'" indexes:"" comment:"审核人名称" json:"user_name,omitempty" bson:",omitempty" msgpack:"user_name,omitempty"`
	Node_name                    *string    `xorm:" varchar(255) null 'node_name'" indexes:"" comment:"节点名称" json:"node_name,omitempty" bson:",omitempty" msgpack:"node_name,omitempty"`
	Node_result_name             *string    `xorm:" varchar(255) null 'node_result_name'" indexes:"" comment:"流程结果名称" json:"node_result_name,omitempty" bson:",omitempty" msgpack:"node_result_name,omitempty"`
	Node_id                      *int64     `xorm:" null 'node_id'" indexes:"" comment:"节点主键" json:"node_id,omitempty" bson:",omitempty" msgpack:"node_id,omitempty"`
	Node_result_id               *int64     `xorm:" null 'node_result_id'" indexes:"" comment:"结果节点主键" json:"node_result_id,omitempty" bson:",omitempty" msgpack:"node_result_id,omitempty" fk:"tab_flow_node_result_info(node_result_id)"`
	Node_start_time              *time.Time `xorm:" null 'node_start_time'" indexes:"" comment:"节点开始时间" json:"node_start_time,omitempty" bson:",omitempty" msgpack:"node_start_time,omitempty"`
	Node_end_time                *time.Time `xorm:" null 'node_end_time'" indexes:"" comment:"节点结束时间" json:"node_end_time,omitempty" bson:",omitempty" msgpack:"node_end_time,omitempty"`
	Mark_commit                  *string    `xorm:" text null 'mark_commit'" indexes:"" comment:"审核意见" json:"mark_commit,omitempty" bson:",omitempty" msgpack:"mark_commit,omitempty"`
	Manual_node                  *Boolean   `xorm:" null default 1 'manual_node'" indexes:"" comment:"是否人工节点" json:"manual_node,omitempty" bson:",omitempty" msgpack:"manual_node,omitempty"`
	Node_type_id                 *int64     `xorm:" null 'node_type_id'" indexes:"" comment:"节点类型" json:"node_type_id,omitempty" bson:",omitempty" msgpack:"node_type_id,omitempty" fk:"dic_node_type(node_type_id)"`
	Standard_code_id             *int64     `xorm:" null 'standard_code_id'" indexes:"" comment:"节点标准编码" json:"standard_code_id,omitempty" bson:",omitempty" msgpack:"standard_code_id,omitempty" fk:"dic_node_standard_code(standard_code_id)"`
	Enable_mark                  *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_1) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Next_user_id                 *int64     `xorm:" null 'next_user_id'" indexes:"" comment:"下一步经办人主键" json:"next_user_id,omitempty" bson:",omitempty" msgpack:"next_user_id,omitempty" fk:"tab_user(user_id)"`
	Next_user_name               *string    `xorm:" varchar(50) null 'next_user_name'" indexes:"" comment:"下一步经办人姓名" json:"next_user_name,omitempty" bson:",omitempty" msgpack:"next_user_name,omitempty"`
	Next_attribution_admin_level *int64     `xorm:" null 'next_attribution_admin_level'" indexes:"" comment:"下一步归属地行政级别" json:"next_attribution_admin_level,omitempty" bson:",omitempty" msgpack:"next_attribution_admin_level,omitempty" fk:"dic_administrative_level(level_id)"`
	Next_attribution_area_id     *int64     `xorm:" null 'next_attribution_area_id'" indexes:"" comment:"下一步归属地行政区域主键" json:"next_attribution_area_id,omitempty" bson:",omitempty" msgpack:"next_attribution_area_id,omitempty"`
	Next_attribution_area_name   *string    `xorm:" varchar(255) null 'next_attribution_area_name'" indexes:"" comment:"下一步归属地行政区域名称" json:"next_attribution_area_name,omitempty" bson:",omitempty" msgpack:"next_attribution_area_name,omitempty"`
	Next_attribution_id          *int64     `xorm:" null 'next_attribution_id'" indexes:"" comment:"下一步归属地主键" json:"next_attribution_id,omitempty" bson:",omitempty" msgpack:"next_attribution_id,omitempty" fk:"tab_region(region_id)"`
	Affairs_info_id              *int64     `xorm:" null 'affairs_info_id'" indexes:"" comment:"受理主键" json:"affairs_info_id,omitempty" bson:",omitempty" msgpack:"affairs_info_id,omitempty" fk:"tab_affairs_info(affairs_info_id)"`
	Next_attribution_name        *string    `xorm:" varchar(255) null 'next_attribution_name'" indexes:"" comment:"下一步归属地名称" json:"next_attribution_name,omitempty" bson:",omitempty" msgpack:"next_attribution_name,omitempty"`
	Delete_mark                  *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_2) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code                    *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description                  *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date                  *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid                *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username              *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date                  *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid                *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username              *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version                      *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: tab_client_version
// 说明: WPF客户端版本
// WPF客户端版本
// 
type tab_client_version struct {
	Version_id      *int64     `xorm:" pk not null 'version_id'" indexes:"" comment:"客户端版本主键" json:"version_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"version_id,omitempty"`
	Version_no      *float32   `xorm:" null 'version_no'" indexes:"" comment:"版本号" json:"version_no,omitempty" bson:",omitempty" msgpack:"version_no,omitempty"`
	Changes         *string    `xorm:" text null 'changes'" indexes:"" comment:"更新内容" json:"changes,omitempty" bson:",omitempty" msgpack:"changes,omitempty"`
	Storage_url     *string    `xorm:" varchar(255) null 'storage_url'" indexes:"" comment:"存储位置" json:"storage_url,omitempty" bson:",omitempty" msgpack:"storage_url,omitempty"`
	Current_version *Boolean   `xorm:" null default 1 'current_version'" indexes:"" comment:"是否当前使用版本" json:"current_version,omitempty" bson:",omitempty" msgpack:"current_version,omitempty"`
	Enable_mark     *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_1) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark     *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_2) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code       *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description     *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date     *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid   *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date     *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid   *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version         *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: tab_coporation_info
// 说明: 法人信息表
// 法人信息表
// 
type tab_coporation_info struct {
	Coporation_id        *int64     `xorm:" pk not null 'coporation_id'" indexes:"" comment:"法人主键" json:"coporation_id,omitempty" bson:",omitempty" msgpack:"coporation_id,omitempty"`
	Credit_code          *string    `xorm:" varchar(255) not null 'credit_code'" indexes:"" comment:"社会统一信用代码" json:"credit_code,omitempty" bson:",omitempty" msgpack:"credit_code,omitempty"`
	Org_code             *string    `xorm:" varchar(50) null 'org_code'" indexes:"" comment:"组织机构代码" json:"org_code,omitempty" bson:",omitempty" msgpack:"org_code,omitempty"`
	Org_name             *string    `xorm:" varchar(50) null 'org_name'" indexes:"" comment:"机构名称" json:"org_name,omitempty" bson:",omitempty" msgpack:"org_name,omitempty"`
	Org_type             *int64     `xorm:" null 'org_type'" indexes:"" comment:"机构类型" json:"org_type,omitempty" bson:",omitempty" msgpack:"org_type,omitempty"`
	Legal_person         *string    `xorm:" varchar(50) null 'legal_person'" indexes:"" comment:"法定代表人" json:"legal_person,omitempty" bson:",omitempty" msgpack:"legal_person,omitempty"`
	Legal_person_card_no *string    `xorm:" varchar(255) null 'legal_person_card_no'" indexes:"" comment:"法定代表人身份证号" json:"legal_person_card_no,omitempty" bson:",omitempty" msgpack:"legal_person_card_no,omitempty"`
	Reg_address          *string    `xorm:" varchar(255) null 'reg_address'" indexes:"" comment:"注册地址" json:"reg_address,omitempty" bson:",omitempty" msgpack:"reg_address,omitempty"`
	Reg_date             *string    `xorm:" varchar(255) null 'reg_date'" indexes:"" comment:"注册日期" json:"reg_date,omitempty" bson:",omitempty" msgpack:"reg_date,omitempty"`
	Org_status           *int64     `xorm:" null 'org_status'" indexes:"" comment:"机构状态" json:"org_status,omitempty" bson:",omitempty" msgpack:"org_status,omitempty"`
	Enable_mark          *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_1) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark          *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_2) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code            *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description          *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date          *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid        *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username      *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date          *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid        *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username      *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version              *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: tab_group
// 说明: 分组表
// 分组表
// 
type tab_group struct {
	Group_id        *int64     `xorm:" pk not null 'group_id'" indexes:"" comment:"组的主键" json:"group_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"group_id,omitempty"`
	Group_name      *string    `xorm:" varchar(255) null 'group_name'" indexes:"" comment:"组名称" json:"group_name,omitempty" bson:",omitempty" msgpack:"group_name,omitempty"`
	Location_id     *int64     `xorm:" null 'location_id'" indexes:"" comment:"归属地" json:"location_id,omitempty" bson:",omitempty" msgpack:"location_id,omitempty"`
	Enable_mark     *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_1) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark     *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_2) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code       *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description     *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date     *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid   *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date     *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid   *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version         *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: tab_materials_check
// 说明: 事项受理-材料签入签出
// 事项受理-材料签入签出
// 
type tab_materials_check struct {
	Check_id            *int64     `xorm:" pk not null 'check_id'" indexes:"" comment:"签入签出主键" json:"check_id,omitempty" bson:",omitempty" msgpack:"check_id,omitempty"`
	Check_type          *Boolean   `xorm:" null 'check_type'" indexes:"" comment:"签入签出类型" json:"check_type,omitempty" bson:",omitempty" msgpack:"check_type,omitempty"`
	Check_bill_no       *string    `xorm:" varchar(50) null 'check_bill_no'" indexes:"" comment:"迁入迁出批次号" json:"check_bill_no,omitempty" bson:",omitempty" msgpack:"check_bill_no,omitempty"`
	Affairs_name        *string    `xorm:" varchar(50) null 'affairs_name'" indexes:"" comment:"事项名称" json:"affairs_name,omitempty" bson:",omitempty" msgpack:"affairs_name,omitempty"`
	Affairs_info_code   *string    `xorm:" varchar(50) null 'affairs_info_code'" indexes:"" comment:"事项受理流水号" json:"affairs_info_code,omitempty" bson:",omitempty" msgpack:"affairs_info_code,omitempty"`
	Attribution_id      *int64     `xorm:" null 'attribution_id'" indexes:"" comment:"操作归属地主键" json:"attribution_id,omitempty" bson:",omitempty" msgpack:"attribution_id,omitempty"`
	Attribution_name    *string    `xorm:" varchar(255) null 'attribution_name'" indexes:"" comment:"操作归属地名称" json:"attribution_name,omitempty" bson:",omitempty" msgpack:"attribution_name,omitempty"`
	Node_id             *int64     `xorm:" null 'node_id'" indexes:"" comment:"操作节点ID" json:"node_id,omitempty" bson:",omitempty" msgpack:"node_id,omitempty" fk:"tab_flow_node_info(node_id)"`
	Status_id           *int64     `xorm:" null 'status_id'" indexes:"" comment:"流转状态主键" json:"status_id,omitempty" bson:",omitempty" msgpack:"status_id,omitempty" fk:"tab_affairs_node_status(status_id)"`
	Is_flow_check       *Boolean   `xorm:" null default 1 'is_flow_check'" indexes:" index(Index_2) " comment:"是否流程签入签出" json:"is_flow_check,omitempty" bson:",omitempty" msgpack:"is_flow_check,omitempty"`
	Unflow_check_status *int64     `xorm:" null 'unflow_check_status'" indexes:"" comment:"非流程签入签出状态" json:"unflow_check_status,omitempty" bson:",omitempty" msgpack:"unflow_check_status,omitempty" fk:"dic_node_status(node_status_id)"`
	Affairs_info_id     *int64     `xorm:" null 'affairs_info_id'" indexes:"" comment:"受理主键" json:"affairs_info_id,omitempty" bson:",omitempty" msgpack:"affairs_info_id,omitempty" fk:"tab_affairs_info(affairs_info_id)"`
	Operator_id         *int64     `xorm:" null 'operator_id'" indexes:"" comment:"操作人主键" json:"operator_id,omitempty" bson:",omitempty" msgpack:"operator_id,omitempty"`
	Operator_name       *string    `xorm:" varchar(50) null 'operator_name'" indexes:"" comment:"操作人名称" json:"operator_name,omitempty" bson:",omitempty" msgpack:"operator_name,omitempty"`
	Operator_time       *time.Time `xorm:" null 'operator_time'" indexes:" index(Index_1) " comment:"操作时间" json:"operator_time,omitempty" bson:",omitempty" msgpack:"operator_time,omitempty"`
	Express_company_id  *int64     `xorm:" null 'express_company_id'" indexes:"" comment:"快递公司主键" json:"express_company_id,omitempty" bson:",omitempty" msgpack:"express_company_id,omitempty"`
	Express_company     *string    `xorm:" varchar(50) null 'express_company'" indexes:"" comment:"快递公司" json:"express_company,omitempty" bson:",omitempty" msgpack:"express_company,omitempty"`
	Express_bill_no     *string    `xorm:" varchar(50) null 'express_bill_no'" indexes:"" comment:"快递单号" json:"express_bill_no,omitempty" bson:",omitempty" msgpack:"express_bill_no,omitempty"`
	Enable_mark         *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_3) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark         *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_4) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code           *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description         *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date         *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid       *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username     *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date         *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid       *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username     *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version             *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: tab_region_events
// 说明: 查询机-大事记
// 查询机-大事记
// 
type tab_region_events struct {
	Event_id        *int64     `xorm:" pk not null 'event_id'" indexes:"" comment:"大事记主键" json:"event_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"event_id,omitempty"`
	Region_id       *int64     `xorm:" null 'region_id'" indexes:"" comment:"归属地主键" json:"region_id,omitempty" bson:",omitempty" msgpack:"region_id,omitempty" fk:"tab_region(region_id)"`
	Event_date      *time.Time `xorm:" null 'event_date'" indexes:"" comment:"大事记日期" json:"event_date,omitempty" bson:",omitempty" msgpack:"event_date,omitempty"`
	Event_content   *string    `xorm:" text null 'event_content'" indexes:"" comment:"大事记内容" json:"event_content,omitempty" bson:",omitempty" msgpack:"event_content,omitempty"`
	Photo_url       *string    `xorm:" varchar(255) null 'photo_url'" indexes:"" comment:"附属照片" json:"photo_url,omitempty" bson:",omitempty" msgpack:"photo_url,omitempty"`
	Delete_mark     *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_1) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code       *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description     *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date     *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid   *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date     *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid   *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version         *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: dic_universal_range
// 说明: 字典-通办范围
// 字典-通办范围
// 
type dic_universal_range struct {
	Universal_range_id     *int64     `xorm:" pk not null 'universal_range_id'" indexes:"" comment:"通办范围主键" json:"universal_range_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"universal_range_id,omitempty"`
	Universal_range_name   *string    `xorm:" varchar(255) null 'universal_range_name'" indexes:"" comment:"通办范围名称" json:"universal_range_name,omitempty" bson:",omitempty" msgpack:"universal_range_name,omitempty"`
	Universal_range_prompt *string    `xorm:" varchar(255) null 'universal_range_prompt'" indexes:"" comment:"通办范围简称" json:"universal_range_prompt,omitempty" bson:",omitempty" msgpack:"universal_range_prompt,omitempty"`
	Universal_range_code   *string    `xorm:" varchar(255) null 'universal_range_code'" indexes:" unique(Index_1) " comment:"通办范围编码" json:"universal_range_code,omitempty" bson:",omitempty" msgpack:"universal_range_code,omitempty"`
	Enable_mark            *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_2) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark            *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_3) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code              *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description            *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date            *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid          *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username        *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date            *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid          *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username        *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version                *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: tab_affairs_node_current
// 说明: 事项受理-当前待审核结点
// 事项受理-当前待审核结点
// 
type tab_affairs_node_current struct {
	Current_id       *int64     `xorm:" pk not null 'current_id'" indexes:"" comment:"当前待审核节点主键" json:"current_id,omitempty" bson:",omitempty" msgpack:"current_id,omitempty"`
	Node_id          *int64     `xorm:" null 'node_id'" indexes:" index(Index_1) " comment:"节点主键" json:"node_id,omitempty" bson:",omitempty" msgpack:"node_id,omitempty" fk:"tab_flow_node_info(node_id)"`
	Special_user_id  *int64     `xorm:" null 'special_user_id'" indexes:"" comment:"经办人" json:"special_user_id,omitempty" bson:",omitempty" msgpack:"special_user_id,omitempty" fk:"tab_user(user_id)"`
	Special_region   *int64     `xorm:" null 'special_region'" indexes:"" comment:"归属地" json:"special_region,omitempty" bson:",omitempty" msgpack:"special_region,omitempty" fk:"tab_region(region_id)"`
	Special_group_id *int64     `xorm:" null 'special_group_id'" indexes:"" comment:"用户组" json:"special_group_id,omitempty" bson:",omitempty" msgpack:"special_group_id,omitempty" fk:"tab_group(group_id)"`
	Affairs_info_id  *int64     `xorm:" null 'affairs_info_id'" indexes:"" comment:"受理主键" json:"affairs_info_id,omitempty" bson:",omitempty" msgpack:"affairs_info_id,omitempty" fk:"tab_affairs_info(affairs_info_id)"`
	Enable_mark      *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_2) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark      *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_3) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code        *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description      *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date      *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid    *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username  *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date      *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid    *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username  *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version          *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: dic_country
// 说明: 自治家园字典-国家
// 自治家园字典-国家
// 
type dic_country struct {
	Country_id      *int64     `xorm:" pk not null 'country_id'" indexes:"" comment:"国家主键" json:"country_id,omitempty" bson:",omitempty" msgpack:"country_id,omitempty"`
	Country_name    *string    `xorm:" varchar(255) null 'country_name'" indexes:"" comment:"国家名称" json:"country_name,omitempty" bson:",omitempty" msgpack:"country_name,omitempty"`
	Country_prompt  *string    `xorm:" varchar(255) null 'country_prompt'" indexes:"" comment:"国家别称" json:"country_prompt,omitempty" bson:",omitempty" msgpack:"country_prompt,omitempty"`
	Country_code    *string    `xorm:" varchar(255) null 'country_code'" indexes:"" comment:"国家编码" json:"country_code,omitempty" bson:",omitempty" msgpack:"country_code,omitempty"`
	Enable_mark     *Boolean   `xorm:" null default 1 'enable_mark'" indexes:"" comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark     *Boolean   `xorm:" null default 0 'delete_mark'" indexes:"" comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code       *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description     *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date     *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid   *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date     *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid   *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version         *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: dic_material_info
// 说明: 字典-材料定义
// 字典-材料定义
// 
type dic_material_info struct {
	Material_info_id    *int64     `xorm:" pk not null 'material_info_id'" indexes:"" comment:"材料主键" json:"material_info_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"material_info_id,omitempty"`
	Material_code       *string    `xorm:" varchar(50) null 'material_code'" indexes:" unique(Index_2) " comment:"材料编号" json:"material_code,omitempty" bson:",omitempty" msgpack:"material_code,omitempty"`
	Material_name       *string    `xorm:" varchar(255) null 'material_name'" indexes:" unique(Index_1) " comment:"材料名称" json:"material_name,omitempty" bson:",omitempty" msgpack:"material_name,omitempty"`
	Material_alias      *string    `xorm:" varchar(255) null 'material_alias'" indexes:"" comment:"材料本地别名" json:"material_alias,omitempty" bson:",omitempty" msgpack:"material_alias,omitempty"`
	Material_prompt     *string    `xorm:" varchar(255) null 'material_prompt'" indexes:"" comment:"材料简称" json:"material_prompt,omitempty" bson:",omitempty" msgpack:"material_prompt,omitempty"`
	Material_doc_url    *string    `xorm:" varchar(255) null 'material_doc_url'" indexes:"" comment:"材料样式" json:"material_doc_url,omitempty" bson:",omitempty" msgpack:"material_doc_url,omitempty"`
	Material_sample_url *string    `xorm:" varchar(255) null 'material_sample_url'" indexes:"" comment:"材料示例" json:"material_sample_url,omitempty" bson:",omitempty" msgpack:"material_sample_url,omitempty"`
	Reuse_type_id       *int64     `xorm:" not null default 5 'reuse_type_id'" indexes:"" comment:"复用类型" json:"reuse_type_id,omitempty" bson:",omitempty" msgpack:"reuse_type_id,omitempty" fk:"dic_reuse_type(reuse_type_id)"`
	Enable_mark         *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_3) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark         *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_4) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code           *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description         *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date         *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid       *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username     *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date         *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid       *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username     *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version             *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: dic_material_upload_type
// 说明: 字典-材料上传类型
// 字典-材料上传类型
// 
type dic_material_upload_type struct {
	Material_upload_id          *int64     `xorm:" pk not null 'material_upload_id'" indexes:"" comment:"材料上传类型主键" json:"material_upload_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"material_upload_id,omitempty"`
	Material_upload_type_name   *string    `xorm:" varchar(50) null 'material_upload_type_name'" indexes:"" comment:"材料上传类型名称" json:"material_upload_type_name,omitempty" bson:",omitempty" msgpack:"material_upload_type_name,omitempty"`
	Material_upload_type_prompt *string    `xorm:" varchar(50) null 'material_upload_type_prompt'" indexes:"" comment:"材料上传类型前台名称" json:"material_upload_type_prompt,omitempty" bson:",omitempty" msgpack:"material_upload_type_prompt,omitempty"`
	Value_code                  *string    `xorm:" varchar(50) null 'value_code'" indexes:" unique(Index_1) " comment:"编码" json:"value_code,omitempty" bson:",omitempty" msgpack:"value_code,omitempty"`
	Enable_mark                 *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_2) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark                 *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_3) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code                   *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description                 *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date                 *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid               *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username             *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date                 *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid               *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username             *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version                     *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: tab_affairs
// 说明: 事项定义表
// 事项定义表
// 
type tab_affairs struct {
	Affairs_id              *int64     `xorm:" pk not null 'affairs_id'" indexes:"" comment:"事项主键" json:"affairs_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"affairs_id,omitempty"`
	Affairs_local_id        *string    `xorm:" varchar(100) null 'affairs_local_id'" indexes:" index(Index_3) " comment:"地方事项编号" json:"affairs_local_id,omitempty" bson:",omitempty" msgpack:"affairs_local_id,omitempty"`
	Affairs_global_id       *string    `xorm:" varchar(100) null 'affairs_global_id'" indexes:" index(Index_4) " comment:"事项全国唯一编号" json:"affairs_global_id,omitempty" bson:",omitempty" msgpack:"affairs_global_id,omitempty"`
	Implement_code          *string    `xorm:" varchar(100) null 'implement_code'" indexes:"" comment:"实施编码" json:"implement_code,omitempty" bson:",omitempty" msgpack:"implement_code,omitempty"`
	Affairs_simple_name     *string    `xorm:" varchar(255) null 'affairs_simple_name'" indexes:"" comment:"事项简称" json:"affairs_simple_name,omitempty" bson:",omitempty" msgpack:"affairs_simple_name,omitempty"`
	Affairs_name            *string    `xorm:" varchar(255) null 'affairs_name'" indexes:"" comment:"事项名称" json:"affairs_name,omitempty" bson:",omitempty" msgpack:"affairs_name,omitempty"`
	Affairs_type            *int64     `xorm:" null 'affairs_type'" indexes:"" comment:"事项类型" json:"affairs_type,omitempty" bson:",omitempty" msgpack:"affairs_type,omitempty" fk:"dic_affairs_type(affairs_type_id)"`
	Setup_basis             *string    `xorm:" text null 'setup_basis'" indexes:"" comment:"设定依据" json:"setup_basis,omitempty" bson:",omitempty" msgpack:"setup_basis,omitempty"`
	Exercise_level          *int64     `xorm:" null 'exercise_level'" indexes:"" comment:"行使层级" json:"exercise_level,omitempty" bson:",omitempty" msgpack:"exercise_level,omitempty"`
	Authority_division      *string    `xorm:" text null 'authority_division'" indexes:"" comment:"权限划分" json:"authority_division,omitempty" bson:",omitempty" msgpack:"authority_division,omitempty"`
	Exercise_content        *string    `xorm:" text null 'exercise_content'" indexes:"" comment:"行使内容" json:"exercise_content,omitempty" bson:",omitempty" msgpack:"exercise_content,omitempty"`
	Implement_sector        *string    `xorm:" text null 'implement_sector'" indexes:"" comment:"实施机构" json:"implement_sector,omitempty" bson:",omitempty" msgpack:"implement_sector,omitempty"`
	Subject_nature          *int64     `xorm:" null 'subject_nature'" indexes:"" comment:"实施主体性质" json:"subject_nature,omitempty" bson:",omitempty" msgpack:"subject_nature,omitempty" fk:"dic_subject_nature(subject_nature_id)"`
	Law_finish_limit        *int64     `xorm:" null 'law_finish_limit'" indexes:"" comment:"法定办结时限" json:"law_finish_limit,omitempty" bson:",omitempty" msgpack:"law_finish_limit,omitempty"`
	Law_accept_condistion   *string    `xorm:" text null 'law_accept_condistion'" indexes:"" comment:"法定受理条件" json:"law_accept_condistion,omitempty" bson:",omitempty" msgpack:"law_accept_condistion,omitempty"`
	Union_sector            *string    `xorm:" text null 'union_sector'" indexes:"" comment:"联办机构" json:"union_sector,omitempty" bson:",omitempty" msgpack:"union_sector,omitempty"`
	Inter_service           *string    `xorm:" text null 'inter_service'" indexes:"" comment:"中介服务" json:"inter_service,omitempty" bson:",omitempty" msgpack:"inter_service,omitempty"`
	Flow_chart_url          *string    `xorm:" varchar(255) null 'flow_chart_url'" indexes:"" comment:"办理流程图" json:"flow_chart_url,omitempty" bson:",omitempty" msgpack:"flow_chart_url,omitempty"`
	Flow_chart              *string    `xorm:" text null 'flow_chart'" indexes:"" comment:"办理流程" json:"flow_chart,omitempty" bson:",omitempty" msgpack:"flow_chart,omitempty"`
	Number_limit            *int64     `xorm:" null 'number_limit'" indexes:"" comment:"数量限制" json:"number_limit,omitempty" bson:",omitempty" msgpack:"number_limit,omitempty"`
	Result_name             *string    `xorm:" varchar(255) null 'result_name'" indexes:"" comment:"结果名称" json:"result_name,omitempty" bson:",omitempty" msgpack:"result_name,omitempty"`
	Result_sample_url       *string    `xorm:" varchar(255) null 'result_sample_url'" indexes:"" comment:"结果样本" json:"result_sample_url,omitempty" bson:",omitempty" msgpack:"result_sample_url,omitempty"`
	Fee_tag                 *Boolean   `xorm:" null default 0 'fee_tag'" indexes:"" comment:"是否收费" json:"fee_tag,omitempty" bson:",omitempty" msgpack:"fee_tag,omitempty"`
	Fee_standard            *float32   `xorm:" null 'fee_standard'" indexes:"" comment:"收费标准" json:"fee_standard,omitempty" bson:",omitempty" msgpack:"fee_standard,omitempty"`
	Fee_standard_text       *string    `xorm:" text null 'fee_standard_text'" indexes:"" comment:"收费标准说明" json:"fee_standard_text,omitempty" bson:",omitempty" msgpack:"fee_standard_text,omitempty"`
	Fee_basis               *string    `xorm:" text null 'fee_basis'" indexes:"" comment:"收费依据" json:"fee_basis,omitempty" bson:",omitempty" msgpack:"fee_basis,omitempty"`
	Service_subject         *int64     `xorm:" not null 'service_subject'" indexes:"" comment:"服务对象" json:"service_subject,omitempty" bson:",omitempty" msgpack:"service_subject,omitempty" fk:"dic_proposer_type(proposer_type_id)"`
	Affair_type             *Boolean   `xorm:" null default 0 'affair_type'" indexes:"" comment:"办件类型" json:"affair_type,omitempty" bson:",omitempty" msgpack:"affair_type,omitempty"`
	Affair_promise_type     *Boolean   `xorm:" null default 0 'affair_promise_type'" indexes:"" comment:"承诺类型" json:"affair_promise_type,omitempty" bson:",omitempty" msgpack:"affair_promise_type,omitempty"`
	Is_consultation         *Boolean   `xorm:" null default 0 'is_consultation'" indexes:"" comment:"是否咨询件" json:"is_consultation,omitempty" bson:",omitempty" msgpack:"is_consultation,omitempty"`
	Promise_finish_limit    *int64     `xorm:" null default 1 'promise_finish_limit'" indexes:"" comment:"承诺办结时限" json:"promise_finish_limit,omitempty" bson:",omitempty" msgpack:"promise_finish_limit,omitempty"`
	Time_limit_type_id      *int64     `xorm:" not null default 1 'time_limit_type_id'" indexes:"" comment:"承诺办结时限类型" json:"time_limit_type_id,omitempty" bson:",omitempty" msgpack:"time_limit_type_id,omitempty" fk:"dic_time_limit_type(time_limit_type_id)"`
	Universal_range         *int64     `xorm:" null default 4 'universal_range'" indexes:"" comment:"通办范围" json:"universal_range,omitempty" bson:",omitempty" msgpack:"universal_range,omitempty" fk:"dic_universal_range(universal_range_id)"`
	Handling_form           *int64     `xorm:" null default 0 'handling_form'" indexes:"" comment:"办理形式" json:"handling_form,omitempty" bson:",omitempty" msgpack:"handling_form,omitempty" fk:"dic_handling_form(handling_form_id)"`
	Online_support          *Boolean   `xorm:" null default 0 'online_support'" indexes:"" comment:"是否网上办理" json:"online_support,omitempty" bson:",omitempty" msgpack:"online_support,omitempty"`
	Material_retention      *Boolean   `xorm:" null default 1 'material_retention'" indexes:"" comment:"物料是否留存" json:"material_retention,omitempty" bson:",omitempty" msgpack:"material_retention,omitempty"`
	Booking_support         *Boolean   `xorm:" null default 0 'booking_support'" indexes:"" comment:"是否支持网上预审" json:"booking_support,omitempty" bson:",omitempty" msgpack:"booking_support,omitempty"`
	Pay_online_support      *Boolean   `xorm:" null default 0 'pay_online_support'" indexes:"" comment:"是否支持网上支付" json:"pay_online_support,omitempty" bson:",omitempty" msgpack:"pay_online_support,omitempty"`
	Express_support         *Boolean   `xorm:" null default 0 'express_support'" indexes:"" comment:"是否支持物流快递" json:"express_support,omitempty" bson:",omitempty" msgpack:"express_support,omitempty"`
	Running_system          *int64     `xorm:" null 'running_system'" indexes:"" comment:"运行系统" json:"running_system,omitempty" bson:",omitempty" msgpack:"running_system,omitempty"`
	Transact_place          *string    `xorm:" text null 'transact_place'" indexes:"" comment:"办理地点" json:"transact_place,omitempty" bson:",omitempty" msgpack:"transact_place,omitempty"`
	Transact_time           *string    `xorm:" text null 'transact_time'" indexes:"" comment:"办理时间" json:"transact_time,omitempty" bson:",omitempty" msgpack:"transact_time,omitempty"`
	Consultation_phone      *string    `xorm:" text null 'consultation_phone'" indexes:"" comment:"咨询电话" json:"consultation_phone,omitempty" bson:",omitempty" msgpack:"consultation_phone,omitempty"`
	Faqs                    *string    `xorm:" text null 'faqs'" indexes:"" comment:"常见问题" json:"faqs,omitempty" bson:",omitempty" msgpack:"faqs,omitempty"`
	Accept_condition        *string    `xorm:" text null 'accept_condition'" indexes:"" comment:"受理条件" json:"accept_condition,omitempty" bson:",omitempty" msgpack:"accept_condition,omitempty"`
	Flow_description        *string    `xorm:" text null 'flow_description'" indexes:"" comment:"内部流程描述" json:"flow_description,omitempty" bson:",omitempty" msgpack:"flow_description,omitempty"`
	Affairs_guide           *string    `xorm:" text null 'affairs_guide'" indexes:"" comment:"办事指南" json:"affairs_guide,omitempty" bson:",omitempty" msgpack:"affairs_guide,omitempty"`
	Supervision_phone       *string    `xorm:" varchar(50) null 'supervision_phone'" indexes:"" comment:"监督电话" json:"supervision_phone,omitempty" bson:",omitempty" msgpack:"supervision_phone,omitempty"`
	Power_update_type       *int64     `xorm:" null 'power_update_type'" indexes:"" comment:"权力更新类型" json:"power_update_type,omitempty" bson:",omitempty" msgpack:"power_update_type,omitempty"`
	Version_code            *string    `xorm:" varchar(50) null 'version_code'" indexes:"" comment:"版本号" json:"version_code,omitempty" bson:",omitempty" msgpack:"version_code,omitempty"`
	Valided_date            *time.Time `xorm:" null 'valided_date'" indexes:"" comment:"版本生效时间" json:"valided_date,omitempty" bson:",omitempty" msgpack:"valided_date,omitempty"`
	Power_status            *int64     `xorm:" null 'power_status'" indexes:"" comment:"权力状态" json:"power_status,omitempty" bson:",omitempty" msgpack:"power_status,omitempty"`
	Notice_list_id          *string    `xorm:" varchar(255) null 'notice_list_id'" indexes:"" comment:"告知单ID" json:"notice_list_id,omitempty" bson:",omitempty" msgpack:"notice_list_id,omitempty"`
	Policy_basis            *string    `xorm:" text null 'policy_basis'" indexes:"" comment:"政策依据" json:"policy_basis,omitempty" bson:",omitempty" msgpack:"policy_basis,omitempty"`
	Material_needed         *string    `xorm:" text null 'material_needed'" indexes:"" comment:"所需材料" json:"material_needed,omitempty" bson:",omitempty" msgpack:"material_needed,omitempty"`
	Department_id           *int64     `xorm:" not null 'department_id'" indexes:"" comment:"部门主键" json:"department_id,omitempty" bson:",omitempty" msgpack:"department_id,omitempty" fk:"dic_department(department_id)"`
	Accept_department       *string    `xorm:" text null 'accept_department'" indexes:"" comment:"受理机构" json:"accept_department,omitempty" bson:",omitempty" msgpack:"accept_department,omitempty"`
	Check_standard          *string    `xorm:" text null 'check_standard'" indexes:"" comment:"审查标准" json:"check_standard,omitempty" bson:",omitempty" msgpack:"check_standard,omitempty"`
	Affairs_subject_id      *int64     `xorm:" null 'affairs_subject_id'" indexes:"" comment:"自然人事项主题" json:"affairs_subject_id,omitempty" bson:",omitempty" msgpack:"affairs_subject_id,omitempty" fk:"dic_affairs_subject(affairs_subject_id)"`
	Corporate_subject_id    *int64     `xorm:" null 'corporate_subject_id'" indexes:"" comment:"法人事项主题" json:"corporate_subject_id,omitempty" bson:",omitempty" msgpack:"corporate_subject_id,omitempty" fk:"dic_corporate_subject(corporate_subject_id)"`
	Affairs_object_id       *int64     `xorm:" null 'affairs_object_id'" indexes:"" comment:"事项对象" json:"affairs_object_id,omitempty" bson:",omitempty" msgpack:"affairs_object_id,omitempty" fk:"dic_affairs_objects(affairs_object_id)"`
	Life_event_id           *int64     `xorm:" null 'life_event_id'" indexes:"" comment:"人生事件" json:"life_event_id,omitempty" bson:",omitempty" msgpack:"life_event_id,omitempty" fk:"dic_life_events(life_event_id)"`
	Template_id             *int64     `xorm:" null 'template_id'" indexes:"" comment:"当前启用模板" json:"template_id,omitempty" bson:",omitempty" msgpack:"template_id,omitempty"`
	Net_level               *int64     `xorm:" null 'net_level'" indexes:"" comment:"网上办理深度" json:"net_level,omitempty" bson:",omitempty" msgpack:"net_level,omitempty" fk:"dic_net_level(net_level_id)"`
	Progress_result_query   *string    `xorm:" text null 'progress_result_query'" indexes:"" comment:"办理进程和结果公开查询" json:"progress_result_query,omitempty" bson:",omitempty" msgpack:"progress_result_query,omitempty"`
	Consultation_path       *string    `xorm:" text null 'consultation_path'" indexes:"" comment:"咨询途径" json:"consultation_path,omitempty" bson:",omitempty" msgpack:"consultation_path,omitempty"`
	Supervision_path        *string    `xorm:" text null 'supervision_path'" indexes:"" comment:"监督投诉渠道" json:"supervision_path,omitempty" bson:",omitempty" msgpack:"supervision_path,omitempty"`
	Person_rights           *string    `xorm:" text null 'person_rights'" indexes:"" comment:"行政相对人权利和义务" json:"person_rights,omitempty" bson:",omitempty" msgpack:"person_rights,omitempty"`
	Confirm_department      *string    `xorm:" text null 'confirm_department'" indexes:"" comment:"决定机构" json:"confirm_department,omitempty" bson:",omitempty" msgpack:"confirm_department,omitempty"`
	Prohibitive_requirement *string    `xorm:" text null 'prohibitive_requirement'" indexes:"" comment:"禁止性要求" json:"prohibitive_requirement,omitempty" bson:",omitempty" msgpack:"prohibitive_requirement,omitempty"`
	Result_receive          *int64     `xorm:" null 'result_receive'" indexes:"" comment:"结果送达" json:"result_receive,omitempty" bson:",omitempty" msgpack:"result_receive,omitempty" fk:"dic_result_receive(result_receive_id)"`
	Correct_timing          *Boolean   `xorm:" null 'correct_timing'" indexes:"" comment:"补正是否计时" json:"correct_timing,omitempty" bson:",omitempty" msgpack:"correct_timing,omitempty"`
	Correct_extend          *Boolean   `xorm:" null 'correct_extend'" indexes:"" comment:"补正是否延时" json:"correct_extend,omitempty" bson:",omitempty" msgpack:"correct_extend,omitempty"`
	Enable_mark             *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_1) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark             *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_2) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code               *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description             *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date             *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid           *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username         *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date             *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid           *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username         *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version                 *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: dic_result_receive
// 说明: 字典-结果送达方式
// 字典-结果送达方式
// 
type dic_result_receive struct {
	Result_receive_id     *int64     `xorm:" pk not null 'result_receive_id'" indexes:"" comment:"送达方式主键" json:"result_receive_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"result_receive_id,omitempty"`
	Result_receive_name   *string    `xorm:" varchar(255) null 'result_receive_name'" indexes:" unique(Index_1) " comment:"送达方式名称" json:"result_receive_name,omitempty" bson:",omitempty" msgpack:"result_receive_name,omitempty"`
	Result_receive_prompt *string    `xorm:" varchar(255) null 'result_receive_prompt'" indexes:"" comment:"送达方式简称" json:"result_receive_prompt,omitempty" bson:",omitempty" msgpack:"result_receive_prompt,omitempty"`
	Result_receive_code   *string    `xorm:" varchar(50) null 'result_receive_code'" indexes:" unique(Index_2) " comment:"送达方式编码" json:"result_receive_code,omitempty" bson:",omitempty" msgpack:"result_receive_code,omitempty"`
	Enable_mark           *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_3) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark           *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_4) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code             *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description           *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date           *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid         *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username       *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date           *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid         *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username       *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version               *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: tab_affairs_node_results
// 说明: 事项受理-已审核结点结果
// 事项受理-已审核结点结果
// 
type tab_affairs_node_results struct {
	Results_id      *int64     `xorm:" pk not null 'results_id'" indexes:"" comment:"已审核节点结果主键" json:"results_id,omitempty" bson:",omitempty" msgpack:"results_id,omitempty"`
	Node_result_id  *int64     `xorm:" null 'node_result_id'" indexes:" index(Index_1) " comment:"流程结果主键" json:"node_result_id,omitempty" bson:",omitempty" msgpack:"node_result_id,omitempty" fk:"tab_flow_node_result_info(node_result_id)"`
	Affairs_info_id *int64     `xorm:" null 'affairs_info_id'" indexes:"" comment:"受理主键" json:"affairs_info_id,omitempty" bson:",omitempty" msgpack:"affairs_info_id,omitempty" fk:"tab_affairs_info(affairs_info_id)"`
	Enable_mark     *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_2) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark     *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_3) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code       *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description     *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date     *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid   *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date     *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid   *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version         *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: tab_material_relation
// 说明: 材料关系表
// 材料关系表
// 
type tab_material_relation struct {
	Material_relation_id    *int64     `xorm:" pk not null 'material_relation_id'" indexes:"" comment:"材料关系主键" json:"material_relation_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"material_relation_id,omitempty"`
	Object_id               *int64     `xorm:" not null 'object_id'" indexes:"" comment:"对象主键" json:"object_id,omitempty" bson:",omitempty" msgpack:"object_id,omitempty" fk:"tab_material_objects(object_id)"`
	Affairs_id              *int64     `xorm:" null 'affairs_id'" indexes:"" comment:"事项主键" json:"affairs_id,omitempty" bson:",omitempty" msgpack:"affairs_id,omitempty" fk:"tab_affairs(affairs_id)"`
	Material_info_id        *int64     `xorm:" not null 'material_info_id'" indexes:"" comment:"材料主键" json:"material_info_id,omitempty" bson:",omitempty" msgpack:"material_info_id,omitempty" fk:"dic_material_info(material_info_id)"`
	Material_group          *int64     `xorm:" null 'material_group'" indexes:"" comment:"材料分组" json:"material_group,omitempty" bson:",omitempty" msgpack:"material_group,omitempty"`
	Reuse_type_id           *int64     `xorm:" null 'reuse_type_id'" indexes:"" comment:"复用类型" json:"reuse_type_id,omitempty" bson:",omitempty" msgpack:"reuse_type_id,omitempty" fk:"dic_reuse_type(reuse_type_id)"`
	Material_type_id        *int64     `xorm:" null 'material_type_id'" indexes:"" comment:"材料类型" json:"material_type_id,omitempty" bson:",omitempty" msgpack:"material_type_id,omitempty" fk:"dic_material_type(material_type_id)"`
	Original_check          *Boolean   `xorm:" null default 0 'original_check'" indexes:"" comment:"是否验原件" json:"original_check,omitempty" bson:",omitempty" msgpack:"original_check,omitempty"`
	Valid_term_measure      *int64     `xorm:" null 'valid_term_measure'" indexes:"" comment:"有效期限类型" json:"valid_term_measure,omitempty" bson:",omitempty" msgpack:"valid_term_measure,omitempty"`
	Valid_term              *int64     `xorm:" null 'valid_term'" indexes:"" comment:"有效期限" json:"valid_term,omitempty" bson:",omitempty" msgpack:"valid_term,omitempty"`
	Must_collect            *int64     `xorm:" null 'must_collect'" indexes:"" comment:"必交数量" json:"must_collect,omitempty" bson:",omitempty" msgpack:"must_collect,omitempty"`
	Material_type_changable *Boolean   `xorm:" null default 0 'material_type_changable'" indexes:"" comment:"材料类型是否可改变" json:"material_type_changable,omitempty" bson:",omitempty" msgpack:"material_type_changable,omitempty"`
	Return_success          *Boolean   `xorm:" null default 0 'return_success'" indexes:"" comment:"成功时退回" json:"return_success,omitempty" bson:",omitempty" msgpack:"return_success,omitempty"`
	Return_failed           *Boolean   `xorm:" null default 0 'return_failed'" indexes:"" comment:"不成功退回" json:"return_failed,omitempty" bson:",omitempty" msgpack:"return_failed,omitempty"`
	Outcome_tag             *Boolean   `xorm:" null default 0 'outcome_tag'" indexes:"" comment:"是否结果物" json:"outcome_tag,omitempty" bson:",omitempty" msgpack:"outcome_tag,omitempty"`
	Need_collected          *Boolean   `xorm:" null default 0 'need_collected'" indexes:"" comment:"需要收取" json:"need_collected,omitempty" bson:",omitempty" msgpack:"need_collected,omitempty"`
	Need_scaned             *Boolean   `xorm:" null default 0 'need_scaned'" indexes:"" comment:"需要扫描" json:"need_scaned,omitempty" bson:",omitempty" msgpack:"need_scaned,omitempty"`
	Material_doc_url        *string    `xorm:" varchar(255) null 'material_doc_url'" indexes:"" comment:"材料样式" json:"material_doc_url,omitempty" bson:",omitempty" msgpack:"material_doc_url,omitempty"`
	Material_sample_url     *string    `xorm:" varchar(255) null 'material_sample_url'" indexes:"" comment:"材料示例" json:"material_sample_url,omitempty" bson:",omitempty" msgpack:"material_sample_url,omitempty"`
	Check_point             *string    `xorm:" text null 'check_point'" indexes:"" comment:"审核要点说明" json:"check_point,omitempty" bson:",omitempty" msgpack:"check_point,omitempty"`
	Check_point_url         *string    `xorm:" varchar(255) null 'check_point_url'" indexes:"" comment:"审核要点ID" json:"check_point_url,omitempty" bson:",omitempty" msgpack:"check_point_url,omitempty"`
	Enable_mark             *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_1) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark             *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_2) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code               *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description             *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date             *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid           *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username         *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date             *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid           *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username         *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version                 *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: tab_prepare_materials
// 说明: 外网-附件及材料表
// 外网-附件及材料表
// 
type tab_prepare_materials struct {
	Prepare_materials_id    *int64     `xorm:" pk not null 'prepare_materials_id'" indexes:"" comment:"附件主键" json:"prepare_materials_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"prepare_materials_id,omitempty"`
	Prepare_consultation_id *int64     `xorm:" null 'prepare_consultation_id'" indexes:"" comment:"预审或咨询主键" json:"prepare_consultation_id,omitempty" bson:",omitempty" msgpack:"prepare_consultation_id,omitempty"`
	Materials_name          *string    `xorm:" varchar(255) null 'materials_name'" indexes:"" comment:"名称" json:"materials_name,omitempty" bson:",omitempty" msgpack:"materials_name,omitempty"`
	Material_info_id        *int64     `xorm:" null 'material_info_id'" indexes:"" comment:"材料主键" json:"material_info_id,omitempty" bson:",omitempty" msgpack:"material_info_id,omitempty"`
	Material_name           *string    `xorm:" varchar(255) null 'material_name'" indexes:"" comment:"材料名称" json:"material_name,omitempty" bson:",omitempty" msgpack:"material_name,omitempty"`
	Materials_url           *string    `xorm:" varchar(255) null 'materials_url'" indexes:"" comment:"文件ID" json:"materials_url,omitempty" bson:",omitempty" msgpack:"materials_url,omitempty"`
	Enable_mark             *Boolean   `xorm:" null default 1 'enable_mark'" indexes:"" comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark             *Boolean   `xorm:" null default 0 'delete_mark'" indexes:"" comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code               *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description             *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date             *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid           *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username         *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date             *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid           *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username         *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version                 *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: tab_check_bill
// 说明: 材料迁入迁出批次表
// 材料迁入迁出批次表
// 
type tab_check_bill struct {
	Check_bill_id      *int64     `xorm:" pk not null 'check_bill_id'" indexes:"" comment:"批次主键" json:"check_bill_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"check_bill_id,omitempty"`
	Check_bill_no      *string    `xorm:" varchar(50) null 'check_bill_no'" indexes:"" comment:"批次号" json:"check_bill_no,omitempty" bson:",omitempty" msgpack:"check_bill_no,omitempty"`
	Check_type         *Boolean   `xorm:" null 'check_type'" indexes:"" comment:"签入签出类型" json:"check_type,omitempty" bson:",omitempty" msgpack:"check_type,omitempty"`
	Attribution_id     *int64     `xorm:" null 'attribution_id'" indexes:"" comment:"操作归属地主键" json:"attribution_id,omitempty" bson:",omitempty" msgpack:"attribution_id,omitempty"`
	Attribution_name   *string    `xorm:" varchar(255) null 'attribution_name'" indexes:"" comment:"操作归属地名称" json:"attribution_name,omitempty" bson:",omitempty" msgpack:"attribution_name,omitempty"`
	Operator_id        *int64     `xorm:" null 'operator_id'" indexes:"" comment:"操作人主键" json:"operator_id,omitempty" bson:",omitempty" msgpack:"operator_id,omitempty"`
	Operator_name      *string    `xorm:" varchar(50) null 'operator_name'" indexes:"" comment:"操作人名称" json:"operator_name,omitempty" bson:",omitempty" msgpack:"operator_name,omitempty"`
	Operator_time      *time.Time `xorm:" null 'operator_time'" indexes:"" comment:"操作时间" json:"operator_time,omitempty" bson:",omitempty" msgpack:"operator_time,omitempty"`
	Express_company_id *int64     `xorm:" null 'express_company_id'" indexes:"" comment:"快递公司主键" json:"express_company_id,omitempty" bson:",omitempty" msgpack:"express_company_id,omitempty"`
	Express_company    *string    `xorm:" varchar(50) null 'express_company'" indexes:"" comment:"快递公司" json:"express_company,omitempty" bson:",omitempty" msgpack:"express_company,omitempty"`
	Express_bill_no    *string    `xorm:" varchar(50) null 'express_bill_no'" indexes:"" comment:"快递单号" json:"express_bill_no,omitempty" bson:",omitempty" msgpack:"express_bill_no,omitempty"`
	Affairs_count      *int64     `xorm:" null 'affairs_count'" indexes:"" comment:"事项受理总数" json:"affairs_count,omitempty" bson:",omitempty" msgpack:"affairs_count,omitempty"`
	Is_flow_check      *Boolean   `xorm:" null default 1 'is_flow_check'" indexes:"" comment:"是否流程签入签出" json:"is_flow_check,omitempty" bson:",omitempty" msgpack:"is_flow_check,omitempty"`
	Materials_count    *int64     `xorm:" null 'materials_count'" indexes:"" comment:"物料总数" json:"materials_count,omitempty" bson:",omitempty" msgpack:"materials_count,omitempty"`
	Enable_mark        *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_1) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark        *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_2) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code          *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description        *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date        *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid      *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username    *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date        *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid      *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username    *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version            *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: tab_flow_template
// 说明: 流程模板-流程模版表
// 流程模板-流程模版表
// 
type tab_flow_template struct {
	Template_id          *int64     `xorm:" pk not null 'template_id'" indexes:"" comment:"模板主键" json:"template_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"template_id,omitempty"`
	Template_name        *string    `xorm:" varchar(60) null 'template_name'" indexes:"" comment:"模板名称" json:"template_name,omitempty" bson:",omitempty" msgpack:"template_name,omitempty"`
	Template_description *string    `xorm:" varchar(255) null 'template_description'" indexes:"" comment:"模版说明" json:"template_description,omitempty" bson:",omitempty" msgpack:"template_description,omitempty"`
	Affairs_id           *int64     `xorm:" null 'affairs_id'" indexes:"" comment:"事项主键" json:"affairs_id,omitempty" bson:",omitempty" msgpack:"affairs_id,omitempty" fk:"tab_affairs(affairs_id)"`
	Is_enabled           *Boolean   `xorm:" null default 0 'is_enabled'" indexes:"" comment:"是否曾经启用" json:"is_enabled,omitempty" bson:",omitempty" msgpack:"is_enabled,omitempty"`
	Use_material_pool    *Boolean   `xorm:" null default 0 'use_material_pool'" indexes:"" comment:"是否使用物料池" json:"use_material_pool,omitempty" bson:",omitempty" msgpack:"use_material_pool,omitempty"`
	Enable_mark          *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_1) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark          *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_2) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code            *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description          *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date          *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid        *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username      *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date          *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid        *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username      *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version              *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: dic_connection_type
// 说明: 字典-专线连接方式
// 字典-专线连接方式
// 
type dic_connection_type struct {
	Connection_type_id     *int64     `xorm:" pk not null 'connection_type_id'" indexes:"" comment:"连接方式主键" json:"connection_type_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"connection_type_id,omitempty"`
	Connection_type_name   *string    `xorm:" varchar(255) null 'connection_type_name'" indexes:" unique(Index_1) " comment:"连接方式名称" json:"connection_type_name,omitempty" bson:",omitempty" msgpack:"connection_type_name,omitempty"`
	Connection_type_code   *string    `xorm:" varchar(255) null 'connection_type_code'" indexes:" unique(Index_2) " comment:"连接方式编码" json:"connection_type_code,omitempty" bson:",omitempty" msgpack:"connection_type_code,omitempty"`
	Connection_type_prompt *string    `xorm:" varchar(255) null 'connection_type_prompt'" indexes:"" comment:"连接方式简称" json:"connection_type_prompt,omitempty" bson:",omitempty" msgpack:"connection_type_prompt,omitempty"`
	Enable_mark            *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_3) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark            *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_4) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code              *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description            *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date            *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid          *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username        *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date            *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid          *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username        *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version                *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: dic_education
// 说明: 自治家园字典-学历
// 自治家园字典-学历
// 
type dic_education struct {
	Education_id     *int64     `xorm:" pk not null 'education_id'" indexes:"" comment:"学历主键" json:"education_id,omitempty" bson:",omitempty" msgpack:"education_id,omitempty"`
	Education_name   *string    `xorm:" varchar(255) null 'education_name'" indexes:"" comment:"学历名称" json:"education_name,omitempty" bson:",omitempty" msgpack:"education_name,omitempty"`
	Education_prompt *string    `xorm:" varchar(255) null 'education_prompt'" indexes:"" comment:"学历别称" json:"education_prompt,omitempty" bson:",omitempty" msgpack:"education_prompt,omitempty"`
	Education_code   *string    `xorm:" varchar(255) null 'education_code'" indexes:"" comment:"学历编码" json:"education_code,omitempty" bson:",omitempty" msgpack:"education_code,omitempty"`
	Enable_mark      *Boolean   `xorm:" null default 1 'enable_mark'" indexes:"" comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark      *Boolean   `xorm:" null default 0 'delete_mark'" indexes:"" comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code        *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description      *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date      *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid    *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username  *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date      *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid    *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username  *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version          *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: dic_handling_form
// 说明: 字典-事项办理形式
// 字典-事项办理形式
// 
type dic_handling_form struct {
	Handling_form_id    *int64     `xorm:" pk not null 'handling_form_id'" indexes:"" comment:"办理形式主键" json:"handling_form_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"handling_form_id,omitempty"`
	Handling_form_name  *string    `xorm:" varchar(255) null 'handling_form_name'" indexes:"" comment:"办理形式名称" json:"handling_form_name,omitempty" bson:",omitempty" msgpack:"handling_form_name,omitempty"`
	Handling_form_promt *string    `xorm:" varchar(50) null 'handling_form_promt'" indexes:"" comment:"办理形式简称" json:"handling_form_promt,omitempty" bson:",omitempty" msgpack:"handling_form_promt,omitempty"`
	Handling_form_code  *string    `xorm:" varchar(255) null 'handling_form_code'" indexes:" unique(Index_1) " comment:"办理形式编码" json:"handling_form_code,omitempty" bson:",omitempty" msgpack:"handling_form_code,omitempty"`
	Enable_mark         *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_2) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark         *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_3) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code           *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description         *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date         *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid       *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username     *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date         *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid       *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username     *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version             *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: dic_proposer_type
// 说明: 字典-申请人类型
// 字典-申请人类型
// 
type dic_proposer_type struct {
	Proposer_type_id     *int64     `xorm:" pk not null 'proposer_type_id'" indexes:"" comment:"申请人类型主键" json:"proposer_type_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"proposer_type_id,omitempty"`
	Proposer_type_value  *string    `xorm:" varchar(50) null 'proposer_type_value'" indexes:"" comment:"申请人类型名称" json:"proposer_type_value,omitempty" bson:",omitempty" msgpack:"proposer_type_value,omitempty"`
	Proposer_type_prompt *string    `xorm:" varchar(50) null 'proposer_type_prompt'" indexes:"" comment:"申请人类型简称" json:"proposer_type_prompt,omitempty" bson:",omitempty" msgpack:"proposer_type_prompt,omitempty"`
	Value_code           *string    `xorm:" varchar(50) null 'value_code'" indexes:" unique(Index_1) " comment:"编码" json:"value_code,omitempty" bson:",omitempty" msgpack:"value_code,omitempty"`
	Enable_mark          *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_2) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark          *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_3) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code            *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description          *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date          *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid        *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username      *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date          *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid        *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username      *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version              *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: dic_satisfy_level
// 说明: 字典-评价结果等级
// 字典-评价结果等级
// 
type dic_satisfy_level struct {
	Satisfy_level_id     *int64     `xorm:" pk not null 'satisfy_level_id'" indexes:"" comment:"评价等级主键" json:"satisfy_level_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"satisfy_level_id,omitempty"`
	Satisfy_level_name   *string    `xorm:" varchar(255) null 'satisfy_level_name'" indexes:"" comment:"评价等级名称" json:"satisfy_level_name,omitempty" bson:",omitempty" msgpack:"satisfy_level_name,omitempty"`
	Satisfy_level_prompt *string    `xorm:" varchar(255) null 'satisfy_level_prompt'" indexes:"" comment:"评价等级简称" json:"satisfy_level_prompt,omitempty" bson:",omitempty" msgpack:"satisfy_level_prompt,omitempty"`
	Satisfy_level_code   *string    `xorm:" varchar(255) null 'satisfy_level_code'" indexes:" unique(Index_1) " comment:"评价等级编码" json:"satisfy_level_code,omitempty" bson:",omitempty" msgpack:"satisfy_level_code,omitempty"`
	Enable_mark          *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_2) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark          *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_3) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code            *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description          *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date          *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid        *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username      *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date          *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid        *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username      *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version              *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: tab_affairs_info_outcome
// 说明: 事项受理-结果物
// 事项受理-结果物
// 
type tab_affairs_info_outcome struct {
	Affairs_outcome_id *int64     `xorm:" pk not null 'affairs_outcome_id'" indexes:"" comment:"事项受理结果物主键" json:"affairs_outcome_id,omitempty" bson:",omitempty" msgpack:"affairs_outcome_id,omitempty"`
	Result_receive_id  *int64     `xorm:" null 'result_receive_id'" indexes:"" comment:"送达方式主键" json:"result_receive_id,omitempty" bson:",omitempty" msgpack:"result_receive_id,omitempty" fk:"dic_result_receive(result_receive_id)"`
	Affairs_info_id    *int64     `xorm:" null 'affairs_info_id'" indexes:"" comment:"受理主键" json:"affairs_info_id,omitempty" bson:",omitempty" msgpack:"affairs_info_id,omitempty" fk:"tab_affairs_info(affairs_info_id)"`
	Affairs_info_code  *string    `xorm:" varchar(255) null 'affairs_info_code'" indexes:"" comment:"受理流水号" json:"affairs_info_code,omitempty" bson:",omitempty" msgpack:"affairs_info_code,omitempty"`
	Affairs_name       *string    `xorm:" varchar(255) null 'affairs_name'" indexes:"" comment:"事项名称" json:"affairs_name,omitempty" bson:",omitempty" msgpack:"affairs_name,omitempty"`
	User_id            *int64     `xorm:" null 'user_id'" indexes:"" comment:"受理人主键" json:"user_id,omitempty" bson:",omitempty" msgpack:"user_id,omitempty"`
	User_name          *string    `xorm:" varchar(50) null 'user_name'" indexes:"" comment:"受理人姓名" json:"user_name,omitempty" bson:",omitempty" msgpack:"user_name,omitempty"`
	Material_id        *int64     `xorm:" null 'material_id'" indexes:"" comment:"结果物主键" json:"material_id,omitempty" bson:",omitempty" msgpack:"material_id,omitempty"`
	Material_name      *string    `xorm:" varchar(255) null 'material_name'" indexes:"" comment:"结果物名称" json:"material_name,omitempty" bson:",omitempty" msgpack:"material_name,omitempty"`
	Is_outcome         *Boolean   `xorm:" null default 1 'is_outcome'" indexes:"" comment:"是否新材料" json:"is_outcome,omitempty" bson:",omitempty" msgpack:"is_outcome,omitempty"`
	Express_bill       *string    `xorm:" varchar(255) null 'express_bill'" indexes:"" comment:"快递单号" json:"express_bill,omitempty" bson:",omitempty" msgpack:"express_bill,omitempty"`
	Telephone          *string    `xorm:" varchar(255) null 'telephone'" indexes:"" comment:"收件人电话" json:"telephone,omitempty" bson:",omitempty" msgpack:"telephone,omitempty"`
	Post_date          *time.Time `xorm:" null 'post_date'" indexes:"" comment:"寄件时间" json:"post_date,omitempty" bson:",omitempty" msgpack:"post_date,omitempty"`
	Pickup_date        *time.Time `xorm:" null 'pickup_date'" indexes:"" comment:"收件时间" json:"pickup_date,omitempty" bson:",omitempty" msgpack:"pickup_date,omitempty"`
	Pickup_man         *string    `xorm:" varchar(50) null 'pickup_man'" indexes:"" comment:"收件人姓名" json:"pickup_man,omitempty" bson:",omitempty" msgpack:"pickup_man,omitempty"`
	Post_address       *string    `xorm:" text null 'post_address'" indexes:"" comment:"收件人地址" json:"post_address,omitempty" bson:",omitempty" msgpack:"post_address,omitempty"`
	Print_bill         *string    `xorm:" varchar(50) null 'print_bill'" indexes:" index(Index_3) " comment:"打印单号" json:"print_bill,omitempty" bson:",omitempty" msgpack:"print_bill,omitempty"`
	Enable_mark        *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_1) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark        *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_2) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code          *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description        *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date        *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid      *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username    *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date        *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid      *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username    *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version            *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: dic_married_type
// 说明: 自治家园字典-已婚分类
// 自治家园字典-已婚分类
// 
type dic_married_type struct {
	Married_type_id     *int64     `xorm:" pk not null 'married_type_id'" indexes:"" comment:"已婚分类主键" json:"married_type_id,omitempty" bson:",omitempty" msgpack:"married_type_id,omitempty"`
	Married_type_name   *string    `xorm:" varchar(255) null 'married_type_name'" indexes:"" comment:"已婚分类名称" json:"married_type_name,omitempty" bson:",omitempty" msgpack:"married_type_name,omitempty"`
	Married_type_prompt *string    `xorm:" varchar(255) null 'married_type_prompt'" indexes:"" comment:"已婚分类别称" json:"married_type_prompt,omitempty" bson:",omitempty" msgpack:"married_type_prompt,omitempty"`
	Married_type_code   *string    `xorm:" varchar(255) null 'married_type_code'" indexes:"" comment:"已婚分类编码" json:"married_type_code,omitempty" bson:",omitempty" msgpack:"married_type_code,omitempty"`
	Enable_mark         *Boolean   `xorm:" null default 1 'enable_mark'" indexes:"" comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark         *Boolean   `xorm:" null default 0 'delete_mark'" indexes:"" comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code           *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description         *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date         *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid       *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username     *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date         *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid       *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username     *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version             *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: tab_material_objects
// 说明: 材料对象表
// 材料对象表
// 
type tab_material_objects struct {
	Object_id       *int64     `xorm:" pk not null 'object_id'" indexes:"" comment:"对象主键" json:"object_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"object_id,omitempty"`
	Affairs_id      *int64     `xorm:" null 'affairs_id'" indexes:"" comment:"事项主键" json:"affairs_id,omitempty" bson:",omitempty" msgpack:"affairs_id,omitempty" fk:"tab_affairs(affairs_id)"`
	Object_name     *string    `xorm:" varchar(255) null 'object_name'" indexes:"" comment:"对象名称" json:"object_name,omitempty" bson:",omitempty" msgpack:"object_name,omitempty"`
	Time_limit      *int64     `xorm:" null 'time_limit'" indexes:"" comment:"时限/单位" json:"time_limit,omitempty" bson:",omitempty" msgpack:"time_limit,omitempty"`
	Time_limit_type *int64     `xorm:" null 'time_limit_type'" indexes:"" comment:"时限类型id" json:"time_limit_type,omitempty" bson:",omitempty" msgpack:"time_limit_type,omitempty" fk:"dic_time_limit_type(time_limit_type_id)"`
	Enable_mark     *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_1) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark     *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_2) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code       *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description     *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date     *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid   *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date     *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid   *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version         *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: dic_affairs_objects
// 说明: 字典-事项对象
// 字典-事项对象
// 
type dic_affairs_objects struct {
	Affairs_object_id     *int64     `xorm:" pk not null 'affairs_object_id'" indexes:"" comment:"事项对象主键" json:"affairs_object_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"affairs_object_id,omitempty"`
	Affairs_object_name   *string    `xorm:" varchar(255) null 'affairs_object_name'" indexes:"" comment:"事项对象名称" json:"affairs_object_name,omitempty" bson:",omitempty" msgpack:"affairs_object_name,omitempty"`
	Affairs_object_prompt *string    `xorm:" varchar(255) null 'affairs_object_prompt'" indexes:"" comment:"事项对象简称" json:"affairs_object_prompt,omitempty" bson:",omitempty" msgpack:"affairs_object_prompt,omitempty"`
	Affairs_object_code   *string    `xorm:" varchar(255) null 'affairs_object_code'" indexes:" unique(Index_1) " comment:"事项对象编号" json:"affairs_object_code,omitempty" bson:",omitempty" msgpack:"affairs_object_code,omitempty"`
	Enable_mark           *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_2) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark           *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_3) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code             *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description           *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date           *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid         *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username       *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date           *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid         *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username       *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version               *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: dic_religion
// 说明: 自治家园字典-宗教信仰
// 自治家园字典-宗教信仰
// 
type dic_religion struct {
	Religion_id     *int64     `xorm:" pk not null 'religion_id'" indexes:"" comment:"宗教信仰" json:"religion_id,omitempty" bson:",omitempty" msgpack:"religion_id,omitempty"`
	Religion_name   *string    `xorm:" varchar(255) null 'religion_name'" indexes:"" comment:"宗教名称" json:"religion_name,omitempty" bson:",omitempty" msgpack:"religion_name,omitempty"`
	Religion_prompt *string    `xorm:" varchar(255) null 'religion_prompt'" indexes:"" comment:"宗教别称" json:"religion_prompt,omitempty" bson:",omitempty" msgpack:"religion_prompt,omitempty"`
	Religion_code   *string    `xorm:" varchar(255) null 'religion_code'" indexes:"" comment:"宗教编码" json:"religion_code,omitempty" bson:",omitempty" msgpack:"religion_code,omitempty"`
	Enable_mark     *Boolean   `xorm:" null default 1 'enable_mark'" indexes:"" comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark     *Boolean   `xorm:" null default 0 'delete_mark'" indexes:"" comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code       *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description     *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date     *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid   *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date     *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid   *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version         *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: dic_work
// 说明: 自治家园字典-职业
// 自治家园字典-职业
// 
type dic_work struct {
	Work_id         *int64     `xorm:" pk not null 'work_id'" indexes:"" comment:"职业主键" json:"work_id,omitempty" bson:",omitempty" msgpack:"work_id,omitempty"`
	Work_name       *string    `xorm:" varchar(255) null 'work_name'" indexes:"" comment:"职业名称" json:"work_name,omitempty" bson:",omitempty" msgpack:"work_name,omitempty"`
	Work_prompt     *string    `xorm:" varchar(255) null 'work_prompt'" indexes:"" comment:"职业别称" json:"work_prompt,omitempty" bson:",omitempty" msgpack:"work_prompt,omitempty"`
	Work_code       *string    `xorm:" varchar(255) null 'work_code'" indexes:"" comment:"职业编码" json:"work_code,omitempty" bson:",omitempty" msgpack:"work_code,omitempty"`
	Enable_mark     *Boolean   `xorm:" null default 1 'enable_mark'" indexes:"" comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark     *Boolean   `xorm:" null default 0 'delete_mark'" indexes:"" comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code       *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description     *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date     *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid   *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date     *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid   *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version         *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: tab_dedicate_account
// 说明: 专线设置-专线登录帐号
// 专线设置-专线登录帐号
// 
type tab_dedicate_account struct {
	Dedicate_account_id *int64     `xorm:" pk not null 'dedicate_account_id'" indexes:"" comment:"专线登录帐号主键" json:"dedicate_account_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"dedicate_account_id,omitempty"`
	User_id             *int64     `xorm:" null 'user_id'" indexes:"" comment:"用户主键" json:"user_id,omitempty" bson:",omitempty" msgpack:"user_id,omitempty" fk:"tab_user(user_id)"`
	Dedicate_account    *string    `xorm:" varchar(255) null 'dedicate_account'" indexes:"" comment:"专线登录帐号" json:"dedicate_account,omitempty" bson:",omitempty" msgpack:"dedicate_account,omitempty"`
	Dedicate_pwd        *string    `xorm:" varchar(255) null 'dedicate_pwd'" indexes:"" comment:"专线登录密码" json:"dedicate_pwd,omitempty" bson:",omitempty" msgpack:"dedicate_pwd,omitempty"`
	Enable_mark         *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_1) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark         *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_2) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code           *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description         *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date         *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid       *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username     *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date         *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid       *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username     *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version             *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: tab_virtual_machines
// 说明: 专线设置-虚拟机列表
// 专线设置-虚拟机列表
// 
type tab_virtual_machines struct {
	Vm_id              *int64     `xorm:" pk not null 'vm_id'" indexes:"" comment:"虚拟机主键" json:"vm_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"vm_id,omitempty"`
	Vm_code            *string    `xorm:" varchar(255) null 'vm_code'" indexes:"" comment:"虚拟机编号" json:"vm_code,omitempty" bson:",omitempty" msgpack:"vm_code,omitempty"`
	Vm_name            *string    `xorm:" varchar(255) null 'vm_name'" indexes:"" comment:"虚拟机名称" json:"vm_name,omitempty" bson:",omitempty" msgpack:"vm_name,omitempty"`
	Vm_description     *string    `xorm:" varchar(500) null 'vm_description'" indexes:"" comment:"虚拟机描述" json:"vm_description,omitempty" bson:",omitempty" msgpack:"vm_description,omitempty"`
	Vm_address         *string    `xorm:" varchar(255) null 'vm_address'" indexes:"" comment:"虚拟机IP地址" json:"vm_address,omitempty" bson:",omitempty" msgpack:"vm_address,omitempty"`
	Vm_special_address *string    `xorm:" varchar(255) null 'vm_special_address'" indexes:"" comment:"虚拟机专网IP地址" json:"vm_special_address,omitempty" bson:",omitempty" msgpack:"vm_special_address,omitempty"`
	Dedicate_protect   *Boolean   `xorm:" null default 0 'dedicate_protect'" indexes:"" comment:"专线保护" json:"dedicate_protect,omitempty" bson:",omitempty" msgpack:"dedicate_protect,omitempty"`
	Online_status      *Boolean   `xorm:" null default 1 'online_status'" indexes:"" comment:"在线状态" json:"online_status,omitempty" bson:",omitempty" msgpack:"online_status,omitempty"`
	Default_start      *string    `xorm:" varchar(255) null 'default_start'" indexes:"" comment:"默认启动程序" json:"default_start,omitempty" bson:",omitempty" msgpack:"default_start,omitempty"`
	Online_user_count  *int64     `xorm:" null default 0 'online_user_count'" indexes:"" comment:"在线用户数" json:"online_user_count,omitempty" bson:",omitempty" msgpack:"online_user_count,omitempty"`
	Enable_mark        *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_1) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark        *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_2) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code          *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description        *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date        *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid      *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username    *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date        *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid      *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username    *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version            *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: tab_prepare_proposer
// 说明: 外网-申请人表
// 外网-申请人表
// 
type tab_prepare_proposer struct {
	Proposer_id        *int64     `xorm:" pk not null 'proposer_id'" indexes:"" comment:"申请人主键" json:"proposer_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"proposer_id,omitempty"`
	Prepare_affairs_id *int64     `xorm:" null 'prepare_affairs_id'" indexes:"" comment:"预审申请主键" json:"prepare_affairs_id,omitempty" bson:",omitempty" msgpack:"prepare_affairs_id,omitempty" fk:"tab_prepare_affairs(prepare_affairs_id)"`
	Proposer_name      *string    `xorm:" varchar(255) null 'proposer_name'" indexes:"" comment:"申请人名" json:"proposer_name,omitempty" bson:",omitempty" msgpack:"proposer_name,omitempty"`
	Sex                *int64     `xorm:" null 'sex'" indexes:"" comment:"性别" json:"sex,omitempty" bson:",omitempty" msgpack:"sex,omitempty"`
	Id_code            *string    `xorm:" varchar(255) not null 'id_code'" indexes:" index(Index_1) " comment:"身份证号码" json:"id_code,omitempty" bson:",omitempty" msgpack:"id_code,omitempty"`
	Mobile             *string    `xorm:" varchar(255) null 'mobile'" indexes:" index(Index_2) " comment:"手机" json:"mobile,omitempty" bson:",omitempty" msgpack:"mobile,omitempty"`
	Email              *string    `xorm:" varchar(255) null 'email'" indexes:" index(Index_3) " comment:"邮箱" json:"email,omitempty" bson:",omitempty" msgpack:"email,omitempty"`
	Photo_url          *string    `xorm:" varchar(255) null 'photo_url'" indexes:"" comment:"照片" json:"photo_url,omitempty" bson:",omitempty" msgpack:"photo_url,omitempty"`
	Enable_mark        *Boolean   `xorm:" null default 1 'enable_mark'" indexes:"" comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark        *Boolean   `xorm:" null default 0 'delete_mark'" indexes:"" comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code          *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description        *string    `xorm:" varchar(255) null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date        *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid      *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username    *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date        *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid      *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username    *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version            *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: tab_system_list
// 说明: 用户中心-系统列表
// 用户中心-系统列表
// 
type tab_system_list struct {
	System_id       *int64     `xorm:" pk not null 'system_id'" indexes:"" comment:"系统主键" json:"system_id,omitempty" bson:",omitempty" msgpack:"system_id,omitempty"`
	System_name     *string    `xorm:" varchar(50) null 'system_name'" indexes:" unique(Index_1) " comment:"系统名称" json:"system_name,omitempty" bson:",omitempty" msgpack:"system_name,omitempty"`
	System_code     *string    `xorm:" varchar(50) null 'system_code'" indexes:" unique(Index_2) " comment:"系统编码" json:"system_code,omitempty" bson:",omitempty" msgpack:"system_code,omitempty"`
	Enable_mark     *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_3) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark     *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_4) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code       *int64     `xorm:" null 'sort_code'" indexes:" index(Index_5) " comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description     *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date     *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid   *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date     *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid   *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version         *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: dic_census_register_status
// 说明: 自治家园字典-户籍状态
// 自治家园字典-户籍状态
// 
type dic_census_register_status struct {
	Census_register_status_id     *int64     `xorm:" pk not null 'census_register_status_id'" indexes:"" comment:"户籍状态主键" json:"census_register_status_id,omitempty" bson:",omitempty" msgpack:"census_register_status_id,omitempty"`
	Census_register_status_name   *string    `xorm:" varchar(255) null 'census_register_status_name'" indexes:"" comment:"户籍状态名称" json:"census_register_status_name,omitempty" bson:",omitempty" msgpack:"census_register_status_name,omitempty"`
	Census_register_status_prompt *string    `xorm:" varchar(255) null 'census_register_status_prompt'" indexes:"" comment:"户籍状态别称" json:"census_register_status_prompt,omitempty" bson:",omitempty" msgpack:"census_register_status_prompt,omitempty"`
	Census_register_status_code   *string    `xorm:" varchar(255) null 'census_register_status_code'" indexes:"" comment:"户籍状态编码" json:"census_register_status_code,omitempty" bson:",omitempty" msgpack:"census_register_status_code,omitempty"`
	Enable_mark                   *Boolean   `xorm:" null default 1 'enable_mark'" indexes:"" comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark                   *Boolean   `xorm:" null default 0 'delete_mark'" indexes:"" comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code                     *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description                   *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date                   *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid                 *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username               *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date                   *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid                 *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username               *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version                       *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: dic_gender
// 说明: 字典-性别
// 字典-性别
// 
type dic_gender struct {
	Gender_id       *int64     `xorm:" pk not null 'gender_id'" indexes:"" comment:"性别主键" json:"gender_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"gender_id,omitempty"`
	Gender_name     *string    `xorm:" varchar(50) null 'gender_name'" indexes:"" comment:"性别名称" json:"gender_name,omitempty" bson:",omitempty" msgpack:"gender_name,omitempty"`
	Gender_prompt   *string    `xorm:" varchar(50) null 'gender_prompt'" indexes:"" comment:"性别简称" json:"gender_prompt,omitempty" bson:",omitempty" msgpack:"gender_prompt,omitempty"`
	Value_code      *string    `xorm:" varchar(50) null 'value_code'" indexes:" unique(Index_1) " comment:"性别编码" json:"value_code,omitempty" bson:",omitempty" msgpack:"value_code,omitempty"`
	Enable_mark     *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_2) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark     *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_3) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code       *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description     *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date     *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid   *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date     *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid   *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version         *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: dic_material_type
// 说明: 字典-材料类型
// 字典-材料类型
// 
type dic_material_type struct {
	Material_type_id     *int64     `xorm:" pk not null 'material_type_id'" indexes:"" comment:"材料类型" json:"material_type_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"material_type_id,omitempty"`
	Material_type_name   *string    `xorm:" varchar(50) null 'material_type_name'" indexes:"" comment:"材料类型名称" json:"material_type_name,omitempty" bson:",omitempty" msgpack:"material_type_name,omitempty"`
	Material_type_prompt *string    `xorm:" varchar(50) null 'material_type_prompt'" indexes:"" comment:"材料类型简称" json:"material_type_prompt,omitempty" bson:",omitempty" msgpack:"material_type_prompt,omitempty"`
	Value_code           *string    `xorm:" varchar(50) null 'value_code'" indexes:" unique(Index_1) " comment:"编码" json:"value_code,omitempty" bson:",omitempty" msgpack:"value_code,omitempty"`
	Enable_mark          *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_2) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark          *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_3) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code            *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description          *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date          *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid        *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username      *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date          *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid        *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username      *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version              *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: dic_speciality
// 说明: 自治家园字典-特长
// 自治家园字典-特长
// 
type dic_speciality struct {
	Speciality_id     *int64     `xorm:" pk not null 'speciality_id'" indexes:"" comment:"特长主键" json:"speciality_id,omitempty" bson:",omitempty" msgpack:"speciality_id,omitempty"`
	Speciality_name   *string    `xorm:" varchar(255) null 'speciality_name'" indexes:"" comment:"特长名称" json:"speciality_name,omitempty" bson:",omitempty" msgpack:"speciality_name,omitempty"`
	Speciality_prompt *string    `xorm:" varchar(255) null 'speciality_prompt'" indexes:"" comment:"特长别称" json:"speciality_prompt,omitempty" bson:",omitempty" msgpack:"speciality_prompt,omitempty"`
	Speciality_code   *string    `xorm:" varchar(255) null 'speciality_code'" indexes:"" comment:"特长编码" json:"speciality_code,omitempty" bson:",omitempty" msgpack:"speciality_code,omitempty"`
	Enable_mark       *Boolean   `xorm:" null default 1 'enable_mark'" indexes:"" comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark       *Boolean   `xorm:" null default 0 'delete_mark'" indexes:"" comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code         *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description       *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date       *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid     *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username   *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date       *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid     *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username   *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version           *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: tab_coporation_off
// 说明: 法人注销信息
// 法人注销信息
// 
type tab_coporation_off struct {
	Coporation_off_id *int64     `xorm:" pk not null 'coporation_off_id'" indexes:"" comment:"法人注销主键" json:"coporation_off_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"coporation_off_id,omitempty"`
	Credit_code       *string    `xorm:" varchar(255) null 'credit_code'" indexes:"" comment:"社会统一信用代码" json:"credit_code,omitempty" bson:",omitempty" msgpack:"credit_code,omitempty"`
	Org_code          *string    `xorm:" varchar(50) null 'org_code'" indexes:"" comment:"组织机构代码" json:"org_code,omitempty" bson:",omitempty" msgpack:"org_code,omitempty"`
	Off_reason        *string    `xorm:" varchar(255) null 'off_reason'" indexes:"" comment:"注销原因" json:"off_reason,omitempty" bson:",omitempty" msgpack:"off_reason,omitempty"`
	Off_date          *time.Time `xorm:" null 'off_date'" indexes:"" comment:"注销日期" json:"off_date,omitempty" bson:",omitempty" msgpack:"off_date,omitempty"`
	Off_sector        *string    `xorm:" varchar(255) null 'off_sector'" indexes:"" comment:"注销机关" json:"off_sector,omitempty" bson:",omitempty" msgpack:"off_sector,omitempty"`
	Enable_mark       *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_1) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark       *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_2) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code         *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description       *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date       *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid     *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username   *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date       *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid     *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username   *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version           *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: tab_flow_node_result_info
// 说明: 流程模板-流程节点结果信息表
// 流程模板-流程节点结果信息表
// 
type tab_flow_node_result_info struct {
	Node_result_id    *int64     `xorm:" pk not null 'node_result_id'" indexes:"" comment:"流程结果主键" json:"node_result_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"node_result_id,omitempty"`
	Node_id           *int64     `xorm:" null 'node_id'" indexes:"" comment:"节点编号" json:"node_id,omitempty" bson:",omitempty" msgpack:"node_id,omitempty" fk:"tab_flow_node_info(node_id)"`
	Node_result_name  *string    `xorm:" varchar(255) null 'node_result_name'" indexes:"" comment:"流程结果名称" json:"node_result_name,omitempty" bson:",omitempty" msgpack:"node_result_name,omitempty"`
	Next_node_id      *int64     `xorm:" null 'next_node_id'" indexes:"" comment:"下一步节点" json:"next_node_id,omitempty" bson:",omitempty" msgpack:"next_node_id,omitempty" fk:"tab_flow_node_info(node_id)"`
	Is_default_result *Boolean   `xorm:" null default 0 'is_default_result'" indexes:"" comment:"是否默认结果" json:"is_default_result,omitempty" bson:",omitempty" msgpack:"is_default_result,omitempty"`
	Enable_mark       *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_1) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark       *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_2) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code         *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description       *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date       *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid     *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username   *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date       *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid     *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username   *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version           *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: dic_census_register_type
// 说明: 自治家园字典-户口类别
// 自治家园字典-户口类别
// 
type dic_census_register_type struct {
	Census_register_type_id     *int64     `xorm:" pk not null 'census_register_type_id'" indexes:"" comment:"户口类别主键" json:"census_register_type_id,omitempty" bson:",omitempty" msgpack:"census_register_type_id,omitempty"`
	Census_register_type_name   *string    `xorm:" varchar(255) null 'census_register_type_name'" indexes:"" comment:"户口类别名称" json:"census_register_type_name,omitempty" bson:",omitempty" msgpack:"census_register_type_name,omitempty"`
	Census_register_type_prompt *string    `xorm:" varchar(255) null 'census_register_type_prompt'" indexes:"" comment:"户口类别别称" json:"census_register_type_prompt,omitempty" bson:",omitempty" msgpack:"census_register_type_prompt,omitempty"`
	Census_register_type_code   *string    `xorm:" varchar(255) null 'census_register_type_code'" indexes:"" comment:"户口类别编码" json:"census_register_type_code,omitempty" bson:",omitempty" msgpack:"census_register_type_code,omitempty"`
	Enable_mark                 *Boolean   `xorm:" null default 1 'enable_mark'" indexes:"" comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark                 *Boolean   `xorm:" null default 0 'delete_mark'" indexes:"" comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code                   *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description                 *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date                 *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid               *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username             *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date                 *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid               *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username             *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version                     *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: tab_credential_field
// 说明: 证件照面字段定义
// 证件照面字段定义
// 
type tab_credential_field struct {
	Field_id        *int64     `xorm:" pk not null 'field_id'" indexes:"" comment:"证面字段主键" json:"field_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"field_id,omitempty"`
	Credential_type *int64     `xorm:" null 'credential_type'" indexes:"" comment:"证照目录编码" json:"credential_type,omitempty" bson:",omitempty" msgpack:"credential_type,omitempty"`
	Field_name      *string    `xorm:" varchar(255) null 'field_name'" indexes:"" comment:"证面字段名称" json:"field_name,omitempty" bson:",omitempty" msgpack:"field_name,omitempty"`
	Field_type      *int64     `xorm:" null 'field_type'" indexes:"" comment:"证面字段类型" json:"field_type,omitempty" bson:",omitempty" msgpack:"field_type,omitempty"`
	Enable_mark     *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_1) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark     *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_2) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code       *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description     *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date     *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid   *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date     *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid   *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version         *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: tab_person
// 说明: 自治家园-居民信表
// 自治家园-居民信表
// 
type tab_person struct {
	Person_id                        *int64     `xorm:" pk not null 'person_id'" indexes:"" comment:"居民表主键" json:"person_id,omitempty" bson:",omitempty" msgpack:"person_id,omitempty"`
	Nation_id                        *int64     `xorm:" null 'nation_id'" indexes:"" comment:"民族主键" json:"nation_id,omitempty" bson:",omitempty" msgpack:"nation_id,omitempty"`
	Health_status_id                 *int64     `xorm:" null 'health_status_id'" indexes:"" comment:"健康状况主键" json:"health_status_id,omitempty" bson:",omitempty" msgpack:"health_status_id,omitempty"`
	Blood_type_id                    *int64     `xorm:" null 'blood_type_id'" indexes:"" comment:"血型主键" json:"blood_type_id,omitempty" bson:",omitempty" msgpack:"blood_type_id,omitempty"`
	Marriage_status_id               *int64     `xorm:" null 'marriage_status_id'" indexes:"" comment:"婚姻状况主键" json:"marriage_status_id,omitempty" bson:",omitempty" msgpack:"marriage_status_id,omitempty"`
	Married_type_id                  *int64     `xorm:" null 'married_type_id'" indexes:"" comment:"已婚分类主键" json:"married_type_id,omitempty" bson:",omitempty" msgpack:"married_type_id,omitempty"`
	Military_service_status_id       *int64     `xorm:" null 'military_service_status_id'" indexes:"" comment:"兵役状况主键" json:"military_service_status_id,omitempty" bson:",omitempty" msgpack:"military_service_status_id,omitempty"`
	Politics_status_id               *int64     `xorm:" null 'politics_status_id'" indexes:"" comment:"政治面貌主键" json:"politics_status_id,omitempty" bson:",omitempty" msgpack:"politics_status_id,omitempty"`
	Education_id                     *int64     `xorm:" null 'education_id'" indexes:"" comment:"学历主键" json:"education_id,omitempty" bson:",omitempty" msgpack:"education_id,omitempty"`
	Family_master_relation_id        *int64     `xorm:" null 'family_master_relation_id'" indexes:"" comment:"与户主关系主键" json:"family_master_relation_id,omitempty" bson:",omitempty" msgpack:"family_master_relation_id,omitempty"`
	Card_type_id                     *int64     `xorm:" null 'card_type_id'" indexes:"" comment:"证件类型主键" json:"card_type_id,omitempty" bson:",omitempty" msgpack:"card_type_id,omitempty"`
	Country_id                       *int64     `xorm:" null 'country_id'" indexes:"" comment:"国家主键" json:"country_id,omitempty" bson:",omitempty" msgpack:"country_id,omitempty"`
	Religion_id                      *int64     `xorm:" null 'religion_id'" indexes:"" comment:"宗教信仰" json:"religion_id,omitempty" bson:",omitempty" msgpack:"religion_id,omitempty"`
	Work_status_id                   *int64     `xorm:" null 'work_status_id'" indexes:"" comment:"从业状况主键" json:"work_status_id,omitempty" bson:",omitempty" msgpack:"work_status_id,omitempty"`
	Part_time_job_id                 *int64     `xorm:" null 'part_time_job_id'" indexes:"" comment:"社会兼职主键" json:"part_time_job_id,omitempty" bson:",omitempty" msgpack:"part_time_job_id,omitempty"`
	Speciality_id                    *int64     `xorm:" null 'speciality_id'" indexes:"" comment:"特长主键" json:"speciality_id,omitempty" bson:",omitempty" msgpack:"speciality_id,omitempty"`
	Census_register_status_id        *int64     `xorm:" null 'census_register_status_id'" indexes:"" comment:"户籍状态主键" json:"census_register_status_id,omitempty" bson:",omitempty" msgpack:"census_register_status_id,omitempty"`
	Not_transfer_registered_cause_id *int64     `xorm:" null 'not_transfer_registered_cause_id'" indexes:"" comment:"未迁户口原因主键" json:"not_transfer_registered_cause_id,omitempty" bson:",omitempty" msgpack:"not_transfer_registered_cause_id,omitempty"`
	Census_register_type_id          *int64     `xorm:" null 'census_register_type_id'" indexes:"" comment:"户口类别主键" json:"census_register_type_id,omitempty" bson:",omitempty" msgpack:"census_register_type_id,omitempty"`
	Work_id                          *int64     `xorm:" null 'work_id'" indexes:"" comment:"职业主键" json:"work_id,omitempty" bson:",omitempty" msgpack:"work_id,omitempty"`
	Unit_property_id                 *int64     `xorm:" null 'unit_property_id'" indexes:"" comment:"单位性质主键" json:"unit_property_id,omitempty" bson:",omitempty" msgpack:"unit_property_id,omitempty"`
	Unit_trade_id                    *int64     `xorm:" null 'unit_trade_id'" indexes:"" comment:"单位所属行业主键" json:"unit_trade_id,omitempty" bson:",omitempty" msgpack:"unit_trade_id,omitempty"`
	Duty_level_id                    *int64     `xorm:" null 'duty_level_id'" indexes:"" comment:"干部职务主键" json:"duty_level_id,omitempty" bson:",omitempty" msgpack:"duty_level_id,omitempty"`
	Technology_duty_id               *int64     `xorm:" null 'technology_duty_id'" indexes:"" comment:"专业技术职务主键" json:"technology_duty_id,omitempty" bson:",omitempty" msgpack:"technology_duty_id,omitempty"`
	Register_region_id               *int64     `xorm:" null 'register_region_id'" indexes:"" comment:"户籍地所属归属地" json:"register_region_id,omitempty" bson:",omitempty" msgpack:"register_region_id,omitempty"`
	Dwell_region_id                  *int64     `xorm:" null 'dwell_region_id'" indexes:"" comment:"居住地所属归属地" json:"dwell_region_id,omitempty" bson:",omitempty" msgpack:"dwell_region_id,omitempty"`
	Card_code                        *string    `xorm:" varchar(20) not null 'card_code'" indexes:"" comment:"证件号码" json:"card_code,omitempty" bson:",omitempty" msgpack:"card_code,omitempty"`
	Name                             *string    `xorm:" varchar(50) not null 'name'" indexes:"" comment:"姓名" json:"name,omitempty" bson:",omitempty" msgpack:"name,omitempty"`
	Former_name                      *string    `xorm:" varchar(50) null 'former_name'" indexes:"" comment:"曾用名" json:"former_name,omitempty" bson:",omitempty" msgpack:"former_name,omitempty"`
	Phone                            *string    `xorm:" varchar(20) null 'phone'" indexes:"" comment:"联系电话" json:"phone,omitempty" bson:",omitempty" msgpack:"phone,omitempty"`
	Mobile_phone                     *string    `xorm:" varchar(20) null 'mobile_phone'" indexes:"" comment:"手机号码" json:"mobile_phone,omitempty" bson:",omitempty" msgpack:"mobile_phone,omitempty"`
	Join_work_date                   *time.Time `xorm:" null 'join_work_date'" indexes:"" comment:"参加工作日期" json:"join_work_date,omitempty" bson:",omitempty" msgpack:"join_work_date,omitempty"`
	Birthday                         *time.Time `xorm:" null 'birthday'" indexes:"" comment:"出生日期" json:"birthday,omitempty" bson:",omitempty" msgpack:"birthday,omitempty"`
	Age                              *int64     `xorm:" null 'age'" indexes:"" comment:"年龄" json:"age,omitempty" bson:",omitempty" msgpack:"age,omitempty"`
	Birthplace                       *string    `xorm:" varchar(255) null 'birthplace'" indexes:"" comment:"出生地" json:"birthplace,omitempty" bson:",omitempty" msgpack:"birthplace,omitempty"`
	Reside_address                   *string    `xorm:" varchar(255) null 'reside_address'" indexes:"" comment:"居住地址" json:"reside_address,omitempty" bson:",omitempty" msgpack:"reside_address,omitempty"`
	Stature                          *int64     `xorm:" null 'stature'" indexes:"" comment:"身高,单位cm" json:"stature,omitempty" bson:",omitempty" msgpack:"stature,omitempty"`
	Email                            *string    `xorm:" varchar(255) null 'email'" indexes:"" comment:"电子邮箱" json:"email,omitempty" bson:",omitempty" msgpack:"email,omitempty"`
	Photo                            *string    `xorm:" varchar(255) null 'photo'" indexes:"" comment:"照片" json:"photo,omitempty" bson:",omitempty" msgpack:"photo,omitempty"`
	Tags                             *int64     `xorm:" null 'tags'" indexes:"" comment:"标签(所有标签用一个字段表示)" json:"tags,omitempty" bson:",omitempty" msgpack:"tags,omitempty"`
	Enable_mark                      *Boolean   `xorm:" null default 1 'enable_mark'" indexes:"" comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark                      *Boolean   `xorm:" null default 0 'delete_mark'" indexes:"" comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code                        *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description                      *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date                      *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid                    *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username                  *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date                      *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid                    *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username                  *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version                          *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: dic_credential_type
// 说明: 字典-证照类型
// 字典-证照类型
// 
type dic_credential_type struct {
	Credential_type_id     *int64     `xorm:" pk not null 'credential_type_id'" indexes:"" comment:"证照类型主键" json:"credential_type_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"credential_type_id,omitempty"`
	Proposer_type_id       *int64     `xorm:" null 'proposer_type_id'" indexes:"" comment:"申请人类型主键" json:"proposer_type_id,omitempty" bson:",omitempty" msgpack:"proposer_type_id,omitempty" fk:"dic_proposer_type(proposer_type_id)"`
	Credential_type_name   *string    `xorm:" varchar(50) null 'credential_type_name'" indexes:"" comment:"证照类型名称" json:"credential_type_name,omitempty" bson:",omitempty" msgpack:"credential_type_name,omitempty"`
	Credential_type_prompt *string    `xorm:" varchar(50) null 'credential_type_prompt'" indexes:"" comment:"证照类型简称" json:"credential_type_prompt,omitempty" bson:",omitempty" msgpack:"credential_type_prompt,omitempty"`
	Credential_type_code   *string    `xorm:" varchar(50) null 'credential_type_code'" indexes:" unique(Index_1) " comment:"证照类型编码" json:"credential_type_code,omitempty" bson:",omitempty" msgpack:"credential_type_code,omitempty"`
	Enable_mark            *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_2) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark            *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_3) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code              *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description            *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date            *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid          *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username        *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date            *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid          *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username        *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version                *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: dic_reuse_type
// 说明: 字典-复用类型
// 字典-复用类型
// 
type dic_reuse_type struct {
	Reuse_type_id    *int64     `xorm:" pk not null 'reuse_type_id'" indexes:"" comment:"复用类型主键" json:"reuse_type_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"reuse_type_id,omitempty"`
	Reuse_type_name  *string    `xorm:" varchar(50) null 'reuse_type_name'" indexes:"" comment:"复用类型名称" json:"reuse_type_name,omitempty" bson:",omitempty" msgpack:"reuse_type_name,omitempty"`
	Reuse_type_promt *string    `xorm:" varchar(50) null 'reuse_type_promt'" indexes:"" comment:"复用类型简称" json:"reuse_type_promt,omitempty" bson:",omitempty" msgpack:"reuse_type_promt,omitempty"`
	Value_code       *string    `xorm:" varchar(50) null 'value_code'" indexes:" unique(Index_1) " comment:"编码" json:"value_code,omitempty" bson:",omitempty" msgpack:"value_code,omitempty"`
	Enable_mark      *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_2) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark      *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_3) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code        *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description      *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date      *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid    *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username  *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date      *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid    *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username  *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version          *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: tab_affairs_info
// 说明: 事项受理
// 事项受理
// 
type tab_affairs_info struct {
	Affairs_info_id              *int64     `xorm:" pk not null 'affairs_info_id'" indexes:"" comment:"受理主键" json:"affairs_info_id,omitempty" bson:",omitempty" msgpack:"affairs_info_id,omitempty"`
	Affairs_info_code            *string    `xorm:" varchar(255) null 'affairs_info_code'" indexes:" index(Index_1) " comment:"受理流水号" json:"affairs_info_code,omitempty" bson:",omitempty" msgpack:"affairs_info_code,omitempty"`
	Accept_file_code             *string    `xorm:" varchar(255) null 'accept_file_code'" indexes:"" comment:"受理文书编号" json:"accept_file_code,omitempty" bson:",omitempty" msgpack:"accept_file_code,omitempty"`
	Affairs_id                   *int64     `xorm:" not null 'affairs_id'" indexes:"" comment:"事项主键" json:"affairs_id,omitempty" bson:",omitempty" msgpack:"affairs_id,omitempty" fk:"tab_affairs(affairs_id)"`
	Affairs_name                 *string    `xorm:" varchar(255) null 'affairs_name'" indexes:"" comment:"事项名称" json:"affairs_name,omitempty" bson:",omitempty" msgpack:"affairs_name,omitempty"`
	User_id                      *int64     `xorm:" not null 'user_id'" indexes:"" comment:"受理人主键" json:"user_id,omitempty" bson:",omitempty" msgpack:"user_id,omitempty" fk:"tab_user(user_id)"`
	User_name                    *string    `xorm:" varchar(50) null 'user_name'" indexes:"" comment:"受理人姓名" json:"user_name,omitempty" bson:",omitempty" msgpack:"user_name,omitempty"`
	User_dept                    *string    `xorm:" varchar(255) null 'user_dept'" indexes:"" comment:"受理人所在部门" json:"user_dept,omitempty" bson:",omitempty" msgpack:"user_dept,omitempty"`
	Next_user_id                 *int64     `xorm:" null 'next_user_id'" indexes:"" comment:"下一步经办人主键" json:"next_user_id,omitempty" bson:",omitempty" msgpack:"next_user_id,omitempty"`
	Next_user_name               *string    `xorm:" varchar(50) null 'next_user_name'" indexes:"" comment:"下一步经办人姓名" json:"next_user_name,omitempty" bson:",omitempty" msgpack:"next_user_name,omitempty"`
	Reg_attribution_id           *int64     `xorm:" not null 'reg_attribution_id'" indexes:"" comment:"受理归属地主键" json:"reg_attribution_id,omitempty" bson:",omitempty" msgpack:"reg_attribution_id,omitempty" fk:"tab_region(region_id)"`
	Reg_attribution_name         *string    `xorm:" varchar(255) null 'reg_attribution_name'" indexes:"" comment:"受理归属地名称" json:"reg_attribution_name,omitempty" bson:",omitempty" msgpack:"reg_attribution_name,omitempty"`
	Next_attribution_admin_level *int64     `xorm:" null 'next_attribution_admin_level'" indexes:"" comment:"下一步归属地行政级别" json:"next_attribution_admin_level,omitempty" bson:",omitempty" msgpack:"next_attribution_admin_level,omitempty" fk:"dic_administrative_level(level_id)"`
	Next_attribution_area_id     *int64     `xorm:" null 'next_attribution_area_id'" indexes:"" comment:"下一步归属地行政区域主键" json:"next_attribution_area_id,omitempty" bson:",omitempty" msgpack:"next_attribution_area_id,omitempty"`
	Next_attribution_id          *int64     `xorm:" null 'next_attribution_id'" indexes:"" comment:"下一步归属地主键" json:"next_attribution_id,omitempty" bson:",omitempty" msgpack:"next_attribution_id,omitempty"`
	Next_attribution_name        *string    `xorm:" varchar(255) null 'next_attribution_name'" indexes:"" comment:"下一步归属地名称" json:"next_attribution_name,omitempty" bson:",omitempty" msgpack:"next_attribution_name,omitempty"`
	Finish_attribution_id        *int64     `xorm:" not null 'finish_attribution_id'" indexes:"" comment:"审结归属地主键" json:"finish_attribution_id,omitempty" bson:",omitempty" msgpack:"finish_attribution_id,omitempty"`
	Finish_attribution_name      *string    `xorm:" varchar(255) null 'finish_attribution_name'" indexes:"" comment:"审结归属地名称" json:"finish_attribution_name,omitempty" bson:",omitempty" msgpack:"finish_attribution_name,omitempty"`
	Node_current                 *string    `xorm:" text null 'node_current'" indexes:"" comment:"当前待审核结点" json:"node_current,omitempty" bson:",omitempty" msgpack:"node_current,omitempty"`
	Node_current_name            *string    `xorm:" text null 'node_current_name'" indexes:"" comment:"当前待审核节点名称" json:"node_current_name,omitempty" bson:",omitempty" msgpack:"node_current_name,omitempty"`
	Node_results                 *string    `xorm:" text null 'node_results'" indexes:"" comment:"已审核结点结果" json:"node_results,omitempty" bson:",omitempty" msgpack:"node_results,omitempty"`
	Queue_time                   *time.Time `xorm:" null 'queue_time'" indexes:"" comment:"取号时间" json:"queue_time,omitempty" bson:",omitempty" msgpack:"queue_time,omitempty"`
	Queue_call_time              *time.Time `xorm:" null 'queue_call_time'" indexes:"" comment:"呼号时间" json:"queue_call_time,omitempty" bson:",omitempty" msgpack:"queue_call_time,omitempty"`
	Apply_from                   *int64     `xorm:" null default 1 'apply_from'" indexes:"" comment:"申请来源" json:"apply_from,omitempty" bson:",omitempty" msgpack:"apply_from,omitempty" fk:"dic_apply_from(apply_from_id)"`
	Apply_from_no                *string    `xorm:" varchar(255) null 'apply_from_no'" indexes:"" comment:"来源唯一识别号" json:"apply_from_no,omitempty" bson:",omitempty" msgpack:"apply_from_no,omitempty"`
	Consult_affair               *Boolean   `xorm:" null default 0 'consult_affair'" indexes:"" comment:"是否咨询件" json:"consult_affair,omitempty" bson:",omitempty" msgpack:"consult_affair,omitempty"`
	Project_code                 *string    `xorm:" varchar(50) null 'project_code'" indexes:"" comment:"项目编号" json:"project_code,omitempty" bson:",omitempty" msgpack:"project_code,omitempty"`
	Window_name                  *string    `xorm:" varchar(255) null 'window_name'" indexes:"" comment:"受理窗口号" json:"window_name,omitempty" bson:",omitempty" msgpack:"window_name,omitempty"`
	Affairs_deadline             *time.Time `xorm:" null 'affairs_deadline'" indexes:"" comment:"事项时限" json:"affairs_deadline,omitempty" bson:",omitempty" msgpack:"affairs_deadline,omitempty"`
	Affairs_start                *time.Time `xorm:" null 'affairs_start'" indexes:" index(Index_2) " comment:"服务开始时间" json:"affairs_start,omitempty" bson:",omitempty" msgpack:"affairs_start,omitempty"`
	Affairs_end                  *time.Time `xorm:" null 'affairs_end'" indexes:" index(Index_3) " comment:"服务结束时间" json:"affairs_end,omitempty" bson:",omitempty" msgpack:"affairs_end,omitempty"`
	Affairs_finish               *time.Time `xorm:" null 'affairs_finish'" indexes:" index(Index_4) " comment:"事项结束时间" json:"affairs_finish,omitempty" bson:",omitempty" msgpack:"affairs_finish,omitempty"`
	Affairs_real_finish          *time.Time `xorm:" null 'affairs_real_finish'" indexes:" index(Index_5) " comment:"实际结束时间" json:"affairs_real_finish,omitempty" bson:",omitempty" msgpack:"affairs_real_finish,omitempty"`
	Satisfy_time                 *time.Time `xorm:" null 'satisfy_time'" indexes:"" comment:"评价时间" json:"satisfy_time,omitempty" bson:",omitempty" msgpack:"satisfy_time,omitempty"`
	Satisfy_result               *string    `xorm:" varchar(255) null 'satisfy_result'" indexes:"" comment:"评价结果" json:"satisfy_result,omitempty" bson:",omitempty" msgpack:"satisfy_result,omitempty"`
	Satisfy_level_id             *int64     `xorm:" null 'satisfy_level_id'" indexes:"" comment:"评价等级" json:"satisfy_level_id,omitempty" bson:",omitempty" msgpack:"satisfy_level_id,omitempty" fk:"dic_satisfy_level(satisfy_level_id)"`
	Template_id                  *int64     `xorm:" null 'template_id'" indexes:"" comment:"模板主键" json:"template_id,omitempty" bson:",omitempty" msgpack:"template_id,omitempty" fk:"tab_flow_template(template_id)"`
	Check_id                     *int64     `xorm:" null 'check_id'" indexes:"" comment:"物料最后签入签出主键" json:"check_id,omitempty" bson:",omitempty" msgpack:"check_id,omitempty"`
	Check_type                   *Boolean   `xorm:" null 'check_type'" indexes:"" comment:"物料最后签入签出标识" json:"check_type,omitempty" bson:",omitempty" msgpack:"check_type,omitempty"`
	Check_date                   *time.Time `xorm:" null 'check_date'" indexes:"" comment:"物料最后签入签出时间" json:"check_date,omitempty" bson:",omitempty" msgpack:"check_date,omitempty"`
	Check_attribution_id         *int64     `xorm:" null 'check_attribution_id'" indexes:"" comment:"物料最后签入签出归属地" json:"check_attribution_id,omitempty" bson:",omitempty" msgpack:"check_attribution_id,omitempty"`
	Check_attribution_name       *string    `xorm:" varchar(255) null 'check_attribution_name'" indexes:"" comment:"物料最后签入签出归属地名称" json:"check_attribution_name,omitempty" bson:",omitempty" msgpack:"check_attribution_name,omitempty"`
	Check_user_id                *int64     `xorm:" null 'check_user_id'" indexes:"" comment:"物料最后签入签出人" json:"check_user_id,omitempty" bson:",omitempty" msgpack:"check_user_id,omitempty"`
	Check_user_name              *string    `xorm:" varchar(50) null 'check_user_name'" indexes:"" comment:"物料最后签入签出人名称" json:"check_user_name,omitempty" bson:",omitempty" msgpack:"check_user_name,omitempty"`
	Enable_mark                  *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_6) index(Index_9) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark                  *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_7) index(Index_10) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code                    *int64     `xorm:" null 'sort_code'" indexes:" index(Index_8) " comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description                  *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date                  *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid                *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username              *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date                  *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid                *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username              *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version                      *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: dic_work_status
// 说明: 自治家园字典-从业状况
// 自治家园字典-从业状况
// 
type dic_work_status struct {
	Work_status_id     *int64     `xorm:" pk not null 'work_status_id'" indexes:"" comment:"从业状况主键" json:"work_status_id,omitempty" bson:",omitempty" msgpack:"work_status_id,omitempty"`
	Work_status_name   *string    `xorm:" varchar(255) null 'work_status_name'" indexes:"" comment:"从业状况名称" json:"work_status_name,omitempty" bson:",omitempty" msgpack:"work_status_name,omitempty"`
	Work_status_prompt *string    `xorm:" varchar(255) null 'work_status_prompt'" indexes:"" comment:"从业状况别称" json:"work_status_prompt,omitempty" bson:",omitempty" msgpack:"work_status_prompt,omitempty"`
	Work_status_code   *string    `xorm:" varchar(255) null 'work_status_code'" indexes:"" comment:"从业状况编码" json:"work_status_code,omitempty" bson:",omitempty" msgpack:"work_status_code,omitempty"`
	Enable_mark        *Boolean   `xorm:" null default 1 'enable_mark'" indexes:"" comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark        *Boolean   `xorm:" null default 0 'delete_mark'" indexes:"" comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code          *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description        *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date        *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid      *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username    *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date        *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid      *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username    *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version            *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: tab_group_user_relation
// 说明: 组别用户关系表
// 组别用户关系表
// 
type tab_group_user_relation struct {
	Relation_id     *int64     `xorm:" pk not null 'relation_id'" indexes:"" comment:"关系主键" json:"relation_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"relation_id,omitempty"`
	User_id         *int64     `xorm:" null 'user_id'" indexes:"" comment:"用户主键" json:"user_id,omitempty" bson:",omitempty" msgpack:"user_id,omitempty" fk:"tab_user(user_id)"`
	Group_id        *int64     `xorm:" null 'group_id'" indexes:"" comment:"用户组主键" json:"group_id,omitempty" bson:",omitempty" msgpack:"group_id,omitempty" fk:"tab_group(group_id)"`
	Enable_mark     *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_1) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark     *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_2) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code       *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description     *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date     *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid   *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date     *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid   *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version         *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: tab_policy_interpretation
// 说明: 政策解读
// 政策解读
// 
type tab_policy_interpretation struct {
	Interpretation_id *int64     `xorm:" pk not null 'interpretation_id'" indexes:"" comment:"政策解读主键" json:"interpretation_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"interpretation_id,omitempty"`
	Title             *string    `xorm:" varchar(255) null 'title'" indexes:"" comment:"标题" json:"title,omitempty" bson:",omitempty" msgpack:"title,omitempty"`
	Content           *string    `xorm:" text null 'content'" indexes:"" comment:"内容" json:"content,omitempty" bson:",omitempty" msgpack:"content,omitempty"`
	Relation_file     *string    `xorm:" varchar(255) null 'relation_file'" indexes:"" comment:"相关政策文号" json:"relation_file,omitempty" bson:",omitempty" msgpack:"relation_file,omitempty"`
	Enable_mark       *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_1) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark       *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_2) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code         *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description       *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date       *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid     *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username   *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date       *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid     *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username   *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version           *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: tab_region_schedule
// 说明: 外网-归属地预约时段表
// 外网-归属地预约时段表
// 
type tab_region_schedule struct {
	Schedule_id       *int64     `xorm:" pk not null 'schedule_id'" indexes:"" comment:"预约时间段主键" json:"schedule_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"schedule_id,omitempty"`
	Region_id         *int64     `xorm:" null 'region_id'" indexes:"" comment:"归属地主键" json:"region_id,omitempty" bson:",omitempty" msgpack:"region_id,omitempty"`
	Schedule_duration *string    `xorm:" varchar(255) null 'schedule_duration'" indexes:"" comment:"预约时间段" json:"schedule_duration,omitempty" bson:",omitempty" msgpack:"schedule_duration,omitempty"`
	Pubnumbers        *int64     `xorm:" null 'pubnumbers'" indexes:"" comment:"放号数量" json:"pubnumbers,omitempty" bson:",omitempty" msgpack:"pubnumbers,omitempty"`
	If_pauzed         *Boolean   `xorm:" null default 0 'if_pauzed'" indexes:"" comment:"是否暂停预约放号" json:"if_pauzed,omitempty" bson:",omitempty" msgpack:"if_pauzed,omitempty"`
	Enable_mark       *Boolean   `xorm:" null default 1 'enable_mark'" indexes:"" comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark       *Boolean   `xorm:" null default 0 'delete_mark'" indexes:"" comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code         *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description       *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date       *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid     *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username   *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date       *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid     *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username   *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version           *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: dic_community
// 说明: 字典-行政区划-社区
// 字典-行政区划-社区
// 
type dic_community struct {
	Community_id     *int64     `xorm:" pk not null 'community_id'" indexes:"" comment:"社区主键" json:"community_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"community_id,omitempty"`
	Community_name   *string    `xorm:" varchar(255) null 'community_name'" indexes:"" comment:"社区名称" json:"community_name,omitempty" bson:",omitempty" msgpack:"community_name,omitempty"`
	Community_prompt *string    `xorm:" varchar(255) null 'community_prompt'" indexes:"" comment:"社区简称" json:"community_prompt,omitempty" bson:",omitempty" msgpack:"community_prompt,omitempty"`
	Community_code   *string    `xorm:" varchar(255) null 'community_code'" indexes:" unique(Index_1) " comment:"社区代码" json:"community_code,omitempty" bson:",omitempty" msgpack:"community_code,omitempty"`
	Street_id        *int64     `xorm:" null 'street_id'" indexes:"" comment:"街道主键" json:"street_id,omitempty" bson:",omitempty" msgpack:"street_id,omitempty" fk:"dic_street(street_id)"`
	Community_type   *int64     `xorm:" null 'community_type'" indexes:"" comment:"社区类型" json:"community_type,omitempty" bson:",omitempty" msgpack:"community_type,omitempty"`
	Enable_mark      *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_2) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark      *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_3) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code        *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description      *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date      *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid    *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username  *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date      *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid    *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username  *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version          *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: dic_corporate_subject
// 说明: 字典-法人事项主题
// 字典-法人事项主题
// 
type dic_corporate_subject struct {
	Corporate_subject_id   *int64     `xorm:" pk not null 'corporate_subject_id'" indexes:"" comment:"事项主题主键" json:"corporate_subject_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"corporate_subject_id,omitempty"`
	Affairs_subject_name   *string    `xorm:" varchar(255) null 'affairs_subject_name'" indexes:"" comment:"事项主题名称" json:"affairs_subject_name,omitempty" bson:",omitempty" msgpack:"affairs_subject_name,omitempty"`
	Affairs_subject_prompt *string    `xorm:" varchar(255) null 'affairs_subject_prompt'" indexes:"" comment:"事项主题简称" json:"affairs_subject_prompt,omitempty" bson:",omitempty" msgpack:"affairs_subject_prompt,omitempty"`
	Affairs_subject_code   *string    `xorm:" varchar(255) null 'affairs_subject_code'" indexes:"" comment:"事项主题编码" json:"affairs_subject_code,omitempty" bson:",omitempty" msgpack:"affairs_subject_code,omitempty"`
	Icon_url               *string    `xorm:" varchar(255) null 'icon_url'" indexes:"" comment:"图标" json:"icon_url,omitempty" bson:",omitempty" msgpack:"icon_url,omitempty"`
	Enable_mark            *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_1) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark            *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_2) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code              *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description            *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date            *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid          *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username        *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date            *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid          *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username        *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version                *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: dic_life_events
// 说明: 字典-人生事件
// 字典-人生事件
// 
type dic_life_events struct {
	Life_event_id     *int64     `xorm:" pk not null 'life_event_id'" indexes:"" comment:"人生事件主键" json:"life_event_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"life_event_id,omitempty"`
	Life_event_name   *string    `xorm:" varchar(255) null 'life_event_name'" indexes:"" comment:"人生事件名称" json:"life_event_name,omitempty" bson:",omitempty" msgpack:"life_event_name,omitempty"`
	Life_event_prompt *string    `xorm:" varchar(255) null 'life_event_prompt'" indexes:"" comment:"人生事件简称" json:"life_event_prompt,omitempty" bson:",omitempty" msgpack:"life_event_prompt,omitempty"`
	Life_event_code   *string    `xorm:" varchar(255) null 'life_event_code'" indexes:" unique(Index_1) " comment:"人生事件编码" json:"life_event_code,omitempty" bson:",omitempty" msgpack:"life_event_code,omitempty"`
	Enable_mark       *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_2) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark       *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_3) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code         *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description       *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date       *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid     *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username   *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date       *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid     *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username   *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version           *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: dic_user_level
// 说明: 字典-用户级别
// 字典-用户级别
// 
type dic_user_level struct {
	User_level_id     *int64     `xorm:" pk not null 'user_level_id'" indexes:"" comment:"用户级别主键" json:"user_level_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"user_level_id,omitempty"`
	User_level_name   *string    `xorm:" varchar(50) null 'user_level_name'" indexes:"" comment:"用户级别名称" json:"user_level_name,omitempty" bson:",omitempty" msgpack:"user_level_name,omitempty"`
	User_level_prompt *string    `xorm:" varchar(50) null 'user_level_prompt'" indexes:"" comment:"用户级别简称" json:"user_level_prompt,omitempty" bson:",omitempty" msgpack:"user_level_prompt,omitempty"`
	Value_code        *string    `xorm:" varchar(50) null 'value_code'" indexes:" unique(Index_1) " comment:"用户级别编码" json:"value_code,omitempty" bson:",omitempty" msgpack:"value_code,omitempty"`
	Enable_mark       *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_2) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark       *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_3) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code         *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description       *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date       *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid     *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username   *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date       *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid     *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username   *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version           *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: tab_client_machine
// 说明: WPF客户机信息
// WPF客户机信息
// 
type tab_client_machine struct {
	Machine_id           *int64     `xorm:" pk not null 'machine_id'" indexes:"" comment:"客户机主键" json:"machine_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"machine_id,omitempty"`
	Mashine_identy       *string    `xorm:" varchar(50) null 'mashine_identy'" indexes:"" comment:"客户机识别码" json:"mashine_identy,omitempty" bson:",omitempty" msgpack:"mashine_identy,omitempty"`
	Home_url             *string    `xorm:" varchar(255) null 'home_url'" indexes:"" comment:"主页URL" json:"home_url,omitempty" bson:",omitempty" msgpack:"home_url,omitempty"`
	Window_no            *int64     `xorm:" null 'window_no'" indexes:"" comment:"注册窗口号" json:"window_no,omitempty" bson:",omitempty" msgpack:"window_no,omitempty"`
	Region_id            *int64     `xorm:" null 'region_id'" indexes:"" comment:"归属地主键" json:"region_id,omitempty" bson:",omitempty" msgpack:"region_id,omitempty" fk:"tab_region(region_id)"`
	Region_name          *string    `xorm:" varchar(255) null 'region_name'" indexes:"" comment:"归属地名称" json:"region_name,omitempty" bson:",omitempty" msgpack:"region_name,omitempty"`
	Client_version_no    *float32   `xorm:" null 'client_version_no'" indexes:"" comment:"客户端版本" json:"client_version_no,omitempty" bson:",omitempty" msgpack:"client_version_no,omitempty"`
	Specified_version_no *float32   `xorm:" null 'specified_version_no'" indexes:"" comment:"指定客户端版本" json:"specified_version_no,omitempty" bson:",omitempty" msgpack:"specified_version_no,omitempty"`
	Browser_mode         *Boolean   `xorm:" null default 1 'browser_mode'" indexes:"" comment:"是否浏览器模式" json:"browser_mode,omitempty" bson:",omitempty" msgpack:"browser_mode,omitempty"`
	Plugin_folder        *string    `xorm:" varchar(255) null 'plugin_folder'" indexes:"" comment:"插件加载目录" json:"plugin_folder,omitempty" bson:",omitempty" msgpack:"plugin_folder,omitempty"`
	Enable_mark          *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_1) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark          *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_2) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code            *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description          *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date          *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid        *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username      *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date          *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid        *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username      *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version              *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: tab_user
// 说明: 用户表
// 用户表
// 
type tab_user struct {
	User_id         *int64     `xorm:" pk not null 'user_id'" indexes:"" comment:"用户主键" json:"user_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"user_id,omitempty"`
	User_level_id   *int64     `xorm:" null default 1 'user_level_id'" indexes:"" comment:"用户级别主键" json:"user_level_id,omitempty" bson:",omitempty" msgpack:"user_level_id,omitempty" fk:"dic_user_level(user_level_id)"`
	Access_token    *string    `xorm:" varchar(255) null 'access_token'" indexes:"" comment:"令牌" json:"access_token,omitempty" bson:",omitempty" msgpack:"access_token,omitempty"`
	Account         *string    `xorm:" varchar(255) not null 'account'" indexes:" unique(Index_1) " comment:"登录账号" json:"account,omitempty" bson:",omitempty" msgpack:"account,omitempty"`
	Password        *string    `xorm:" varchar(255) not null 'password'" indexes:"" comment:"密码" json:"password,omitempty" bson:",omitempty" msgpack:"password,omitempty"`
	Idcard          *string    `xorm:" varchar(255) null 'idcard'" indexes:"" comment:"身份证号" json:"idcard,omitempty" bson:",omitempty" msgpack:"idcard,omitempty"`
	User_name       *string    `xorm:" varchar(255) null 'user_name'" indexes:"" comment:"用户姓名" json:"user_name,omitempty" bson:",omitempty" msgpack:"user_name,omitempty"`
	Gender          *int64     `xorm:" null 'gender'" indexes:"" comment:"用户性别" json:"gender,omitempty" bson:",omitempty" msgpack:"gender,omitempty" fk:"dic_gender(gender_id)"`
	Position_id     *int64     `xorm:" null 'position_id'" indexes:"" comment:"主要岗位" json:"position_id,omitempty" bson:",omitempty" msgpack:"position_id,omitempty" fk:"dic_positions(position_id)"`
	User_role       *int64     `xorm:" null 'user_role'" indexes:"" comment:"用户角色" json:"user_role,omitempty" bson:",omitempty" msgpack:"user_role,omitempty"`
	Phone           *string    `xorm:" varchar(255) null 'phone'" indexes:"" comment:"手机号" json:"phone,omitempty" bson:",omitempty" msgpack:"phone,omitempty"`
	Worker_no       *string    `xorm:" varchar(255) not null 'worker_no'" indexes:"" comment:"工号" json:"worker_no,omitempty" bson:",omitempty" msgpack:"worker_no,omitempty"`
	Department      *int64     `xorm:" null 'department'" indexes:"" comment:"所属业务部门" json:"department,omitempty" bson:",omitempty" msgpack:"department,omitempty"`
	Location_id     *int64     `xorm:" null 'location_id'" indexes:"" comment:"归属地" json:"location_id,omitempty" bson:",omitempty" msgpack:"location_id,omitempty"`
	Finger_print    *string    `xorm:" varchar(255) null 'finger_print'" indexes:"" comment:"指纹模版" json:"finger_print,omitempty" bson:",omitempty" msgpack:"finger_print,omitempty"`
	Error_times     *int64     `xorm:" null 'error_times'" indexes:"" comment:"密码错误登录的次数" json:"error_times,omitempty" bson:",omitempty" msgpack:"error_times,omitempty"`
	Error_date      *time.Time `xorm:" null 'error_date'" indexes:"" comment:"密码错误登录的时间" json:"error_date,omitempty" bson:",omitempty" msgpack:"error_date,omitempty"`
	Potrait_url     *string    `xorm:" varchar(255) null 'potrait_url'" indexes:"" comment:"照片" json:"potrait_url,omitempty" bson:",omitempty" msgpack:"potrait_url,omitempty"`
	Position        *int64     `xorm:" null 'position'" indexes:"" comment:"大厅岗位" json:"position,omitempty" bson:",omitempty" msgpack:"position,omitempty"`
	Enable_mark     *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_2) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark     *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_3) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code       *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description     *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date     *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid   *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date     *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid   *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version         *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: dic_node_type
// 说明: 字典-流程节点类型
// 字典-流程节点类型
// 
type dic_node_type struct {
	Node_type_id     *int64     `xorm:" pk not null 'node_type_id'" indexes:"" comment:"节点类型主键" json:"node_type_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"node_type_id,omitempty"`
	Node_type_name   *string    `xorm:" varchar(50) null 'node_type_name'" indexes:"" comment:"节点类型名称" json:"node_type_name,omitempty" bson:",omitempty" msgpack:"node_type_name,omitempty"`
	Node_type_prompt *string    `xorm:" varchar(50) null 'node_type_prompt'" indexes:"" comment:"节点类型简称" json:"node_type_prompt,omitempty" bson:",omitempty" msgpack:"node_type_prompt,omitempty"`
	Value_code       *string    `xorm:" varchar(50) null 'value_code'" indexes:" unique(Index_1) " comment:"编码" json:"value_code,omitempty" bson:",omitempty" msgpack:"value_code,omitempty"`
	Enable_mark      *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_2) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark      *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_3) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code        *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description      *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date      *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid    *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username  *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date      *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid    *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username  *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version          *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: tab_affair_tags
// 说明: 事项标签
// 事项标签
// 
type tab_affair_tags struct {
	Tag_id          *int64     `xorm:" pk not null 'tag_id'" indexes:"" comment:"标签主键" json:"tag_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"tag_id,omitempty"`
	Tag_name        *string    `xorm:" varchar(255) null 'tag_name'" indexes:" unique(Index_1) " comment:"标签名称" json:"tag_name,omitempty" bson:",omitempty" msgpack:"tag_name,omitempty"`
	Is_public       *Boolean   `xorm:" null default 0 'is_public'" indexes:"" comment:"是否公共标签" json:"is_public,omitempty" bson:",omitempty" msgpack:"is_public,omitempty"`
	Offset          *int64     `xorm:" null default 0 'offset'" indexes:"" comment:"偏移量" json:"offset,omitempty" bson:",omitempty" msgpack:"offset,omitempty"`
	Bitmap_content  []uint8    `xorm:" null 'bitmap_content'" indexes:"" comment:"位图内容" json:"bitmap_content,omitempty" bson:",omitempty" msgpack:"bitmap_content,omitempty"`
	Tag_type        *string    `xorm:" varchar(255) null 'tag_type'" indexes:" unique(Index_1) " comment:"标签分类" json:"tag_type,omitempty" bson:",omitempty" msgpack:"tag_type,omitempty"`
	Enable_mark     *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_2) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark     *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_3) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code       *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description     *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date     *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid   *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date     *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid   *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version         *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: tab_booking
// 说明: 外网-网上预约表
// 外网-网上预约表
// 
type tab_booking struct {
	Booking_id        *int64     `xorm:" pk not null 'booking_id'" indexes:"" comment:"预约主键" json:"booking_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"booking_id,omitempty"`
	Region_id         *int64     `xorm:" null 'region_id'" indexes:"" comment:"归属地主键" json:"region_id,omitempty" bson:",omitempty" msgpack:"region_id,omitempty"`
	Region_name       *string    `xorm:" varchar(255) null 'region_name'" indexes:"" comment:"归属地名称" json:"region_name,omitempty" bson:",omitempty" msgpack:"region_name,omitempty"`
	Out_user_id       *int64     `xorm:" null 'out_user_id'" indexes:"" comment:"外网用户主键" json:"out_user_id,omitempty" bson:",omitempty" msgpack:"out_user_id,omitempty" fk:"tab_outer_user(out_user_id)"`
	Affairs_id        *int64     `xorm:" null 'affairs_id'" indexes:"" comment:"事项主键" json:"affairs_id,omitempty" bson:",omitempty" msgpack:"affairs_id,omitempty"`
	Affairs_name      *string    `xorm:" varchar(255) null 'affairs_name'" indexes:"" comment:"事项名称" json:"affairs_name,omitempty" bson:",omitempty" msgpack:"affairs_name,omitempty"`
	Schedule_id       *int64     `xorm:" null 'schedule_id'" indexes:"" comment:"预约时间段主键" json:"schedule_id,omitempty" bson:",omitempty" msgpack:"schedule_id,omitempty" fk:"tab_region_schedule(schedule_id)"`
	Booking_date      *string    `xorm:" varchar(255) null 'booking_date'" indexes:"" comment:"预约日期" json:"booking_date,omitempty" bson:",omitempty" msgpack:"booking_date,omitempty"`
	Booking_seq       *string    `xorm:" varchar(255) null 'booking_seq'" indexes:"" comment:"预约号" json:"booking_seq,omitempty" bson:",omitempty" msgpack:"booking_seq,omitempty"`
	Finished_tag      *Boolean   `xorm:" null default 0 'finished_tag'" indexes:"" comment:"是否完成" json:"finished_tag,omitempty" bson:",omitempty" msgpack:"finished_tag,omitempty"`
	Finished_time     *time.Time `xorm:" null 'finished_time'" indexes:"" comment:"完成时间" json:"finished_time,omitempty" bson:",omitempty" msgpack:"finished_time,omitempty"`
	Finshed_window_no *int64     `xorm:" null 'finshed_window_no'" indexes:"" comment:"完成窗口号" json:"finshed_window_no,omitempty" bson:",omitempty" msgpack:"finshed_window_no,omitempty"`
	Finished_operator *int64     `xorm:" null 'finished_operator'" indexes:"" comment:"完成操作人" json:"finished_operator,omitempty" bson:",omitempty" msgpack:"finished_operator,omitempty"`
	Enable_mark       *Boolean   `xorm:" null default 1 'enable_mark'" indexes:"" comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark       *Boolean   `xorm:" null default 0 'delete_mark'" indexes:"" comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code         *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description       *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date       *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid     *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username   *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date       *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid     *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username   *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version           *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: dic_administrative_level
// 说明: 字典-行政级别
// 字典-行政级别
// 
type dic_administrative_level struct {
	Level_id        *int64     `xorm:" pk not null 'level_id'" indexes:"" comment:"行政等级主键" json:"level_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"level_id,omitempty"`
	Level_name      *string    `xorm:" varchar(255) null 'level_name'" indexes:" unique(Index_1) " comment:"行政等级名称" json:"level_name,omitempty" bson:",omitempty" msgpack:"level_name,omitempty"`
	Level_prompt    *string    `xorm:" varchar(255) null 'level_prompt'" indexes:"" comment:"行政等级简称" json:"level_prompt,omitempty" bson:",omitempty" msgpack:"level_prompt,omitempty"`
	Value_code      *string    `xorm:" varchar(255) null 'value_code'" indexes:" unique(Index_2) " comment:"行政等级编码" json:"value_code,omitempty" bson:",omitempty" msgpack:"value_code,omitempty"`
	Enable_mark     *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_3) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark     *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_4) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code       *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description     *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date     *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid   *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date     *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid   *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version         *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: dic_city
// 说明: 字典-行政区划-地市
// 字典-行政区划-地市
// 
type dic_city struct {
	City_id         *int64     `xorm:" pk not null 'city_id'" indexes:"" comment:"地市主键" json:"city_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"city_id,omitempty"`
	City_name       *string    `xorm:" varchar(255) null 'city_name'" indexes:"" comment:"地市名称" json:"city_name,omitempty" bson:",omitempty" msgpack:"city_name,omitempty"`
	City_prompt     *string    `xorm:" varchar(255) null 'city_prompt'" indexes:"" comment:"地市简称" json:"city_prompt,omitempty" bson:",omitempty" msgpack:"city_prompt,omitempty"`
	City_code       *string    `xorm:" varchar(255) null 'city_code'" indexes:" unique(Index_1) " comment:"地市代码" json:"city_code,omitempty" bson:",omitempty" msgpack:"city_code,omitempty"`
	Enable_mark     *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_2) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark     *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_3) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code       *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description     *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date     *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid   *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date     *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid   *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version         *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: tab_region_photos
// 说明: 查询机-机构图片
// 查询机-机构图片
// 
type tab_region_photos struct {
	Region_photo_id *int64     `xorm:" pk not null 'region_photo_id'" indexes:"" comment:"机构图片主键" json:"region_photo_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"region_photo_id,omitempty"`
	Region_id       *int64     `xorm:" null 'region_id'" indexes:"" comment:"归属地主键" json:"region_id,omitempty" bson:",omitempty" msgpack:"region_id,omitempty" fk:"tab_region(region_id)"`
	Photo_url       *string    `xorm:" varchar(255) null 'photo_url'" indexes:"" comment:"图片地址" json:"photo_url,omitempty" bson:",omitempty" msgpack:"photo_url,omitempty"`
	Enable_mark     *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_1) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark     *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_2) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code       *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description     *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date     *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid   *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date     *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid   *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version         *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: tab_credential_catalog
// 说明: 证照目录信息
// 证照目录信息
// 
type tab_credential_catalog struct {
	Credential_catalog_id   *int64     `xorm:" pk not null 'credential_catalog_id'" indexes:"" comment:"证照目录主键" json:"credential_catalog_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"credential_catalog_id,omitempty"`
	Credential_catalog_code *string    `xorm:" varchar(50) null 'credential_catalog_code'" indexes:"" comment:"证照目录编码" json:"credential_catalog_code,omitempty" bson:",omitempty" msgpack:"credential_catalog_code,omitempty"`
	Credential_name         *string    `xorm:" varchar(255) null 'credential_name'" indexes:"" comment:"证照名称" json:"credential_name,omitempty" bson:",omitempty" msgpack:"credential_name,omitempty"`
	Credential_type         *int64     `xorm:" null 'credential_type'" indexes:"" comment:"证照类别" json:"credential_type,omitempty" bson:",omitempty" msgpack:"credential_type,omitempty"`
	Grant_object            *int64     `xorm:" null 'grant_object'" indexes:"" comment:"证照授予对象" json:"grant_object,omitempty" bson:",omitempty" msgpack:"grant_object,omitempty"`
	Enable_mark             *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_1) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark             *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_2) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code               *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description             *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date             *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid           *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username         *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date             *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid           *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username         *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version                 *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: tab_files
// 说明: 自治家园-文件表
// 自治家园-文件表
// 
type tab_files struct {
	File_id         *int64     `xorm:" pk not null 'file_id'" indexes:"" comment:"文件主键" json:"file_id,omitempty" bson:",omitempty" msgpack:"file_id,omitempty"`
	Extend_id       *int64     `xorm:" not null 'extend_id'" indexes:"" comment:"相关的主键" json:"extend_id,omitempty" bson:",omitempty" msgpack:"extend_id,omitempty"`
	File_show_name  *string    `xorm:" varchar(255) not null 'file_show_name'" indexes:"" comment:"文件名称" json:"file_show_name,omitempty" bson:",omitempty" msgpack:"file_show_name,omitempty"`
	File_ext        *string    `xorm:" varchar(255) not null 'file_ext'" indexes:"" comment:"文件扩展名" json:"file_ext,omitempty" bson:",omitempty" msgpack:"file_ext,omitempty"`
	File_path       *string    `xorm:" varchar(255) not null 'file_path'" indexes:"" comment:"存贮路径" json:"file_path,omitempty" bson:",omitempty" msgpack:"file_path,omitempty"`
	File_type       *int64     `xorm:" not null 'file_type'" indexes:"" comment:"标志是哪张表的附件" json:"file_type,omitempty" bson:",omitempty" msgpack:"file_type,omitempty"`
	File_kind       *string    `xorm:" varchar(255) null 'file_kind'" indexes:"" comment:"文件种类" json:"file_kind,omitempty" bson:",omitempty" msgpack:"file_kind,omitempty"`
	Enable_mark     *Boolean   `xorm:" null default 1 'enable_mark'" indexes:"" comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark     *Boolean   `xorm:" null default 0 'delete_mark'" indexes:"" comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code       *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description     *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date     *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid   *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date     *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid   *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version         *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: tab_region_honors
// 说明: 查询机-所获荣誉
// 查询机-所获荣誉
// 
type tab_region_honors struct {
	Honor_id        *int64     `xorm:" pk not null 'honor_id'" indexes:"" comment:"荣誉主键" json:"honor_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"honor_id,omitempty"`
	Region_id       *int64     `xorm:" null 'region_id'" indexes:"" comment:"归属地主键" json:"region_id,omitempty" bson:",omitempty" msgpack:"region_id,omitempty" fk:"tab_region(region_id)"`
	Issue_date      *time.Time `xorm:" null 'issue_date'" indexes:"" comment:"颁发时间" json:"issue_date,omitempty" bson:",omitempty" msgpack:"issue_date,omitempty"`
	Issuer          *string    `xorm:" varchar(50) null 'issuer'" indexes:"" comment:"颁发机构" json:"issuer,omitempty" bson:",omitempty" msgpack:"issuer,omitempty"`
	Honor_desc      *string    `xorm:" varchar(255) null 'honor_desc'" indexes:"" comment:"荣誉描述" json:"honor_desc,omitempty" bson:",omitempty" msgpack:"honor_desc,omitempty"`
	Photo_url       *string    `xorm:" varchar(255) null 'photo_url'" indexes:"" comment:"附属照片" json:"photo_url,omitempty" bson:",omitempty" msgpack:"photo_url,omitempty"`
	Enable_mark     *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_1) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark     *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_2) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code       *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description     *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date     *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid   *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date     *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid   *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version         *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: tab_application_db_config
// 说明: 应用数据配置
// 应用数据配置
// 
type tab_application_db_config struct {
	Dal_id          *int64     `xorm:" null 'dal_id'" indexes:"" comment:"DAL主键" json:"dal_id,omitempty" bson:",omitempty" msgpack:"dal_id,omitempty" fk:"tab_dal_deploy(dal_id)"`
	App_code        *string    `xorm:" varchar(50) null 'app_code'" indexes:"" comment:"应用代码" json:"app_code,omitempty" bson:",omitempty" msgpack:"app_code,omitempty"`
	Dbtype          *string    `xorm:" varchar(50) null 'dbtype'" indexes:"" comment:"数据库类型" json:"dbtype,omitempty" bson:",omitempty" msgpack:"dbtype,omitempty"`
	Dbip            *string    `xorm:" varchar(50) null 'dbip'" indexes:"" comment:"数据库地址" json:"dbip,omitempty" bson:",omitempty" msgpack:"dbip,omitempty"`
	Dbport          *int64     `xorm:" null 'dbport'" indexes:"" comment:"数据库端口" json:"dbport,omitempty" bson:",omitempty" msgpack:"dbport,omitempty"`
	Dbuser          *string    `xorm:" varchar(50) null 'dbuser'" indexes:"" comment:"数据库用户" json:"dbuser,omitempty" bson:",omitempty" msgpack:"dbuser,omitempty"`
	Dbpwd           *string    `xorm:" varchar(50) null 'dbpwd'" indexes:"" comment:"数据库密码" json:"dbpwd,omitempty" bson:",omitempty" msgpack:"dbpwd,omitempty"`
	Dbname          *string    `xorm:" varchar(50) null 'dbname'" indexes:"" comment:"数据库名称" json:"dbname,omitempty" bson:",omitempty" msgpack:"dbname,omitempty"`
	Autoincr        *Boolean   `xorm:" null 'autoincr'" indexes:"" comment:"主键是否自增长" json:"autoincr,omitempty" bson:",omitempty" msgpack:"autoincr,omitempty"`
	Initialized     *Boolean   `xorm:" null default 0 'initialized'" indexes:"" comment:"是否已初始化" json:"initialized,omitempty" bson:",omitempty" msgpack:"initialized,omitempty"`
	Enable_mark     *Boolean   `xorm:" null default 1 'enable_mark'" indexes:"" comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark     *Boolean   `xorm:" null default 0 'delete_mark'" indexes:"" comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code       *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description     *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date     *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid   *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date     *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid   *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version         *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: tab_dal_deploy
// 说明: DAL部署表
// DAL部署表
// 
type tab_dal_deploy struct {
	Dal_id             *int64     `xorm:" pk not null 'dal_id'" indexes:"" comment:"DAL主键" json:"dal_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"dal_id,omitempty"`
	Dal_key            *string    `xorm:" varchar(50) not null 'dal_key'" indexes:" unique(Index_1) " comment:"DAL部署标识" json:"dal_key,omitempty" bson:",omitempty" msgpack:"dal_key,omitempty"`
	Dal_district_name  *string    `xorm:" varchar(255) null 'dal_district_name'" indexes:"" comment:"DAL行政区划名" json:"dal_district_name,omitempty" bson:",omitempty" msgpack:"dal_district_name,omitempty"`
	Dal_name           *string    `xorm:" varchar(255) not null 'dal_name'" indexes:"" comment:"DAL名称" json:"dal_name,omitempty" bson:",omitempty" msgpack:"dal_name,omitempty"`
	Dal_position       *string    `xorm:" text not null 'dal_position'" indexes:"" comment:"DAL部署位置" json:"dal_position,omitempty" bson:",omitempty" msgpack:"dal_position,omitempty"`
	Pid                *int64     `xorm:" not null 'pid'" indexes:"" comment:"省区代码" json:"pid,omitempty" bson:",omitempty" msgpack:"pid,omitempty"`
	Cid                *int64     `xorm:" not null 'cid'" indexes:"" comment:"地市代码" json:"cid,omitempty" bson:",omitempty" msgpack:"cid,omitempty"`
	Did                *int64     `xorm:" not null 'did'" indexes:"" comment:"区县代码" json:"did,omitempty" bson:",omitempty" msgpack:"did,omitempty"`
	Chipid             *int64     `xorm:" not null 'chipid'" indexes:" unique(Index_2) " comment:"CHIPID" json:"chipid,omitempty" bson:",omitempty" msgpack:"chipid,omitempty"`
	Tm                 *string    `xorm:" varchar(20) not null default '0 0 1 * * ?' 'tm'" indexes:"" comment:"行政监察定时定规则" json:"tm,omitempty" bson:",omitempty" msgpack:"tm,omitempty"`
	Dbtype             *string    `xorm:" varchar(50) not null 'dbtype'" indexes:"" comment:"数据库类型" json:"dbtype,omitempty" bson:",omitempty" msgpack:"dbtype,omitempty"`
	Dbuser             *string    `xorm:" varchar(50) not null 'dbuser'" indexes:"" comment:"数据库用户" json:"dbuser,omitempty" bson:",omitempty" msgpack:"dbuser,omitempty"`
	Dbpwd              *string    `xorm:" varchar(50) not null 'dbpwd'" indexes:"" comment:"数据库密码" json:"dbpwd,omitempty" bson:",omitempty" msgpack:"dbpwd,omitempty"`
	Dbport             *int64     `xorm:" not null 'dbport'" indexes:"" comment:"数据库端口" json:"dbport,omitempty" bson:",omitempty" msgpack:"dbport,omitempty"`
	Dbip               *string    `xorm:" varchar(50) not null default 'dal.rel.egov-china.com' 'dbip'" indexes:"" comment:"数据库地址" json:"dbip,omitempty" bson:",omitempty" msgpack:"dbip,omitempty"`
	Dbname             *string    `xorm:" varchar(50) not null default 'egov' 'dbname'" indexes:"" comment:"数据库库名" json:"dbname,omitempty" bson:",omitempty" msgpack:"dbname,omitempty"`
	Commit_count       *int64     `xorm:" not null default 1000 'commit_count'" indexes:"" comment:"初始化数据时最多提交记录数" json:"commit_count,omitempty" bson:",omitempty" msgpack:"commit_count,omitempty"`
	Sync               *int64     `xorm:" null default 1 'sync'" indexes:"" comment:"是否同步结构和外键" json:"sync,omitempty" bson:",omitempty" msgpack:"sync,omitempty"`
	Debug              *int64     `xorm:" null default 1 'debug'" indexes:"" comment:"是否打开调试信息" json:"debug,omitempty" bson:",omitempty" msgpack:"debug,omitempty"`
	Docs               *int64     `xorm:" null default 1 'docs'" indexes:"" comment:"是否打开文档模式" json:"docs,omitempty" bson:",omitempty" msgpack:"docs,omitempty"`
	Port               *int64     `xorm:" null default 9008 'port'" indexes:"" comment:"DAL端口" json:"port,omitempty" bson:",omitempty" msgpack:"port,omitempty"`
	Separate_dbtype    *string    `xorm:" varchar(50) null 'separate_dbtype'" indexes:"" comment:"分离数据库类型" json:"separate_dbtype,omitempty" bson:",omitempty" msgpack:"separate_dbtype,omitempty"`
	Separate_dbuser    *string    `xorm:" varchar(50) null 'separate_dbuser'" indexes:"" comment:"分离数据库用户" json:"separate_dbuser,omitempty" bson:",omitempty" msgpack:"separate_dbuser,omitempty"`
	Separate_dbpwd     *string    `xorm:" varchar(50) null 'separate_dbpwd'" indexes:"" comment:"分离数据库密码" json:"separate_dbpwd,omitempty" bson:",omitempty" msgpack:"separate_dbpwd,omitempty"`
	Separate_dbport    *int64     `xorm:" null 'separate_dbport'" indexes:"" comment:"分离数据库端口" json:"separate_dbport,omitempty" bson:",omitempty" msgpack:"separate_dbport,omitempty"`
	Separate_dbip      *string    `xorm:" varchar(50) null default 'dal.rel.egov-china.com' 'separate_dbip'" indexes:"" comment:"分离数据库地址" json:"separate_dbip,omitempty" bson:",omitempty" msgpack:"separate_dbip,omitempty"`
	Separate_dbname    *string    `xorm:" varchar(50) null default 'egov_separate' 'separate_dbname'" indexes:"" comment:"分离数据库库名" json:"separate_dbname,omitempty" bson:",omitempty" msgpack:"separate_dbname,omitempty"`
	First_login_time   *time.Time `xorm:" null 'first_login_time'" indexes:"" comment:"首次上线时间" json:"first_login_time,omitempty" bson:",omitempty" msgpack:"first_login_time,omitempty"`
	First_login_ip     *string    `xorm:" varchar(50) null 'first_login_ip'" indexes:"" comment:"首次上线地址" json:"first_login_ip,omitempty" bson:",omitempty" msgpack:"first_login_ip,omitempty"`
	Last_login_time    *time.Time `xorm:" null 'last_login_time'" indexes:"" comment:"上次上线时间" json:"last_login_time,omitempty" bson:",omitempty" msgpack:"last_login_time,omitempty"`
	Last_login_ip      *string    `xorm:" varchar(50) null 'last_login_ip'" indexes:"" comment:"上次上线地址" json:"last_login_ip,omitempty" bson:",omitempty" msgpack:"last_login_ip,omitempty"`
	Is_online          *Boolean   `xorm:" null default 0 'is_online'" indexes:"" comment:"DAL是否在线" json:"is_online,omitempty" bson:",omitempty" msgpack:"is_online,omitempty"`
	Msg_server         *string    `xorm:" varchar(50) null default 'dalud.egov-china.com' 'msg_server'" indexes:"" comment:"消息服务器地址" json:"msg_server,omitempty" bson:",omitempty" msgpack:"msg_server,omitempty"`
	Msg_port           *int64     `xorm:" null default 9020 'msg_port'" indexes:"" comment:"消息服务器端口" json:"msg_port,omitempty" bson:",omitempty" msgpack:"msg_port,omitempty"`
	Sync_data          *int64     `xorm:" null default 1 'sync_data'" indexes:"" comment:"同步所有字典" json:"sync_data,omitempty" bson:",omitempty" msgpack:"sync_data,omitempty"`
	Local_msg_server   *string    `xorm:" varchar(50) null default 'dalud.egov-china.com' 'local_msg_server'" indexes:"" comment:"本地消息服务器地址" json:"local_msg_server,omitempty" bson:",omitempty" msgpack:"local_msg_server,omitempty"`
	Local_msg_port     *int64     `xorm:" null default 9020 'local_msg_port'" indexes:"" comment:"本地消息服务器端口" json:"local_msg_port,omitempty" bson:",omitempty" msgpack:"local_msg_port,omitempty"`
	Sync_online        *Boolean   `xorm:" null 'sync_online'" indexes:"" comment:"同步服务在线状态" json:"sync_online,omitempty" bson:",omitempty" msgpack:"sync_online,omitempty"`
	Sync_status        *int64     `xorm:" null 'sync_status'" indexes:"" comment:"同步服务运行状态" json:"sync_status,omitempty" bson:",omitempty" msgpack:"sync_status,omitempty"`
	Sync_error         *string    `xorm:" text null 'sync_error'" indexes:"" comment:"同步服务运行错误" json:"sync_error,omitempty" bson:",omitempty" msgpack:"sync_error,omitempty"`
	Replication_status *Boolean   `xorm:" null 'replication_status'" indexes:"" comment:"复制运行状态" json:"replication_status,omitempty" bson:",omitempty" msgpack:"replication_status,omitempty"`
	Replication_error  *string    `xorm:" text null 'replication_error'" indexes:"" comment:"复制运行错误" json:"replication_error,omitempty" bson:",omitempty" msgpack:"replication_error,omitempty"`
	Log_server         *string    `xorm:" varchar(50) null 'log_server'" indexes:"" comment:"日志服务地址" json:"log_server,omitempty" bson:",omitempty" msgpack:"log_server,omitempty"`
	Log_level          *int64     `xorm:" null default 1 'log_level'" indexes:"" comment:"日志级别" json:"log_level,omitempty" bson:",omitempty" msgpack:"log_level,omitempty"`
	Enable_mark        *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_3) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark        *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_4) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code          *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description        *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date        *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid      *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username    *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date        *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid      *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username    *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version            *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: tab_legal_person_info
// 说明: 法定代表人信息
// 法定代表人信息
// 
type tab_legal_person_info struct {
	Legal_person_id   *int64     `xorm:" pk not null 'legal_person_id'" indexes:"" comment:"法定代表人主键" json:"legal_person_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"legal_person_id,omitempty"`
	Credit_code       *string    `xorm:" varchar(255) null 'credit_code'" indexes:"" comment:"社会统一信用代码" json:"credit_code,omitempty" bson:",omitempty" msgpack:"credit_code,omitempty"`
	Legal_person_name *string    `xorm:" varchar(50) null 'legal_person_name'" indexes:"" comment:"法定代表人姓名" json:"legal_person_name,omitempty" bson:",omitempty" msgpack:"legal_person_name,omitempty"`
	Document_type     *int64     `xorm:" null 'document_type'" indexes:"" comment:"法定代表人证件类型" json:"document_type,omitempty" bson:",omitempty" msgpack:"document_type,omitempty"`
	Document_code     *string    `xorm:" varchar(50) null 'document_code'" indexes:"" comment:"法定代表人证件号码" json:"document_code,omitempty" bson:",omitempty" msgpack:"document_code,omitempty"`
	Enable_mark       *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_1) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark       *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_2) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code         *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description       *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date       *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid     *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username   *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date       *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid     *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username   *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version           *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: tab_monitors
// 说明: 行政监察表
// 行政监察表
// 
type tab_monitors struct {
	Monitor_id               *int64     `xorm:" pk not null 'monitor_id'" indexes:"" comment:"监察主键" json:"monitor_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"monitor_id,omitempty"`
	Is_affair                *Boolean   `xorm:" null default 1 'is_affair'" indexes:" index(Index_7) " comment:"节点或事项" json:"is_affair,omitempty" bson:",omitempty" msgpack:"is_affair,omitempty"`
	Red_or_yellow            *int64     `xorm:" not null default 1 'red_or_yellow'" indexes:" index(Index_3) " comment:"红黄牌类型" json:"red_or_yellow,omitempty" bson:",omitempty" msgpack:"red_or_yellow,omitempty"`
	Is_finished              *Boolean   `xorm:" null default 0 'is_finished'" indexes:" index(Index_2) " comment:"完成状态" json:"is_finished,omitempty" bson:",omitempty" msgpack:"is_finished,omitempty"`
	Begin_date               *time.Time `xorm:" null 'begin_date'" indexes:"" comment:"开始时间" json:"begin_date,omitempty" bson:",omitempty" msgpack:"begin_date,omitempty"`
	End_date                 *time.Time `xorm:" null 'end_date'" indexes:"" comment:"结束时间" json:"end_date,omitempty" bson:",omitempty" msgpack:"end_date,omitempty"`
	Include_received_day     *Boolean   `xorm:" null 'include_received_day'" indexes:"" comment:"处理时间是否包括到达当天" json:"include_received_day,omitempty" bson:",omitempty" msgpack:"include_received_day,omitempty"`
	Time_limit_type_id       *int64     `xorm:" null 'time_limit_type_id'" indexes:"" comment:"时限类型主键" json:"time_limit_type_id,omitempty" bson:",omitempty" msgpack:"time_limit_type_id,omitempty" fk:"dic_time_limit_type(time_limit_type_id)"`
	Time_limit               *int64     `xorm:" null default 1 'time_limit'" indexes:"" comment:"时限" json:"time_limit,omitempty" bson:",omitempty" msgpack:"time_limit,omitempty"`
	Limit_date               *time.Time `xorm:" null 'limit_date'" indexes:"" comment:"时限日期" json:"limit_date,omitempty" bson:",omitempty" msgpack:"limit_date,omitempty"`
	Mark_id                  *int64     `xorm:" null 'mark_id'" indexes:"" comment:"审核主键" json:"mark_id,omitempty" bson:",omitempty" msgpack:"mark_id,omitempty" fk:"tab_affairs_mark_result(mark_id)"`
	Node_id                  *int64     `xorm:" null 'node_id'" indexes:"" comment:"节点编号" json:"node_id,omitempty" bson:",omitempty" msgpack:"node_id,omitempty" fk:"tab_flow_node_info(node_id)"`
	Standard_code_id         *int64     `xorm:" null 'standard_code_id'" indexes:"" comment:"节点标准编码主键" json:"standard_code_id,omitempty" bson:",omitempty" msgpack:"standard_code_id,omitempty" fk:"dic_node_standard_code(standard_code_id)"`
	Affairs_info_id          *int64     `xorm:" null 'affairs_info_id'" indexes:"" comment:"受理主键" json:"affairs_info_id,omitempty" bson:",omitempty" msgpack:"affairs_info_id,omitempty" fk:"tab_affairs_info(affairs_info_id)"`
	Is_hook                  *int64     `xorm:" null default 0 'is_hook'" indexes:"" comment:"是否挂起" json:"is_hook,omitempty" bson:",omitempty" msgpack:"is_hook,omitempty"`
	Op_user_id               *int64     `xorm:" null 'op_user_id'" indexes:"" comment:"最后操作人主键" json:"op_user_id,omitempty" bson:",omitempty" msgpack:"op_user_id,omitempty"`
	Op_user_name             *string    `xorm:" varchar(50) null 'op_user_name'" indexes:"" comment:"最后操作人名称" json:"op_user_name,omitempty" bson:",omitempty" msgpack:"op_user_name,omitempty"`
	Warnning_deadline_date   *time.Time `xorm:" null 'warnning_deadline_date'" indexes:"" comment:"黄牌开始时间" json:"warnning_deadline_date,omitempty" bson:",omitempty" msgpack:"warnning_deadline_date,omitempty"`
	Warnning_deadline        *int64     `xorm:" null 'warnning_deadline'" indexes:"" comment:"工作日黄牌倒计时" json:"warnning_deadline,omitempty" bson:",omitempty" msgpack:"warnning_deadline,omitempty"`
	Nature_warnning_deadline *int64     `xorm:" null 'nature_warnning_deadline'" indexes:"" comment:"自然日黄牌倒计时" json:"nature_warnning_deadline,omitempty" bson:",omitempty" msgpack:"nature_warnning_deadline,omitempty"`
	Deadline_date            *time.Time `xorm:" null 'deadline_date'" indexes:"" comment:"红牌开始时间" json:"deadline_date,omitempty" bson:",omitempty" msgpack:"deadline_date,omitempty"`
	Deadline                 *int64     `xorm:" null 'deadline'" indexes:"" comment:"工作日红牌倒计时" json:"deadline,omitempty" bson:",omitempty" msgpack:"deadline,omitempty"`
	Nature_deadline          *int64     `xorm:" null 'nature_deadline'" indexes:"" comment:"自然日红牌倒计时" json:"nature_deadline,omitempty" bson:",omitempty" msgpack:"nature_deadline,omitempty"`
	Last_compute_time        *time.Time `xorm:" null 'last_compute_time'" indexes:"" comment:"最近计算时间" json:"last_compute_time,omitempty" bson:",omitempty" msgpack:"last_compute_time,omitempty"`
	Red_card_status          *int64     `xorm:" null default 0 'red_card_status'" indexes:"" comment:"红牌状态" json:"red_card_status,omitempty" bson:",omitempty" msgpack:"red_card_status,omitempty"`
	Yellow_card_status       *int64     `xorm:" null default 0 'yellow_card_status'" indexes:"" comment:"黄牌状态" json:"yellow_card_status,omitempty" bson:",omitempty" msgpack:"yellow_card_status,omitempty"`
	Last_hook_time           *time.Time `xorm:" null 'last_hook_time'" indexes:"" comment:"最后挂起时间" json:"last_hook_time,omitempty" bson:",omitempty" msgpack:"last_hook_time,omitempty"`
	Timing_reset             *Boolean   `xorm:" null 'timing_reset'" indexes:"" comment:"事项是否重新计时" json:"timing_reset,omitempty" bson:",omitempty" msgpack:"timing_reset,omitempty"`
	Hook_comment             *string    `xorm:" text null 'hook_comment'" indexes:"" comment:"挂起备注" json:"hook_comment,omitempty" bson:",omitempty" msgpack:"hook_comment,omitempty"`
	Transfer_days            *int64     `xorm:" null default 3 'transfer_days'" indexes:"" comment:"黄转红超时天数" json:"transfer_days,omitempty" bson:",omitempty" msgpack:"transfer_days,omitempty"`
	History_yellow_status    *int64     `xorm:" null default 0 'history_yellow_status'" indexes:"" comment:"历史节点黄牌数" json:"history_yellow_status,omitempty" bson:",omitempty" msgpack:"history_yellow_status,omitempty"`
	History_red_status       *int64     `xorm:" null default 0 'history_red_status'" indexes:"" comment:"历史节点红牌数" json:"history_red_status,omitempty" bson:",omitempty" msgpack:"history_red_status,omitempty"`
	Current_yellow_status    *int64     `xorm:" null default 0 'current_yellow_status'" indexes:"" comment:"当前节点黄牌数" json:"current_yellow_status,omitempty" bson:",omitempty" msgpack:"current_yellow_status,omitempty"`
	Current_red_status       *int64     `xorm:" null default 0 'current_red_status'" indexes:"" comment:"当前节点红牌数" json:"current_red_status,omitempty" bson:",omitempty" msgpack:"current_red_status,omitempty"`
	Current_red_or_yellow    *int64     `xorm:" null default 0 'current_red_or_yellow'" indexes:" index(Index_4) " comment:"当前节点红黄牌类型" json:"current_red_or_yellow,omitempty" bson:",omitempty" msgpack:"current_red_or_yellow,omitempty"`
	History_red_or_yellow    *int64     `xorm:" null default 0 'history_red_or_yellow'" indexes:" index(Index_6) " comment:"历史节点红黄牌类型" json:"history_red_or_yellow,omitempty" bson:",omitempty" msgpack:"history_red_or_yellow,omitempty"`
	Enable_mark              *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_8) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark              *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_9) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code                *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description              *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date              *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid            *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username          *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date              *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid            *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username          *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version                  *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: dic_affairs_type
// 说明: 字典-事项类型
// 字典-事项类型
// 
type dic_affairs_type struct {
	Affairs_type_id     *int64     `xorm:" pk not null 'affairs_type_id'" indexes:"" comment:"事项类型主键" json:"affairs_type_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"affairs_type_id,omitempty"`
	Affairs_type_name   *string    `xorm:" varchar(255) null 'affairs_type_name'" indexes:"" comment:"事项类型名称" json:"affairs_type_name,omitempty" bson:",omitempty" msgpack:"affairs_type_name,omitempty"`
	Affairs_type_prompt *string    `xorm:" varchar(255) null 'affairs_type_prompt'" indexes:"" comment:"事项类型简" json:"affairs_type_prompt,omitempty" bson:",omitempty" msgpack:"affairs_type_prompt,omitempty"`
	Affairs_type_code   *string    `xorm:" varchar(255) null 'affairs_type_code'" indexes:" unique(Index_1) " comment:"事项类型编码" json:"affairs_type_code,omitempty" bson:",omitempty" msgpack:"affairs_type_code,omitempty"`
	Enable_mark         *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_2) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark         *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_3) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code           *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description         *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date         *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid       *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username     *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date         *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid       *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username     *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version             *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: dic_time_limit_type
// 说明: 字典-时限类型表
// 字典-时限类型表
// 
type dic_time_limit_type struct {
	Time_limit_type_id     *int64     `xorm:" pk not null 'time_limit_type_id'" indexes:"" comment:"时限类型主键" json:"time_limit_type_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"time_limit_type_id,omitempty"`
	Time_limit_type_name   *string    `xorm:" varchar(50) null 'time_limit_type_name'" indexes:"" comment:"时限类型名称" json:"time_limit_type_name,omitempty" bson:",omitempty" msgpack:"time_limit_type_name,omitempty"`
	Time_limit_type_prompt *string    `xorm:" varchar(50) null 'time_limit_type_prompt'" indexes:"" comment:"时限类型简称" json:"time_limit_type_prompt,omitempty" bson:",omitempty" msgpack:"time_limit_type_prompt,omitempty"`
	Value_code             *string    `xorm:" varchar(50) null 'value_code'" indexes:" unique(Index_1) " comment:"编码" json:"value_code,omitempty" bson:",omitempty" msgpack:"value_code,omitempty"`
	Enable_mark            *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_2) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark            *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_3) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code              *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description            *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date            *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid          *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username        *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date            *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid          *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username        *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version                *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: tab_authorization
// 说明: 权限表
// 权限表
// 
type tab_authorization struct {
	Authorization_id *int64     `xorm:" pk not null 'authorization_id'" indexes:"" comment:"权限主键" json:"authorization_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"authorization_id,omitempty"`
	User_id          *int64     `xorm:" null 'user_id'" indexes:"" comment:"用户主键" json:"user_id,omitempty" bson:",omitempty" msgpack:"user_id,omitempty" fk:"tab_user(user_id)"`
	Group_id         *int64     `xorm:" null 'group_id'" indexes:"" comment:"组(角色)主键" json:"group_id,omitempty" bson:",omitempty" msgpack:"group_id,omitempty" fk:"tab_group(group_id)"`
	Node_id          *int64     `xorm:" not null 'node_id'" indexes:"" comment:"节点主键" json:"node_id,omitempty" bson:",omitempty" msgpack:"node_id,omitempty" fk:"tab_flow_node_info(node_id)"`
	Affairs_id       *int64     `xorm:" null 'affairs_id'" indexes:"" comment:"事项主键" json:"affairs_id,omitempty" bson:",omitempty" msgpack:"affairs_id,omitempty" fk:"tab_affairs(affairs_id)"`
	Enable_mark      *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_1) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark      *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_2) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code        *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description      *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date      *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid    *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username  *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date      *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid    *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username  *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version          *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: tab_prepare_agent
// 说明: 外网-代理人表
// 外网-代理人表
// 
type tab_prepare_agent struct {
	Agent_id           *int64     `xorm:" pk not null 'agent_id'" indexes:"" comment:"代理人主键" json:"agent_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"agent_id,omitempty"`
	Prepare_affairs_id *int64     `xorm:" null 'prepare_affairs_id'" indexes:"" comment:"预审申请主键" json:"prepare_affairs_id,omitempty" bson:",omitempty" msgpack:"prepare_affairs_id,omitempty" fk:"tab_prepare_affairs(prepare_affairs_id)"`
	Agent_name         *string    `xorm:" varchar(255) null 'agent_name'" indexes:"" comment:"代理人名" json:"agent_name,omitempty" bson:",omitempty" msgpack:"agent_name,omitempty"`
	Sex                *int64     `xorm:" null 'sex'" indexes:"" comment:"性别" json:"sex,omitempty" bson:",omitempty" msgpack:"sex,omitempty"`
	Id_code            *string    `xorm:" varchar(255) not null 'id_code'" indexes:" index(Index_1) " comment:"身份证号码" json:"id_code,omitempty" bson:",omitempty" msgpack:"id_code,omitempty"`
	Mobile             *string    `xorm:" varchar(255) not null 'mobile'" indexes:" index(Index_2) " comment:"手机" json:"mobile,omitempty" bson:",omitempty" msgpack:"mobile,omitempty"`
	Email              *string    `xorm:" varchar(255) null 'email'" indexes:" index(Index_3) " comment:"邮箱" json:"email,omitempty" bson:",omitempty" msgpack:"email,omitempty"`
	Photo_url          *string    `xorm:" varchar(255) null 'photo_url'" indexes:"" comment:"照片" json:"photo_url,omitempty" bson:",omitempty" msgpack:"photo_url,omitempty"`
	Enable_mark        *Boolean   `xorm:" null default 1 'enable_mark'" indexes:"" comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark        *Boolean   `xorm:" null default 0 'delete_mark'" indexes:"" comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code          *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description        *string    `xorm:" varchar(255) null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date        *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid      *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username    *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date        *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid      *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username    *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version            *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: dic_node_status
// 说明: 字典-节点流转状态
// 字典-节点流转状态
// 
type dic_node_status struct {
	Node_status_id     *int64     `xorm:" pk not null 'node_status_id'" indexes:"" comment:"流转状态主键" json:"node_status_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"node_status_id,omitempty"`
	Node_status_name   *string    `xorm:" varchar(255) null 'node_status_name'" indexes:"" comment:"流转状态名称" json:"node_status_name,omitempty" bson:",omitempty" msgpack:"node_status_name,omitempty"`
	Node_status_prompt *string    `xorm:" varchar(255) null 'node_status_prompt'" indexes:"" comment:"流转状态简称" json:"node_status_prompt,omitempty" bson:",omitempty" msgpack:"node_status_prompt,omitempty"`
	Node_status_code   *string    `xorm:" varchar(255) null 'node_status_code'" indexes:" unique(Index_1) " comment:"流转状态编码" json:"node_status_code,omitempty" bson:",omitempty" msgpack:"node_status_code,omitempty"`
	Enable_mark        *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_2) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark        *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_3) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code          *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description        *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date        *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid      *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username    *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date        *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid      *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username    *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version            *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: tab_affairs_delicate
// 说明: 事项定义-事项专线关系
// 事项定义-事项专线关系
// 
type tab_affairs_delicate struct {
	Affairs_delicate_id *int64     `xorm:" pk not null 'affairs_delicate_id'" indexes:"" comment:"事项专线主键" json:"affairs_delicate_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"affairs_delicate_id,omitempty"`
	Affairs_id          *int64     `xorm:" null 'affairs_id'" indexes:"" comment:"事项主键" json:"affairs_id,omitempty" bson:",omitempty" msgpack:"affairs_id,omitempty" fk:"tab_affairs(affairs_id)"`
	Enable_mark         *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_1) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark         *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_2) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code           *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description         *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date         *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid       *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username     *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date         *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid       *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username     *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version             *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: tab_autoinput
// 说明: 自动填充表
// 自动填充表
// 
type tab_autoinput struct {
	Autoinput_id       *int64     `xorm:" pk not null 'autoinput_id'" indexes:"" comment:"自动填充主键" json:"autoinput_id,omitempty" bson:",omitempty" msgpack:"autoinput_id,omitempty"`
	Affairs_id         *int64     `xorm:" null 'affairs_id'" indexes:"" comment:"事项主键" json:"affairs_id,omitempty" bson:",omitempty" msgpack:"affairs_id,omitempty" fk:"tab_affairs(affairs_id)"`
	Name_id            *string    `xorm:" varchar(255) null 'name_id'" indexes:"" comment:"名称" json:"name_id,omitempty" bson:",omitempty" msgpack:"name_id,omitempty"`
	Gender_id          *int64     `xorm:" null 'gender_id'" indexes:"" comment:"性别" json:"gender_id,omitempty" bson:",omitempty" msgpack:"gender_id,omitempty"`
	Type_id            *int64     `xorm:" null 'type_id'" indexes:"" comment:"类型" json:"type_id,omitempty" bson:",omitempty" msgpack:"type_id,omitempty"`
	Credential_type_id *int64     `xorm:" null 'credential_type_id'" indexes:"" comment:"证件类型" json:"credential_type_id,omitempty" bson:",omitempty" msgpack:"credential_type_id,omitempty"`
	Credential_code_id *string    `xorm:" varchar(50) null 'credential_code_id'" indexes:" index(Index_1) " comment:"证件号码" json:"credential_code_id,omitempty" bson:",omitempty" msgpack:"credential_code_id,omitempty"`
	Phone_id           *string    `xorm:" varchar(50) null 'phone_id'" indexes:"" comment:"联系电话" json:"phone_id,omitempty" bson:",omitempty" msgpack:"phone_id,omitempty"`
	Post_address_id    *string    `xorm:" varchar(255) null 'post_address_id'" indexes:"" comment:"通信地址" json:"post_address_id,omitempty" bson:",omitempty" msgpack:"post_address_id,omitempty"`
	Fixed_phone_id     *string    `xorm:" varchar(50) null 'fixed_phone_id'" indexes:"" comment:"固定电话" json:"fixed_phone_id,omitempty" bson:",omitempty" msgpack:"fixed_phone_id,omitempty"`
	Reg_address_id     *string    `xorm:" varchar(50) null 'reg_address_id'" indexes:"" comment:"户籍地址" json:"reg_address_id,omitempty" bson:",omitempty" msgpack:"reg_address_id,omitempty"`
	Bithday_id         *time.Time `xorm:" null 'bithday_id'" indexes:"" comment:"出生日期" json:"bithday_id,omitempty" bson:",omitempty" msgpack:"bithday_id,omitempty"`
	Age_id             *int64     `xorm:" null 'age_id'" indexes:"" comment:"年龄" json:"age_id,omitempty" bson:",omitempty" msgpack:"age_id,omitempty"`
	Postcode_id        *string    `xorm:" varchar(50) null 'postcode_id'" indexes:"" comment:"邮编" json:"postcode_id,omitempty" bson:",omitempty" msgpack:"postcode_id,omitempty"`
	Enable_mark        *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_2) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark        *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_3) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code          *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description        *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date        *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid      *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username    *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date        *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid      *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username    *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version            *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: tab_document_info
// 说明: 证件信息表
// 证件信息表
// 
type tab_document_info struct {
	Document_info_id   *int64     `xorm:" pk not null 'document_info_id'" indexes:"" comment:"证照标识" json:"document_info_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"document_info_id,omitempty"`
	Document_type      *int64     `xorm:" null 'document_type'" indexes:"" comment:"证照目录编码" json:"document_type,omitempty" bson:",omitempty" msgpack:"document_type,omitempty"`
	Documentl_code     *string    `xorm:" varchar(255) null 'documentl_code'" indexes:"" comment:"证照编号" json:"documentl_code,omitempty" bson:",omitempty" msgpack:"documentl_code,omitempty"`
	Documentl_name     *string    `xorm:" varchar(255) null 'documentl_name'" indexes:"" comment:"证件名" json:"documentl_name,omitempty" bson:",omitempty" msgpack:"documentl_name,omitempty"`
	Confer_date        *time.Time `xorm:" null 'confer_date'" indexes:"" comment:"颁证时间" json:"confer_date,omitempty" bson:",omitempty" msgpack:"confer_date,omitempty"`
	Confer_sector      *string    `xorm:" varchar(255) null 'confer_sector'" indexes:"" comment:"颁证单位" json:"confer_sector,omitempty" bson:",omitempty" msgpack:"confer_sector,omitempty"`
	Valid_period_start *time.Time `xorm:" null 'valid_period_start'" indexes:"" comment:"有效期开始时间" json:"valid_period_start,omitempty" bson:",omitempty" msgpack:"valid_period_start,omitempty"`
	Valid_period_end   *time.Time `xorm:" null 'valid_period_end'" indexes:"" comment:"有效期结束时间" json:"valid_period_end,omitempty" bson:",omitempty" msgpack:"valid_period_end,omitempty"`
	Holder             *string    `xorm:" varchar(255) null 'holder'" indexes:"" comment:"持证者" json:"holder,omitempty" bson:",omitempty" msgpack:"holder,omitempty"`
	Change_record      *string    `xorm:" text null 'change_record'" indexes:"" comment:"证照变更记录" json:"change_record,omitempty" bson:",omitempty" msgpack:"change_record,omitempty"`
	Image_url          *string    `xorm:" text null 'image_url'" indexes:"" comment:"证照图像" json:"image_url,omitempty" bson:",omitempty" msgpack:"image_url,omitempty"`
	Document_url       *string    `xorm:" text null 'document_url'" indexes:"" comment:"证照电子文书" json:"document_url,omitempty" bson:",omitempty" msgpack:"document_url,omitempty"`
	Enable_mark        *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_1) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark        *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_2) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code          *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description        *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date        *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid      *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username    *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date        *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid      *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username    *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version            *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: tab_real_name_materials
// 说明: 外网-实名认证资料
// 外网-实名认证资料
// 
type tab_real_name_materials struct {
	Real_materials_id *int64     `xorm:" pk not null 'real_materials_id'" indexes:"" comment:"附件主键" json:"real_materials_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"real_materials_id,omitempty"`
	User_id           *int64     `xorm:" null 'user_id'" indexes:"" comment:"用户主键" json:"user_id,omitempty" bson:",omitempty" msgpack:"user_id,omitempty"`
	Material_name     *string    `xorm:" varchar(255) null 'material_name'" indexes:"" comment:"材料名称" json:"material_name,omitempty" bson:",omitempty" msgpack:"material_name,omitempty"`
	Materials_url     *string    `xorm:" varchar(255) null 'materials_url'" indexes:"" comment:"文件ID" json:"materials_url,omitempty" bson:",omitempty" msgpack:"materials_url,omitempty"`
	Enable_mark       *Boolean   `xorm:" null default 1 'enable_mark'" indexes:"" comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark       *Boolean   `xorm:" null default 0 'delete_mark'" indexes:"" comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code         *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description       *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date       *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid     *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username   *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date       *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid     *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username   *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version           *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: tab_affair_material
// 说明: 事项受理-材料-包含结果物
// 事项受理-材料-包含结果物
// 
type tab_affair_material struct {
	Affairs_material_id     *int64     `xorm:" pk not null 'affairs_material_id'" indexes:"" comment:"事项受理材料主键" json:"affairs_material_id,omitempty" bson:",omitempty" msgpack:"affairs_material_id,omitempty"`
	Affairs_proposers_id    *int64     `xorm:" null 'affairs_proposers_id'" indexes:"" comment:"事项申请人主键" json:"affairs_proposers_id,omitempty" bson:",omitempty" msgpack:"affairs_proposers_id,omitempty" fk:"tab_affairs_proposers(affairs_proposers_id)"`
	Material_provider_id    *int64     `xorm:" null 'material_provider_id'" indexes:"" comment:"材料提供人主键" json:"material_provider_id,omitempty" bson:",omitempty" msgpack:"material_provider_id,omitempty"`
	Material_info_id        *int64     `xorm:" null 'material_info_id'" indexes:"" comment:"材料主键" json:"material_info_id,omitempty" bson:",omitempty" msgpack:"material_info_id,omitempty"`
	Material_name           *string    `xorm:" varchar(255) null 'material_name'" indexes:"" comment:"材料名称" json:"material_name,omitempty" bson:",omitempty" msgpack:"material_name,omitempty"`
	Material_file_id        *string    `xorm:" varchar(255) null 'material_file_id'" indexes:"" comment:"材料文件ID" json:"material_file_id,omitempty" bson:",omitempty" msgpack:"material_file_id,omitempty"`
	Material_file_url       *string    `xorm:" varchar(255) null 'material_file_url'" indexes:"" comment:"材料文件位置" json:"material_file_url,omitempty" bson:",omitempty" msgpack:"material_file_url,omitempty"`
	Material_upload_type_id *int64     `xorm:" not null 'material_upload_type_id'" indexes:"" comment:"上传操作类型" json:"material_upload_type_id,omitempty" bson:",omitempty" msgpack:"material_upload_type_id,omitempty" fk:"dic_material_upload_type(material_upload_id)"`
	Material_reuse_type_id  *int64     `xorm:" null 'material_reuse_type_id'" indexes:"" comment:"材料复用类型" json:"material_reuse_type_id,omitempty" bson:",omitempty" msgpack:"material_reuse_type_id,omitempty" fk:"dic_reuse_type(reuse_type_id)"`
	Material_type_id        *int64     `xorm:" null 'material_type_id'" indexes:"" comment:"材料类型" json:"material_type_id,omitempty" bson:",omitempty" msgpack:"material_type_id,omitempty" fk:"dic_material_type(material_type_id)"`
	Affairs_info_id         *int64     `xorm:" null 'affairs_info_id'" indexes:"" comment:"受理主键" json:"affairs_info_id,omitempty" bson:",omitempty" msgpack:"affairs_info_id,omitempty" fk:"tab_affairs_info(affairs_info_id)"`
	Material_region_id      *int64     `xorm:" null 'material_region_id'" indexes:"" comment:"材料提交归属地主键" json:"material_region_id,omitempty" bson:",omitempty" msgpack:"material_region_id,omitempty"`
	Material_region_name    *string    `xorm:" varchar(255) null 'material_region_name'" indexes:"" comment:"材料提交归属地名称" json:"material_region_name,omitempty" bson:",omitempty" msgpack:"material_region_name,omitempty"`
	Material_valid_term     *time.Time `xorm:" null 'material_valid_term'" indexes:"" comment:"材料到期时间" json:"material_valid_term,omitempty" bson:",omitempty" msgpack:"material_valid_term,omitempty"`
	Material_relative_field *string    `xorm:" varchar(255) null 'material_relative_field'" indexes:"" comment:"材料相关字段" json:"material_relative_field,omitempty" bson:",omitempty" msgpack:"material_relative_field,omitempty"`
	Gathered                *Boolean   `xorm:" null 'gathered'" indexes:"" comment:"是否已收取" json:"gathered,omitempty" bson:",omitempty" msgpack:"gathered,omitempty"`
	Gathered_mode           *int64     `xorm:" null 'gathered_mode'" indexes:"" comment:"收取方式" json:"gathered_mode,omitempty" bson:",omitempty" msgpack:"gathered_mode,omitempty"`
	Gathered_number         *int64     `xorm:" null 'gathered_number'" indexes:"" comment:"收取数量" json:"gathered_number,omitempty" bson:",omitempty" msgpack:"gathered_number,omitempty"`
	Gathered_time           *time.Time `xorm:" null 'gathered_time'" indexes:"" comment:"收取时间" json:"gathered_time,omitempty" bson:",omitempty" msgpack:"gathered_time,omitempty"`
	Return_success          *Boolean   `xorm:" null default 0 'return_success'" indexes:"" comment:"成功时退回" json:"return_success,omitempty" bson:",omitempty" msgpack:"return_success,omitempty"`
	Return_failed           *Boolean   `xorm:" null default 0 'return_failed'" indexes:"" comment:"不成功退回" json:"return_failed,omitempty" bson:",omitempty" msgpack:"return_failed,omitempty"`
	Outcome_tag             *Boolean   `xorm:" not null default 0 'outcome_tag'" indexes:"" comment:"是否结果物" json:"outcome_tag,omitempty" bson:",omitempty" msgpack:"outcome_tag,omitempty"`
	Signer_date             *time.Time `xorm:" null 'signer_date'" indexes:"" comment:"取回时间" json:"signer_date,omitempty" bson:",omitempty" msgpack:"signer_date,omitempty"`
	Signer_name             *string    `xorm:" varchar(50) null 'signer_name'" indexes:"" comment:"取回人签字" json:"signer_name,omitempty" bson:",omitempty" msgpack:"signer_name,omitempty"`
	Signer_photo            *string    `xorm:" varchar(255) null 'signer_photo'" indexes:"" comment:"签字人照片" json:"signer_photo,omitempty" bson:",omitempty" msgpack:"signer_photo,omitempty"`
	Enable_mark             *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_1) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark             *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_2) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code               *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description             *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date             *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid           *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username         *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date             *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid           *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username         *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version                 *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: tab_affairs_node_status
// 说明: 事项受理-节点流转状态
// 事项受理-节点流转状态
// 
type tab_affairs_node_status struct {
	Status_id       *int64     `xorm:" pk not null 'status_id'" indexes:"" comment:"流转状态主键" json:"status_id,omitempty" bson:",omitempty" msgpack:"status_id,omitempty"`
	Node_id         *int64     `xorm:" null 'node_id'" indexes:"" comment:"节点编号" json:"node_id,omitempty" bson:",omitempty" msgpack:"node_id,omitempty" fk:"tab_flow_node_info(node_id)"`
	Next_node_id    *int64     `xorm:" null 'next_node_id'" indexes:"" comment:"下个节点编号" json:"next_node_id,omitempty" bson:",omitempty" msgpack:"next_node_id,omitempty" fk:"tab_flow_node_info(node_id)"`
	Status          *int64     `xorm:" null 'status'" indexes:"" comment:"流转状态" json:"status,omitempty" bson:",omitempty" msgpack:"status,omitempty" fk:"dic_node_status(node_status_id)"`
	Region_id       *int64     `xorm:" null 'region_id'" indexes:"" comment:"流转归属地" json:"region_id,omitempty" bson:",omitempty" msgpack:"region_id,omitempty" fk:"tab_region(region_id)"`
	Affairs_info_id *int64     `xorm:" null 'affairs_info_id'" indexes:"" comment:"受理主键" json:"affairs_info_id,omitempty" bson:",omitempty" msgpack:"affairs_info_id,omitempty" fk:"tab_affairs_info(affairs_info_id)"`
	Region_name     *string    `xorm:" varchar(255) null 'region_name'" indexes:"" comment:"流转归属地名称" json:"region_name,omitempty" bson:",omitempty" msgpack:"region_name,omitempty"`
	Enable_mark     *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_1) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark     *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_2) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code       *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description     *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date     *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid   *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date     *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid   *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version         *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: dic_district
// 说明: 字典-行政区划-区县
// 字典-行政区划-区县
// 
type dic_district struct {
	District_id     *int64     `xorm:" pk not null 'district_id'" indexes:"" comment:"区县主键" json:"district_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"district_id,omitempty"`
	District_name   *string    `xorm:" varchar(255) null 'district_name'" indexes:"" comment:"区县名称" json:"district_name,omitempty" bson:",omitempty" msgpack:"district_name,omitempty"`
	District_prompt *string    `xorm:" varchar(255) null 'district_prompt'" indexes:"" comment:"区县简称" json:"district_prompt,omitempty" bson:",omitempty" msgpack:"district_prompt,omitempty"`
	District_code   *string    `xorm:" varchar(255) null 'district_code'" indexes:" unique(Index_1) " comment:"区县代码" json:"district_code,omitempty" bson:",omitempty" msgpack:"district_code,omitempty"`
	City_id         *int64     `xorm:" null 'city_id'" indexes:"" comment:"地市主键" json:"city_id,omitempty" bson:",omitempty" msgpack:"city_id,omitempty" fk:"dic_city(city_id)"`
	Enable_mark     *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_2) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark     *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_3) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code       *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description     *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date     *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid   *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date     *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid   *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version         *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: dic_node_standard_code
// 说明: 字典-流程节点标准编码
// 字典-流程节点标准编码
// 
type dic_node_standard_code struct {
	Standard_code_id     *int64     `xorm:" pk not null 'standard_code_id'" indexes:"" comment:"节点标准编码主键" json:"standard_code_id,omitempty" bson:",omitempty" msgpack:"standard_code_id,omitempty"`
	Standard_code_name   *string    `xorm:" varchar(255) null 'standard_code_name'" indexes:" unique(Index_1) " comment:"节点标准编码名称" json:"standard_code_name,omitempty" bson:",omitempty" msgpack:"standard_code_name,omitempty"`
	Standard_code_prompt *string    `xorm:" varchar(255) null 'standard_code_prompt'" indexes:"" comment:"节点标准编码简称" json:"standard_code_prompt,omitempty" bson:",omitempty" msgpack:"standard_code_prompt,omitempty"`
	Standard_code        *string    `xorm:" varchar(255) null 'standard_code'" indexes:"" comment:"节点标准编码" json:"standard_code,omitempty" bson:",omitempty" msgpack:"standard_code,omitempty"`
	Enable_mark          *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_2) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark          *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_3) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code            *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description          *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date          *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid        *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username      *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date          *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid        *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username      *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version              *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: dic_subject_nature
// 说明: 字典-实施主体性质
// 字典-实施主体性质
// 
type dic_subject_nature struct {
	Subject_nature_id     *int64     `xorm:" pk not null 'subject_nature_id'" indexes:"" comment:"实施主体性质主键" json:"subject_nature_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"subject_nature_id,omitempty"`
	Subject_nature_name   *string    `xorm:" varchar(255) null 'subject_nature_name'" indexes:"" comment:"实施主体性质名称" json:"subject_nature_name,omitempty" bson:",omitempty" msgpack:"subject_nature_name,omitempty"`
	Subject_nature_code   *string    `xorm:" varchar(50) null 'subject_nature_code'" indexes:"" comment:"实施主体性质编码" json:"subject_nature_code,omitempty" bson:",omitempty" msgpack:"subject_nature_code,omitempty"`
	Subject_nature_prompt *string    `xorm:" varchar(50) null 'subject_nature_prompt'" indexes:"" comment:"实施主体性质简称" json:"subject_nature_prompt,omitempty" bson:",omitempty" msgpack:"subject_nature_prompt,omitempty"`
	Enable_mark           *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_1) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark           *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_2) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code             *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description           *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date           *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid         *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username       *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date           *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid         *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username       *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version               *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: tab_affairs_attachment
// 说明: 事项受理-附件表
// 事项受理-附件表
// 
type tab_affairs_attachment struct {
	Attachment_id   *int64     `xorm:" pk not null 'attachment_id'" indexes:"" comment:"附件主键" json:"attachment_id,omitempty" bson:",omitempty" msgpack:"attachment_id,omitempty"`
	Affairs_info_id *int64     `xorm:" null 'affairs_info_id'" indexes:"" comment:"受理主键" json:"affairs_info_id,omitempty" bson:",omitempty" msgpack:"affairs_info_id,omitempty" fk:"tab_affairs_info(affairs_info_id)"`
	Attachment_url  *string    `xorm:" varchar(255) null 'attachment_url'" indexes:"" comment:"附件位置" json:"attachment_url,omitempty" bson:",omitempty" msgpack:"attachment_url,omitempty"`
	Attachment_name *string    `xorm:" varchar(255) null 'attachment_name'" indexes:"" comment:"附件名称" json:"attachment_name,omitempty" bson:",omitempty" msgpack:"attachment_name,omitempty"`
	Enable_mark     *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_1) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark     *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_2) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code       *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description     *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date     *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid   *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date     *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid   *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version         *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}

//
// 表名: tab_region
// 说明: 归属地
// 归属地
// 
type tab_region struct {
	Region_id            *int64     `xorm:" pk not null 'region_id'" indexes:"" comment:"归属地主键" json:"region_id,omitempty" startwith:"0" bson:",omitempty" msgpack:"region_id,omitempty"`
	Region_name          *string    `xorm:" varchar(255) null 'region_name'" indexes:"" comment:"归属地名称" json:"region_name,omitempty" bson:",omitempty" msgpack:"region_name,omitempty"`
	Administrative_level *int64     `xorm:" null 'administrative_level'" indexes:"" comment:"管理层级" json:"administrative_level,omitempty" bson:",omitempty" msgpack:"administrative_level,omitempty"`
	Country_id           *int64     `xorm:" null 'country_id'" indexes:"" comment:"国家主键" json:"country_id,omitempty" bson:",omitempty" msgpack:"country_id,omitempty"`
	City_id              *int64     `xorm:" null 'city_id'" indexes:"" comment:"地市主键" json:"city_id,omitempty" bson:",omitempty" msgpack:"city_id,omitempty" fk:"dic_city(city_id)"`
	District_id          *int64     `xorm:" null 'district_id'" indexes:"" comment:"区县主键" json:"district_id,omitempty" bson:",omitempty" msgpack:"district_id,omitempty" fk:"dic_district(district_id)"`
	Street_id            *int64     `xorm:" null 'street_id'" indexes:"" comment:"街道主键" json:"street_id,omitempty" bson:",omitempty" msgpack:"street_id,omitempty" fk:"dic_street(street_id)"`
	Community_id         *int64     `xorm:" null 'community_id'" indexes:"" comment:"社区主键" json:"community_id,omitempty" bson:",omitempty" msgpack:"community_id,omitempty" fk:"dic_community(community_id)"`
	Is_department        *Boolean   `xorm:" null default 0 'is_department'" indexes:"" comment:"是否部门" json:"is_department,omitempty" bson:",omitempty" msgpack:"is_department,omitempty"`
	Region_brief         *string    `xorm:" text null 'region_brief'" indexes:"" comment:"机构简介" json:"region_brief,omitempty" bson:",omitempty" msgpack:"region_brief,omitempty"`
	Address              *string    `xorm:" text null 'address'" indexes:"" comment:"详细地址" json:"address,omitempty" bson:",omitempty" msgpack:"address,omitempty"`
	Region_history       *string    `xorm:" text null 'region_history'" indexes:"" comment:"历史沿革" json:"region_history,omitempty" bson:",omitempty" msgpack:"region_history,omitempty"`
	Supervise_phone      *string    `xorm:" varchar(50) null 'supervise_phone'" indexes:"" comment:"监督电话" json:"supervise_phone,omitempty" bson:",omitempty" msgpack:"supervise_phone,omitempty"`
	Consult_phone        *string    `xorm:" varchar(50) null 'consult_phone'" indexes:"" comment:"咨询电话" json:"consult_phone,omitempty" bson:",omitempty" msgpack:"consult_phone,omitempty"`
	Booking_perday       *int64     `xorm:" null default 0 'booking_perday'" indexes:"" comment:"每日最大预约数" json:"booking_perday,omitempty" bson:",omitempty" msgpack:"booking_perday,omitempty"`
	Prepare_perday       *int64     `xorm:" null default 0 'prepare_perday'" indexes:"" comment:"每日最大预审数" json:"prepare_perday,omitempty" bson:",omitempty" msgpack:"prepare_perday,omitempty"`
	Service_time         *string    `xorm:" text null 'service_time'" indexes:"" comment:"服务时间" json:"service_time,omitempty" bson:",omitempty" msgpack:"service_time,omitempty"`
	Longitude            *string    `xorm:" varchar(50) null 'longitude'" indexes:"" comment:"纬度" json:"longitude,omitempty" bson:",omitempty" msgpack:"longitude,omitempty"`
	Latitude             *string    `xorm:" varchar(50) null 'latitude'" indexes:"" comment:"经度" json:"latitude,omitempty" bson:",omitempty" msgpack:"latitude,omitempty"`
	Enable_mark          *Boolean   `xorm:" null default 1 'enable_mark'" indexes:" index(Index_1) " comment:"有效标记" json:"enable_mark,omitempty" bson:",omitempty" msgpack:"enable_mark,omitempty"`
	Delete_mark          *Boolean   `xorm:" null default 0 'delete_mark'" indexes:" index(Index_2) " comment:"删除标记" json:"delete_mark,omitempty" bson:",omitempty" msgpack:"delete_mark,omitempty"`
	Sort_code            *int64     `xorm:" null 'sort_code'" indexes:"" comment:"排序码" json:"sort_code,omitempty" bson:",omitempty" msgpack:"sort_code,omitempty"`
	Description          *string    `xorm:" text null 'description'" indexes:"" comment:"备注" json:"description,omitempty" bson:",omitempty" msgpack:"description,omitempty"`
	Create_date          *time.Time `xorm:" null created 'create_date'" indexes:"" comment:"创建时间" json:"create_date,omitempty" bson:",omitempty" msgpack:"create_date,omitempty"`
	Create_userid        *int64     `xorm:" null 'create_userid'" indexes:"" comment:"创建用户Id" json:"create_userid,omitempty" bson:",omitempty" msgpack:"create_userid,omitempty"`
	Create_username      *string    `xorm:" varchar(50) null 'create_username'" indexes:"" comment:"创建人姓名" json:"create_username,omitempty" bson:",omitempty" msgpack:"create_username,omitempty"`
	Modify_date          *time.Time `xorm:" null updated 'modify_date'" indexes:"" comment:"修改时间" json:"modify_date,omitempty" bson:",omitempty" msgpack:"modify_date,omitempty"`
	Modify_userid        *int64     `xorm:" null 'modify_userid'" indexes:"" comment:"修改用户Id" json:"modify_userid,omitempty" bson:",omitempty" msgpack:"modify_userid,omitempty"`
	Modify_username      *string    `xorm:" varchar(50) null 'modify_username'" indexes:"" comment:"修改用户姓名" json:"modify_username,omitempty" bson:",omitempty" msgpack:"modify_username,omitempty"`
	Version              *int64     `xorm:" null default 1 'version'" indexes:"" comment:"版本" json:"version,omitempty" bson:",omitempty" msgpack:"version,omitempty"`
}
