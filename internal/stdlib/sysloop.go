package stdlib

import (
	"fmt"

	"github.com/gammazero/deque"

	"github.com/vulogov/Bund/internal/signal"
	vmmod "github.com/vulogov/Bund/internal/vm"
)

func RealLoopElement(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	if e.Type == "CALL" {
		for {
			vm.Debug("is Exit requested: %v", signal.ExitRequested())
			if signal.ExitRequested() {
				vm.Debug("Exit requested in 'loop' loop")
				break
			}
			vmmod.EvalCmd(vm, e)
		}
		vm.Debug("main thread also requesting an exit in 'loop' loop")
		signal.ExitRequest()
	} else {
		return nil, fmt.Errorf("For 'IN' loop, we expecting CALL on top of the stack, not this: %v", e.Type)
	}
	return &vmmod.Elem{Type: "bool", Value: true}, nil
}

func LoopElement(vm *vmmod.VM, e *vmmod.Elem, criteria bool) (*vmmod.Elem, error) {
	if e.Type == "CALL" {
		for {
			vmmod.EvalCmd(vm, e)
			if vm.CanGet() {
				res := vm.Take()
				if res.Type == "bool" {
					if res.Value.(bool) == criteria {
						return nil, nil
					}
				} else {
					return res, nil
				}
			}
		}
	} else {
		return nil, fmt.Errorf("For 'IN' loop, we expecting CALL on top of the stack, not this: %v", e.Type)
	}
	return nil, nil
}

func WhileElement(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	return LoopElement(vm, e, false)
}

func UntilElement(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	return LoopElement(vm, e, true)
}

func InElement(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	if e.Type == "dblock" {
		fmt.Println(e)
		vm.Debug("IN LOOP: %v:%v", e.Name, e.Type)
		q := e.Value.(*deque.Deque)
		for i := 0; i < q.Len(); i++ {
			val := q.At(i).(*vmmod.Elem)
			vm.Debug("'IN' %v", val)
			switch val.Type {
			case "int", "flt", "str", "bool", "cpx", "glob":
				vm.Put(val)
			case "CALL":
				ns := vm.AsNS(e.Name)
				if ns == nil {
					return nil, fmt.Errorf("We can not find NS %v for CALL", val.Name)
				}
				if ns.HasLambda(val.Value.(string)) {
					err := vmmod.ApplyInNS(val.Value.(string), ns, vm)
					if err != nil {
						return nil, err
					}
				} else {
					vmmod.EvalCmd(vm, val)
				}
			default:
				vmmod.EvalCmd(vm, val)
			}
		}
	} else {
		return nil, fmt.Errorf("For 'IN' loop, we expecting DBLOCK on top of the stack, not this: %v", e.Type)
	}
	return nil, nil
}

func InitLoopFunctions(vm *vmmod.VM) {
	vm.Debug("[ BUND ] bund.InitLoopFunctions() reached")
	vm.AddFunction("in", InElement)
	vm.AddFunction("while", WhileElement)
	vm.AddFunction("until", UntilElement)
	vm.AddFunction("loop", RealLoopElement)
}
