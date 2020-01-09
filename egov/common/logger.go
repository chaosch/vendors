package common

import (
	"egov/log"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
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
	Logger = NewSimpleLogger2(os.Stdout, pId+" ["+pName+"]", log.Ldate|log.Ltime|log.Lmicroseconds)

}

type LogLevel int

const (
	// !nashtsai! following level also match syslog.Priority value
	LOG_EXTREME LogLevel = iota
	LOG_DEBUG
	LOG_INFO
	LOG_WARNING
	LOG_ERR
	LOG_OFF
)

// logger interface
type ILogger interface {
	Debug(v ...interface{})
	Debugf(format string, v ...interface{})
	Error(v ...interface{})
	Errorf(format string, v ...interface{})
	Info(v ...interface{})
	Infof(format string, v ...interface{})
	Warn(v ...interface{})
	Warnf(format string, v ...interface{})

	Level() LogLevel
	SetLevel(l LogLevel)

	ShowSQL(show ...bool)
	IsShowSQL() bool
}

const (
	DEFAULT_LOG_PREFIX = ""
	DEFAULT_LOG_FLAG   = log.Ldate | log.Lmicroseconds
	DEFAULT_LOG_LEVEL  = LOG_INFO
)

var _ ILogger = DiscardLogger{}

// DiscardLogger don't log implementation for core.ILogger
type DiscardLogger struct{}

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

// Level empty implementation
func (DiscardLogger) Level() LogLevel {
	return LOG_OFF
}

// SetLevel empty implementation
func (DiscardLogger) SetLevel(l LogLevel) {}

// ShowSQL empty implementation
func (DiscardLogger) ShowSQL(show ...bool) {}

// IsShowSQL empty implementation
func (DiscardLogger) IsShowSQL() bool {
	return false
}

// SimpleLogger is the default implment of core.ILogger
type SimpleLogger struct {
	DEBUG   *log.Logger
	ERR     *log.Logger
	INFO    *log.Logger
	WARN    *log.Logger
	level   LogLevel
	showSQL bool
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
		WARN:  log.New(out, fmt.Sprintf("WARN  %s ", prefix), flag),
		level: l,
	}
}

// Error implement core.ILogger
func (s *SimpleLogger) Error(v ...interface{}) {
	s.OpenExtreme(&v)
	if s.level <= LOG_ERR {
		s.ERR.Output(2, fmt.Sprint(v...))
	}
	return
}

// Errorf implement core.ILogger
func (s *SimpleLogger) Errorf(format string, v ...interface{}) {
	s.OpenExtremeF(&format, &v)
	if s.level <= LOG_ERR {
		s.ERR.Output(2, fmt.Sprintf(format, v...))
	}
	return
}

// Debug implement core.ILogger
func (s *SimpleLogger) Debug(v ...interface{}) {
	s.OpenExtreme(&v)
	if s.level <= LOG_DEBUG {
		s.DEBUG.Output(2, fmt.Sprint(v...))
	}
	return
}

// Debugf implement core.ILogger
func (s *SimpleLogger) Debugf(format string, v ...interface{}) {
	s.OpenExtremeF(&format, &v)
	if s.level <= LOG_DEBUG {
		s.DEBUG.Output(2, fmt.Sprintf(format, v...))
	}
	return
}

// Info implement core.ILogger
func (s *SimpleLogger) Info(v ...interface{}) {
	s.OpenExtreme(&v)
	if s.level <= LOG_INFO {
		s.INFO.Output(2, fmt.Sprint(v...))
	}
	return
}

// Infof implement core.ILogger
func (s *SimpleLogger) Infof(format string, v ...interface{}) {
	s.OpenExtremeF(&format, &v)
	if s.level <= LOG_INFO {
		s.INFO.Output(2, fmt.Sprintf(format, v...))
	}
	return
}

// Warn implement core.ILogger
func (s *SimpleLogger) Warn(v ...interface{}) {
	s.OpenExtreme(&v)
	if s.level <= LOG_WARNING {
		s.WARN.Output(2, fmt.Sprint(v...))
	}
	return
}

// Warnf implement core.ILogger
func (s *SimpleLogger) Warnf(format string, v ...interface{}) {
	s.OpenExtremeF(&format, &v)
	if s.level <= LOG_WARNING {
		s.WARN.Output(2, fmt.Sprintf(format, v...))
	}
	return
}

// Level implement core.ILogger
func (s *SimpleLogger) Level() LogLevel {
	return s.level
}

// SetLevel implement core.ILogger
func (s *SimpleLogger) SetLevel(l LogLevel) {
	s.level = l
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

func (s *SimpleLogger) OpenExtremeF(format *string, v *[]interface{}) {
	if s.level <= LOG_EXTREME {
		_, file, line, _ := runtime.Caller(2)
		if idx := strings.Index(file, "/src/"); idx >= 0 {
			file = file[idx+5:]
		}
		//vr := make([]interface{}, 0)
		//vr = append(vr, file)
		//vr = append(vr, line)
		*v = append(*v, file)
		*v = append(*v, line)
		*format = *format + " <--file:%s line:%d-->"
	}
}

func (s *SimpleLogger) OpenExtreme(v *[]interface{}) {
	if s.level <= LOG_EXTREME {
		_, file, line, _ := runtime.Caller(2)
		if idx := strings.Index(file, "/src/"); idx >= 0 {
			file = file[idx+5:]
		}
		//vr := make([]interface{}, 0)
		//vr = append(vr, fmt.Sprintf("file:%s ",file))
		//vr = append(vr, fmt.Sprintf("line:%d ",line))
		*v = append(*v, fmt.Sprintf(" <--file:%s ", file))
		*v = append(*v, fmt.Sprintf("line:%d-->", line))
	}
}
