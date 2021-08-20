package vm

import (
	"fmt"
	"math/big"

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

func BlockLen(d *Elem) int64 {
	switch d.Type {
	case "dblock":
		q := d.Value.(*deque.Deque)
		return int64(q.Len())
	}
	return 0
}

func BlockAt(d *Elem, t string, n int64) (interface{}, error) {
	if n >= BlockLen(d) {
		return nil, fmt.Errorf("Index %v is out of bound", n)
	}
	switch d.Type {
	case "dblock":
		q := d.Value.(*deque.Deque)
		v := q.At(int(n)).(*Elem)
		if v.Type != t {
			return nil, fmt.Errorf("Invalid type at BlockAt %v <> %v", t, v.Type)
		}
		switch v.Type {
		case "int":
			return v.Value.(*big.Int).Int64(), nil
		}
		return v.Value, nil
	}
	return nil, fmt.Errorf("I do not know how to extract At value from %v", d.Type)
}

func Block_At(d *Elem, n int64) (interface{}, error) {
	if n >= BlockLen(d) {
		return nil, fmt.Errorf("Index %v is out of bound", n)
	}
	switch d.Type {
	case "dblock":
		q := d.Value.(*deque.Deque)
		v := q.At(int(n)).(*Elem)
		switch v.Type {
		case "int":
			return v.Value.(*big.Int).Int64(), nil
		}
		return v.Value, nil
	}
	return nil, fmt.Errorf("I do not know how to extract At value from %v", d.Type)
}
