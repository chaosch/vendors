package postlog

import (
	"time"
	"reflect"
	"github.com/jacoblai/yiyidb"
	"log"
	"fmt"
	"os"
	"encoding/json"
	"path/filepath"
	. "egov/common"
)
var p=NewPool(128,1,1)
type LogsEngine struct {
	Connurl 	string
	Ticker	  *time.Ticker
	StopQueue chan bool
	LogQueue  *yiyidb.Queue //日志队列
}
func NewlogsEngine() *LogsEngine{
	le:=&LogsEngine{
		Ticker:time.NewTicker(10000 * time.Millisecond),
		StopQueue: make(chan bool, 1),
	}
	return le
}
func(l *LogsEngine)Init(log_server string){
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	dir=dir+"/data"
	if err != nil {
		log.Fatal(err)
	}
	LogQueue,err2:=yiyidb.OpenQueue(dir+"/LogQueue")
	if err!=nil{
		panic(err2)
	}
	l.LogQueue=LogQueue
	l.Connurl=log_server
}
func (l *LogsEngine)InQueue(val []byte)*ResultTemplate{
	_,err:=l.LogQueue.Enqueue(val)
	if err!=nil{
		log.Println("add logdata queue err", err)
		return RetErr(MixError(err))
	}
	l.ResumRunQueue()
	return RetChanges(1)
}

func (l *LogsEngine)RunLogQueue(){
	go func(){
		for {
			select{
			case <-l.Ticker.C:
				l.DealLogQueue()
			case <-l.StopQueue:
				fmt.Println("logqueue stop")
			}
		}
	}()
}

func (l *LogsEngine)DealLogQueue() uint64{
	count:=l.LogQueue.Length()
	if count>0{
		peekItem,err:=l.LogQueue.Peek()
		if err!=nil{
			log.Println("不能从队列中取出任务,剩余%d个操作",count)
			l.ResumRunQueue()
		}
		defer l.LogQueue.Dequeue()

		err1:=p.HttpPostLog2(peekItem.Value,l.Connurl)
		if err1!=nil{
			log.Println("处理队列数据出错,剩余%d个操作",count)
			l.ResumRunQueue()
			return count
		}
	}
	l.PauseRunQueue()
	return 0
}

func (l *LogsEngine) ResumRunQueue() {
	//fmt.Println("resume")
	l.Ticker = time.NewTicker(100 * time.Millisecond)
}

func (l *LogsEngine) PauseRunQueue() {
	//fmt.Println("pause")
	l.Ticker = time.NewTicker(10000 * time.Millisecond)
}

func (l *LogsEngine) StopRunQueue() {
	l.StopQueue <- true
}

func(l*LogsEngine)Sendlog(s string,k string,c interface{}){
	var cBytes []byte
	if reflect.TypeOf(reflect.ValueOf(c)).Kind() != reflect.Map && reflect.TypeOf(reflect.ValueOf(c)).Kind() != reflect.Struct {
		if reflect.TypeOf(reflect.ValueOf(c)).Kind() == reflect.String {
			sc := simpleContent{Message: c}
			cBytes,_=json.Marshal(sc)
		} else {
			return
		}
	}else{
		cBytes,_=json.Marshal(c)
	}
	cStr:=string(cBytes)

	msg := map[string]string{"system": s, "kind": k, "content": cStr}
	b, _ := json.Marshal(msg)
	l.InQueue(b)
}


