package vm

import (
	"fmt"
)

func HttpParser(vm *VM, args ...interface{}) {
	vm.Info("HTTP %v", args)
}

func HttpEval(vm *VM, args ...interface{}) (*Elem, error) {
	var val *Elem
	vm.Debug("HTTP: %v", args)
	eh, err := vm.GetType("http")
	if err != nil {
		return nil, fmt.Errorf("BUND type 'http' not defined: %v", err)
	}
	if len(args) > 0 {
		val = eh.FromString(vm, args[0].(string))
	} else {
		return nil, fmt.Errorf("Do not know how to convert 'http'")
	}
	return val, nil
}

func HttpLambda(vm *VM, args ...interface{}) (*Elem, error) {
	val, err := HttpEval(vm, args...)
	if val != nil {
		return val, nil
	}
	return nil, err
}

func HttpImport(vm *VM, args ...interface{}) {
	vm.Fatal("HTTP Import not implemented")
}

func HttpExport(vm *VM, args ...interface{}) {
	vm.Fatal("HTTP Export not implemented")
}

func InitOpcodeHttp(vm *VM) {
	vm.RegisterOpcode("http", HttpParser, HttpLambda, HttpEval, HttpImport, HttpExport)
}
