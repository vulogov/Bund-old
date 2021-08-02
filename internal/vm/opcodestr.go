package vm

import (
	"fmt"
)

func StrParser(vm *VM, args ...interface{}) {
	vm.Info("STRING %v", args)
}

func StrEval(vm *VM, args ...interface{}) (*Elem, error) {
	var val *Elem
	vm.Debug("STRING: %v", args)
	eh, err := vm.GetType("str")
	if err != nil {
		return nil, fmt.Errorf("BUND type 'str' not defined: %v", err)
	}
	if len(args) > 0 {
		val = eh.FromString(vm, args[0].(string))
	} else {
		return nil, fmt.Errorf("Do not know how to convert 'str'")
	}
	return val, nil
}

func StrLambda(vm *VM, args ...interface{}) (*Elem, error) {
	val, err := StrEval(vm, args...)
	if val != nil {
		return val, nil
	}
	return nil, err
}

func StrImport(vm *VM, args ...interface{}) {
	vm.Fatal("String Import not implemented")
}

func StrExport(vm *VM, args ...interface{}) {
	vm.Fatal("String Export not implemented")
}

func InitOpcodeString(vm *VM) {
	vm.RegisterOpcode("str", StrParser, StrLambda, StrEval, StrImport, StrExport)
}
