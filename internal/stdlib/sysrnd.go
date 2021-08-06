package stdlib

import (
	"crypto/rand"
	"fmt"
	"math/big"

	vmmod "github.com/vulogov/Bund/internal/vm"
)

func IntRandomElement(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	if vm.Current.Len() < 1 {
		return nil, fmt.Errorf("Namespace stack is too shallow for rnd/Integer operation: %v:%v", vm.CurrentNS.Name, vm.Current.Len())
	}
	_max := vm.Take()
	if e.Type != "int" && _max.Type != "int" {
		return nil, fmt.Errorf("Integers required for rnd/Integer operation: %v:%v", _max.Type, e.Type)
	}
	max := _max.Value.(*big.Int)
	min := e.Value.(*big.Int)
	max.Sub(max, min)
	if max.Cmp(min) != 1 {
		return nil, fmt.Errorf("max(%v) > min(%v) for rnd/Integer operation/Integer", max.String(), min.String())
	}
	res, err := rand.Int(rand.Reader, max)
	vm.OnError(err, "Error in generating random number rand/Integer")
	res.Add(res, min)
	return &vmmod.Elem{Type: "int", Value: res}, nil
}

func PrimeRandomElement(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	vm.Debug("rnd/Prime: %v", e.Type)
	if e.Type != "int" {
		return nil, fmt.Errorf("Integer required for rnd/Prime operation: %v", e.Type)
	}
	bits := int(e.Value.(*big.Int).Int64())
	if bits <= 2 {
		return nil, fmt.Errorf("You must request more than two bits for rnd/Prime operation: %v", bits)
	}
	res, err := rand.Prime(rand.Reader, bits)
	vm.OnError(err, "Error in generating prime rand/Prime")
	return &vmmod.Elem{Type: "int", Value: res}, nil
}

func InitRndFunctions(vm *vmmod.VM) {
	vm.Debug("[ BUND ] bund.InitRndFunctions() reached")
	vm.AddFunction("rnd/Integer", IntRandomElement)
	vm.AddFunction("rnd/Prime", PrimeRandomElement)
}
