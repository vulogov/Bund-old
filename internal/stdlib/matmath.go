package stdlib

import (
	"fmt"

	"gonum.org/v1/gonum/mat"

	vmmod "github.com/vulogov/Bund/internal/vm"
)

func MatMathOp(vm *vmmod.VM, e1 *vmmod.Elem, e2 *vmmod.Elem, op int) (*vmmod.Elem, error) {
	if e1.Type == "MAT" && e2.Type == "MAT" {
		vm.Debug("MAT operation %v requested", op)
		switch v1 := e1.Value.(type) {
		case *mat.Dense:
			switch v2 := e2.Value.(type) {
			case *mat.Dense:
				x1, y1 := v1.Dims()
				x2, y2 := v2.Dims()
				if x1 != x2 && y1 != y2 {
					return nil, fmt.Errorf("For the matrix math operations, matrixes must be the same arity: %v <-> %v %v <-> %v", x1, x2, y1, y2)
				}
				vm.Debug("MAT dimensions: %v=%v, %v=%v", x1, x2, y1, y2)
				mres := mat.NewDense(x1, y1, nil)
				switch op {
				case add:
					mres.Add(v1, v2)
				case sub:
					mres.Sub(v1, v2)
				case mul:
					mres.MulElem(v1, v2)
				case div:
					mres.DivElem(v1, v2)
				default:
					return nil, fmt.Errorf("Matrix operation was requested, and I do know for sure that's you got OP code wrong.")
				}
				return &vmmod.Elem{Type: "MAT", Value: mres}, nil
			}
		}
	}
	return nil, fmt.Errorf("Matrix operation was requested, but I do not know how to do it, due error in data, or OP code")
}
