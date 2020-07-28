package server

import (
	"os"
	"os/signal"
	"syscall"
)

func WaitSignal() os.Signal {
	sig := make(chan os.Signal, 1)
	defer close(sig)

	signal.Notify(sig,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	return <-sig
}
