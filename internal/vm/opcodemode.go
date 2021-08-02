package vm

import (
	"fmt"
)

func toMode(vm *VM, m string) (bool, error) {
	switch m {
	case ";":
		vm.Mode = true
		vm.Debug("STACK: pushing to END")
		return true, nil
	case ":":
		vm.Mode = false
		vm.Debug("STACK: pushing to FRONT")
		return false, nil
	}
	return false, fmt.Errorf("I can not determin which mode I shall put stack to: %v", m)
}

func ModeParser(vm *VM, args ...interface{}) {
	if len(args) == 0 {
		vm.Error("I can not determin which mode I shall put stack to")
		return
	}
	res, err := toMode(vm, args[0].(string))
	if err != nil {
		vm.Error("MODE Error: %v", err)
		return
	}
	vm.Info("MODE %v", res)
}

func ModeEval(vm *VM, args ...interface{}) (*Elem, error) {
	if len(args) == 0 {
		return nil, fmt.Errorf("I can not determin which mode I shall put stack to")
	}
	_, err := toMode(vm, args[0].(string))
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func ModeLambda(vm *VM, args ...interface{}) (*Elem, error) {
	vm.Debug("MODE(in lambda) %v", args)
	return &Elem{Type: "MODE", Value: args[0].(string)}, nil
}

func ModeImport(vm *VM, args ...interface{}) {
	vm.Fatal("MODE Import not implemented")
}

func ModeExport(vm *VM, args ...interface{}) {
	vm.Fatal("MODE Export not implemented")
}

func InitOpcodeMode(vm *VM) {
	vm.RegisterOpcode("MODE", ModeParser, ModeLambda, ModeEval, ModeImport, ModeExport)
}
