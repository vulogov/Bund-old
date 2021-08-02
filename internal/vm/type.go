package vm

func (vm *VM) RegisterTypes() {
	vm.Debug("BUND: registering dynamic types")
	RegisterBoolean(vm)
	RegisterInt(vm)
	RegisterFloat(vm)
	RegisterString(vm)
	RegisterCpx(vm)
	RegisterCALL(vm)
}
