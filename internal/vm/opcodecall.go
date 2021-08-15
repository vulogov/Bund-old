package vm

import (
	"fmt"
)

func CallParser(vm *VM, args ...interface{}) {
	vm.Info("CALL %v", args)
}

func CallEval(vm *VM, args ...interface{}) (*Elem, error) {
	if len(args) == 0 {
		return nil, fmt.Errorf("I do not know what to CALL")
	}
	vm.Debug("CALL(%v): %v", vm.CurrentNS.Name, args)
	val, err := vm.Exec(args[0].(string))
	if err != nil {
		return nil, fmt.Errorf("F(%v) returns: %v", args[0].(string), err)
	}
	return val, nil
}

func CallLambda(vm *VM, args ...interface{}) (*Elem, error) {
	if len(args) == 0 {
		return nil, fmt.Errorf("CALL: I do not know what to call")
	}
	return &Elem{Type: "CALL", Value: args[0].(string)}, nil
}

func CallImport(vm *VM, args ...interface{}) {
	vm.Fatal("Call Import not implemented")
}

func CallExport(vm *VM, args ...interface{}) {
	vm.Fatal("Call Export not implemented")
}

func InitOpcodeCall(vm *VM) {
	vm.RegisterOpcode("CALL", CallParser, CallLambda, CallEval, CallImport, CallExport)
}
