package vm

import (
	"fmt"
)

func (vm *VM) Exec(name string) (*Elem, error) {
	if vm.HasUserFunction(name) {
		err := vm.Apply(name)
		return nil, err
	}
	if !vm.CanGet() {
		return nil, fmt.Errorf("Attempt to call a function: %v on empty stack", name)
	}
	f, err := vm.GetFunction(name)
	if err != nil {
		return nil, err
	}
	if f == nil {
		return nil, fmt.Errorf("Attempt to locate function: %v failed", name)
	}
	val := vm.Take()
	res, err := f(vm, val)
	return res, err
}