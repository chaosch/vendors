package postlog

import (
	"fmt"
	"encoding/json"
	"bytes"
	"time"
	"net/http"
	"reflect"
	"errors"
	"github.com/pquerna/ffjson/ffjson"
)

const (
	E string = "错误"
	L string = "日志"
	W string = "警告"
)

const (
	D     string = "dal"
	Sync  string = "sync"
	Msg   string = "msg"
	File  string = "file"
	Maint string = "maint"
	Cron  string = "corn"
)

type Dlog struct {
	Pl *Pool
}

var Log *Pool

func init() {
	Log = NewPl(128, 1, 1)
}

func NewPl(s, q, sp int) *Pool {
	return NewPool(s, q, sp)
}

var ErrScheduleTimeout = fmt.Errorf("schedule error: timed out")

// Pool contains logic of goroutine reuse.
type Pool struct {
	sig   chan struct{}
	workf chan func()
}

// NewPool creates new goroutine pool with given size. It also creates a work
// queue of given size. Finally, it spawns given amount of goroutines
// immediately.
func NewPool(size, queue, spawn int) *Pool {
	if spawn <= 0 && queue > 0 {
		panic("dead queue configuration detected")
	}
	if spawn > size {
		panic("spawn > workers")
	}
	p := &Pool{
		sig:   make(chan struct{}, size),
		workf: make(chan func(), queue),
	}
	for i := 0; i < spawn; i++ {
		p.sig <- struct{}{}
		go p.worker(func() {})
	}

	return p
}

// Schedule schedules task to be executed over pool's workers.
func (p *Pool) Schedule(task func()) {
	p.schedule(task, nil)
}

// ScheduleTimeout schedules task to be executed over pool's workers.
// It returns ErrScheduleTimeout when no free workers met during given timeout.
func (p *Pool) ScheduleTimeout(timeout time.Duration, task func()) error {
	return p.schedule(task, time.After(timeout))
}

func (p *Pool) schedule(task func(), timeout <-chan time.Time) error {
	select {
	case <-timeout:
		return ErrScheduleTimeout
	case p.workf <- task:
		return nil
	case p.sig <- struct{}{}:
		go p.worker(task)
		return nil
	}
}

func (p *Pool) worker(task func()) {
	defer func() { <-p.sig }()

	task()

	for task2 := range p.workf {
		task2()
	}
}
func (p *Pool) HttpPostLog(msg map[string]interface{}, logsUrl string) error {
	err := p.ScheduleTimeout(5*time.Second, func() {
		b, _ := json.Marshal(msg)
		req, err := http.NewRequest("POST", "http://"+logsUrl+"/api/logs", bytes.NewBuffer(b))
		if err != nil {
			return
		}
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		client.Do(req)
	})
	if err != nil {
		return err
	}
	return nil
}

func (p *Pool) HttpPostLogSlice(msg LogStructSlice, logsUrl string) error {
	err := p.ScheduleTimeout(5*time.Second, func() {
		b, err := ffjson.Marshal(msg)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(fmt.Sprintf("%+v", string(b)))
		req, err := http.NewRequest("POST", "http://"+logsUrl+"/api/logSlice", bytes.NewBuffer(b))
		if err != nil {
			return
		}
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		//x := string(b)
		//fmt.Println(x)
		client.Do(req)
	})
	if err != nil {
		return err
	}
	return nil
}

func (p *Pool) HttpPostLog2(b []byte, logsUrl string) error {
	err := p.ScheduleTimeout(5*time.Second, func() {
		req, err := http.NewRequest("POST", "http://"+logsUrl+"/api/logs", bytes.NewBuffer(b))
		if err != nil {
			fmt.Println(err)
			return
		}
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		client.Do(req)
	})
	if err != nil {
		return err
	}
	return nil
}

type simpleContent struct {
	Message interface{} `json:"message"`
}

