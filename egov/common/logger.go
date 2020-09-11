package common

import (
	"egov/log"
	"egov/postlog"
	"fmt"
	"github.com/pquerna/ffjson/ffjson"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"time"
)

// Copyright 2015 The Xorm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// default log options

var Logger ILogger

func init() {
	pId := strconv.FormatInt(int64(os.Getpid()), 10)
	pId = fmt.Sprintf("%5s", pId)
	_, pName := filepath.Split(os.Args[0])
	pName = fmt.Sprintf("%s", pName)
	Logger = NewSimpleLogger2(os.Stdout, pId+" ["+pName+"] /*", log.Ldate|log.Ltime|log.Lmicroseconds|log.Llongfile)

}

type LogLevel int

const (
	// !nashtsai! following level also match syslog.Priority value
	LOG_EXTREME LogLevel = iota
	LOG_SQL
	LOG_DEBUG
	LOG_INFO
	LOG_WATCH
	LOG_WARNING
	LOG_ERR
	LOG_OFF
)

// logger interface
type ILogger interface {
	Sql(v ...interface{})
	Sqlf(format string, v ...interface{})
	Debug(v ...interface{})
	Debugf(format string, v ...interface{})
	Error(v ...interface{})
	Errorf(format string, v ...interface{})
	Info(v ...interface{})
	Infof(format string, v ...interface{})
	Warn(v ...interface{})
	Warnf(format string, v ...interface{})
	Watch(v ...interface{})
	Watchf(format string, v ...interface{})

	Level() LogLevel
	SetLevel(l LogLevel, system string, logServer string, output string)

	ShowSQL(show ...bool)
	IsShowSQL() bool
	SetSensitiveKeys(keys []string)
}

const (
	DEFAULT_LOG_PREFIX = ""
	DEFAULT_LOG_FLAG   = log.Ldate | log.Lmicroseconds | log.Lshortfile
	DEFAULT_LOG_LEVEL  = LOG_EXTREME
	DEFAULT_OUT_PUT    = "stdout"
)

var _ ILogger = DiscardLogger{}

// DiscardLogger don't log implementation for core.ILogger
type DiscardLogger struct{}

func (DiscardLogger) SetSensitiveKeys(keys []string) {}

// Sql empty implementation
func (DiscardLogger) Sql(v ...interface{}) {}

// Sqlf empty implementation
func (DiscardLogger) Sqlf(format string, v ...interface{}) {}

// Debug empty implementation
func (DiscardLogger) Debug(v ...interface{}) {}

// Debugf empty implementation
func (DiscardLogger) Debugf(format string, v ...interface{}) {}

// Error empty implementation
func (DiscardLogger) Error(v ...interface{}) {}

// Errorf empty implementation
func (DiscardLogger) Errorf(format string, v ...interface{}) {}

// Info empty implementation
func (DiscardLogger) Info(v ...interface{}) {}

// Infof empty implementation
func (DiscardLogger) Infof(format string, v ...interface{}) {}

// Warn empty implementation
func (DiscardLogger) Warn(v ...interface{}) {}

// Warnf empty implementation
func (DiscardLogger) Warnf(format string, v ...interface{}) {}

// Watch empty implementation
func (DiscardLogger) Watch(v ...interface{}) {}

// Watchf empty implementation
func (DiscardLogger) Watchf(format string, v ...interface{}) {}

// Level empty implementation
func (DiscardLogger) Level() LogLevel {
	return LOG_OFF
}

// SetLevel empty implementation
func (DiscardLogger) SetLevel(l LogLevel, system string, logServer string, output string) {}

// ShowSQL empty implementation
func (DiscardLogger) ShowSQL(show ...bool) {}

// IsShowSQL empty implementation
func (DiscardLogger) IsShowSQL() bool {
	return false
}

// SimpleLogger is the default implment of core.ILogger
type SimpleLogger struct {
	DEBUG         *log.Logger
	ERR           *log.Logger
	WATCH         *log.Logger
	INFO          *log.Logger
	WARN          *log.Logger
	SQL           *log.Logger
	level         LogLevel
	showSQL       bool
	postLog       bool
	system        string
	ExtraFunction func(p ...interface{})
	logUrl        string
	SensitiveKeys []string
}

