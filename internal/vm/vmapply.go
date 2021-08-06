package vm

import (
	"fmt"
)

func EvalCmd(vm *VM, cmd *Elem) {
	oh := vm.Opcode(cmd.Type)
	if oh != nil {
		if cmd.Value != nil {
			res, err := oh.InEval(vm, cmd.Value.(string))
			vm.OnError(err, "Error in EVAL")
			if res != nil {
				vm.Put(res)
			}
		} else {
			res, err := oh.InEval(vm)
			vm.OnError(err, "Error in EVAL")
			if res != nil {
				vm.Put(res)
			}
		}
	} else {
		vm.Fatal("Unknown OPCODE: %v", cmd.Type)
	}
}

func SimpleEvalCmd(vm *VM, cmd *Elem) {
	oh := vm.Opcode("MODE")
	if oh != nil {
		oh.InEval(vm, cmd.Value.(string))
	} else {
		vm.Fatal("Unknown OPCODE: %v", cmd.Type)
	}
}

func Apply(name string, vm *VM) error {
	if !vm.IsStack() {
		return fmt.Errorf("Attempt to execute lambda %v on empty context", name)
	}
	ls := vm.CurrentNS.GetLambda(name)
	if ls == nil {
		return fmt.Errorf("Lambda %v not exist in %v", name, vm.Name)
	}
	vm.Debug("Executing lambda %v in %v", name, vm.Name)
	for ls.Len() > 0 {
		cmd := ls.PopFront().(*Elem)
		fmt.Println(cmd)
		switch cmd.Type {
		case "int", "flt", "str", "bool", "cpx", "glob":
			vm.Put(cmd)
		case "NS", "exitNS":
			EvalCmd(vm, cmd)
		case "BLOCK", "exitBLOCK":
			EvalCmd(vm, cmd)
		case "DBLOCK", "exitDBLOCK":
			EvalCmd(vm, cmd)
		case "CALL":
			EvalCmd(vm, cmd)
		case "OP":
			EvalCmd(vm, cmd)
		case "MODE":
			SimpleEvalCmd(vm, cmd)
		case "SEPARATE":
			EvalCmd(vm, cmd)
		case "MBLOCK", "exitMBLOCK":
			EvalCmd(vm, cmd)
		case "TRUEBLOCK", "FALSEBLOCK":
			EvalCmd(vm, cmd)
		case "exitTRUEBLOCK", "exitFALSEBLOCK":
			EvalCmd(vm, cmd)
		case "FUNCTION", "exitFUNCTION":
			EvalCmd(vm, cmd)
		default:
			vm.Error("Unknown command in APPLY: %v", cmd)
		}
	}
	return nil
}

func (vm *VM) Apply(name string) error {
	err := Apply(name, vm)
	if err != nil {
		vm.OnError(err, "Error in Apply(%v)", name)
	}
	return err
}
