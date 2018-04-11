package common

import "time"

type ProcessStatus struct {
	InBuf         string     `json:"inbuf"`
	Starttime     time.Time  `json:"starttime"`
	EndTime       time.Time  `json:"endtime"`
	Duration      int64      `json:"duration"`
	OK            bool       `json:"ok"`
	Err           ErrContext `json:"err"`
	SqlDuration   int64      `json:"sqlduration"`
	Changes       int64      `json:"changes"`
	ChipId        int64      `json:"chipid"`
	DalVersion    string     `json:"dalversion"`
	IpAddress     string     `json:"ipaddress"`
	InTransaction bool       `json:"intransaction"`
	Prompt        string     `json:"prompt"`
	Version       string     `json:"version"`
}

type LogData struct {
	system  string
	kind    string
	content interface{}
}
