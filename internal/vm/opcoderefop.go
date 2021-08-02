package vm

import (
	"fmt"
)

func ROpParser(vm *VM, args ...interface{}) {
	vm.Info("ROP %v", args)
}

func ROpEval(vm *VM, args ...interface{}) (*Elem, error) {
	if len(args) == 0 {
		return nil, fmt.Errorf("I do not know what to ROP")
	}
	vm.Debug("ROP: %v", args)
	return &Elem{Type: "OP", Value: args[0].(string)}, nil
}

func ROpLambda(vm *VM, args ...interface{}) (*Elem, error) {
	return CallEval(vm, args...)
}

func ROpImport(vm *VM, args ...interface{}) {
	vm.Fatal("ROP Import not implemented")
}

func ROpExport(vm *VM, args ...interface{}) {
	vm.Fatal("ROP Export not implemented")
}

func InitOpcodeRop(vm *VM) {
	vm.RegisterOpcode("ROP", ROpParser, ROpLambda, ROpEval, ROpImport, ROpExport)
}
