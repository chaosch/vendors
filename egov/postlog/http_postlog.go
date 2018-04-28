package postlog

import (
	"fmt"
	"encoding/json"
	"bytes"
	"time"
	"net/http"
	"reflect"
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
func (p *Pool) HttpPostLog2(b []byte, logsUrl string) error {
	err := p.ScheduleTimeout(5*time.Second, func(){
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
	if reflect.TypeOf(reflect.ValueOf(c)).Kind() == reflect.String{
		c=simpleContent{Message: c}
	}
	msg := map[string]interface{}{"system": s, "kind": k, "content": c}
	err := p.HttpPostLog(msg, logsUrl)
	if err != nil {
		fmt.Println("log send fail", err.Error())
	}
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
