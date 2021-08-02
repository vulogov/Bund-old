package vm

import (
	"fmt"

	"strconv"
)

func ComplexFactory(vm *VM) *Elem {
	return &Elem{Type: "cpx", Value: complex(0, 0)}
}

func ComplexToString(vm *VM, e *Elem) string {
	if e.Type == "cpx" {
		switch v := e.Value.(type) {
		case complex128:
			return fmt.Sprint(v)
		}
	}
	vm.Error("trying to convert an Complex and it is not an Complex: %v %T", e.Type, e.Value)
	return "<?>"
}

func ComplexFromString(vm *VM, d string) *Elem {
	res, err := strconv.ParseComplex(d, 128)
	if err != nil {
		vm.OnError(err, "Error parsing complex number %v", d)
		return nil
	}
	return &Elem{Type: "cpx", Value: res}
}

func ComplexCompare(vm *VM, e1 *Elem, e2 *Elem) int {
	var fe1, fe2 complex128
	if e1.Type == "cpx" {
		fe1 = e1.Value.(complex128)
	} else {
		return IDK
	}
	if e2.Type == "cpx" {
		fe2 = e2.Value.(complex128)
	} else {
		return IDK
	}
	if fe1 == fe2 {
		return Eq
	} else {
		return IDK
	}
}

func ComplexDup(vm *VM, e *Elem) *Elem {
	if e.Type != "cpx" {
		return nil
	}
	return &Elem{Type: "cpx", Value: e.Value}
}

func RegisterCpx(vm *VM) {
	vm.RegisterType("cpx", ComplexFactory, ComplexToString, ComplexFromString, ComplexCompare, ComplexDup)
}
