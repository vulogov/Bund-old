package vm

import (
	"fmt"
)

func LBlockParser(vm *VM, args ...interface{}) {
	vm.Info("LBLOCK %v", args)
}

func LBlockEval(vm *VM, args ...interface{}) (*Elem, error) {
	if len(args) < 3 {
		return nil, fmt.Errorf("There is no information from the LBLOCK header")
	}
	vm.Debug("LBLOCK: %v", args)
	blockname := args[0].(string)
	lbmode := args[2].(bool)
	if vm.CanGet() {
		e := vm.Get()
		if e.Type == "bool" {
			if e.Value.(bool) == false && lbmode == false {
				vm.NotIgnore()
				vm.GetNS(blockname)
				vm.Debug("%v Block will be executed", lbmode)
				return nil, nil
			} else if e.Value.(bool) == true && lbmode == true {
				vm.NotIgnore()
				vm.GetNS(blockname)
				vm.Debug("%v Block will be executed", lbmode)
				return nil, nil
			}
			vm.Debug("%v Block will not be executed", lbmode)
			vm.Ignore()
		} else {
			vm.Ignore()
			vm.Debug("%v Block will not be executed because there is no bool data in stack", lbmode)
		}
		return nil, nil
	}
	return nil, fmt.Errorf("Do not have enough data about LBLOCK()")
}

func LBlockLambda(vm *VM, args ...interface{}) (*Elem, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("There is no information from the LBLOCK header")
	}
	vm.Debug("LBLOCK: %v", args)
	blockname := args[0].(string)
	lbmode := args[1].(bool)
	vm.Debug("LBLOCK(in lambda) %v", args)
	if lbmode == true {
		return &Elem{Type: "TRUEBLOCK", Value: blockname}, nil
	} else {
		return &Elem{Type: "FALSEBLOCK", Value: blockname}, nil
	}
}

func LBlockImport(vm *VM, args ...interface{}) {
	vm.Fatal("LBLOCK Import not implemented")
}

func LBlockExport(vm *VM, args ...interface{}) {
	vm.Fatal("LBLOCK Export not implemented")
}

func ExitLBlockParser(vm *VM, args ...interface{}) {
	vm.Info("exit LBLOCK %v", args)
}

func ExitLBlockEval(vm *VM, args ...interface{}) (*Elem, error) {
	if len(args) < 1 {
		return nil, fmt.Errorf("There is no information from the LBLOCK header")
	}
	lbmode := args[0].(bool)
	vm.Debug("EXITING %v Block", lbmode)
	if !vm.MustIgnore() {
		if vm.CanGet() {
			vm.EndNS()
		}
	}
	return nil, nil
}

func ExitLBlockLambda(vm *VM, args ...interface{}) (*Elem, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("There is no information from the LBLOCK header")
	}
	vm.Debug("LBLOCK: %v", args)
	blockname := args[0].(string)
	lbmode := args[1].(bool)
	vm.Debug("exit LBLOCK(in lambda) %v", args)
	if lbmode == true {
		return &Elem{Type: "exitTRUEBLOCK", Value: blockname}, nil
	} else {
		return &Elem{Type: "exitFALSEBLOCK", Value: blockname}, nil
	}
}

func ExitLBlockImport(vm *VM, args ...interface{}) {
	vm.Fatal("exit LBLOCK Import not implemented")
}

func ExitLBlockExport(vm *VM, args ...interface{}) {
	vm.Fatal("exit LBLOCK Export not implemented")
}

func InitOpcodeLBlock(vm *VM) {
	vm.RegisterOpcode("LBLOCK", LBlockParser, LBlockLambda, LBlockEval, LBlockImport, LBlockExport)
	vm.RegisterOpcode("exitLBLOCK", ExitLBlockParser, ExitLBlockLambda, ExitLBlockEval, ExitLBlockImport, ExitLBlockExport)
}
