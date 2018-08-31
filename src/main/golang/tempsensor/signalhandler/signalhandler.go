package signalhandler

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func SetupSignalHandler() {
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, syscall.SIGTERM)
	signal.Notify(signalChan, syscall.SIGINT)
	go func() {
		sig := <-signalChan
		log.Printf("caught signal: %+v, terminating in 2 seconds\n", sig)
		time.Sleep(2 * time.Second)
		os.Exit(0)
	}()
}
