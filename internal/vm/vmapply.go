package vm

import (
	"fmt"
)

func Apply(name string, vm *VM) error {
	if !vm.IsStack() {
		return fmt.Errorf("Attempt to execute lambda %v on empty context", name)
	}
	ls := vm.CurrentNS.GetLambda(name)
	if ls == nil {
		return fmt.Errorf("Lambda %v not exist in %v", name, vm.Name)
	}
	vm.Debug("Executing lambda %v in %v", name, vm.Name)
	for ls.Len() > 0 {
		cmd := ls.PopFront().(*Elem)
		fmt.Println(cmd)
	}
	return nil
}

func (vm *VM) Apply(name string) error {
	err := Apply(name, vm)
	if err != nil {
		vm.OnError(err, "Error in Apply(%v)", name)
	}
	return err
}
