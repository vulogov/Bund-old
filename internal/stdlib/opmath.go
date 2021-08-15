package stdlib

import (
	"fmt"
	"math"
	"math/big"

	"gonum.org/v1/gonum/mat"

	vmmod "github.com/vulogov/Bund/internal/vm"
)

func AddOperator(vm *vmmod.VM, e1 *vmmod.Elem, e2 *vmmod.Elem) (*vmmod.Elem, error) {
	if e1.Type == e2.Type {
		switch e1.Type {
		case "int":
			res := big.NewInt(0)
			d1 := e1.Value.(*big.Int)
			d2 := e2.Value.(*big.Int)
			res.Add(d1, d2)
			return &vmmod.Elem{Type: "int", Value: res}, nil
		case "flt":
			return &vmmod.Elem{Type: "flt", Value: (e1.Value.(float64) + e2.Value.(float64))}, nil
		case "str":
			return &vmmod.Elem{Type: "str", Value: (e1.Value.(string) + e2.Value.(string))}, nil
		case "dblock":
			return DblocksMathOp(vm, e1, e2, add)
		case "MAT":
			return MatMathOp(vm, e1, e2, add)
		}
	} else {
		switch e1.Type {
		case "int":
			switch e2.Type {
			case "flt":
				return &vmmod.Elem{Type: "flt", Value: (float64(e1.Value.(*big.Int).Int64()) + e2.Value.(float64))}, nil
			case "dblock":
				return DblockMathSingleOp(vm, e2, e1, add)
			}
		case "flt":
			switch e2.Type {
			case "int":
				return &vmmod.Elem{Type: "flt", Value: (float64(e1.Value.(*big.Int).Int64()) + e2.Value.(float64))}, nil
			case "dblock":
				return DblockMathSingleOp(vm, e2, e1, add)
			}
		case "dblock":
			switch e2.Type {
			case "int":
				return DblockMathSingleOp(vm, e1, e2, add)
			case "flt":
				return DblockMathSingleOp(vm, e1, e2, add)
			}
		}
	}
	return nil, fmt.Errorf("I do not know how to perform '+' for this data %v %v", e1.Type, e2.Type)
}

func MulOperator(vm *vmmod.VM, e1 *vmmod.Elem, e2 *vmmod.Elem) (*vmmod.Elem, error) {
	if e1.Type == e2.Type {
		switch e1.Type {
		case "int":
			res := big.NewInt(0)
			d1 := e1.Value.(*big.Int)
			d2 := e2.Value.(*big.Int)
			res.Mul(d1, d2)
			return &vmmod.Elem{Type: "int", Value: res}, nil
		case "flt":
			return &vmmod.Elem{Type: "flt", Value: (e1.Value.(float64) * e2.Value.(float64))}, nil
		case "dblock":
			return DblocksMathOp(vm, e1, e2, mul)
		case "MAT":
			return MatMathOp(vm, e1, e2, mul)
		}
	} else {
		switch e1.Type {
		case "int":
			switch e2.Type {
			case "flt":
				return &vmmod.Elem{Type: "flt", Value: (float64(e1.Value.(*big.Int).Int64()) * e2.Value.(float64))}, nil
			case "dblock":
				return DblockMathSingleOp(vm, e2, e1, mul)
			}
		case "flt":
			switch e2.Type {
			case "int":
				return &vmmod.Elem{Type: "flt", Value: (float64(e1.Value.(*big.Int).Int64()) * e2.Value.(float64))}, nil
			case "dblock":
				return DblockMathSingleOp(vm, e2, e1, mul)
			}
		case "dblock":
			switch e2.Type {
			case "int":
				return DblockMathSingleOp(vm, e1, e2, mul)
			case "flt":
				return DblockMathSingleOp(vm, e1, e2, mul)
			}
		}
	}
	return nil, fmt.Errorf("I do not know how to perform '*' for this data")
}

func DivOperator(vm *vmmod.VM, e1 *vmmod.Elem, e2 *vmmod.Elem) (*vmmod.Elem, error) {
	if e1.Type == e2.Type {
		switch e1.Type {
		case "int":
			res := big.NewInt(0)
			d1 := e1.Value.(*big.Int)
			d2 := e2.Value.(*big.Int)
			res.Div(d1, d2)
			return &vmmod.Elem{Type: "int", Value: res}, nil
		case "flt":
			return &vmmod.Elem{Type: "flt", Value: (e1.Value.(float64) / e2.Value.(float64))}, nil
		case "dblock":
			return DblocksMathOp(vm, e1, e2, div)
		case "MAT":
			return MatMathOp(vm, e1, e2, div)
		}
	} else {
		switch e1.Type {
		case "int":
			switch e2.Type {
			case "flt":
				return &vmmod.Elem{Type: "flt", Value: (float64(e1.Value.(*big.Int).Int64()) / e2.Value.(float64))}, nil
			case "dblock":
				return DblockMathSingleOp(vm, e2, e1, div)
			}
		case "flt":
			switch e2.Type {
			case "int":
				return &vmmod.Elem{Type: "flt", Value: (float64(e1.Value.(*big.Int).Int64()) / e2.Value.(float64))}, nil
			case "dblock":
				return DblockMathSingleOp(vm, e2, e1, div)
			}
		case "dblock":
			switch e2.Type {
			case "int":
				return DblockMathSingleOp(vm, e1, e2, div)
			case "flt":
				return DblockMathSingleOp(vm, e1, e2, div)
			}
		}
	}
	return nil, fmt.Errorf("I do not know how to perform '÷' for this data")
}

