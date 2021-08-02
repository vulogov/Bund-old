package vm

import (
	"fmt"
)

func IntegerParser(vm *VM, args ...interface{}) {
	vm.Info("INTEGER %v", args)
}

func IntegerEval(vm *VM, args ...interface{}) (*Elem, error) {
	var val *Elem
	vm.Debug("INTEGER: %v", args)
	eh, err := vm.GetType("int")
	if err != nil {
		return nil, fmt.Errorf("BUND type 'int' not defined: %v", err)
	}
	if len(args) > 0 {
		val = eh.FromString(vm, args[0].(string))
	} else {
		return nil, fmt.Errorf("Do not know how to convert 'int'")
	}
	return val, nil
}

func IntegerLambda(vm *VM, args ...interface{}) (*Elem, error) {
	val, err := IntegerEval(vm, args...)
	if val != nil {
		return val, nil
	}
	return nil, err
}

func IntegerImport(vm *VM, args ...interface{}) {
	vm.Fatal("Integer Import not implemented")
}

func IntegerExport(vm *VM, args ...interface{}) {
	vm.Fatal("Integer Export not implemented")
}

func InitOpcodeInteger(vm *VM) {
	vm.RegisterOpcode("int", IntegerParser, IntegerLambda, IntegerEval, IntegerImport, IntegerExport)
}
