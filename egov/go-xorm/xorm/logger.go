// Copyright 2015 The Xorm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xorm

import (
	"egov/common"
	"fmt"
	"io"
	"log"
	"time"
)

// default log options
const (
	DEFAULT_LOG_PREFIX = "[xorm]"
	DEFAULT_LOG_FLAG   = log.Ldate | log.Lmicroseconds
	DEFAULT_LOG_LEVEL  = common.LOG_DEBUG
)

var _ common.ILogger = DiscardLogger{}

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
func (DiscardLogger) Level() common.LogLevel {
	return common.LOG_OFF
}

// SetLevel empty implementation
func (DiscardLogger) SetLevel(l common.LogLevel) {}

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
	level   common.LogLevel
	showSQL bool
}

var _ common.ILogger = &SimpleLogger{}

// NewSimpleLogger use a special io.Writer as logger output
func NewSimpleLogger(out io.Writer) *SimpleLogger {
	return NewSimpleLogger2(out, DEFAULT_LOG_PREFIX, DEFAULT_LOG_FLAG)
}

// NewSimpleLogger2 let you customrize your logger prefix and flag
func NewSimpleLogger2(out io.Writer, prefix string, flag int) *SimpleLogger {
	return NewSimpleLogger3(out, prefix, flag, common.DEFAULT_LOG_LEVEL)
}

// NewSimpleLogger3 let you customrize your logger prefix and flag and logLevel
func NewSimpleLogger3(out io.Writer, prefix string, flag int, l common.LogLevel) *SimpleLogger {
	return &SimpleLogger{
		DEBUG: log.New(out, fmt.Sprintf(time.Now().Local().Format("2006-01-02 15:04:05 ")+"DEBUG %s ",prefix), flag),
		ERR:   log.New(out, fmt.Sprintf(time.Now().Local().Format("2006-01-02 15:04:05 ")+"ERROR %s ", prefix), flag),
		INFO:  log.New(out, fmt.Sprintf(time.Now().Local().Format("2006-01-02 15:04:05 ")+"INFO  %s ", prefix), flag),
		WARN:  log.New(out, fmt.Sprintf(time.Now().Local().Format("2006-01-02 15:04:05 ")+"WARN  %s ", prefix), flag),
		level: l,
	}
}

// Error implement core.ILogger
func (s *SimpleLogger) Error(v ...interface{}) {
	if s.level <= common.LOG_ERR {
		s.ERR.Output(2, fmt.Sprint(v...))
	}
	return
}

// Errorf implement core.ILogger
func (s *SimpleLogger) Errorf(format string, v ...interface{}) {
	if s.level <= common.LOG_ERR {
		s.ERR.Output(2, fmt.Sprintf(format, v...))
	}
	return
}

// Debug implement core.ILogger
func (s *SimpleLogger) Debug(v ...interface{}) {
	if s.level <= common.LOG_DEBUG {
		s.DEBUG.Output(2, fmt.Sprint(v...))
	}
	return
}

// Debugf implement core.ILogger
func (s *SimpleLogger) Debugf(format string, v ...interface{}) {
	if s.level <= common.LOG_DEBUG {
		s.DEBUG.Output(2, fmt.Sprintf(format, v...))
	}
	return
}

// Info implement core.ILogger
func (s *SimpleLogger) Info(v ...interface{}) {
	if s.level <= common.LOG_INFO {
		s.INFO.Output(2, fmt.Sprint(v...))
	}
	return
}

// Infof implement core.ILogger
func (s *SimpleLogger) Infof(format string, v ...interface{}) {
	if s.level <= common.LOG_INFO {
		s.INFO.Output(2, fmt.Sprintf(format, v...))
	}
	return
}

// Warn implement core.ILogger
func (s *SimpleLogger) Warn(v ...interface{}) {
	if s.level <= common.LOG_WARNING {
		s.WARN.Output(2, fmt.Sprint(v...))
	}
	return
}

// Warnf implement core.ILogger
func (s *SimpleLogger) Warnf(format string, v ...interface{}) {
	if s.level <= common.LOG_WARNING {
		s.WARN.Output(2, fmt.Sprintf(format, v...))
	}
	return
}

// Level implement core.ILogger
func (s *SimpleLogger) Level() common.LogLevel {
	return s.level
}

// SetLevel implement core.ILogger
func (s *SimpleLogger) SetLevel(l common.LogLevel) {
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
