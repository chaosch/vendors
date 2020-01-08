package common

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

const (
	INFO  = "INFO"
	WARN  = "WARN"
	ERROR = "ERROR"
	DEBUG = "DEBUG"
)

type LogLevel string

func PrintScreen(LogLevel LogLevel, LogContent string) {
	var LogTime, Pid, Pname string
	LogTime = time.Now().Local().Format("2006-01-02 15:04:05")
	Pid = strconv.FormatInt(int64(os.Getpid()), 10)
	Pname = filepath.Clean(os.Args[0])
	fmt.Printf("%s %s %s [%s]: /*%s*/\n", LogTime, LogLevel, Pid, Pname, LogContent)
}

func ScreenLog(SystemLogLeveL int, LogLevel LogLevel, LogContent string) {
	switch SystemLogLeveL {
	case 0:
		return
	case 1:
		if LogLevel == ERROR {
			PrintScreen(LogLevel, LogContent)
		}
	case 2:
		if LogLevel == ERROR || LogLevel == WARN {
			PrintScreen(LogLevel, LogContent)
		}
	case 3:
		if LogLevel == ERROR || LogLevel == WARN || LogLevel == INFO {
			PrintScreen(LogLevel, LogContent)
		}
	case 4:
		if LogLevel == ERROR || LogLevel == WARN || LogLevel == INFO || LogLevel == DEBUG {
			PrintScreen(LogLevel, LogContent)
		}
	default:
		PrintScreen(LogLevel, LogContent)
	}
}
