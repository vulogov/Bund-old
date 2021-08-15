package vm

import (
	"fmt"
)

func OpParser(vm *VM, args ...interface{}) {
	vm.Info("OP %v", args)
}

func OpEval(vm *VM, args ...interface{}) (*Elem, error) {
	if len(args) == 0 {
		return nil, fmt.Errorf("I do not know what to OPERATOR")
	}
	vm.Debug("OP: %v", args)
	val, err := vm.Op(args[0].(string))
	if err != nil {
		return nil, fmt.Errorf("O(%v) returns: %v", args[0].(string), err)
	}
	return val, nil
}

func OpLambda(vm *VM, args ...interface{}) (*Elem, error) {
	if len(args) == 0 {
		return nil, fmt.Errorf("OP: I do not know what to call")
	}
	vm.Debug("OP(in lambda): %v", args)
	return &Elem{Type: "OP", Value: args[0].(string)}, nil
}

func OpImport(vm *VM, args ...interface{}) {
	vm.Fatal("Operator Import not implemented")
}

func OpExport(vm *VM, args ...interface{}) {
	vm.Fatal("Operator Export not implemented")
}

func InitOpcodeOperator(vm *VM) {
	vm.RegisterOpcode("OP", OpParser, OpLambda, OpEval, OpImport, OpExport)
}
