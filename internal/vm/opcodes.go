package vm

type OpcodeParserFun func(v *VM, args ...interface{})
type OpcodeInLambdaFun func(v *VM, args ...interface{}) (*Elem, error)
type OpcodeEvalFun func(v *VM, args ...interface{}) (*Elem, error)
type OpcodeExportFun func(v *VM, args ...interface{})
type OpcodeImportFun func(v *VM, args ...interface{})

type OPCODE struct {
	Name     string
	InParser OpcodeParserFun
	InLambda OpcodeInLambdaFun
	InEval   OpcodeEvalFun
	InEx     OpcodeExportFun
	InIm     OpcodeImportFun
}

func (vm *VM) RegisterOpcodes() {
	vm.Debug("Registering VM opcodes")
	InitOpcodeNs(vm)
	InitOpcodeExitNs(vm)
	InitOpcodeBlock(vm)
	InitOpcodeBoolean(vm)
}

func (vm *VM) RegisterOpcode(t string, ip OpcodeParserFun, il OpcodeInLambdaFun, ef OpcodeEvalFun, ex OpcodeExportFun, im OpcodeImportFun) bool {
	if _, ok := vm.Opcodes.Load(t); ok {
		return true
	}
	vm.Opcodes.Store(t, &OPCODE{Name: t, InParser: ip, InLambda: il, InEval: ef, InIm: im, InEx: ex})
	vm.Debug("Register BUND opcode: %v", t)
	return true
}

func (vm *VM) Opcode(t string) *OPCODE {
	if res, ok := vm.Opcodes.Load(t); ok {
		return res.(*OPCODE)
	}
	vm.Fatal("BUND do not have opcode: %v", t)
	return nil
}
