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
	oh := vm.Opcode(cmd.Type)
	if oh != nil {
		oh.InEval(vm, cmd.Value.(string))
	} else {
		vm.Fatal("Unknown OPCODE: %v", cmd.Type)
	}
}

func ApplyInNS(name string, ns *NS, vm *VM) error {
	var cmd *Elem
	if !vm.IsStack() {
		return fmt.Errorf("Attempt to execute lambda %v on empty context", name)
	}
	if !ns.InLCache(name) {
		vm.Debug("Adding lambda %v to lcache", name)
		ls := ns.GetLambda(name)
		if ls == nil {
			return fmt.Errorf("Lambda %v not exist in %v", name, vm.Name)
		}
		for i := 0; i < ls.Len(); i++ {
			cmd = ls.At(i).(*Elem)
			ns.AddToLCache(name, cmd)
		}
	} else {
		vm.Debug("Lambda %v already in lcache", name)
	}
	cl := ns.LCacheGet(name)
	if cl == nil {
		return fmt.Errorf("Lambda %v is missed in LCache", name)
	}
	for i := 0; i < len(cl); i++ {
		cmd = &cl[i]
		switch cmd.Type {
		case "TRUEBLOCK", "FALSEBLOCK":
			EvalCmd(vm, cmd)
			continue
		case "exitTRUEBLOCK", "exitFALSEBLOCK":
			EvalCmd(vm, cmd)
			continue
		}
		if vm.CheckIgnore() {
			continue
		}
		switch cmd.Type {
		case "int", "flt", "str", "bool", "cpx", "glob":
			vm.Debug("DATA: %v", cmd)
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
		case "FUNCTION", "exitFUNCTION":
			EvalCmd(vm, cmd)
		case "UNIXCMD":
			EvalCmd(vm, cmd)
		default:
			vm.Error("Unknown command in APPLY: %v", cmd)
		}
		vm.DbgFun(i, cmd)
	}
	ns.LCacheDel(name)
	return nil
}

func Apply(name string, vm *VM) error {
	var cmd *Elem
	if !vm.IsStack() {
		return fmt.Errorf("Attempt to execute lambda %v on empty context", name)
	}
	if !vm.CurrentNS.InLCache(name) {
		vm.Debug("Adding lambda %v to lcache", name)
		ls := vm.CurrentNS.GetLambda(name)
		if ls == nil {
			return fmt.Errorf("Lambda %v not exist in %v", name, vm.Name)
		}
		for i := 0; i < ls.Len(); i++ {
			cmd = ls.At(i).(*Elem)
			vm.CurrentNS.AddToLCache(name, cmd)
		}
	} else {
		vm.Debug("Lambda %v already in lcache", name)
	}
	cl := vm.CurrentNS.LCacheGet(name)
	if cl == nil {
		return fmt.Errorf("Lambda %v is missed in LCache", name)
	}
	for i := 0; i < len(cl); i++ {
		cmd = &cl[i]
		switch cmd.Type {
		case "TRUEBLOCK", "FALSEBLOCK":
			EvalCmd(vm, cmd)
			continue
		case "exitTRUEBLOCK", "exitFALSEBLOCK":
			EvalCmd(vm, cmd)
			continue
		}
		if vm.CheckIgnore() {
			continue
		}
		switch cmd.Type {
		case "int", "flt", "str", "bool", "cpx", "glob":
			vm.Debug("DATA: %v", cmd)
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
		case "FUNCTION", "exitFUNCTION":
			EvalCmd(vm, cmd)
		default:
			vm.Error("Unknown command in APPLY: %v", cmd)
		}
		vm.DbgFun(i, cmd)
	}
	vm.CurrentNS.LCacheDel(name)
	return nil
}

func (vm *VM) Apply(name string) error {
	err := Apply(name, vm)
	if err != nil {
		vm.OnError(err, "Error in Apply(%v)", name)
	}
	return err
}
