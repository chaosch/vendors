package consul_api

import (
"errors"
"github.com/hashicorp/consul/api"
"strings"
)

func GetKey(consulUrl *string, key string) ([]byte, error) {
	if *consulUrl == "" {
		return nil, errors.New("连接consul服务失败，没有指定consul Url")
	}
	apiConfig := api.Config{}
	//形如http://192.168.4.116:8500的consulUrl
	apiConfig.Address = strings.Split(*consulUrl, "://")[1]
	apiConfig.Scheme = strings.Split(*consulUrl, "://")[0]
	client, err := api.NewClient(&apiConfig)
	if err != nil {
		return nil, errors.New("consul 客户端建立失败" + ":" + err.Error())
	}
	kv := client.KV()
	pair, _, err := kv.Get(key, nil)
	if err != nil {
		return nil, errors.New("consul 获取key失败" + ":" + err.Error())
	}

	//	fmt.Printf("KV: %v %s\n", pair.Key, pair.Value)
	buf := pair.Value

	if len(buf) == 0 {
		return nil, errors.New("consul 获取" + key + "内容为空")
	}
	return buf, nil
}
