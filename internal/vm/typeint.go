package vm

import (
	"fmt"

	"math/big"
)

func BigFactory(vm *VM) *Elem {
	return &Elem{Type: "int", Value: big.NewInt(0)}
}

func BigToString(vm *VM, e *Elem) string {
	if e.Type == "int" {
		i := e.Value.(*big.Int)
		return fmt.Sprintf("%v", i.String())
	}
	vm.Error("trying to convert an Int and it is not a Int: %v %T", e.Type, e.Value)
	return "<?>"
}

func BigFromString(vm *VM, d string) *Elem {
	i := big.NewInt(0)
	i.SetString(d, 0)
	res := Elem{Type: "int", Value: i}
	return &res
}

func BigCompare(vm *VM, e1 *Elem, e2 *Elem) int {
	if e1.Type == "int" && e2.Type == "int" {
		r1 := e1.Value.(*big.Int)
		r2 := e2.Value.(*big.Int)
		res := r1.Cmp(r2)
		if res == 0 {
			return Eq
		} else if res == -1 {
			return Ls
		} else if res == 1 {
			return Gt
		}
	}
	return IDK
}

func BigDup(vm *VM, e *Elem) *Elem {
	if e.Type != "int" {
		return nil
	}
	return BigFromString(vm, BigToString(vm, e))
}

func RegisterInt(vm *VM) {
	vm.RegisterType("int", BigFactory, BigToString, BigFromString, BigCompare, BigDup)
}
