package vm

func BlockParser(vm *VM, args ...interface{}) {
	vm.Info("BLOCK %v", args)
}

func BlockEval(vm *VM, args ...interface{}) {
	vm.Debug("BLOCK")
}

func BlockLambda(vm *VM, args ...interface{}) {
	vm.Debug("BLOCK(in lambda)")
}

func BlockImport(vm *VM, args ...interface{}) {
	vm.Fatal("BLOCK Import not implemented")
}

func BlockExport(vm *VM, args ...interface{}) {
	vm.Fatal("BLOCK Export not implemented")
}

func ExitBlockParser(vm *VM, args ...interface{}) {
	vm.Info("exit BLOCK %v", args)
}

func ExitBlockEval(vm *VM, args ...interface{}) {
	vm.Debug("exit BLOCK")
}

func ExitBlockLambda(vm *VM, args ...interface{}) {
	vm.Debug("exit BLOCK(in lambda)")
}

func ExitBlockImport(vm *VM, args ...interface{}) {
	vm.Fatal("exit BLOCK Import not implemented")
}

func ExitBlockExport(vm *VM, args ...interface{}) {
	vm.Fatal("exit BLOCK Export not implemented")
}

func InitOpcodeBlock(vm *VM) {
	vm.RegisterOpcode("BLOCK", BlockParser, BlockLambda, BlockEval, BlockImport, BlockExport)
	vm.RegisterOpcode("exitBLOCK", ExitBlockParser, ExitBlockLambda, ExitBlockEval, ExitBlockImport, ExitBlockExport)
}
