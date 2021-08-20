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
	InitRndFunctions(vm)
	InitMathFunctions(vm)
	InitLoopFunctions(vm)
	InitPctFunctions(vm)
	InitVFSFunctions(vm)
	InitStringFunctions(vm)
	InitFunctorFunctions(vm)
	InitJsonFunctions(vm)
	InitOpMath(vm)
	InitOpCmp(vm)
	InitGPMOperators(vm)
	InitOpSys(vm)
	InitOpZip(vm)
	InitOpPercent(vm)
}
