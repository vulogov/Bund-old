package stdlib

import (
	"fmt"

	vmmod "github.com/vulogov/Bund/internal/vm"
)

func PrintElement(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	eh, err := vm.GetType(e.Type)
	if err != nil {
		vm.OnError(err, "Error in println functor")
		return nil, err
	}
	fmt.Printf("%s", eh.ToString(vm, e))
	return e, nil
}

func PrintlnElement(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	eh, err := vm.GetType(e.Type)
	if err != nil {
		vm.OnError(err, "Error in println functor")
		return nil, err
	}
	fmt.Printf("%s\n", eh.ToString(vm, e))
	return e, nil
}

func InitPrintFunctions(vm *vmmod.VM) {
	vm.Debug("[ BUND ] bund.InitPrintFunctions() reached")
	vm.AddFunction("print", PrintElement)
	vm.AddFunction("println", PrintlnElement)
}
