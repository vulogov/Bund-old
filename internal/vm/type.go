package vm

func (vm *VM) RegisterTypes() {
	vm.Debug("BUND: registering dynamic types")
	RegisterBoolean(vm)
}
