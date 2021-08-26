package stdlib

import (
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/gammazero/deque"

	vmmod "github.com/vulogov/Bund/internal/vm"
)

func PassthrougElement(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	vm.Debug("PASSTHROUGH: %v", e.Type)
	return e, nil
}

func RotateFrontElement(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	vm.Put(e)
	q := vm.Current
	q.Rotate(1)
	return nil, nil
}

func RotateBackElement(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	vm.Put(e)
	q := vm.Current
	q.Rotate(-1)
	return nil, nil
}

func PrintStack(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	vm.Debug("PRINTSTACK: %v", e.Type)
	vm.Put(e)
	res := fmt.Sprintf("(%v)[ ", vm.CurrentNS.Name)
	for i := vm.Current.Len() - 1; i >= 0; i-- {
		_e := vm.Current.At(i)
		eh, err := vm.GetType(_e.(*vmmod.Elem).Type)
		if err != nil {
			vm.Error("Can not get type: %v", _e.(*vmmod.Elem).Type)
			continue
		}
		res = res + eh.ToString(vm, _e.(*vmmod.Elem))
		res = res + ", "
	}
	res = res + " ]"
	fmt.Println(res)
	return nil, nil
}

func DropElement(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	vm.Debug("DROP: %v", e.Type)
	return nil, nil
}

func DropOppositeElement(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	if vm.Current.Len() > 0 {
		vm.Debug("DROP OPPOSITE: %v", e.Type)
		vm.TakeOpposite()
	}
	return e, nil
}

func DupElement(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	eh, err := vm.GetType(e.Type)
	if err != nil {
		return nil, err
	}
	vm.Put(e)
	res := eh.Duplicate(vm, e)
	vm.Debug("DUP: %v %v", e.Type, eh.ToString(vm, e))
	return res, nil
}

func ExecuteElement(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	vm.Debug("EXECUTE: %v", e.Type)
	if e.Type == "CALL" {
		res, err := vm.Exec(e.Value.(string))
		return res, err
	} else if e.Type == "OP" {
		res, err := vm.Op(e.Value.(string))
		return res, err
	}
	return nil, fmt.Errorf("Request to EXECUTE on wrong context: %v", e.Type)
}

func SetAlias(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	if vm.Current.Len() < 1 {
		return nil, fmt.Errorf("Namespace stack is too shallow for ALIAS operation: %v:%v", vm.CurrentNS.Name, vm.Current.Len())
	}
	if e.Type != "str" {
		return nil, fmt.Errorf("String value is required for an ALIAS operation: %v", e.Type)
	}
	e2 := vm.Take()
	vm.CurrentNS.SetAlias(e.Value.(string), e2)
	return nil, nil
}

func GetAlias(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	if e.Type != "str" {
		return nil, fmt.Errorf("String value is required for an ALIAS operation: %v", e.Type)
	}
	val := vm.CurrentNS.GetAlias(e.Value.(string))
	if val == nil {
		vm.Error("There is no alias for: %v", e.Value.(string))
		return nil, nil
	}
	return val.(*vmmod.Elem), nil
}

func RefElement(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	if e.Type == "REF" {
		switch v := e.Value.(type) {
		case *vmmod.Elem:
			vm.Debug("Dereferencing %v", v.Type)
			return v, nil
		}
		return nil, fmt.Errorf("REF: Attempt of de-referencing the data, but there is no data there")
	} else {
		vm.Debug("Making REF for %v", e.Type)
		return &vmmod.Elem{Type: "REF", Value: e}, nil
	}
}

func SleepFun(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	var d int
	switch e.Type {
	case "int":
		res := big.NewInt(0)
		d = int(res.Mul(e.Value.(*big.Int), big.NewInt(1000000000)).Int64())
	case "flt":
		d = int(e.Value.(float64) * float64(1000000000))
	default:
		return nil, fmt.Errorf("Parameter for 'sleep' is not float or int: %v", e.Type)
	}
	vm.Debug("Sleeping for a %v 𝝶-seconds", d)
	time.Sleep(time.Duration(d) * time.Nanosecond)
	return e, nil
}

func ExitFun(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	if e.Type != "int" {
		return nil, fmt.Errorf("Exit function require a number in stack: %v", e.Type)
	}
	res := int(e.Value.(*big.Int).Int64())
	vm.Debug("EXIT is called with %v", res)
	os.Exit(res)
	return nil, nil
}

func NameFun(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	vm.Put(e)
	return &vmmod.Elem{Type: "str", Value: vm.CurrentNS.Name}, nil
}

func TypeFun(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	vm.Put(e)
	return &vmmod.Elem{Type: "str", Value: e.Type}, nil
}

func SeqFun(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	if e.Type == "int" {
		eh, err := vm.GetType("dblock")
		vm.OnError(err, "Error in 'seq'")
		res := eh.Factory(vm)
		q := res.Value.(*deque.Deque)
		c := int(e.Value.(*big.Int).Int64())
		for i := 0; i < c; i++ {
			q.PushBack(&vmmod.Elem{Type: "int", Value: big.NewInt(int64(i))})
		}
		return res, nil
	}
	return nil, fmt.Errorf("Exit function require a number in stack: %v", e.Type)
}

func InitSystemFunctions(vm *vmmod.VM) {
	vm.Debug("[ BUND ] bund.InitSystemFunctions() reached")
	vm.AddFunction("passthrough", PassthrougElement)
	vm.AddFunction("dumpstack", PrintStack)
	vm.AddFunction(",", DropElement)
	vm.AddFunction("_,", DropElement)
	vm.AddFunction(",_", DropOppositeElement)
	vm.AddFunction("^", DupElement)
	vm.AddFunction("!", ExecuteElement)
	vm.AddFunction("setAlias", SetAlias)
	vm.AddFunction("≡", SetAlias)
	vm.AddFunction("alias", GetAlias)
	vm.AddFunction("#", RefElement)
	vm.AddFunction("sleep", SleepFun)
	vm.AddFunction("name", NameFun)
	vm.AddFunction("type", TypeFun)
	vm.AddFunction("seq", SeqFun)
	vm.AddFunction("exit", ExitFun)
	vm.AddFunction("rotateFront", RotateFrontElement)
	vm.AddFunction("rotate", RotateFrontElement)
	vm.AddFunction("rotateBack", RotateBackElement)
}
