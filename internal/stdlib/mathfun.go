package stdlib

import (
	"math"
	"math/big"

	vmmod "github.com/vulogov/Bund/internal/vm"
)

func MathSinElement(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	vm.Debug("SIN(): %v", e.Type)
	switch e.Type {
	case "flt":
		return &vmmod.Elem{Type: "flt", Value: math.Sin(e.Value.(float64))}, nil
	case "int":
		return &vmmod.Elem{Type: "flt", Value: math.Sin(float64(e.Value.(*big.Int).Int64()))}, nil
	case "dblock":
		return DblockWalk(vm, e, math.Sin)
	case "MAT":
		return MatrixWalk(vm, e, math.Sin)
	}
	return e, nil
}

func MathCosElement(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	vm.Debug("COS(): %v", e.Type)
	switch e.Type {
	case "flt":
		return &vmmod.Elem{Type: "flt", Value: math.Cos(e.Value.(float64))}, nil
	case "int":
		return &vmmod.Elem{Type: "flt", Value: math.Cos(float64(e.Value.(*big.Int).Int64()))}, nil
	case "dblock":
		return DblockWalk(vm, e, math.Cos)
	case "MAT":
		return MatrixWalk(vm, e, math.Cos)
	}
	return e, nil
}

func MathTanElement(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	vm.Debug("TAN(): %v", e.Type)
	switch e.Type {
	case "flt":
		return &vmmod.Elem{Type: "flt", Value: math.Tan(e.Value.(float64))}, nil
	case "int":
		return &vmmod.Elem{Type: "flt", Value: math.Tan(float64(e.Value.(*big.Int).Int64()))}, nil
	case "dblock":
		return DblockWalk(vm, e, math.Tan)
	case "MAT":
		return MatrixWalk(vm, e, math.Tan)
	}
	return e, nil
}

func MathSqrtElement(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	vm.Debug("SQRT(): %v", e.Type)
	switch e.Type {
	case "flt":
		return &vmmod.Elem{Type: "flt", Value: math.Sqrt(e.Value.(float64))}, nil
	case "int":
		return &vmmod.Elem{Type: "flt", Value: math.Sqrt(float64(e.Value.(*big.Int).Int64()))}, nil
	case "dblock":
		return DblockWalk(vm, e, math.Sqrt)
	case "MAT":
		return MatrixWalk(vm, e, math.Sqrt)
	}
	return e, nil
}

func MathExpElement(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	vm.Debug("EXP(): %v", e.Type)
	switch e.Type {
	case "flt":
		return &vmmod.Elem{Type: "flt", Value: math.Exp(e.Value.(float64))}, nil
	case "int":
		return &vmmod.Elem{Type: "flt", Value: math.Exp(float64(e.Value.(*big.Int).Int64()))}, nil
	case "dblock":
		return DblockWalk(vm, e, math.Exp)
	case "MAT":
		return MatrixWalk(vm, e, math.Exp)
	}
	return e, nil
}

func MathLogElement(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	vm.Debug("LOG(): %v", e.Type)
	switch e.Type {
	case "flt":
		return &vmmod.Elem{Type: "flt", Value: math.Log(e.Value.(float64))}, nil
	case "int":
		return &vmmod.Elem{Type: "flt", Value: math.Log(float64(e.Value.(*big.Int).Int64()))}, nil
	case "dblock":
		return DblockWalk(vm, e, math.Log)
	case "MAT":
		return MatrixWalk(vm, e, math.Log)
	}
	return e, nil
}

func MathLog10Element(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	vm.Debug("LOG10(): %v", e.Type)
	switch e.Type {
	case "flt":
		return &vmmod.Elem{Type: "flt", Value: math.Log10(e.Value.(float64))}, nil
	case "int":
		return &vmmod.Elem{Type: "flt", Value: math.Log10(float64(e.Value.(*big.Int).Int64()))}, nil
	case "dblock":
		return DblockWalk(vm, e, math.Log10)
	case "MAT":
		return MatrixWalk(vm, e, math.Log10)
	}
	return e, nil
}

func InitMathFunctions(vm *vmmod.VM) {
	vm.Debug("[ BUND ] bund.InitMathFunctions() reached")
	vm.AddFunction("math/Sin", MathSinElement)
	vm.AddFunction("math/Cos", MathCosElement)
	vm.AddFunction("math/Tan", MathTanElement)
	vm.AddFunction("math/Sqrt", MathSqrtElement)
	vm.AddFunction("math/Exp", MathExpElement)
	vm.AddFunction("math/Log", MathLogElement)
	vm.AddFunction("math/Log10", MathLog10Element)
}
