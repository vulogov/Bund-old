package vm

import (
	"github.com/gammazero/deque"
)

func DblockFactory(vm *VM) *Elem {
	return &Elem{Type: "dblock", Value: new(deque.Deque)}
}

func DblockToString(vm *VM, e *Elem) string {
	if e.Type == "dblock" {
		res := "(* "
		for i := e.Value.(*deque.Deque).Len() - 1; i >= 0; i-- {
			_e := e.Value.(*deque.Deque).At(i)
			eh, err := vm.GetType(_e.(*Elem).Type)
			if err != nil {
				vm.Error("Can not get type: %v", _e.(*Elem).Type)
				continue
			}
			res = res + eh.ToString(vm, _e.(*Elem))
			res = res + " "
		}
		res = res + " )"
		return res
	}
	vm.Error("trying to convert an Dblock and it is not an Dblock: %v %T", e.Type, e.Value)
	return "<?>"
}

func DblockFromString(vm *VM, d string) *Elem {
	res := new(deque.Deque)
	return &Elem{Type: "dblock", Value: &res}
}

func DblockCompare(vm *VM, e1 *Elem, e2 *Elem) int {
	if DblockSame(vm, e1, e2) {
		return Eq
	}
	return Ne
}

func DblockDup(vm *VM, e *Elem) *Elem {
	if e.Type != "dblock" {
		return nil
	}
	res := DblockFactory(vm)
	q1 := e.Value.(*deque.Deque)
	q2 := res.Value.(*deque.Deque)
	for i := 0; i < q1.Len(); i++ {
		e := q1.At(i).(*Elem)
		q2.PushBack(e)
	}
	return res
}

func RegisterDblock(vm *VM) {
	vm.RegisterType("dblock", DblockFactory, DblockToString, DblockFromString, DblockCompare, DblockDup)
}
