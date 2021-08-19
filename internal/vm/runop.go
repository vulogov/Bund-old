package vm

func (vm *VM) runop(t string, args ...interface{}) {
	var res *Elem
	var err error

	if !vm.InLambda() {
		if len(args) == 2 && len(args[1].(string)) > 0 {
			functor := args[1].(string)
			res_eval, err := vm.Opcode(t).InEval(vm, args...)
			if err != nil {
				vm.OnError(err, "Error in function application before functor: %v", args[0])
			}
			fun, err := vm.GetFunction(functor)
			if err == nil {
				vm.Debug("Functor execution for %v(%v)", t, functor)
				res, err = fun(vm, res_eval)
				if err != nil {
					vm.OnError(err, "Error returned by functor: %v", functor)
				}
			} else {
				vm.OnError(err, "Error executing functor: %v", functor)
			}
		} else {
			res, err = vm.Opcode(t).InEval(vm, args...)
		}
		if err != nil {
			vm.OnError(err, "Error evaluating: %v", t)
		}
		if res != nil {
			vm.Put(res)
		}
	} else {
		res, err = vm.Opcode(t).InLambda(vm, args...)
		if err != nil {
			vm.OnError(err, "Error placing in lambda: %v", t)
		}
		if len(args) == 2 && len(args[1].(string)) > 0 {
			res.Functor = args[1].(string)
		}
		ls := vm.CurrentLambda()
		if ls != nil {
			vm.Debug("TO f(%v) %v", vm.CurrentNS.NameOfCurrentLambda(), res)
			ls.PushBack(res)
		} else {
			vm.Debug("Following OPCODE not been pushed to Lambda stack: %v", res.Type)
		}
	}
}

func (vm *VM) RunOp(t string, args ...interface{}) {
	if vm.CheckIgnore() {
		return
	}
	vm.runop(t, args...)
}

func (vm *VM) AlwaysRunOp(t string, args ...interface{}) {
	vm.runop(t, args...)
}

func (vm *VM) LambdaRunOp(t string, args ...interface{}) {
	res, err := vm.Opcode(t).InEval(vm, args...)
	if err != nil {
		vm.OnError(err, "Error evaluating: %v", t)
	}
	if res != nil {
		vm.Put(res)
	}
}
