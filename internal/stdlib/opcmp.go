package stdlib

import (
	"fmt"

	vmmod "github.com/vulogov/Bund/internal/vm"
)

func runcmp(vm *vmmod.VM, e1 *vmmod.Elem, e2 *vmmod.Elem, cmpf func(*vmmod.VM, int) bool) (*vmmod.Elem, error) {
	var cmpv *vmmod.Elem
	cmpv = e2
	if e1.Type != e2.Type {
		if e1.Type == "glob" {
			eh, err := vm.GetType(e2.Type)
			if err != nil {
				vm.OnError(err, "Error in logical operator for datatype: %v", e2.Type)
			}
			cmpv = &vmmod.Elem{Type: "str", Value: eh.ToString(vm, e2)}
		} else {
			return nil, fmt.Errorf("Datatypes for the glob operations are not valid: %v %v", e1.Type, e2.Type)
		}
	}
	eh, err := vm.GetType(e1.Type)
	if err != nil {
		return nil, err
	}
	res := eh.Compare(vm, e1, cmpv)
	return &vmmod.Elem{Type: "bool", Value: cmpf(vm, res)}, nil
}

func iseq(vm *vmmod.VM, res int) bool {
	if res == vmmod.Eq {
		return true
	} else {
		return false
	}
}

func isneq(vm *vmmod.VM, res int) bool {
	vm.Debug("CMP result is %v", res)
	if res != vmmod.Eq {
		return true
	} else {
		return false
	}
}

func isls(vm *vmmod.VM, res int) bool {
	vm.Debug("CMP result is %v", res)
	if res == vmmod.Ls {
		return true
	} else {
		return false
	}
}

func isgt(vm *vmmod.VM, res int) bool {
	vm.Debug("CMP result is %v", res)
	if res == vmmod.Gt {
		return true
	} else {
		return false
	}
}

func EqOperator(vm *vmmod.VM, e1 *vmmod.Elem, e2 *vmmod.Elem) (*vmmod.Elem, error) {
	return runcmp(vm, e1, e2, iseq)
}

func NEqOperator(vm *vmmod.VM, e1 *vmmod.Elem, e2 *vmmod.Elem) (*vmmod.Elem, error) {
	return runcmp(vm, e1, e2, isneq)
}

func LSqOperator(vm *vmmod.VM, e1 *vmmod.Elem, e2 *vmmod.Elem) (*vmmod.Elem, error) {
	return runcmp(vm, e1, e2, isls)
}

func GTqOperator(vm *vmmod.VM, e1 *vmmod.Elem, e2 *vmmod.Elem) (*vmmod.Elem, error) {
	return runcmp(vm, e1, e2, isgt)
}

func InitOpCmp(vm *vmmod.VM) {
	vm.Debug("[ BUND ] bund.InitOpCmp() reached")
	vm.AddOperator("=", EqOperator)
	vm.AddOperator("<>", NEqOperator)
	vm.AddOperator("<", GTqOperator)
	vm.AddOperator(">", LSqOperator)
}
