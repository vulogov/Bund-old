package vm

import (
	"fmt"
)

func (vm *VM) Op(name string) (*Elem, error) {
	var val1, val2 *Elem
	if vm.HasUserFunction(name) {
		err := vm.Apply(name)
		return nil, err
	}
	if vm.Current.Len() < 1 {
		return nil, fmt.Errorf("Attempt to call an operation: %v on empty stack", name)
	}
	f, err := vm.GetOperator(name)
	if err != nil {
		return nil, err
	}

	if vm.Current.Len() >= 1 {
		val1 = vm.Take()
	} else {
		return nil, fmt.Errorf("Attempt to call an operation: %v on empty stack", name)
	}
	if vm.Current.Len() >= 1 {
		val2 = vm.Take()
	} else {
		val2 = &Elem{Type: "NONE"}
	}
	res, err := f(vm, val1, val2)
	if err != nil {
		return nil, err
	}
	return res, nil
}
