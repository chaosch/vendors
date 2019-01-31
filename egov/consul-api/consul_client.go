package consul_api

import (
	"egov/common"
	"egov/object-id"
	"errors"
	"fmt"
	. "github.com/hashicorp/consul/api"
	"github.com/pquerna/ffjson/ffjson"
	"reflect"
	"strings"
	"time"
)

var (
	CClient *ConsulClient
)

type ConsulClient struct {
	AgentClient   *EngineAgent
	KVClient      *KV
	Name          string
	CheckFunction func() *common.ResultTemplate
	Key           string
	DefaultValue  []byte
	CheckRunning  bool
	Ticker        *time.Ticker
	TickerStop    chan bool
}

type EngineAgent struct {
	Agent   *Agent
	AgentId string
}

func NewConsulClient(consulUrl *string, checkfx func() *common.ResultTemplate, key string, value interface{}, checkDuration time.Duration, serviceName string) error {
	typ := reflect.ValueOf(value).Kind().String()
	var buffer []byte
	if typ == "slice" {
		buffer = value.([]byte)
	} else {
		buffer, _ = ffjson.Marshal(value)
	}
	CClient = &ConsulClient{
		nil, nil, serviceName, checkfx, key, buffer, false, time.NewTicker(checkDuration * time.Second), make(chan bool, 0),
	}
	cfg := &Config{}
	cfg.Address = strings.Split(*consulUrl, "://")[1]
	cfg.Scheme = strings.Split(*consulUrl, "://")[0]
	c, err := NewClient(cfg)
	if err != nil {
		return err
	}
	CClient.AgentClient = &EngineAgent{
		c.Agent(), object_id.NewObjectId().Hex(),
	}
	CClient.KVClient = c.KV()

	if checkfx != nil {
		err = CClient.ServiceRegister()
	}
	return err
}

func (CC *ConsulClient) ServiceRegister() error {
	reg := &AgentServiceRegistration{
		Name: CC.Name,
		ID:   CC.AgentClient.AgentId,
		Tags: []string{CC.Name},
		Port: 9007,
		Check: &AgentServiceCheck{
			TTL:     "15s",
			CheckID: CC.AgentClient.AgentId,
			Status:  HealthPassing,
			Notes:   "启动成功",
		},
	}
	err := CC.AgentClient.Agent.ServiceRegister(reg)
	if err != nil {
		return err
	}
	if CC.CheckFunction != nil {
		go func() {
			defer CC.Ticker.Stop()
			if CC.CheckRunning {
				return
			}
			CC.CheckRunning = true
			for {
				select {
				case <-CC.Ticker.C:
					res := &common.ResultTemplate{}
					if CC.CheckFunction != nil {
						res = CC.CheckFunction()
					} else {
						res = CC.DefaultCheckFun()
					}
					if res.Ok {
						CC.AgentClient.Agent.UpdateTTL(CC.AgentClient.AgentId, common.RetOkStr(res.Data), HealthPassing)
					} else {
						if res.Err.Err().ErrCode == 0 { //错误代码为0时，返回警告
							CC.AgentClient.Agent.UpdateTTL(CC.AgentClient.AgentId, common.RetErrStr(res.Err), HealthWarning)
						} else {
							CC.AgentClient.Agent.UpdateTTL(CC.AgentClient.AgentId, common.RetErrStr(res.Err), HealthCritical)
						}
					}
				case <-CC.TickerStop:
					fmt.Println("健康检查已停止！")
					CC.CheckRunning = false
					return
				}
			}
		}()
	}
	return err
}

func (CC *ConsulClient) DefaultCheckFun() *common.ResultTemplate {
	return common.RetOk("默认健康检查通过!")
}

func (CC *ConsulClient) SetCheckFun(checkfx func() *common.ResultTemplate) {
	if checkfx == nil {
		return
	}
	CC.CheckFunction = checkfx
	CC.ServiceRegister()

}

func (CC *ConsulClient) ServiceDeRegister() {
	fmt.Println("服务已注销！")
	CC.AgentClient.Agent.ServiceDeregister(CC.AgentClient.AgentId)
	if CC.CheckRunning {
		CC.TickerStop <- true
		CC.CheckRunning = false
	}
}

//key
func (CC *ConsulClient) GetKey() ([]byte, error) {
	pair, _, err := CC.KVClient.Get("backend/"+CC.Name+"/"+CC.Key, nil)
	if err != nil {
		return nil, CC.PutKey(CC.DefaultValue)
	}
	if pair == nil {
		return nil, CC.PutKey(CC.DefaultValue)
	}
	if pair.Value == nil || len(pair.Value) == 0 {
		return nil, CC.PutKey(CC.DefaultValue)
	}
	//	fmt.Printf("KV: %v %s\n", pair.Key, pair.Value)
	buf := pair.Value
	return buf, nil
}

func (CC *ConsulClient) PutKey(value []byte) error {
	key := "backend/" + CC.Name + "/" + CC.Key
	kvPair := &KVPair{}

	kvPair.Value = value
	kvPair.Key = key

	_, err := CC.KVClient.Put(kvPair, nil)
	//	fmt.Printf("KV: %v %s\n", pair.Key, pair.Value)
	if err != nil {
		return errors.New("写入 " + kvPair.Key + "出错:" + err.Error())
	}
	return nil
}
