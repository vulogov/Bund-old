package stdlib

import (
	"fmt"

	vmmod "github.com/vulogov/Bund/internal/vm"
)

func runcmp(vm *vmmod.VM, e1 *vmmod.Elem, e2 *vmmod.Elem, cmpf func(*vmmod.VM, int) bool) (*vmmod.Elem, error) {
	if e1.Type != e2.Type {
		return nil, fmt.Errorf("Datatypes for operation '=' must be the same")
	}
	eh, err := vm.GetType(e1.Type)
	if err != nil {
		return nil, err
	}
	res := eh.Compare(vm, e1, e2)
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

func EqOperator(vm *vmmod.VM, e1 *vmmod.Elem, e2 *vmmod.Elem) (*vmmod.Elem, error) {
	return runcmp(vm, e1, e2, iseq)
}

func NEqOperator(vm *vmmod.VM, e1 *vmmod.Elem, e2 *vmmod.Elem) (*vmmod.Elem, error) {
	return runcmp(vm, e1, e2, isneq)
}

func InitOpCmp(vm *vmmod.VM) {
	vm.Debug("[ BUND ] bund.InitOpCmp() reached")
	vm.AddOperator("=", EqOperator)
	vm.AddOperator("<>", NEqOperator)
}
