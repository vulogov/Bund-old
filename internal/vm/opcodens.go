package vm

func NSParser(vm *VM, args ...interface{}) {
	vm.Info("NAMESPACE %v", args)
}

func NSEval(vm *VM, args ...interface{}) {
	vm.Debug("NAMESPACE")
}

func NSLambda(vm *VM, args ...interface{}) {
	vm.Debug("NAMESPACE(in lambda)")
}

func NSImport(vm *VM, args ...interface{}) {
	vm.Fatal("NS Import not implemented")
}

func NSExport(vm *VM, args ...interface{}) {
	vm.Fatal("NS Export not implemented")
}

func InitOpcodeNs(vm *VM) {
	vm.RegisterOpcode("NS", NSParser, NSLambda, NSEval, NSImport, NSExport)
}
