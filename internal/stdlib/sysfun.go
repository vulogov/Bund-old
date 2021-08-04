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

func SetAlias(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	if vm.Current.Len() < 1 {
		return nil, fmt.Errorf("Namespace stack is too shallow for ALIAS operation: %v:%v", vm.CurrentNS.Name, vm.Current.Len())
	}
	if e.Type != "str" {
		return nil, fmt.Errorf("String value is required for an ALIAS operation: %v", e.Type)
	}
	e2 := vm.Take()
	vm.CurrentNS.SetAlias(e.Value.(string), e2)
	return nil, nil
}

func GetAlias(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	if e.Type != "str" {
		return nil, fmt.Errorf("String value is required for an ALIAS operation: %v", e.Type)
	}
	val := vm.CurrentNS.GetAlias(e.Value.(string))
	if val == nil {
		vm.Error("There is no alias for: %v", e.Value.(string))
		return nil, nil
	}
	return val.(*vmmod.Elem), nil
}

func RefElement(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	if e.Type == "REF" {
		switch v := e.Value.(type) {
		case *vmmod.Elem:
			vm.Debug("Dereferencing %v", v.Type)
			return v, nil
		}
		return nil, fmt.Errorf("REF: Attempt of de-referencing the data, but there is no data there")
	} else {
		vm.Debug("Making REF for %v", e.Type)
		return &vmmod.Elem{Type: "REF", Value: e}, nil
	}
}

func InitSystemFunctions(vm *vmmod.VM) {
	vm.Debug("[ BUND ] bund.InitSystemFunctions() reached")
	vm.AddFunction("passthrough", PassthrougElement)
	vm.AddFunction(",", DropElement)
	vm.AddFunction("^", DupElement)
	vm.AddFunction("!", ExecuteElement)
	vm.AddFunction("setAlias", SetAlias)
	vm.AddFunction("alias", GetAlias)
	vm.AddFunction("#", RefElement)
}
