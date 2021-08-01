package stdlib

import (
	vmmod "github.com/vulogov/Bund/internal/vm"
)

func PassthrougElement(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	vm.Debug("PASSTHROUGH: %v", e.Type)
	return e, nil
}

func InitSystemFunctions(vm *vmmod.VM) {
	vm.Debug("[ BUND ] bund.InitSystemFunctions() reached")
	vm.AddFunction("passthrough", PassthrougElement)
}