func SubOperator(vm *vmmod.VM, e1 *vmmod.Elem, e2 *vmmod.Elem) (*vmmod.Elem, error) {
	if e1.Type == e2.Type {
		switch e1.Type {
		case "int":
			res := big.NewInt(0)
			d1 := e1.Value.(*big.Int)
			d2 := e2.Value.(*big.Int)
			res.Sub(d1, d2)
			return &vmmod.Elem{Type: "int", Value: res}, nil
		case "flt":
			return &vmmod.Elem{Type: "flt", Value: (e1.Value.(float64) - e2.Value.(float64))}, nil
		case "dblock":
			return DblocksMathOp(vm, e1, e2, sub)
		case "MAT":
			return MatMathOp(vm, e1, e2, sub)
		}
	} else {
		switch e1.Type {
		case "int":
			switch e2.Type {
			case "flt":
				return &vmmod.Elem{Type: "flt", Value: (float64(e1.Value.(*big.Int).Int64()) - e2.Value.(float64))}, nil
			case "dblock":
				return DblockMathSingleOp(vm, e2, e1, sub)
			}
		case "flt":
			switch e2.Type {
			case "int":
				return &vmmod.Elem{Type: "flt", Value: (float64(e1.Value.(*big.Int).Int64()) - e2.Value.(float64))}, nil
			case "dblock":
				return DblockMathSingleOp(vm, e2, e1, sub)
			}
		case "dblock":
			switch e2.Type {
			case "int":
				return DblockMathSingleOp(vm, e1, e2, sub)
			case "flt":
				return DblockMathSingleOp(vm, e1, e2, sub)
			}
		}
	}
	return nil, fmt.Errorf("I do not know how to perform '-' for this data: %v %v", e1.Type, e2.Type)
}

func PowOperator(vm *vmmod.VM, e1 *vmmod.Elem, e2 *vmmod.Elem) (*vmmod.Elem, error) {
	if e1.Type == e2.Type {
		switch e1.Type {
		case "int":
			d1 := e1.Value.(*big.Int)
			d2 := e2.Value.(*big.Int)
			res := big.NewInt(0)
			res.Exp(d1, d2, nil)
			return &vmmod.Elem{Type: "int", Value: res}, nil
		case "flt":
			return &vmmod.Elem{Type: "flt", Value: math.Pow(e1.Value.(float64), e2.Value.(float64))}, nil
		}
	} else {
		switch e1.Type {
		case "int":
			p := int(e1.Value.(*big.Int).Int64())
			switch e2.Type {
			case "MAT":
				res := new(mat.Dense)
				x, y := e2.Value.(*mat.Dense).Dims()
				if x != y {
					return nil, fmt.Errorf("I can not perform ** for those dimentions of MAT: %vx%v", x, y)
				}
				vm.Debug("Performing %vx%v ** %v", x, y, p)
				res.Pow(e2.Value.(*mat.Dense), p)
				return &vmmod.Elem{Type: "MAT", Value: res}, nil
			}
		case "MAT":
			res := new(mat.Dense)
			switch e2.Type {
			case "int":
				p := int(e2.Value.(*big.Int).Int64())
				x, y := e1.Value.(*mat.Dense).Dims()
				if x != y {
					return nil, fmt.Errorf("I can not perform ** for those dimentions of MAT: %vx%v", x, y)
				}
				vm.Debug("Performing %vx%v ** %v", x, y, p)
				res.Pow(e1.Value.(*mat.Dense), p)
				return &vmmod.Elem{Type: "MAT", Value: res}, nil
			}
		}
	}
	return nil, fmt.Errorf("I do not know how to perform '**' for this data: %v %v", e1.Type, e2.Type)
}

func InitOpMath(vm *vmmod.VM) {
	vm.Debug("[ BUND ] bund.InitOpMath() reached")
	vm.AddOperator("+", AddOperator)
	vm.AddOperator("*", MulOperator)
	vm.AddOperator("×", MulOperator)
	vm.AddOperator("\\", DivOperator)
	vm.AddOperator("÷", DivOperator)
	vm.AddOperator("-", SubOperator)
	vm.AddOperator("**", PowOperator)
}
