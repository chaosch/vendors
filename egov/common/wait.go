package common

import (
	"os"
	"os/signal"
	"syscall"
)

func WaitForSignal() {
	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGTERM)
	signal.Notify(sigs, syscall.SIGINT)
	<-sigs
}

