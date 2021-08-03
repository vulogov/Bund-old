package stdlib

import (
	vmmod "github.com/vulogov/Bund/internal/vm"
)

func SwapOperator(vm *vmmod.VM, e1 *vmmod.Elem, e2 *vmmod.Elem) (*vmmod.Elem, error) {
	vm.Put(e1)
	return e2, nil
}

func InitOpSys(vm *vmmod.VM) {
	vm.Debug("[ BUND ] bund.InitOpSys() reached")
	vm.AddOperator("<=>", SwapOperator)
}
