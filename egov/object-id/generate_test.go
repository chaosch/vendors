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


func BenchmarkDb_engine_Insert2(b *testing.B) {
	m:=make(map[string]bool)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next(){
			v:=NewObjectId().Hex()
			lock.Lock()
			if _,ok:=m[v];ok{
				panic("duplication id:"+v)
			}else {
				m[v] = true
			}
			lock.Unlock()
		}
	})
}


