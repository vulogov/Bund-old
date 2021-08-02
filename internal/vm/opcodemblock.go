package vm

import (
	"fmt"

	"github.com/gammazero/deque"
)

func MBlockParser(vm *VM, args ...interface{}) {
	vm.Info("MBLOCK %v", args)
}

func MBlockEval(vm *VM, args ...interface{}) (*Elem, error) {
	vm.Debug("MBLOCK: %v", args)
	if len(args) > 0 {
		vm.GetNS(args[0].(string))
		return nil, nil
	}
	return nil, fmt.Errorf("Do not have enough data about MBLOCK()")
}

func MBlockLambda(vm *VM, args ...interface{}) (*Elem, error) {
	vm.Debug("MBLOCK(in lambda) %v", args)
	return &Elem{Type: "MBLOCK", Value: args[0].(string)}, nil
}

func MBlockImport(vm *VM, args ...interface{}) {
	vm.Fatal("MBLOCK Import not implemented")
}

func MBlockExport(vm *VM, args ...interface{}) {
	vm.Fatal("MBLOCK Export not implemented")
}

func ExitMBlockParser(vm *VM, args ...interface{}) {
	vm.Info("exit MBLOCK %v", args)
}

func ExitMBlockEval(vm *VM, args ...interface{}) (*Elem, error) {
	var val *Elem
	var q deque.Deque
	vm.Debug("exit MBLOCK: %v", args)
	if vm.NSStack.Len() < 1 {
		x := vm.NSStack.Len()
		vm.EndNS()
		return nil, fmt.Errorf("Stack is too shallow for a block matching: %v", x)
	}
	if vm.NSStack.Len() == 1 {
		q = vm.RootNS.Stack
	} else {
		q = vm.NSStack.At(1).(deque.Deque)
	}
	if vm.Mode {
		val = q.Back().(*Elem)
	} else {
		val = q.Front().(*Elem)
	}
	res := vm.CblockMatch(vm.Current, val)
	vm.EndNS()
	vm.Debug("Match block result: %v", res)
	return &Elem{Type: "bool", Value: res}, nil
}

func ExitMBlockLambda(vm *VM, args ...interface{}) (*Elem, error) {
	vm.Debug("exit MBLOCK(in lambda) %v", args)
	return &Elem{Type: "exitMBLOCK", Value: nil}, nil
}

func ExitMBlockImport(vm *VM, args ...interface{}) {
	vm.Fatal("exit MBLOCK Import not implemented")
}

func ExitMBlockExport(vm *VM, args ...interface{}) {
	vm.Fatal("exit MBLOCK Export not implemented")
}

func InitOpcodeMBlock(vm *VM) {
	vm.RegisterOpcode("MBLOCK", MBlockParser, MBlockLambda, MBlockEval, MBlockImport, MBlockExport)
	vm.RegisterOpcode("exitMBLOCK", ExitMBlockParser, ExitMBlockLambda, ExitMBlockEval, ExitMBlockImport, ExitMBlockExport)
}
