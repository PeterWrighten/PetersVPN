package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func waitSignal() {
	var sigChan = make(chan os.Signal, 1)
	signal.Notify(sigChan)
	for sig := range sigChan {
		if sig == syscall.SIGINT || sig == syscall.SIGTERM {
			log.Printf("terminated by signal %v\n", sig)
			return
		} else {
			log.Printf("received signal: %v, ignore\n", sig)
		}
	}
}

func main() {
	var faddr, baddr, cryptoMethod, secret, logTo string
	var clientMode bool
	flag.StringVar(&logTo, "logto", "stdout", "stdout or syslog")
}
