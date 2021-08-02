package vm

import (
	"fmt"
)

func ComplexParser(vm *VM, args ...interface{}) {
	vm.Info("COMPLEX %v", args)
}

func ComplexEval(vm *VM, args ...interface{}) (*Elem, error) {
	var val *Elem
	vm.Debug("COMPLEX: %v", args)
	eh, err := vm.GetType("cpx")
	if err != nil {
		return nil, fmt.Errorf("BUND type 'cpx' not defined: %v", err)
	}
	if len(args) > 0 {
		val = eh.FromString(vm, args[0].(string))
	} else {
		return nil, fmt.Errorf("Do not know how to convert 'cpx'")
	}
	return val, nil
}

func ComplexLambda(vm *VM, args ...interface{}) (*Elem, error) {
	val, err := IntegerEval(vm, args...)
	if val != nil {
		return val, nil
	}
	return nil, err
}

func ComplexImport(vm *VM, args ...interface{}) {
	vm.Fatal("Integer Import not implemented")
}

func ComplexExport(vm *VM, args ...interface{}) {
	vm.Fatal("Complex128 Export not implemented")
}

func InitOpcodeComplex(vm *VM) {
	vm.RegisterOpcode("cpx", ComplexParser, ComplexLambda, ComplexEval, ComplexImport, ComplexExport)
}