var _ ILogger = &SimpleLogger{}

// NewSimpleLogger use a special io.Writer as logger output
func NewSimpleLogger(out io.Writer) *SimpleLogger {
	return NewSimpleLogger2(out, DEFAULT_LOG_PREFIX, DEFAULT_LOG_FLAG)
}

// NewSimpleLogger2 let you customrize your logger prefix and flag
func NewSimpleLogger2(out io.Writer, prefix string, flag int) *SimpleLogger {
	return NewSimpleLogger3(out, prefix, flag, DEFAULT_LOG_LEVEL)
}

// NewSimpleLogger3 let you customrize your logger prefix and flag and logLevel
func NewSimpleLogger3(out io.Writer, prefix string, flag int, l LogLevel) *SimpleLogger {
	return &SimpleLogger{
		DEBUG: log.New(out, fmt.Sprintf("DEBUG %s ", prefix), flag),
		ERR:   log.New(out, fmt.Sprintf("ERROR %s ", prefix), flag),
		INFO:  log.New(out, fmt.Sprintf("INFO  %s ", prefix), flag),
		WATCH: log.New(out, fmt.Sprintf("WATCH %s ", prefix), flag),
		WARN:  log.New(out, fmt.Sprintf("WARN  %s ", prefix), flag),
		SQL:   log.New(out, fmt.Sprintf("SQL   %s ", prefix), flag),
		level: l,
	}
}

func (s *SimpleLogger) SetSensitiveKeys(keys []string) {
	s.SensitiveKeys = keys
	s.DEBUG.SetSensitiveKeys(keys)
	s.INFO.SetSensitiveKeys(keys)
	s.WATCH.SetSensitiveKeys(keys)
	s.WARN.SetSensitiveKeys(keys)
	s.ERR.SetSensitiveKeys(keys)
	s.SQL.SetSensitiveKeys(keys)
	return

}

// Sql implement core.ILogger
func (s *SimpleLogger) Sql(v ...interface{}) {
	s.IfOpenExtreme(LOG_SQL, &v)
	if s.level <= LOG_SQL {
		s.SQL.Output(2, fmt.Sprint(v...))
		//s.SendLog(LOG_ERR, fmt.Sprint(v...))
	}
	if s.ExtraFunction != nil {
		s.ExtraFunction()
	}
	return
}

// Sqlf implement core.ILogger
func (s *SimpleLogger) Sqlf(format string, v ...interface{}) {
	s.IfOpenExtremeF(LOG_SQL, &format, &v)
	if s.level <= LOG_SQL {
		s.SQL.Output(2, fmt.Sprintf(format, v...))
		//s.SendLog(LOG_ERR, fmt.Sprintf(format, v...))
	}
	if s.ExtraFunction != nil {
		s.ExtraFunction()
	}
	return
}

// Error implement core.ILogger
func (s *SimpleLogger) Error(v ...interface{}) {
	s.IfOpenExtreme(LOG_ERR, &v)
	if s.level <= LOG_ERR {
		s.ERR.Output(2, fmt.Sprint(v...))
		//s.SendLog(LOG_ERR, fmt.Sprint(v...))
	}
	if s.ExtraFunction != nil {
		s.ExtraFunction()
	}
	return
}

func (s *SimpleLogger) ErrorNoPost(v ...interface{}) {
	if s.level <= LOG_ERR {
		s.ERR.Output(2, fmt.Sprint(v...))
	}
	if s.ExtraFunction != nil {
		s.ExtraFunction()
	}
	return
}

// Errorf implement core.ILogger
func (s *SimpleLogger) Errorf(format string, v ...interface{}) {
	s.IfOpenExtremeF(LOG_ERR, &format, &v)
	if s.level <= LOG_ERR {
		s.ERR.Output(2, fmt.Sprintf(format, v...))
		//s.SendLog(LOG_ERR, fmt.Sprintf(format, v...))
	}
	if s.ExtraFunction != nil {
		s.ExtraFunction()
	}
	return
}

