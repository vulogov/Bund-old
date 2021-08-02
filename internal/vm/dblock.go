package vm

import (
	"github.com/gammazero/deque"
)

func DblockSameArity(vm *VM, e1 *Elem, e2 *Elem) bool {
	if e1.Type == "dblock" && e2.Type == "dblock" {
		q1 := e1.Value.(*deque.Deque)
		q2 := e2.Value.(*deque.Deque)
		if q1.Len() != q2.Len() {
			return false
		}
		for i := 0; i < q1.Len(); i++ {
			e1 := q1.At(i).(*Elem)
			e2 := q2.At(i).(*Elem)
			if e1.Type != e2.Type {
				return false
			}
		}
		return true
	}
	return false
}
