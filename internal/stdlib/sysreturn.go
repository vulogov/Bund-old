package stdlib

import (
	"fmt"

	vmmod "github.com/vulogov/Bund/internal/vm"
)

func ReturnElement(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	if vm.NSStack.Len() < 1 {
		return nil, fmt.Errorf("Namespace stack is too shallow for RETURN operation: %v", vm.NSStack.Len())
	}
	nsr := vm.NSStack.Back().(*vmmod.NS)
	vm.Debug("RETURN %v to %v", e.Type, nsr.Name)
	if vm.Mode {
		nsr.Stack.PushBack(e)
	} else {
		nsr.Stack.PushFront(e)
	}
	return nil, nil
}

func ReturnElementAndKeepOriginal(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	if vm.NSStack.Len() < 1 {
		return nil, fmt.Errorf("Namespace stack is too shallow for RETURN operation: %v", vm.NSStack.Len())
	}
	nsr := vm.NSStack.Back().(*vmmod.NS)
	vm.Debug("RETURN %v to %v", e.Type, nsr.Name)
	if vm.Mode {
		nsr.Stack.PushBack(e)
	} else {
		nsr.Stack.PushFront(e)
	}
	vm.Debug("Keeing %v in current context", e.Type)
	return e, nil
}

func ReturnElementTo(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	if vm.Current.Len() < 1 {
		return nil, fmt.Errorf("Namespace stack is too shallow for RETURNTO operation: %v:%v", vm.CurrentNS.Name, vm.Current.Len())
	}
	e2 := vm.Take()
	if e2.Type != "str" {
		return nil, fmt.Errorf("We did not find string for RETURNTO operation: %v", e2.Type)
	}
	vm.GetNS(e2.Value.(string))
	vm.Put(e)
	vm.EndNS()
	return nil, nil
}

func ReturnElementFrom(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	if e.Type != "str" {
		return nil, fmt.Errorf("We did not find string for RETURNFROM operation: %v", e.Type)
	}
	vm.GetNS(e.Value.(string))
	if vm.Current.Len() < 1 {
		return nil, fmt.Errorf("Namespace stack is too shallow for RETURNFROM operation: %v:%v", vm.CurrentNS.Name, vm.Current.Len())
	}
	e2 := vm.Take()
	vm.Debug("Taking %v from %v %v elemets left", e2.Type, vm.CurrentNS.Name, vm.Current.Len())
	vm.EndNS()
	return e2, nil
}

func ReturnElementFromAndKeepOriginal(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	if e.Type != "str" {
		return nil, fmt.Errorf("We did not find string for RETURNFROM operation: %v", e.Type)
	}
	vm.GetNS(e.Value.(string))
	if vm.Current.Len() < 1 {
		return nil, fmt.Errorf("Namespace stack is too shallow for RETURNFROM operation: %v:%v", vm.CurrentNS.Name, vm.Current.Len())
	}
	e2 := vm.Get()
	vm.Debug("Getting %v from %v %v elemets left", e2.Type, vm.CurrentNS.Name, vm.Current.Len())
	vm.EndNS()
	return e2, nil
}

func InitReturnFunctions(vm *vmmod.VM) {
	vm.Debug("[ BUND ] bund.InitReturnFunctions() reached")
	vm.AddFunction("$", ReturnElement)
	vm.AddFunction("$$", ReturnElementAndKeepOriginal)
	vm.AddFunction("$_", ReturnElementTo)
	vm.AddFunction("returnTo", ReturnElementTo)
	vm.AddFunction("_$", ReturnElementFrom)
	vm.AddFunction("returnFrom", ReturnElementFrom)
	vm.AddFunction("_$$", ReturnElementFromAndKeepOriginal)
}
