package vm

import (
	"github.com/pieterclaerhout/go-log"
)

func (vm *VM) Debug(msg string, args ...interface{}) {
	log.Debugf(msg, args...)
}

func (vm *VM) Error(msg string, args ...interface{}) {
	log.Errorf(msg, args...)
}
