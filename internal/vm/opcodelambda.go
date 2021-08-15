package vm

func LambdaParser(vm *VM, args ...interface{}) {
	vm.Info("LAMBDA %v", args)
}

func LambdaImport(vm *VM, args ...interface{}) {
	vm.Fatal("LAMBDA Import not implemented")
}

func LambdaExport(vm *VM, args ...interface{}) {
	vm.Fatal("LAMBDA Export not implemented")
}

func ExitLambdaParser(vm *VM, args ...interface{}) {
	vm.Info("exit LAMBDA %v", args)
}

func ExitLambdaEval(vm *VM, args ...interface{}) (*Elem, error) {
	funcname := vm.CurrentNS.NameOfCurrentLambda()
	x, y := ExitFunctionEval(vm, funcname)
	vm.Put(&Elem{Type: "CALL", Value: funcname})
	return x, y
}

func ExitLambdaLambda(vm *VM, args ...interface{}) (*Elem, error) {
	vm.Debug("exit FUNCTION(in lambda) %v", args)
	return &Elem{Type: "exitFUNCTION", Value: nil}, nil
}

func ExitLambdaImport(vm *VM, args ...interface{}) {
	vm.Fatal("exit LAMBDA Import not implemented")
}

func ExitLambdaExport(vm *VM, args ...interface{}) {
	vm.Fatal("exit LAMBDA Export not implemented")
}

func InitOpcodeLambda(vm *VM) {
	vm.RegisterOpcode("LAMBDA", LambdaParser, FunctionLambda, FunctionEval, LambdaImport, LambdaExport)
	vm.RegisterOpcode("exitLAMBDA", ExitLambdaParser, ExitLambdaLambda, ExitLambdaEval, ExitLambdaImport, ExitLambdaExport)
}
