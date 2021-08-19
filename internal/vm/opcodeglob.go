package vm

import (
	"fmt"
)

func GlobParser(vm *VM, args ...interface{}) {
	vm.Info("GLOB %v", args)
}

func GlobEval(vm *VM, args ...interface{}) (*Elem, error) {
	var val *Elem
	vm.Debug("GLOB: %v", args)
	eh, err := vm.GetType("glob")
	if err != nil {
		return nil, fmt.Errorf("BUND type 'glob' not defined: %v", err)
	}
	if len(args) > 0 {
		val = eh.FromString(vm, args[0].(string))
	} else {
		return nil, fmt.Errorf("Do not know how to convert 'glob'")
	}
	return val, nil
}

func GlobLambda(vm *VM, args ...interface{}) (*Elem, error) {
	val, err := GlobEval(vm, args...)
	if val != nil {
		return val, nil
	}
	return nil, err
}

func GlobImport(vm *VM, args ...interface{}) {
	vm.Fatal("Glob Import not implemented")
}

func GlobExport(vm *VM, args ...interface{}) {
	vm.Fatal("Glob Export not implemented")
}

func InitOpcodeGlob(vm *VM) {
	vm.RegisterOpcode("glob", GlobParser, GlobLambda, GlobEval, GlobImport, GlobExport)
}
