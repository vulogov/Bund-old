package stdlib

import (
	"fmt"

	vmmod "github.com/vulogov/Bund/internal/vm"
)

func GlobFunctor(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	vm.Debug("GLOB FUNCTOR: %v", e.Type)
	if e.Type == "str" {
		eh, err := vm.GetType("glob")
		vm.OnError(err, "Error in glob functor")
		return eh.FromString(vm, e.Value.(string)), nil
	}
	return e, nil
}

func FileFunctor(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	vm.Debug("FILE FUNCTOR: %v", e.Type)
	if e.Type == "str" {
		eh, err := vm.GetType("file")
		vm.OnError(err, "Error in file functor")
		return eh.FromString(vm, e.Value.(string)), nil
	}
	return e, nil
}

func JsonFunctor(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	vm.Debug("JSON FUNCTOR: %v", e.Type)
	if e.Type == "str" {
		eh, err := vm.GetType("json")
		vm.OnError(err, "Error in json functor")
		return eh.FromString(vm, e.Value.(string)), nil
	}
	return e, nil
}

func UnixFunctor(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	vm.Debug("UNIXCMD FUNCTOR: %v", e.Type)
	if e.Type == "str" {
		oc := vm.Opcode("unixcmd")
		if oc == nil {
			return nil, fmt.Errorf("There is no opcode unixcmd and I do not know why.")
		}
		return oc.InEval(vm, e.Value.(string))
	}
	return e, nil
}

func InitFunctorFunctions(vm *vmmod.VM) {
	vm.Debug("[ BUND ] bund.InitFunctorFunctions() reached")
	vm.AddFunction("glob", GlobFunctor)
	vm.AddFunction("file", FileFunctor)
	vm.AddFunction("unix", UnixFunctor)
	vm.AddFunction("json", JsonFunctor)
}
