package stdlib

import (
	"fmt"

	"github.com/gammazero/deque"

	vmmod "github.com/vulogov/Bund/internal/vm"
)

func mergeQueue(q *deque.Deque, e1 *vmmod.Elem, e2 *vmmod.Elem) {
	q1 := e1.Value.(*deque.Deque)
	q2 := e2.Value.(*deque.Deque)
	for i := 0; i < q2.Len(); i++ {
		q.PushBack(q2.At(i).(*vmmod.Elem))
	}
	for i := 0; i < q1.Len(); i++ {
		q.PushBack(q1.At(i).(*vmmod.Elem))
	}

}

func ZipOperator(vm *vmmod.VM, e1 *vmmod.Elem, e2 *vmmod.Elem) (*vmmod.Elem, error) {
	if e1.Type == e2.Type {
		eh, err := vm.GetType("dblock")
		vm.OnError(err, "Error in ++ operator")
		e := eh.Factory(vm)
		q := e.Value.(*deque.Deque)
		switch e1.Type {
		case "int", "flt", "bool", "str":
			q.PushFront(e1)
			q.PushFront(e2)
			return e, nil
		case "dblock":
			mergeQueue(q, e1, e2)
			return e, nil
		}
	} else {
		switch e1.Type {
		case "dblock":
			switch e2.Type {
			case "int", "flt", "bool", "str":
				e1.Value.(*deque.Deque).PushBack(e2)
				return e1, nil
			default:
				return nil, fmt.Errorf("Non-mergeable types for '++': %v %v", e1.Type, e2.Type)
			}
		case "int", "flt", "bool", "str":
			switch e2.Type {
			case "dblock":
				e2.Value.(*deque.Deque).PushBack(e1)
				return e2, nil
			default:
				return nil, fmt.Errorf("Non-mergeable types for '++': %v %v", e1.Type, e2.Type)
			}
		case "json":
			switch e2.Type {
			case "dblock":
				vm.Put(e2)
				return JsonSetElement(vm, e1)
			}
		default:
			return nil, fmt.Errorf("I finally do not know how to perform '++' for this data: %v %v", e1.Type, e2.Type)
		}
		return e1, nil
	}
	return nil, fmt.Errorf("I am really and finally do not know how to perform '++' for this data: %v %v", e1.Type, e2.Type)
}

func InitOpZip(vm *vmmod.VM) {
	vm.Debug("[ BUND ] bund.InitZipOperators() reached")
	vm.AddOperator("++", ZipOperator)
}