// Debug implement core.ILogger
func (s *SimpleLogger) Debug(v ...interface{}) {
	s.IfOpenExtreme(LOG_DEBUG, &v)
	if s.level <= LOG_DEBUG {
		s.DEBUG.Output(2, fmt.Sprint(v...))
		//s.SendLog(LOG_DEBUG, fmt.Sprint(v...))
	}
	if s.ExtraFunction != nil {
		s.ExtraFunction()
	}
	return
}

// Debugf implement core.ILogger
func (s *SimpleLogger) Debugf(format string, v ...interface{}) {
	s.IfOpenExtremeF(LOG_DEBUG, &format, &v)
	if s.level <= LOG_DEBUG {
		s.DEBUG.Output(2, fmt.Sprintf(format, v...))
		//s.SendLog(LOG_DEBUG, v)
	}
	if s.ExtraFunction != nil {
		s.ExtraFunction()
	}
	return
}

// Info implement core.ILogger
func (s *SimpleLogger) Info(v ...interface{}) {
	s.IfOpenExtreme(LOG_INFO, &v)
	if s.level <= LOG_INFO {
		s.INFO.Output(2, fmt.Sprint(v...))
		//s.SendLog(LOG_INFO, fmt.Sprint(v...))
	}
	if s.ExtraFunction != nil {
		s.ExtraFunction()
	}
	return
}

// Infof implement core.ILogger
func (s *SimpleLogger) Infof(format string, v ...interface{}) {
	s.IfOpenExtremeF(LOG_INFO, &format, &v)
	if s.level <= LOG_INFO {
		s.INFO.Output(2, fmt.Sprintf(format, v...))
		//s.SendLog(LOG_INFO, fmt.Sprintf(format, v...))
	}
	if s.ExtraFunction != nil {
		s.ExtraFunction()
	}
	return
}

// Warn implement core.ILogger
func (s *SimpleLogger) Warn(v ...interface{}) {
	s.IfOpenExtreme(LOG_WARNING, &v)
	if s.level <= LOG_WARNING {
		s.WARN.Output(2, fmt.Sprint(v...))
	}
	if s.ExtraFunction != nil {
		s.ExtraFunction()
	}
	return
}

// Warnf implement core.ILogger
func (s *SimpleLogger) Warnf(format string, v ...interface{}) {
	s.IfOpenExtremeF(LOG_WARNING, &format, &v)
	if s.level <= LOG_WARNING {
		s.WARN.Output(2, fmt.Sprintf(format, v...))
		//s.SendLog(LOG_WARNING, fmt.Sprintf(format, v...))
	}
	if s.ExtraFunction != nil {
		s.ExtraFunction()
	}
	return
}

// Warn implement core.ILogger
func (s *SimpleLogger) Watch(v ...interface{}) {
	s.IfOpenExtreme(LOG_WATCH, &v)
	if s.level <= LOG_WATCH {
		s.WATCH.Output(2, fmt.Sprint(v...))
	}
	if s.ExtraFunction != nil {
		s.ExtraFunction()
	}
	return
}

// Warnf implement core.ILogger
func (s *SimpleLogger) Watchf(format string, v ...interface{}) {
	s.IfOpenExtremeF(LOG_WATCH, &format, &v)
	if s.level <= LOG_WATCH {
		s.WATCH.Output(2, fmt.Sprintf(format, v...))
		//s.SendLog(LOG_WARNING, fmt.Sprintf(format, v...))
	}
	if s.ExtraFunction != nil {
		s.ExtraFunction()
	}
	return
}

// Level implement core.ILogger
func (s *SimpleLogger) Level() LogLevel {
	return s.level
}

// SetLevel implement core.ILogger
func (s *SimpleLogger) SetLevel(l LogLevel, system string, logUrl string, out string) {
	s.level = l
	s.system = system
	s.logUrl = logUrl
	output := os.Stdout
	if out != DEFAULT_OUT_PUT {
		output, _ = os.OpenFile(system+"_"+time.Now().Local().Format("2006_01_02_15_04_05.log"), os.O_CREATE|os.O_APPEND, os.ModePerm)
	}
	s.DEBUG.SetParas(system, logUrl, output)
	s.INFO.SetParas(system, logUrl, output)
	s.WARN.SetParas(system, logUrl, output)
	s.ERR.SetParas(system, logUrl, output)
	s.SQL.SetParas(system, logUrl, output)
	return
}

