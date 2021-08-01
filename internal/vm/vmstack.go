package vm

func (vm *VM) Put(e *Elem) bool {
	if !vm.IsStack() {
		vm.Fatal("Attempt to Put() but Stack doesn't exists")
		return false
	}
	if vm.Mode {
		vm.Current.PushBack(e)
	} else {
		vm.Current.PushFront(e)
	}
	return true
}

func (vm *VM) Get() *Elem {
	var res interface{}

	if !vm.CanGet() {
		vm.Error("Attempt to Get() but Stack doesn't exists or empty")
		return nil
	}
	if vm.Mode {
		res = vm.Current.Back()
	} else {
		res = vm.Current.Front()
	}
	return res.(*Elem)
}

func (vm *VM) Take() *Elem {
	var res interface{}

	if !vm.CanGet() {
		vm.Error("Attempt to Take() but Stack doesn't exists or empty")
		return nil
	}
	if vm.Mode {
		res = vm.Current.PopBack()
	} else {
		res = vm.Current.PopFront()
	}
	return res.(*Elem)
}
