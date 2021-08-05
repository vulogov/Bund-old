package stdlib

import (
	vmmod "github.com/vulogov/Bund/internal/vm"
)

func InitFUNCTIONS(vm *vmmod.VM) {
	vm.Debug("[ BUND ] bund.InitFUNCTIONS() reached")
	InitPrintFunctions(vm)
	InitSystemFunctions(vm)
	InitReturnFunctions(vm)
	InitConsoleFunctions(vm)
	InitMatFunctions(vm)
	InitOpMath(vm)
	InitOpCmp(vm)
	InitGPMOperators(vm)
	InitOpSys(vm)
	InitOpZip(vm)
}
