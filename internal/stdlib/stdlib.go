package stdlib

import (
	vmmod "github.com/vulogov/Bund/internal/vm"
)

func InitFUNCTIONS(vm *vmmod.VM) {
	vm.Debug("[ BUND ] bund.InitFUNCTIONS() reached")
	InitPrintFunctions(vm)
	InitSystemFunctions(vm)
	InitReturnFunctions(vm)
	InitOpMath(vm)
}
