package stdlib

import (
	"fmt"
	"math/big"

	"github.com/gammazero/deque"

	vmmod "github.com/vulogov/Bund/internal/vm"
)

const (
	add = 1
	sub = 2
	mul = 3
	div = 4
)

func DblocksMathOp(vm *vmmod.VM, e1 *vmmod.Elem, e2 *vmmod.Elem, op int) (*vmmod.Elem, error) {
	if !vmmod.DblockSameArity(vm, e1, e2) {
		return nil, fmt.Errorf("Elements for the math operation are not of the same arity")
	}
	eh, err := vm.GetType("dblock")
	vm.OnError(err, "Error during the math operation over dblocks")
	ehint, err := vm.GetType("int")
	vm.OnError(err, "Error during the math operation over dblocks")
	res := eh.Factory(vm)
	resq := res.Value.(*deque.Deque)
	q1 := e1.Value.(*deque.Deque)
	q2 := e2.Value.(*deque.Deque)
	for i := q1.Len() - 1; i >= 0; i-- {
		v1 := q1.At(i).(*vmmod.Elem)
		v2 := q2.At(i).(*vmmod.Elem)
		if v1.Type == "SEPARATE" {
			continue
		}
		switch v1.Type {
		case "int":
			v3 := ehint.Factory(vm)
			v3e := v3.Value.(*big.Int)
			switch op {
			case add:
				v3e.Add(v1.Value.(*big.Int), v2.Value.(*big.Int))
			case sub:
				v3e.Sub(v1.Value.(*big.Int), v2.Value.(*big.Int))
			case mul:
				v3e.Mul(v1.Value.(*big.Int), v2.Value.(*big.Int))
			case div:
				v3e.Div(v1.Value.(*big.Int), v2.Value.(*big.Int))
			}
			resq.PushFront(v3)
		case "flt":
			v3 := 0.0
			switch op {
			case add:
				v3 = v1.Value.(float64) + v2.Value.(float64)
			case sub:
				v3 = v1.Value.(float64) - v2.Value.(float64)
			case mul:
				v3 = v1.Value.(float64) * v2.Value.(float64)
			case div:
				v3 = v1.Value.(float64) / v2.Value.(float64)
			}
			resq.PushFront(&vmmod.Elem{Type: "flt", Value: v3})
		default:
			vm.Fatal("Attempt of performing a math operation in dblock with values of incorrect type %v", v1.Type)
		}
	}
	return res, nil
}

func DblockMathSingleOp(vm *vmmod.VM, e1 *vmmod.Elem, e2 *vmmod.Elem, op int) (*vmmod.Elem, error) {
	if e1.Type != "dblock" {
		return nil, fmt.Errorf("Operation over 'dblock' been requested, but there is no 'dblock': %v", e1.Type)
	}
	q1 := e1.Value.(*deque.Deque)
	for i := 0; i < q1.Len(); i++ {
		v1 := q1.At(i).(*vmmod.Elem)
		switch v1.Type {
		case "SEPARATE":
			continue
		case "int":
			switch e2.Type {
			case "int":
				switch op {
				case add:
					v1.Value.(*big.Int).Add(v1.Value.(*big.Int), e2.Value.(*big.Int))
				case mul:
					v1.Value.(*big.Int).Mul(v1.Value.(*big.Int), e2.Value.(*big.Int))
				case sub:
					v1.Value.(*big.Int).Sub(v1.Value.(*big.Int), e2.Value.(*big.Int))
				case div:
					v1.Value.(*big.Int).Div(v1.Value.(*big.Int), e2.Value.(*big.Int))
				}
				q1.Set(i, v1)
			case "flt":
				um := big.NewInt(int64(e2.Value.(float64)))
				switch op {
				case add:
					v1.Value.(*big.Int).Add(v1.Value.(*big.Int), um)
				case mul:
					v1.Value.(*big.Int).Mul(v1.Value.(*big.Int), um)
				case sub:
					v1.Value.(*big.Int).Sub(v1.Value.(*big.Int), um)
				case div:
					v1.Value.(*big.Int).Div(v1.Value.(*big.Int), um)
				}
				q1.Set(i, v1)
			}
		case "flt":
			switch e2.Type {
			case "int":
				um := float64(e2.Value.(*big.Int).Int64())
				switch op {
				case add:
					q1.Set(i, &vmmod.Elem{Type: "flt", Value: (v1.Value.(float64) + um)})
				case mul:
					q1.Set(i, &vmmod.Elem{Type: "flt", Value: (v1.Value.(float64) * um)})
				case sub:
					q1.Set(i, &vmmod.Elem{Type: "flt", Value: (v1.Value.(float64) - um)})
				case div:
					q1.Set(i, &vmmod.Elem{Type: "flt", Value: (v1.Value.(float64) / um)})
				}
			case "flt":
				switch op {
				case add:
					q1.Set(i, &vmmod.Elem{Type: "flt", Value: (v1.Value.(float64) + e2.Value.(float64))})
				case mul:
					q1.Set(i, &vmmod.Elem{Type: "flt", Value: (v1.Value.(float64) * e2.Value.(float64))})
				case sub:
					q1.Set(i, &vmmod.Elem{Type: "flt", Value: (v1.Value.(float64) - e2.Value.(float64))})
				case div:
					q1.Set(i, &vmmod.Elem{Type: "flt", Value: (v1.Value.(float64) / e2.Value.(float64))})
				}
			}
		}
	}
	return e1, nil
}
