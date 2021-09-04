package vm

import (
	"fmt"
)

func JsonParser(vm *VM, args ...interface{}) {
	vm.Info("JSON %v", args)
}

func JsonEval(vm *VM, args ...interface{}) (*Elem, error) {
	var val *Elem
	vm.Debug("JSON: %v", args)
	eh, err := vm.GetType("json")
	if err != nil {
		return nil, fmt.Errorf("BUND type 'json' not defined: %v", err)
	}
	if len(args) > 0 {
		val = eh.FromString(vm, args[0].(string))
	} else {
		return nil, fmt.Errorf("Do not know how to convert 'json'")
	}
	return val, nil
}

func JsonLambda(vm *VM, args ...interface{}) (*Elem, error) {
	val, err := StrEval(vm, args...)
	if val != nil {
		return val, nil
	}
	return nil, err
}

func JsonImport(vm *VM, args ...interface{}) {
	vm.Fatal("String Import not implemented")
}

func JsonExport(vm *VM, args ...interface{}) {
	vm.Fatal("String Export not implemented")
}

func InitOpcodeJson(vm *VM) {
	vm.RegisterOpcode("json", JsonParser, JsonLambda, JsonEval, JsonImport, JsonExport)
}
