package signal

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/pieterclaerhout/go-log"
)

var wg sync.WaitGroup
var exitChan = make(chan bool, 128)
var ng = 0

func signalHandler() {
	log.Debug("Running signal handler")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)
	<-c
	log.Info(fmt.Sprintf("Interruption signal detected. N=%d", ng))
	ExitRequest()
}

func ExitRequest() {
	log.Debug(fmt.Sprintf("Exit requested. N=%d", ng))
	for i := 0; i < ng; i++ {
		exitChan <- true
	}
}

func ExitRequested() bool {
	if len(exitChan) > 0 {
		log.Debug(fmt.Sprintf("Exit requested: %[1]d", len(exitChan)))
	}
	select {
	case _, ok := <-exitChan:
		if ok {
			ExitRequest()
			return true
		} else {
			ExitRequest()
			return true
		}
	default:
		return false
	}
}

func WG() *sync.WaitGroup {
	return &wg
}

func Len() int {
	return len(exitChan)
}

func InitSignal() {
	log.Debug("Configuring signals")
	go signalHandler()
}

func Reserve(n int) int {
	ng = ng + n
	wg.Add(n)
	return ng
}

func Release(n int) int {
	ng = ng - n
	for i := 0; i < n; i++ {
		wg.Done()
	}
	return ng
}

func Loop() {
	wg.Wait()
}
