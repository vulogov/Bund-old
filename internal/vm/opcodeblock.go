package vm

import (
	"fmt"
)

func BlockParser(vm *VM, args ...interface{}) {
	vm.Info("BLOCK %v", args)
}

func BlockEval(vm *VM, args ...interface{}) (*Elem, error) {
	vm.Debug("BLOCK %v", args)
	if len(args) > 0 {
		vm.GetNS(args[0].(string))
		return nil, nil
	}
	return nil, fmt.Errorf("Do not have enough data about BLOCK()")
}

func BlockLambda(vm *VM, args ...interface{}) (*Elem, error) {
	vm.Debug("BLOCK(in lambda) %v", args)
	return &Elem{Type: "BLOCK", Value: args[0].(string)}, nil
}

func BlockImport(vm *VM, args ...interface{}) {
	vm.Fatal("BLOCK Import not implemented")
}

func BlockExport(vm *VM, args ...interface{}) {
	vm.Fatal("BLOCK Export not implemented")
}

func ExitBlockParser(vm *VM, args ...interface{}) {
	vm.Info("exit BLOCK %v", args)
}

func ExitBlockEval(vm *VM, args ...interface{}) (*Elem, error) {
	vm.Debug("exit BLOCK %v", args)
	name := vm.CurrentNS.Name
	vm.EndNS()
	vm.DelNS(name)
	return nil, nil
}

func ExitBlockLambda(vm *VM, args ...interface{}) (*Elem, error) {
	vm.Debug("exit BLOCK(in lambda) %v", args)
	return &Elem{Type: "BLOCK", Value: nil}, nil
}

func ExitBlockImport(vm *VM, args ...interface{}) {
	vm.Fatal("exit BLOCK Import not implemented")
}

func ExitBlockExport(vm *VM, args ...interface{}) {
	vm.Fatal("exit BLOCK Export not implemented")
}

func InitOpcodeBlock(vm *VM) {
	vm.RegisterOpcode("BLOCK", BlockParser, BlockLambda, BlockEval, BlockImport, BlockExport)
	vm.RegisterOpcode("exitBLOCK", ExitBlockParser, ExitBlockLambda, ExitBlockEval, ExitBlockImport, ExitBlockExport)
}