func (p *Pool) SendLog(s string, k string, c interface{}, logsUrl string) {
	//var cBytes []byte
	//if reflect.TypeOf(reflect.ValueOf(c)).Kind() != reflect.Map && reflect.TypeOf(reflect.ValueOf(c)).Kind() != reflect.Struct {
	//	if reflect.TypeOf(reflect.ValueOf(c)).Kind() == reflect.String {
	//		sc := simpleContent{message: c}
	//		cBytes, _ = json.Marshal(sc)
	//	} else {
	//		return
	//	}
	//} else {
	//	cBytes, _ = json.Marshal(c)
	//}
	//cStr := string(cBytes)
	if reflect.TypeOf(reflect.ValueOf(c)).Kind() == reflect.String {
		c = simpleContent{Message: c}
	}
	msg := map[string]interface{}{"system": s, "kind": k, "content": c}
	err := p.HttpPostLog(msg, logsUrl)
	if err != nil {
		fmt.Println("log send fail", err.Error())
	}
}

func (p *Pool) SendLogSlice(s LogStructSlice, logsUrl string) error {
	//var cBytes []byte
	//if reflect.TypeOf(reflect.ValueOf(c)).Kind() != reflect.Map && reflect.TypeOf(reflect.ValueOf(c)).Kind() != reflect.Struct {
	//	if reflect.TypeOf(reflect.ValueOf(c)).Kind() == reflect.String {
	//		sc := simpleContent{message: c}
	//		cBytes, _ = json.Marshal(sc)
	//	} else {
	//		return
	//	}
	//} else {
	//	cBytes, _ = json.Marshal(c)
	//}
	//cStr := string(cBytes)
	if reflect.TypeOf(s.LogStructs).Kind() != reflect.Slice {
		return errors.New("inputed log is not slice")
	}
	msg := s
	err := p.HttpPostLogSlice(msg, logsUrl)
	if err != nil {
		fmt.Println("log send fail", err.Error())
	}
	return nil
}

func (p *Pool) HttpPostLogWithSignal(msg map[string]string, logsUrl string, signal chan bool) error {
	err := p.ScheduleTimeout(5*time.Second, func() {
		b, _ := json.Marshal(msg)
		req, err := http.NewRequest("POST", "http://"+logsUrl+"/api/logs", bytes.NewBuffer(b))
		if err != nil {
			return
		}
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		client.Do(req)
	})
	if err != nil {
		return err
	}
	signal <- true
	return nil
}

func (p *Pool) SendLogWithSignal(s string, k string, c interface{}, logsUrl string, signal chan bool) {
	var cBytes []byte
	if reflect.TypeOf(reflect.ValueOf(c)).Kind() != reflect.Map && reflect.TypeOf(reflect.ValueOf(c)).Kind() != reflect.Struct {
		if reflect.TypeOf(reflect.ValueOf(c)).Kind() == reflect.String {
			sc := simpleContent{Message: c}
			cBytes, _ = json.Marshal(sc)
		} else {
			return
		}
	} else {
		cBytes, _ = json.Marshal(c)
	}
	cStr := string(cBytes)
	msg := map[string]string{"system": s, "kind": k, "content": cStr}
	err := p.HttpPostLogWithSignal(msg, logsUrl, signal)
	if err != nil {
		fmt.Println("log send fail", err.Error())
	}
}

type LogStruct struct {
	Time    time.Time     `json:"time"`
	Content ProcessStatus `json:"content"`
	Kind    string        `json:"kind"`
	System  string        `json:"system"`
}

type LogStructSlice struct {
	LogStructs []LogStruct `json:"log_structs"`
}

type ProcessStatus struct {
	InBuf string `json:"inbuf"`
	InBufObject   map[string]interface{} `json:"inBufObject" `
	Starttime     time.Time   `json:"starttime"`
	EndTime       time.Time   `json:"endtime"`
	CreateTime    time.Time   `json:"create_time"` //必须
	Duration      int64       `json:"duration"`
	OK            bool        `json:"ok"`
	Err           interface{} `json:"err"`
	SqlDuration   int64       `json:"sqlduration"`
	Changes       int64       `json:"changes"`
	ChipId        int64       `json:"chipid"`        //唯一标识
	Version       string      `json:"version"`       //版本
	IpAddress     string      `json:"ipaddress"`     //ip地址
	InTransaction bool        `json:"intransaction"` //必须
	Prompt        string      `json:"prompt"`
}
