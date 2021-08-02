package vm

import (
	"github.com/gammazero/deque"
)

func (vm *VM) CblockMatch(q *deque.Deque, e *Elem) bool {
	for i := 0; i < q.Len(); i++ {
		val := q.At(i).(*Elem)
		switch val.Type {
		case "int", "str", "bool":
			if val.Type != e.Type {
				continue
			}
			eh, err := vm.GetType(val.Type)
			if err != nil {
				vm.OnError(err, "Error in matching block. With matching.")
				return false
			}
			if eh.Compare(vm, val, e) == Eq {
				return true
			}
		case "glob":
			if e.Type == "str" {
				eh, err := vm.GetType("glob")
				if err != nil {
					vm.Error("Matching error: %v", err)
					return false
				}
				if eh.Compare(vm, val, e) == Eq {
					return true
				}
			}
		default:
			return false
		}
	}
	return false
}
