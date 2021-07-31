package vm

import (
	"os"

	"github.com/pieterclaerhout/go-log"
)

func (vm *VM) Debug(msg string, args ...interface{}) {
	log.Debugf(msg, args...)
}

func (vm *VM) Info(msg string, args ...interface{}) {
	log.Infof(msg, args...)
}

func (vm *VM) Warning(msg string, args ...interface{}) {
	log.Warnf(msg, args...)
}

func (vm *VM) Error(msg string, args ...interface{}) {
	log.Errorf(msg, args...)
}

func (vm *VM) Fatal(msg string, args ...interface{}) {
	log.Fatalf(msg, args...)
}

func (vm *VM) OnError(err error, msg string, args ...interface{}) {
	if err != nil {
		vm.Error(msg, args...)
		log.StackTrace(err)
		os.Exit(99)
	}
}
