package vm

import (
	"fmt"
)

type BundFunction func(vm *VM, e *Elem) (*Elem, error)
type BundOperator func(vm *VM, e1 *Elem, e2 *Elem) (*Elem, error)
type BundSysfun func(vm *VM, n string) error

func (vm *VM) AddFunction(name string, fn BundFunction) bool {
	if _, ok := vm.Functions.Load(name); ok {
		vm.Warning("Function %v already registered", name)
		return true
	}
	vm.Functions.Store(name, fn)
	vm.Debug("Register BUND function: %v", name)
	return true
}

func (vm *VM) AddOperator(name string, fn BundOperator) bool {
	if _, ok := vm.Operators.Load(name); ok {
		vm.Warning("Operator %v already registered", name)
		return true
	}
	vm.Operators.Store(name, fn)
	vm.Debug("[ BUND ] Register BUND operator: %v", name)
	return true
}

func (vm *VM) AddSys(name string, fn BundSysfun) bool {
	if _, ok := vm.Sys.Load(name); ok {
		vm.Warning("System function %v already registered", name)
		return true
	}
	vm.Sys.Store(name, fn)
	vm.Debug("Register BUND system function: %v", name)
	return true
}

func (vm *VM) HasUserFunction(name string) bool {
	if !vm.IsStack() {
		vm.Error("Attempt to HasUserFunction(%v) on empty context", name)
		return false
	}
	return vm.CurrentNS.HasLambda(name)
}

func (vm *VM) GetFunction(name string) (BundFunction, error) {
	if res, ok := vm.Functions.Load(name); ok {
		return res.(BundFunction), nil
	}
	return nil, fmt.Errorf("BUND do not have function: %v", name)
}

func (vm *VM) GetOperator(name string) (BundOperator, error) {
	if res, ok := vm.Operators.Load(name); ok {
		return res.(BundOperator), nil
	}
	return nil, fmt.Errorf("BUND do not have operator: %v", name)
}

func (vm *VM) GetSys(name string) (BundSysfun, error) {
	if res, ok := vm.Sys.Load(name); ok {
		return res.(BundSysfun), nil
	}
	return nil, fmt.Errorf("BUND do not have system function: %v", name)
}
