package redis
import (
	"errors"
	"github.com/go-redis/redis"
)
func NewRedisClusterClient(addr []string,password string)(redis.Cmdable,error){
	client:=redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:              addr,
		MaxRedirects:       0,
		ReadOnly:           false,
		RouteByLatency:     false,
		RouteRandomly:      false,
		ClusterSlots:       nil,
		OnNewNode:          nil,
		Dialer:             nil,
		OnConnect:          nil,
		Password:           password,
		MaxRetries:         0,
		MinRetryBackoff:    0,
		MaxRetryBackoff:    0,
		DialTimeout:        0,
		ReadTimeout:        0,
		WriteTimeout:       0,
		PoolSize:           0,
		MinIdleConns:       0,
		MaxConnAge:         0,
		PoolTimeout:        0,
		IdleTimeout:        0,
		IdleCheckFrequency: 0,
		TLSConfig:          nil,
	})
	return client ,nil
}
func NewRedisSentinelClient(addr []string,masterName,password string,db int)(redis.Cmdable,error){
	client:=redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:         masterName,
		SentinelAddrs:      addr,
		SentinelPassword:   "",
		Dialer:             nil,
		OnConnect:          nil,
		Password:           password,
		DB:                 db,
		MaxRetries:         0,
		MinRetryBackoff:    0,
		MaxRetryBackoff:    0,
		DialTimeout:        0,
		ReadTimeout:        0,
		WriteTimeout:       0,
		PoolSize:           0,
		MinIdleConns:       0,
		MaxConnAge:         0,
		PoolTimeout:        0,
		IdleTimeout:        0,
		IdleCheckFrequency: 0,
		TLSConfig:          nil,
	})
	return client ,nil
}
func NewSingleRedisClient(addr, password string,db int) (redis.Cmdable, error) {
	client := redis.NewClient(&redis.Options{
		Network:            "",
		Addr:               addr,
		Dialer:             nil,
		OnConnect:          nil,
		Password:           password,
		DB:                 db,
		MaxRetries:         0,
		MinRetryBackoff:    0,
		MaxRetryBackoff:    0,
		DialTimeout:        0,
		ReadTimeout:        0,
		WriteTimeout:       0,
		PoolSize:           0,
		MinIdleConns:       0,
		MaxConnAge:         0,
		PoolTimeout:        0,
		IdleTimeout:        0,
		IdleCheckFrequency: 0,
		TLSConfig:          nil,
	},
	)
	_, err := client.Ping().Result()
	//fmt.Println(pong, err)
	return client, err
}
func NewRedisClient(addr []string, masterName,password string,db int) (redis.Cmdable, error) {
	if len(addr)==0{
		return nil, errors.New("redis addr is null")
	}
	if masterName!=""{
		client,err := NewRedisSentinelClient(addr,masterName,password,db)
		return client,err
	}
	if len(addr)==1{
		client,err := NewSingleRedisClient(addr[0],password,db)
		return client,err
	}else{
		client,err :=NewRedisClusterClient(addr,password)
		return client, err
	}
}