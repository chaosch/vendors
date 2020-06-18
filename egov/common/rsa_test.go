package common

import (
	"fmt"
	"testing"
)

func TestCoverPassword(t *testing.T) {
	password := "mongodb://root:123456@192.168.4.124:9009,192.168.4.125:9009/?replicaSet=mongo-cluster&authSource=yizheng"
	p, err := GetRsaPassword([]byte(password))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(p))
	//pk:="UfvyS1oPj9eN/aXbfIFGVCQ1rzROaiATwQQoYhjyfYkpJk9TuDG8bs+v4nPthEB2JRtAG/u1qKfLX2zXj1m9U3PBnemxrT4zyHjPD0y5LuOGDfsp52Rd18dZwa8/tADRs9OXjBu2pLzFt6CIUbIno+F01AkOXBKsRJ/JK2Ixo8s="
	password, err = RsaPassword(string(p))
	fmt.Println(password, err)
	if password != "mongodb://root:123456@192.168.4.124:9009,192.168.4.125:9009/?replicaSet=mongo-cluster&authSource=yizheng" {
		t.Error("解密错误")
	}
}
