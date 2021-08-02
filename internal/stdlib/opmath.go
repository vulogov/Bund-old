package stdlib

import (
	"fmt"
	"math/big"

	vmmod "github.com/vulogov/Bund/internal/vm"
)

func AddOperator(vm *vmmod.VM, e1 *vmmod.Elem, e2 *vmmod.Elem) (*vmmod.Elem, error) {
	if e1.Type != e2.Type {
		return nil, fmt.Errorf("Datatypes for operation '+' must be the same")
	}
	switch e1.Type {
	case "int":
		d1 := e1.Value.(*big.Int)
		d2 := e2.Value.(*big.Int)
		d1.Add(d1, d2)
		return &vmmod.Elem{Type: "int", Value: d1}, nil
	case "flt":
		return &vmmod.Elem{Type: "flt", Value: (e1.Value.(float64) + e2.Value.(float64))}, nil
	case "str":
		return &vmmod.Elem{Type: "str", Value: (e1.Value.(string) + e2.Value.(string))}, nil
	}
	return nil, fmt.Errorf("I do not know how to perform '+' for this data")
}

func MulOperator(vm *vmmod.VM, e1 *vmmod.Elem, e2 *vmmod.Elem) (*vmmod.Elem, error) {
	if e1.Type != e2.Type {
		return nil, fmt.Errorf("Datatypes for operation '*' must be the same")
	}
	switch e1.Type {
	case "int":
		d1 := e1.Value.(*big.Int)
		d2 := e2.Value.(*big.Int)
		d1.Mul(d1, d2)
		return &vmmod.Elem{Type: "int", Value: d1}, nil
	case "flt":
		return &vmmod.Elem{Type: "flt", Value: (e1.Value.(float64) * e2.Value.(float64))}, nil
	}
	return nil, fmt.Errorf("I do not know how to perform '*' for this data")
}

func DivOperator(vm *vmmod.VM, e1 *vmmod.Elem, e2 *vmmod.Elem) (*vmmod.Elem, error) {
	if e1.Type != e2.Type {
		return nil, fmt.Errorf("Datatypes for operation '÷' must be the same")
	}
	switch e1.Type {
	case "int":
		d1 := e1.Value.(*big.Int)
		d2 := e2.Value.(*big.Int)
		d1.Div(d1, d2)
		return &vmmod.Elem{Type: "int", Value: d1}, nil
	case "flt":
		return &vmmod.Elem{Type: "flt", Value: (e1.Value.(float64) / e2.Value.(float64))}, nil
	}
	return nil, fmt.Errorf("I do not know how to perform '÷' for this data")
}

func SubOperator(vm *vmmod.VM, e1 *vmmod.Elem, e2 *vmmod.Elem) (*vmmod.Elem, error) {
	if e1.Type != e2.Type {
		return nil, fmt.Errorf("Datatypes for operation '-' must be the same")
	}
	switch e1.Type {
	case "int":
		d1 := e1.Value.(*big.Int)
		d2 := e2.Value.(*big.Int)
		d1.Sub(d1, d2)
		return &vmmod.Elem{Type: "int", Value: d1}, nil
	case "flt":
		return &vmmod.Elem{Type: "flt", Value: (e1.Value.(float64) - e2.Value.(float64))}, nil
	}
	return nil, fmt.Errorf("I do not know how to perform '-' for this data")
}

func InitOpMath(vm *vmmod.VM) {
	vm.Debug("[ BUND ] bund.InitOpMath() reached")
	vm.AddOperator("+", AddOperator)
	vm.AddOperator("*", MulOperator)
	vm.AddOperator("×", MulOperator)
	vm.AddOperator("\\", DivOperator)
	vm.AddOperator("÷", DivOperator)
	vm.AddOperator("-", SubOperator)
}
