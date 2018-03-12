package main

import (
	"egov/postLog"
)

func main() {

	var p = postlog.NewPl(128, 1, 1)
	p.SendLog(postlog.D, postlog.E, "aaa","")
}
