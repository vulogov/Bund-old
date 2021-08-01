package vm

func (vm *VM) GetNS(name string) *NS {
	var res *NS
	if _res, ok := vm.NS.Load(name); ok {
		vm.Debug("Returning NAMESPACE: %v", name)
		res = _res.(*NS)
	} else {
		res = NewNS(vm, name)
		vm.NS.Store(name, res)
	}
	if vm.CurrentNS != nil {
		vm.NSStack.PushBack(vm.CurrentNS)
	} else {
		vm.Debug("Setting a Root Namespace")
		vm.RootNS = res
	}
	vm.CurrentNS = res
	vm.Current = &res.Stack
	return res
}

func (vm *VM) EndNS() *NS {
	var res *NS
	if vm.NSStack.Len() > 0 {
		_res := vm.NSStack.PopBack()
		if _res != nil {
			res = _res.(*NS)
			vm.CurrentNS = res
			vm.Current = &res.Stack
		} else {
			vm.Debug("NAMESPACE stack returns End_of_Stack")
			res = nil
			vm.CurrentNS = nil
			vm.Current = nil
		}
		vm.Debug("NAMESPACE stack %v size: %v", vm.CurrentNS.Name, vm.NSStack.Len())
	} else {
		vm.Debug("NAMESPACE stack is empty")
		res = nil
		vm.CurrentNS = nil
		vm.Current = nil
	}
	return res
}

func (vm *VM) DelNS(name string) {
	if _, ok := vm.NS.Load(name); ok {
		vm.Debug("Found NAMESPACE: %v", name)
		vm.NS.Delete(name)
		if _, ok := vm.NS.Load(name); !ok {
			vm.Debug("NAMESPACE removed: %v", name)
		}
	}
}
