package vm

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func MatFactory(vm *VM) *Elem {
	return &Elem{Type: "MAT", Value: new(mat.Dense)}
}

func MatToString(vm *VM, e *Elem) string {
	if e.Type == "MAT" {
		m := e.Value.(*mat.Dense)
		fm := mat.Formatted(m, mat.Prefix(""), mat.Squeeze())
		return fmt.Sprintf("%v", fm)
	}
	vm.Error("trying to convert a MAT and it is not a MAT: %v %T", e.Type, e.Value)
	return "<?>"
}

func MatFromString(vm *VM, d string) *Elem {
	return &Elem{Type: "MAT", Value: new(mat.Dense)}
}

func MatCompare(vm *VM, e1 *Elem, e2 *Elem) int {
	if e1.Type == "MAT" && e2.Type == "MAT" {
		switch v1 := e1.Value.(type) {
		case *mat.Dense:
			switch v2 := e2.Value.(type) {
			case *mat.Dense:
				if mat.Equal(v1, v2) {
					return Eq
				} else {
					return Ne
				}
			}
		}
	}
	return IDK
}

func MatDup(vm *VM, e *Elem) *Elem {
	if e.Type != "MAT" {
		return nil
	}
	m := new(mat.Dense)
	switch v := e.Value.(type) {
	case *mat.Dense:
		m.CloneFrom(v)
		return &Elem{Type: "MAT", Value: m}
	}
	return &Elem{Type: "MAT", Value: m}
}

func RegisterMatrix(vm *VM) {
	vm.RegisterType("MAT", MatFactory, MatToString, MatFromString, MatCompare, MatDup)
}
