package consul_api

import (
	"egov/object-id"
	"errors"
	. "github.com/hashicorp/consul/api"
	"github.com/pquerna/ffjson/ffjson"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"time"
)

var (
	CClient = &ConsulClient{}
)

type ConsulClient struct {
	AgentClient   *EngineAgent
	KVClient      *KV
	Name          string
	CheckFunction func() string
	Key           string
	DefaultValue  []byte
	CheckRunning  bool
}

type EngineAgent struct {
	Agent      *Agent
	AgentId    string
	TickerStop chan bool
}

func NewConsulClient(consul_url *string, checkfx func() string, key string, value interface{}) error {
	typ := reflect.ValueOf(value).Kind().String()
	var buffer []byte
	if typ == "slice" {
		buffer = value.([]byte)
	} else {
		buffer, _ = ffjson.Marshal(value)
	}
	CClient = &ConsulClient{
		nil, nil, filepath.Base(os.Args[0]), checkfx, key, buffer, false,
	}
	cfg := &Config{}
	cfg.Address = strings.Split(*consul_url, "://")[1]
	cfg.Scheme = strings.Split(*consul_url, "://")[0]
	c, err := NewClient(cfg)
	if err != nil {
		return err
	}
	CClient.AgentClient = &EngineAgent{
		c.Agent(), object_id.NewObjectId().Hex(), make(chan bool, 0),
	}
	CClient.KVClient = c.KV()
	err = CClient.ServiceRegister()
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
		},
	}
	err := CC.AgentClient.Agent.ServiceRegister(reg)
	if err != nil {
		return err
	}
	go func() {
		CC.CheckRunning = true
		ticker := time.NewTicker(15 * time.Second)
		for {
			select {
			case <-ticker.C:
				CC.AgentClient.Agent.UpdateTTL(CC.AgentClient.AgentId, CC.CheckFunction(), HealthPassing)
			case <-CC.AgentClient.TickerStop:
				return
			}
		}
	}()
	return err
}

func (CC *ConsulClient) ServiceDeRegister() {
	if CC.CheckRunning {
		CC.AgentClient.TickerStop <- true
	}
	CC.AgentClient.Agent.ServiceDeregister(CC.AgentClient.AgentId)
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
