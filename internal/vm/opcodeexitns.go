package vm

func ExitNSParser(vm *VM, args ...interface{}) {
	vm.Info("exit NAMESPACE %v", args)
}

func ExitNSEval(vm *VM, args ...interface{}) {
	vm.Debug("exit NAMESPACE")
}

func ExitNSLambda(vm *VM, args ...interface{}) {
	vm.Debug("exit NAMESPACE(in lambda)")
}

func ExitNSImport(vm *VM, args ...interface{}) {
	vm.Fatal("ExitNS Import not implemented")
}

func ExitNSExport(vm *VM, args ...interface{}) {
	vm.Fatal("ExitNS Export not implemented")
}

func InitOpcodeExitNs(vm *VM) {
	vm.RegisterOpcode("exitNS", ExitNSParser, ExitNSLambda, ExitNSEval, ExitNSImport, ExitNSExport)
}
