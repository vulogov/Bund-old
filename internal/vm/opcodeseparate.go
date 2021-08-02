package vm

func SeparateParser(vm *VM, args ...interface{}) {
	vm.Info("SEPARATE %v", args)
}

func SeparateEval(vm *VM, args ...interface{}) (*Elem, error) {
	var val *Elem
	if vm.CurrentNS.GetOption("separateinline", false).(bool) {
		vm.Debug("SEPARATE inline")
		val = &Elem{Type: "SEPARATE", Value: nil}
	} else {
		vm.Debug("SEPARATE command")
		val = vm.Separate()
	}
	return val, nil
}

func SeparateLambda(vm *VM, args ...interface{}) (*Elem, error) {
	return &Elem{Type: "SEPARATE", Value: args[0].(string)}, nil
}

func SeparateImport(vm *VM, args ...interface{}) {
	vm.Fatal("SEPARATE Import not implemented")
}

func SeparateExport(vm *VM, args ...interface{}) {
	vm.Fatal("SEPARATE Export not implemented")
}

func InitOpcodeSeparate(vm *VM) {
	vm.RegisterOpcode("SEPARATE", SeparateParser, SeparateLambda, SeparateEval, SeparateImport, SeparateExport)
}
