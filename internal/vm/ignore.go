package vm

func (vm *VM) Ignore() {
	vm.IsIgnore.PushBack(true)
}

func (vm *VM) NotIgnore() {
	vm.IsIgnore.PushBack(false)
}

func (vm *VM) MustIgnore() bool {
	if vm.IsIgnore.Len() == 0 {
		return false
	}
	res := vm.IsIgnore.PopBack()
	return res.(bool)
}

func (vm *VM) CheckIgnore() bool {
	if vm.IsIgnore.Len() == 0 {
		return false
	}
	res := vm.IsIgnore.Back()
	return res.(bool)
}
