package vm

import (
	"fmt"

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

func DblockSame(vm *VM, e1 *Elem, e2 *Elem) bool {
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
			eh, err := vm.GetType(e1.Type)
			if err != nil {
				return false
			}
			if eh.Compare(vm, e1, e2) != Eq {
				return false
			}
		}
		return true
	}
	return false
}

func DblockSprintf(e *Elem) *Elem {
	var args []interface{}
	if e.Type == "dblock" {
		q := e.Value.(*deque.Deque)
		if q.Len() > 0 {
			emsg := q.PopBack().(*Elem)
			if emsg.Type == "str" {
				msg := emsg.Value.(string)
				for q.Len() > 0 {
					e1 := q.PopBack().(*Elem)
					args = append(args, e1.Value)
				}
				return &Elem{Type: "str", Value: fmt.Sprintf(msg, args...)}
			}
		}
	}
	return nil
}
