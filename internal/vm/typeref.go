package vm

import (
	"fmt"
)

func RefFactory(vm *VM) *Elem {
	return &Elem{Type: "REF", Value: nil}
}

func RefToString(vm *VM, e *Elem) string {
	if e.Type == "REF" {
		if e.Value == nil {
			return "#REF(NOWHERE)"
		}
		switch v := e.Value.(type) {
		case *Elem:
			eh, err := vm.GetType(v.Type)
			vm.OnError(err, "REF to string Error")
			return fmt.Sprintf("#REF(%v)=%v", v.Type, eh.ToString(vm, v))
		}
		return "#REF(UNKNOWN)"
	}
	vm.Error("trying to convert a REF and it is not a REF: %v %T", e.Type, e.Value)
	return "<?>"
}

func RefFromString(vm *VM, d string) *Elem {
	eh, err := vm.GetType(d)
	vm.OnError(err, "Error in creating REF object")
	return &Elem{Type: "REF", Value: eh.Factory(vm)}
}

func RefCompare(vm *VM, e1 *Elem, e2 *Elem) int {
	if e1.Type == "REF" && e2.Type == "REF" {
		if e1.Value != nil && e2.Value != nil {
			switch v1 := e1.Value.(type) {
			case *Elem:
				switch v2 := e2.Value.(type) {
				case *Elem:
					if v1.Type == v2.Type {
						eh, err := vm.GetType(v1.Type)
						vm.OnError(err, "Error in REF compare")
						return eh.Compare(vm, v1, v2)
					}
				}
			}
		}
	}
	return IDK
}

func RefDup(vm *VM, e *Elem) *Elem {
	if e.Type != "REF" {
		return nil
	}
	switch v := e.Value.(type) {
	case *Elem:
		eh, err := vm.GetType(v.Type)
		vm.OnError(err, "Error in REF DUP")
		res := eh.Duplicate(vm, e.Value.(*Elem))
		return &Elem{Type: "REF", Value: res}
	}
	return &Elem{Type: "REF", Value: nil}
}

func RegisterRef(vm *VM) {
	vm.RegisterType("REF", RefFactory, RefToString, RefFromString, RefCompare, RefDup)
}
