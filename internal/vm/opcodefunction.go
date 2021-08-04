package vm

import (
	"fmt"
)

func FunctionParser(vm *VM, args ...interface{}) {
	vm.Info("FUNCTION %v", args)
}

func FunctionEval(vm *VM, args ...interface{}) (*Elem, error) {
	if len(args) == 0 {
		return nil, fmt.Errorf("Do not have enough data about FUNCTION()")
	}
	funcname := args[0].(string)
	vm.Debug("FUNCTION(start): %v", funcname)
	if !vm.IsStack() {
		return nil, fmt.Errorf("Attempt of defining Lambda function with empty context")
	}
	vm.CurrentNS.GetLambda(funcname)
	vm.CurrentNS.InLambda(funcname)
	return nil, nil
}

func FunctionLambda(vm *VM, args ...interface{}) (*Elem, error) {
	vm.Debug("FUNCTION(in lambda) %v", args)
	return &Elem{Type: "FUNCTION", Value: args[0].(string)}, nil
}

func FunctionImport(vm *VM, args ...interface{}) {
	vm.Fatal("FUNCTION Import not implemented")
}

func FunctionExport(vm *VM, args ...interface{}) {
	vm.Fatal("FUNCTION Export not implemented")
}

func ExitFunctionParser(vm *VM, args ...interface{}) {
	vm.Info("exit FUNCTION %v", args)
}

func ExitFunctionEval(vm *VM, args ...interface{}) (*Elem, error) {
	if len(args) == 0 {
		return nil, fmt.Errorf("Do not have enough data about FUNCTION()")
	}
	if !vm.IsStack() {
		return nil, fmt.Errorf("Attempt to close Lambda function with empty context")
	}
	funcname := args[0].(string)
	ls := vm.CurrentLambda()
	vm.CurrentNS.CloseLambda()
	vm.Debug("LAMBDA(fin): %v, size: %v", funcname, ls.Len())
	return nil, nil
}

func ExitFunctionLambda(vm *VM, args ...interface{}) (*Elem, error) {
	vm.Debug("exit FUNCTION(in lambda) %v", args)
	return &Elem{Type: "exitFUNCTION", Value: nil}, nil
}

func ExitFunctionImport(vm *VM, args ...interface{}) {
	vm.Fatal("exit FUNCTION Import not implemented")
}

func ExitFunctionExport(vm *VM, args ...interface{}) {
	vm.Fatal("exit FUNCTION Export not implemented")
}

func InitOpcodeFunction(vm *VM) {
	vm.RegisterOpcode("FUNCTION", FunctionParser, FunctionLambda, FunctionEval, FunctionImport, FunctionExport)
	vm.RegisterOpcode("exitFUNCTION", ExitFunctionParser, ExitFunctionLambda, ExitFunctionEval, ExitFunctionImport, ExitFunctionExport)
}
