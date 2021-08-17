package stdlib

import (
	"fmt"
	"math/big"

	vmmod "github.com/vulogov/Bund/internal/vm"
)

func PercentOperator(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	if vm.CanGet() {
		e1 := vm.Take()
		if e.Type == e1.Type {
			switch e.Type {
			case "int":
				res := big.NewInt(0)
				res.Mul(e.Value.(*big.Int), e1.Value.(*big.Int))
				res.Div(res, big.NewInt(100))
				return &vmmod.Elem{Type: "int", Value: res}, nil
			case "flt":
				return &vmmod.Elem{Type: "flt", Value: ((e.Value.(float64) * e1.Value.(float64)) / float64(100))}, nil
			default:
				return nil, fmt.Errorf("Unknown data types for percent operator: %v", e.Type)
			}
		} else {
			return nil, fmt.Errorf("Two elements  of same type required in stack for percent operator")
		}
	}
	return nil, fmt.Errorf("Two elements required in stack for percent operator")
}

func PercentOfOperator(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	if vm.CanGet() {
		e1 := vm.Take()
		if e.Type == e1.Type {
			switch e.Type {
			case "int":
				res := big.NewInt(0)
				res.Mul(e.Value.(*big.Int), big.NewInt(100))
				res.Div(res, e1.Value.(*big.Int))
				return &vmmod.Elem{Type: "int", Value: res}, nil
			case "flt":
				return &vmmod.Elem{Type: "flt", Value: ((e.Value.(float64) * float64(100)) / e1.Value.(float64))}, nil
			default:
				return nil, fmt.Errorf("Unknown data types for percent operator: %v", e.Type)
			}
		} else {
			return nil, fmt.Errorf("Two elements  of same type required in stack for percent operator")
		}
	}
	return nil, fmt.Errorf("Two elements required in stack for percent operator")
}

func ChangeOperator(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	if vm.CanGet() {
		e1 := vm.Take()
		if e.Type == e1.Type {
			switch e.Type {
			case "int":
				res := big.NewInt(0)
				res.Sub(e.Value.(*big.Int), e1.Value.(*big.Int))
				res.Div(res, e1.Value.(*big.Int))
				res.Mul(res, big.NewInt(100))
				return &vmmod.Elem{Type: "int", Value: res}, nil
			case "flt":
				return &vmmod.Elem{Type: "flt", Value: float64(100) * ((e.Value.(float64) - e1.Value.(float64)) / e1.Value.(float64))}, nil
			default:
				return nil, fmt.Errorf("Unknown data types for percent operator: %v", e.Type)
			}
		} else {
			return nil, fmt.Errorf("Two elements  of same type required in stack for percent operator")
		}
	}
	return nil, fmt.Errorf("Two elements required in stack for percent operator")
}

func InitOpPercent(vm *vmmod.VM) {
	vm.Debug("[ BUND ] bund.InitOpPercent() reached")
	vm.AddFunction("percent", PercentOperator)
	vm.AddFunction("percentOf", PercentOfOperator)
	vm.AddFunction("changeOf", ChangeOperator)
}
