package vm

import (
	"fmt"
)

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
		case "int", "flt", "str", "bool", "cpx":
			vm.Put(cmd)
		case "DBLOCK":
			oh := vm.Opcode("DBLOCK")
			if oh != nil {
				res, err := oh.InEval(vm, cmd.Value.(string))
				vm.OnError(err, "Error in EVAL")
				if res != nil {
					vm.Put(res)
				}
			} else {
				vm.Fatal("Unknown OPCODE: %v", cmd.Type)
			}
		case "exitDBLOCK":
			oh := vm.Opcode("exitDBLOCK")
			if oh != nil {
				res, err := oh.InEval(vm)
				vm.OnError(err, "Error in EVAL")
				if res != nil {
					vm.Put(res)
				}
			} else {
				vm.Fatal("Unknown OPCODE: %v", cmd.Type)
			}
		case "CALL":
			oh := vm.Opcode("CALL")
			if oh != nil {
				res, err := oh.InEval(vm, cmd.Value.(string))
				vm.OnError(err, "Error in EVAL")
				if res != nil {
					vm.Put(res)
				}
			} else {
				vm.Fatal("Unknown OPCODE: %v", cmd.Type)
			}
		case "OP":
			oh := vm.Opcode("OP")
			if oh != nil {
				res, err := oh.InEval(vm, cmd.Value.(string))
				vm.OnError(err, "Error in EVAL")
				if res != nil {
					vm.Put(res)
				}
			} else {
				vm.Fatal("Unknown OPCODE: %v", cmd.Type)
			}
		case "MODE":
			oh := vm.Opcode("MODE")
			if oh != nil {
				oh.InEval(vm, cmd.Value.(string))
			} else {
				vm.Fatal("Unknown OPCODE: %v", cmd.Type)
			}
		case "SEPARATE":
			oh := vm.Opcode("SEPARATE")
			if oh != nil {
				res, err := oh.InEval(vm)
				vm.OnError(err, "Error in EVAL")
				if res != nil {
					vm.Put(res)
				}
			} else {
				vm.Fatal("Unknown OPCODE: %v", cmd.Type)
			}
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
