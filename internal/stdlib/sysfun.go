package stdlib

import (
	"fmt"

	vmmod "github.com/vulogov/Bund/internal/vm"
)

func PassthrougElement(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	vm.Debug("PASSTHROUGH: %v", e.Type)
	return e, nil
}

func DropElement(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	vm.Debug("DROP: %v", e.Type)
	return nil, nil
}

func DupElement(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	vm.Debug("DUP: %v", e.Type)
	eh, err := vm.GetType(e.Type)
	if err != nil {
		return nil, err
	}
	vm.Put(e)
	res := eh.Duplicate(vm, e)
	return res, nil
}

func ExecuteElement(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	vm.Debug("EXECUTE: %v", e.Type)
	if e.Type == "CALL" {
		res, err := vm.Exec(e.Value.(string))
		return res, err
	} else if e.Type == "OP" {
		res, err := vm.Op(e.Value.(string))
		return res, err
	}
	return nil, fmt.Errorf("Request to EXECUTE on wrong context: %v", e.Type)
}

func InitSystemFunctions(vm *vmmod.VM) {
	vm.Debug("[ BUND ] bund.InitSystemFunctions() reached")
	vm.AddFunction("passthrough", PassthrougElement)
	vm.AddFunction(",", DropElement)
	vm.AddFunction("^", DupElement)
	vm.AddFunction("!", ExecuteElement)
}
