package stdlib

import (
	"fmt"
	"math/big"

	"github.com/Jeffail/gabs/v2"
	"github.com/gammazero/deque"

	vmmod "github.com/vulogov/Bund/internal/vm"
)

func JsonArrayGetElement(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	vm.Debug("json/Array/Get: %v", e.Type)
	if e.Type == "json" {
		if !vm.CanGet() {
			return nil, fmt.Errorf("Stack is too shallow for json/Array/Get")
		}
		j := e.Value.(*gabs.Container)
		d := vm.Take()
		if d.Type != "str" {
			return nil, fmt.Errorf("string is expected for json/Array/Set")
		}
		if !j.ExistsP(d.Value.(string)) {
			return nil, fmt.Errorf("Key %v not found in JSON", d.Value.(string))
		}
		vm.Put(e)
		eh, err := vm.GetType("dblock")
		vm.OnError(err, "Error in json/Array/Get")
		res := eh.Factory(vm)
		q := res.Value.(*deque.Deque)
		for _, child := range j.Path(d.Value.(string)).Children() {
			v := child.Data()
			switch v.(type) {
			case int:
				q.PushBack(&vmmod.Elem{Type: "int", Value: big.NewInt(int64(v.(int)))})
			case int64:
				q.PushBack(&vmmod.Elem{Type: "int", Value: big.NewInt(v.(int64))})
			case float32:
				q.PushBack(&vmmod.Elem{Type: "flt", Value: float64(v.(float32))})
			case float64:
				q.PushBack(&vmmod.Elem{Type: "flt", Value: float64(v.(float64))})
			case string:
				q.PushBack(&vmmod.Elem{Type: "str", Value: string(v.(string))})
			default:
				return nil, fmt.Errorf("Unsupported data in JSON while in json/Value/Get")
			}
		}
		return res, nil
	}
	return e, nil
}

func JsonArrayAddElement(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	vm.Debug("json/Array/Add: %v", e.Type)
	if e.Type == "json" {
		if !vm.CanGet() {
			return nil, fmt.Errorf("Stack is too shallow for json/Array/Add")
		}
		j := e.Value.(*gabs.Container)
		d := vm.Take()
		if d.Type != "dblock" {
			return nil, fmt.Errorf("(* ) is expected for json/Array/Add")
		}
		if vmmod.BlockLen(d) == 0 {
			return nil, fmt.Errorf("(* ) is incorrect arity for json/Array/Add")
		}
		k, err := vmmod.Block_At(d, int64(0))
		vm.OnError(err, "Error in json/Array/Add")
		switch key := k.(type) {
		case string:
			if !j.ExistsP(key) {
				j.ArrayP(key)
			}
			for i := 1; i < int(vmmod.BlockLen(d)); i++ {
				v, err := vmmod.Block_At(d, int64(i))
				vm.OnError(err, "Error in json/Array/Add")
				j.ArrayAppendP(v, key)
			}
		}
	}
	return e, nil
}

func JsonArrayNewElement(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	vm.Debug("json/Array/New: %v", e.Type)
	if e.Type == "json" {
		if !vm.CanGet() {
			return nil, fmt.Errorf("Stack is too shallow for json/Array/New")
		}
		j := e.Value.(*gabs.Container)
		d := vm.Take()
		if d.Type != "dblock" {
			return nil, fmt.Errorf("(* ) is expected for json/Array/New")
		}
		if vmmod.BlockLen(d) == 0 {
			return nil, fmt.Errorf("(* ) is incorrect arity for json/Array/New")
		}
		for i := 0; i < int(vmmod.BlockLen(d)); i++ {
			k, err := vmmod.Block_At(d, int64(i))
			vm.OnError(err, "Error in json/Array/New")
			switch key := k.(type) {
			case string:
				j.ArrayP(key)
			}
		}
	}
	return e, nil
}

func JsonGetElement(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	vm.Debug("json/Value/Get: %v", e.Type)
	if e.Type == "json" {
		if !vm.CanGet() {
			return nil, fmt.Errorf("Stack is too shallow for json/Value/Get")
		}
		j := e.Value.(*gabs.Container)
		d := vm.Take()
		if d.Type != "str" {
			return nil, fmt.Errorf("string is expected for json/Value/Set")
		}
		if !j.ExistsP(d.Value.(string)) {
			return nil, fmt.Errorf("Key %v not found in JSON", d.Value.(string))
		}
		v := j.Path(d.Value.(string)).Data()
		vm.Put(e)
		switch v.(type) {
		case int:
			return &vmmod.Elem{Type: "int", Value: big.NewInt(int64(v.(int)))}, nil
		case int64:
			return &vmmod.Elem{Type: "int", Value: big.NewInt(v.(int64))}, nil
		case float32:
			return &vmmod.Elem{Type: "flt", Value: float64(v.(float32))}, nil
		case float64:
			return &vmmod.Elem{Type: "flt", Value: float64(v.(float64))}, nil
		case string:
			return &vmmod.Elem{Type: "str", Value: string(v.(string))}, nil
		}
		return nil, fmt.Errorf("Unsupported data in JSON while in json/Value/Get")
	}
	return e, nil
}

func JsonSetElement(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	vm.Debug("json/Value/Set: %v", e.Type)
	if e.Type == "json" {
		if !vm.CanGet() {
			return nil, fmt.Errorf("Stack is too shallow for json/Value/Set")
		}
		j := e.Value.(*gabs.Container)
		d := vm.Take()
		if d.Type != "dblock" {
			return nil, fmt.Errorf("(* ) is expected for json/Value/Set")
		}
		if vmmod.BlockLen(d) == 0 || (vmmod.BlockLen(d)%2) != 0 {
			return nil, fmt.Errorf("(* ) is incorrect arity for json/Value/Set")
		}
		i := 0
		for {
			if i >= int(vmmod.BlockLen(d)) {
				break
			}
			k, err := vmmod.Block_At(d, int64(i))
			vm.OnError(err, "Error in json/Value/Set")
			v, err := vmmod.Block_At(d, int64(i+1))
			vm.OnError(err, "Error in json/Value/Set")
			switch key := k.(type) {
			case string:
				j.SetP(v, key)
			}
			i = i + 2
		}
	}
	return e, nil
}

func JsonNewElement(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	vm.Debug("json/New: %v", e.Type)
	if e.Type == "str" {
		eh, err := vm.GetType("json")
		vm.OnError(err, "error on json/New")
		res := eh.FromString(vm, e.Value.(string))
		return res, nil
	}
	return e, nil
}

func InitJsonFunctions(vm *vmmod.VM) {
	vm.Debug("[ BUND ] bund.InitJsonFunctions() reached")
	vm.AddFunction("json/New", JsonNewElement)
	vm.AddFunction("json/Value/Set", JsonSetElement)
	vm.AddFunction("json/Value/Get", JsonGetElement)
	vm.AddFunction("json/Array/New", JsonArrayNewElement)
	vm.AddFunction("json/Array/Add", JsonArrayAddElement)
	vm.AddFunction("json/Array/Get", JsonArrayGetElement)
}
