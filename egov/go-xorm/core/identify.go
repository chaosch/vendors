package core

import (
	"sync"
	"time"
)

var index int64 = 1
var mutex sync.Mutex
var (
	Idt = &Identify{}
)

type Identify struct {
	Count       uint64
	Idprefix    int64
	Provice_id  uint64
	City_id     uint64
	District_id uint64
}

func (Id *Identify) Construct(pid, cid, did uint64) {
	Id.Provice_id = pid  //
	Id.City_id = cid     //
	Id.District_id = did //
	Id.Idprefix = 0

	if !(pid<127&&pid>0){
		panic("pid不合法")
	}

	if cid==0&&did!=0{
		panic("did非零则cid必须非零")
	}
	
	Id.Provice_id = Id.Provice_id << 14
	Id.City_id = Id.City_id << 7
	Id.Idprefix = int64(Id.Provice_id+Id.City_id+Id.District_id) << 32
	//Id.Idprefix=int64(Id.Provice_id*(uint64(math.Pow10(14)))+Id.City_id*uint64(math.Pow10(12))+(Id.District_id*uint64(math.Pow10(10))))
}

func (ID *Identify) GetId() int64 {
	var indexTemp int64 = 1
	var ticksTemp int64 = 1
	mutex.Lock()
	index++
	indexTemp = index
	ticksTemp = time.Now().Unix()
	mutex.Unlock()
	var id = (ticksTemp & 0x7FFFFFFFFFF) << 20
	id += indexTemp
	return id
}