// ShowSQL implement core.ILogger
func (s *SimpleLogger) ShowSQL(show ...bool) {
	if len(show) == 0 {
		s.showSQL = true
		return
	}
	s.showSQL = show[0]
}

// IsShowSQL implement core.ILogger
func (s *SimpleLogger) IsShowSQL() bool {
	return s.showSQL
}

func (s *SimpleLogger) OpenExtremeF(format *string, v *[]interface{}, skip int) {
	//_, file, line, _ := runtime.Caller(skip)
	//if idx := strings.Index(file, "/src/"); idx >= 0 {
	//	file = file[idx+5:]
	//}
	////_,file=filepath.Split(file)
	////vr := make([]interface{}, 0)
	////vr = append(vr, file)
	////vr = append(vr, line)
	//*v = append(*v, file)
	//*v = append(*v, line)
	//*format = *format + " <-- %s:%d:-->"
}

func (s *SimpleLogger) OpenExtreme(v *[]interface{}, skip int) {
	//_, file, line, _ := runtime.Caller(skip)
	//if idx := strings.Index(file, "/src/"); idx >= 0 {
	//	file = file[idx+5:]
	//}
	////_,file=filepath.Split(file)
	////vr := make([]interface{}, 0)
	////vr = append(vr, fmt.Sprintf("file:%s ",file))
	////vr = append(vr, fmt.Sprintf("line:%d ",line))
	//*v = append(*v, fmt.Sprintf(" <-- %s:", file))
	//*v = append(*v, fmt.Sprintf("%d:-->", line))
}

//*buf = append(*buf, file...)
//*buf = append(*buf, ':')
//itoa(buf, line, -1)
//*buf = append(*buf, ":"...)

func (s *SimpleLogger) IfOpenExtremeF(lT LogLevel, format *string, v *[]interface{}) {
	ostr := fmt.Sprintf(*format, v)
	s.SendLog(lT, ostr)
}

func (s *SimpleLogger) IfOpenExtreme(lT LogLevel, v *[]interface{}) {
	//ostr := ""
	if len(*v) == 1 {
		x := reflect.Indirect(reflect.ValueOf(v))
		y := x.Index(0).Interface()
		//fmt.Println(reflect.ValueOf(y).Type().Name())
		if y != nil {
			var buff []byte
			if reflect.ValueOf(y).Type().Name() == "ProcessStatus" {
				buff, _ = ffjson.Marshal(y)
				s.SendLog(lT, y)
				ostr := string(buff)
				x.Index(0).Set(reflect.ValueOf(ostr))
			}else{
				pc:=postlog.ProcessStatus{}
				pc.Prompt=fmt.Sprintf("%v",y.(interface{}))
				s.SendLog(lT, pc.Prompt)
			}
		}
	} else if len(*v) > 1 {
		pc := postlog.ProcessStatus{}
		for _,val :=range *v{
			pc.Prompt+=fmt.Sprintf("%v ",val)
		}
		s.SendLog(lT, pc.Prompt)
		//ostr := pc.Prompt
		//x.Index(0).Set(reflect.ValueOf(ostr))
	}
}

func (s *SimpleLogger) SendLog(lT LogLevel, logStr interface{}) {
	if lT != LOG_WARNING && lT != LOG_ERR && lT != LOG_WATCH {
		return
	}
	go func() {
		logKind := ""
		switch lT {
		case LOG_INFO:
			logKind = "日志"
		case LOG_WATCH:
			logKind = "日志"
		case LOG_DEBUG:
			logKind = "日志"
		case LOG_ERR:
			logKind = "错误"
		case LOG_WARNING:
			logKind = "警告"
		default:
			return
		}

		if s.logUrl != "" {
			err := postlog.Log.SendLog(s.system, logKind, logStr, s.logUrl)
			if err != nil {
				s.ErrorNoPost(err)
			}
		}
	}()
}
