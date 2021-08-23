package stdlib

import (
	"fmt"

	"github.com/gammazero/deque"

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

func AllcallFunctor(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	vm.Debug("Calling '∀' CALL for the stack")
	q := new(deque.Deque)
	oh := vm.Opcode(e.Type)
	if oh == nil {
		return nil, fmt.Errorf("No proper opcode for '∀' functor")
	}
	for vm.Current.Len() > 0 {
		if e.Value != nil {
			res, err := oh.InEval(vm, e.Value.(string))
			vm.OnError(err, "Error in EVAL during inside '∀' functor")
			q.PushBack(res)
		}
	}
	for q.Len() > 0 {
		v := q.PopFront()
		vm.Put(v.(*vmmod.Elem))
	}
	return nil, nil
}

func AllopFunctor(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	vm.Debug("Calling '∀' OP for the stack")
	for vm.Current.Len() > 1 {
		vmmod.EvalCmd(vm, e)
	}
	return nil, nil
}

func AllFunctor(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	switch e.Type {
	case "CALL":
		return AllcallFunctor(vm, e)
	case "OP":
		return AllopFunctor(vm, e)
	}
	vm.Warning("Unknown datatype for functor '∀': %v", e.Type)
	return e, nil
}

func InitFunctorFunctions(vm *vmmod.VM) {
	vm.Debug("[ BUND ] bund.InitFunctorFunctions() reached")
	vm.AddFunction("glob", GlobFunctor)
	vm.AddFunction("file", FileFunctor)
	vm.AddFunction("unix", UnixFunctor)
	vm.AddFunction("json", JsonFunctor)
	vm.AddFunction("all", AllFunctor)
	vm.AddFunction("∀", AllFunctor)
}
