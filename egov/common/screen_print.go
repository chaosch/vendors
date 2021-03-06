package common

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const (
	INFO  LogType = "INFO "
	WARN  LogType = "WARN "
	ERROR LogType = "ERROR"
	DEBUG LogType = "DEBUG"
)

type SystemLogLevel int

type LogType string

func PrintScreen(LogLevel LogLevel, LogContent string) {
	var LogTime, Pid, Pname string
	LogTime = time.Now().Local().Format("2006-01-02 15:04:05")
	Pid = strconv.FormatInt(int64(os.Getpid()), 10)
	_, Pname = filepath.Split(os.Args[0])
	fmt.Printf("%s %v %s [%s]: /*%s*/\n", LogTime, LogLevel, Pid, Pname, LogContent)
}

func ScreenLog1(SystemLogLeveL LogLevel, LogLevel LogType, LogContent string, paras ...interface{}) {
	pCount := strings.Count(LogContent, "%s")

	var LogContentString string
	if pCount > 0 {
		if len(paras) > 0 {
			pCount = int(math.Min(float64(pCount), float64(len(paras))))
			LogContentString = fmt.Sprintf(LogContent, paras[0:pCount-1]...)
		} else {
			LogContentString = fmt.Sprintf(LogContent, paras...)
		}
	} else {
		LogContentString = fmt.Sprintf(LogContent, paras...)
	}
	fmt.Println(LogContentString)
	//switch SystemLogLevel(SystemLogLeveL) {
	//case LOG_OFF:
	//	return
	//case LOG_INFO:
	//	if LogLevel == ERROR || LogLevel == WARN || LogLevel == INFO {
	//		PrintScreen(LogLevel, LogContentString)
	//	}
	//case LOG_WARNING:
	//	if LogLevel == ERROR || LogLevel == WARN {
	//		PrintScreen(LogLevel, LogContentString)
	//	}
	//case LOG_ERR:
	//	if LogLevel == ERROR {
	//		PrintScreen(LogLevel, LogContentString)
	//	}
	//case LOG_DEBUG:
	//	if LogLevel == ERROR || LogLevel == WARN || LogLevel == INFO || LogLevel == DEBUG {
	//		PrintScreen(LogLevel, LogContentString)
	//	}
	//default:
	//	PrintScreen(LogLevel, LogContentString)
	//}
}
