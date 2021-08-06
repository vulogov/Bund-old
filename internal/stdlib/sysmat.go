package stdlib

import (
	"fmt"
	"math/big"

	"github.com/gammazero/deque"
	"gonum.org/v1/gonum/mat"

	vmmod "github.com/vulogov/Bund/internal/vm"
)

func MatrixElement(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	var x, y int
	var fill float64
	if e.Type != "dblock" {
		return nil, fmt.Errorf("You have to pass DATABLOCK as (* ) for M (Matrix) creation. I have: %v", e.Type)
	}
	if vmmod.BlockLen(e) == 2 {
		_x, err := vmmod.BlockAt(e, "int", 1)
		vm.OnError(err, "Error during Matrix creation in M")
		_y, err := vmmod.BlockAt(e, "int", 0)
		vm.OnError(err, "Error during Matrix creation in M")
		x = int(_x.(int64))
		y = int(_y.(int64))
		fill = 0.0
	} else if vmmod.BlockLen(e) == 3 {
		_fill, err := vmmod.BlockAt(e, "flt", 0)
		vm.OnError(err, "Error during Matrix creation in M")
		_x, err := vmmod.BlockAt(e, "int", 2)
		vm.OnError(err, "Error during Matrix creation in M")
		_y, err := vmmod.BlockAt(e, "int", 1)
		vm.OnError(err, "Error during Matrix creation in M")
		x = int(_x.(int64))
		y = int(_y.(int64))
		fill = _fill.(float64)
	} else {
		return nil, fmt.Errorf("You have to pass (* x y <init value> y x ) to M")
	}
	vm.Debug("Creating matrix X=%v Y=%v FILL=%v", x, y, fill)
	farray := make([]float64, (x * y))
	for i := 0; i < len(farray); i++ {
		farray[i] = fill
	}
	return &vmmod.Elem{Type: "MAT", Value: mat.NewDense(x, y, farray)}, nil
}

func matrixSet(vm *vmmod.VM, mtx *vmmod.Elem, e *vmmod.Elem) {
	var x, y int
	var v float64

	if vmmod.BlockLen(e) == 3 {
		_v, err := vmmod.BlockAt(e, "flt", 0)
		vm.OnError(err, "Error during M/Set")
		_x, err := vmmod.BlockAt(e, "int", 2)
		vm.OnError(err, "Error during M/Set")
		_y, err := vmmod.BlockAt(e, "int", 1)
		vm.OnError(err, "Error during M/Set")
		x = int(_x.(int64))
		y = int(_y.(int64))
		v = _v.(float64)
		switch m := mtx.Value.(type) {
		case *mat.Dense:
			vm.Debug("M/Set X=%v Y=%v V=%v", x, y, v)
			m.Set(x, y, v)
		default:
			vm.Fatal("You do not have a proper MAT value")
		}
		return
	}
	vm.Fatal("You have to pass (* value y x ) to M/Set")
}

func MatrixElementSet(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	if vm.Current.Len() < 1 {
		return nil, fmt.Errorf("Namespace stack is too shallow for M/Set operation: %v:%v", vm.CurrentNS.Name, vm.Current.Len())
	}
	e1 := vm.Take()
	switch e.Type {
	case "MAT":
		switch e1.Type {
		case "dblock":
			matrixSet(vm, e, e1)
			return e, nil
		}
	case "dblock":
		switch e1.Type {
		case "MAT":
			matrixSet(vm, e1, e)
			return e1, nil
		}
	}
	return nil, fmt.Errorf("You must have MAT and dblock to perform M/Set operation")
}

func matrixGet(vm *vmmod.VM, mtx *vmmod.Elem, e *vmmod.Elem) *vmmod.Elem {
	var x, y int

	if vmmod.BlockLen(e) == 2 {
		_x, err := vmmod.BlockAt(e, "int", 1)
		vm.OnError(err, "Error during M/Get")
		_y, err := vmmod.BlockAt(e, "int", 0)
		vm.OnError(err, "Error during M/Get")
		x = int(_x.(int64))
		y = int(_y.(int64))
		switch m := mtx.Value.(type) {
		case *mat.Dense:
			vm.Debug("M/Get X=%v Y=%v", x, y)
			val := m.At(x, y)
			return &vmmod.Elem{Type: "flt", Value: val}
		default:
			vm.Fatal("You do not have a proper MAT value")
		}
		vm.Fatal("I do not know what to do in M/Get")
	}
	vm.Fatal("You have to pass (* value y x ) to M/Set")
	return nil
}

func MatrixElementGet(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	if vm.Current.Len() < 1 {
		return nil, fmt.Errorf("Namespace stack is too shallow for M/Get operation: %v:%v", vm.CurrentNS.Name, vm.Current.Len())
	}
	e1 := vm.Take()
	switch e.Type {
	case "MAT":
		switch e1.Type {
		case "dblock":
			res := matrixGet(vm, e, e1)
			vm.Put(e)
			return res, nil
		}
	case "dblock":
		switch e1.Type {
		case "MAT":
			res := matrixGet(vm, e1, e)
			vm.Put(e1)
			return res, nil
		}
	}
	return nil, fmt.Errorf("You must have MAT and dblock to perform M/Set operation")
}

func MatrixElementMake(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	var data []float64
	if e.Type != "dblock" {
		return nil, fmt.Errorf("Expecting to get DBLOCk with data for matrix creation: %v", e.Type)
	}
	q := e.Value.(*deque.Deque)
	rows := 0
	cols := 0
	colsize := 0
	for i := q.Len() - 1; i >= 0; i-- {
		e1 := q.At(i).(*vmmod.Elem)
		if e1.Type == "SEPARATE" {
			rows += 1
			if colsize == 0 {
				colsize = cols
			}
			if cols != colsize {
				return nil, fmt.Errorf("M/Make matrix is not uniform: %v <> %v", colsize, cols)
			}
			cols = 0
			continue
		}
		if e1.Type == "flt" {
			cols += 1
			data = append(data, e1.Value.(float64))
		}
		if e1.Type == "int" {
			cols += 1
			data = append(data, float64(e1.Value.(*big.Int).Int64()))
		}
	}
	res := mat.NewDense(rows, colsize, data)
	return &vmmod.Elem{Type: "MAT", Value: res}, nil
}

func MatrixElementBlock(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	if e.Type != "MAT" {
		return nil, fmt.Errorf("Expecting to get MAT with data for BLOCK creation: %v", e.Type)
	}
	eh, err := vm.GetType("dblock")
	vm.OnError(err, "Error in M/ToBlock")
	res := eh.Factory(vm)
	q := res.Value.(*deque.Deque)
	x, y := e.Value.(*mat.Dense).Dims()
	for i := 0; i < x; i++ {
		q.PushFront(&vmmod.Elem{Type: "SEPARATE", Value: nil})
		for j := 0; j < y; j++ {
			e1 := e.Value.(*mat.Dense).At(i, j)
			q.PushFront(&vmmod.Elem{Type: "flt", Value: e1})
		}
	}
	return res, nil
}

func InitMatFunctions(vm *vmmod.VM) {
	vm.Debug("[ BUND ] bund.InitMatFunctions() reached")
	vm.AddFunction("M", MatrixElement)
	vm.AddFunction("M/Make", MatrixElementMake)
	vm.AddFunction("M/ToBlock", MatrixElementBlock)
	vm.AddFunction("M/Set", MatrixElementSet)
	vm.AddFunction("M/Get", MatrixElementGet)
}
