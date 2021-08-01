package vm

import (
	"github.com/cstockton/go-conv"
)

func BoolFactory(vm *VM) *Elem {
	return &Elem{Type: "bool", Value: true}
}

func BoolToString(vm *VM, e *Elem) string {
	if e.Type == "bool" {
		switch v := e.Value.(type) {
		case bool:
			if v == true {
				return "True"
			} else {
				return "False"
			}
		}
	}
	vm.Error("trying to convert a Boolean and it is not a Boolean: %v %T", e.Type, e.Value)
	return "<?>"
}

func BoolFromString(vm *VM, d string) *Elem {
	res, err := conv.Bool(d)
	if err != nil {
		return nil
	}
	return &Elem{Type: "bool", Value: res}
}

func BoolCompare(vm *VM, e1 *Elem, e2 *Elem) int {
	if e1.Type == "bool" && e2.Type == "bool" {
		r1 := e1.Value.(bool)
		r2 := e2.Value.(bool)
		if r1 == r2 {
			return Eq
		} else if r1 == true && r2 == false {
			return Gt
		} else {
			return Ls
		}
	}
	return IDK
}

func BoolDup(vm *VM, e *Elem) *Elem {
	if e.Type != "bool" {
		return nil
	}
	return &Elem{Type: "bool", Value: e.Value}
}

func RegisterBoolean(vm *VM) {
	vm.RegisterType("bool", BoolFactory, BoolToString, BoolFromString, BoolCompare, BoolDup)
}
