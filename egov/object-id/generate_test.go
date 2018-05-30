package object_id

import (
	"testing"
	//"fmt"
	"sync"
)

var lock sync.Mutex
var mapString map[string]string

func Benchmark_Get_Id(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		mapString=make(map[string]string,0)
		for pb.Next() {
			v:=NewObjectId().Hex()
			lock.Lock()
			if _,ok:=mapString[v];ok{
				panic("duplication id:"+v)
			}else {
				mapString[v] = v
			}
			lock.Unlock()
			//fmt.Println(tempId.Hex())
		}
	})
}