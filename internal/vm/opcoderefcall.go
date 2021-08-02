package vm

import (
	"fmt"
)

func RCallParser(vm *VM, args ...interface{}) {
	vm.Info("RCALL %v", args)
}

func RCallEval(vm *VM, args ...interface{}) (*Elem, error) {
	if len(args) == 0 {
		return nil, fmt.Errorf("I do not know what to RCALL")
	}
	vm.Debug("RCALL: %v", args)
	return &Elem{Type: "CALL", Value: args[0].(string)}, nil
}

func RCallLambda(vm *VM, args ...interface{}) (*Elem, error) {
	return CallEval(vm, args...)
}

func RCallImport(vm *VM, args ...interface{}) {
	vm.Fatal("RCall Import not implemented")
}

func RCallExport(vm *VM, args ...interface{}) {
	vm.Fatal("RCall Export not implemented")
}

func InitOpcodeRCall(vm *VM) {
	vm.RegisterOpcode("RCALL", RCallParser, RCallLambda, RCallEval, RCallImport, RCallExport)
}
