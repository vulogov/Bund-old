package vm

import (
	"github.com/gammazero/deque"
)

func (vm *VM) CurrentLambda() *deque.Deque {
	if !vm.IsStack() {
		vm.Fatal("Attempt to CurrentLambda() but NS doesn't exists")
		return nil
	}
	return vm.CurrentNS.CurrentLambda()
}

func (vm *VM) InLambda() bool {
	if !vm.IsStack() {
		vm.Debug("Attempt to InLambda() but Stack doesn't exists")
		return false
	}
	if vm.CurrentNS.LambdasStack.Len() > 0 {
		vm.Debug("We are in InLambda(%v)", vm.CurrentNS.LambdasStack.Back().(string))
		return true
	}
	return false
}
