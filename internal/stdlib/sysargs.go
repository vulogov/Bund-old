package stdlib

import (
	"fmt"
	"math/big"

	"github.com/gammazero/deque"

	"github.com/vulogov/Bund/internal/conf"
	vmmod "github.com/vulogov/Bund/internal/vm"
)

func ArgsLenElement(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	vm.Put(e)
	return &vmmod.Elem{Type: "int", Value: big.NewInt(int64(len(conf.Argv)))}, nil
}

func ArgsLenElementLen(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	if e.Type == "int" {
		i := int(e.Value.(*big.Int).Int64())
		if i > (len(conf.Argv) - 1) {
			return nil, fmt.Errorf("Index %v for args/Len/Slice out of bounds", i)
		}
		return &vmmod.Elem{Type: "int", Value: big.NewInt(int64(len(conf.Argv[i])))}, nil
	}
	return nil, fmt.Errorf("Expecting an index as int for args/Slice/Len")
}

func ArgsElement(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	if e.Type == "int" {
		i := int(e.Value.(*big.Int).Int64())
		if i > (len(conf.Argv) - 1) {
			return nil, fmt.Errorf("Index %v for args/Len/Element out of bounds", i)
		}
		eh, err := vm.GetType("dblock")
		vm.OnError(err, "Error in arg/Slice")
		res := eh.Factory(vm)
		q := res.Value.(*deque.Deque)
		for c := 0; c < len(conf.Argv[i]); c++ {
			q.PushBack(&vmmod.Elem{Type: "str", Value: conf.Argv[i][c]})
		}
		return res, nil
	}
	return nil, fmt.Errorf("Expecting an index as int for args/Slice")
}

func InitArgsFunctions(vm *vmmod.VM) {
	vm.Debug("[ BUND ] bund.InitArgsFunctions() reached")
	vm.AddFunction("arg/Len", ArgsLenElement)
	vm.AddFunction("arg/Slice/Len", ArgsLenElementLen)
	vm.AddFunction("arglen", ArgsLenElementLen)
	vm.AddFunction("arg/Slice", ArgsElement)
	vm.AddFunction("arg", ArgsElement)
}
