package stdlib

import (
	"fmt"
	"math/big"

	"github.com/gammazero/deque"
	"gonum.org/v1/gonum/mat"

	vmmod "github.com/vulogov/Bund/internal/vm"
)

func MatrixWalk(vm *vmmod.VM, e *vmmod.Elem, f func(float64) float64) (*vmmod.Elem, error) {
	if e.Type != "MAT" {
		return nil, fmt.Errorf("Expecting to get MAT with data : %v", e.Type)
	}
	x, y := e.Value.(*mat.Dense).Dims()
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			e1 := e.Value.(*mat.Dense).At(i, j)
			e.Value.(*mat.Dense).Set(i, j, f(e1))
		}
	}
	return e, nil
}

func DblockWalk(vm *vmmod.VM, e *vmmod.Elem, f func(float64) float64) (*vmmod.Elem, error) {
	if e.Type != "dblock" {
		return nil, fmt.Errorf("Expecting to get DBLOCK with data : %v", e.Type)
	}
	q1 := e.Value.(*deque.Deque)
	eh, err := vm.GetType("dblock")
	vm.OnError(err, "Error in M/ToBlock")
	res := eh.Factory(vm)
	q2 := res.Value.(*deque.Deque)
	for q1.Len() > 0 {
		e1 := q1.PopFront().(*vmmod.Elem)
		switch e1.Type {
		case "flt":
			q2.PushBack(&vmmod.Elem{Type: "flt", Value: f(e1.Value.(float64))})
		case "int":
			q2.PushBack(&vmmod.Elem{Type: "flt", Value: f(float64(e1.Value.(*big.Int).Int64()))})
		default:
			q2.PushBack(e1)
		}
	}
	return res, nil
}
