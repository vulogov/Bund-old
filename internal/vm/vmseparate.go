package vm

import (
	"github.com/gammazero/deque"
	"github.com/pieterclaerhout/go-log"
)

func mergeBlocksWithSeparator(e1 *Elem, e2 *Elem, dir bool) {
	log.Debugf("SEPARATING %v and %v", e1.Type, e2.Type)
	switch e1.Type {
	case "dblock":
		switch e2.Type {
		case "dblock":
			q1 := e1.Value.(*deque.Deque)
			q2 := e2.Value.(*deque.Deque)
			if dir {
				q1.PushBack(&Elem{Type: "SEPARATE", Value: nil})
			}
			for q2.Len() > 0 {
				e := q2.PopFront()
				q1.PushBack(e)
			}
			if !dir {
				q1.PushBack(&Elem{Type: "SEPARATE", Value: nil})
			}
			return
		case "int", "uint", "flt", "bool", "str":
			q1 := e1.Value.(*deque.Deque)
			if dir {
				q1.PushBack(&Elem{Type: "SEPARATE", Value: nil})
			}
			q1.PushBack(e2)
			if !dir {
				q1.PushBack(&Elem{Type: "SEPARATE", Value: nil})
			}
			return
		}
		log.Errorf("I can not merge %v to dblock", e2.Type)
		return
	}
	log.Errorf("%v is not mergeable", e1.Type)
}

func (vm *VM) Separate() *Elem {
	var e1, e2 *Elem
	if vm.Current.Len() < 1 {
		vm.Fatal("Your stack for the | command is too shallow")
		return nil
	}
	deh, err := vm.GetType("dblock")
	if err != nil {
		vm.OnError(err, "Error in applying SEPARATE command")
	}
	res := deh.Factory(vm)
	q := res.Value.(*deque.Deque)
	e1 = vm.Take()
	if vm.Current.Len() >= 1 {
		e2 = vm.Take()
	} else {
		switch e1.Type {
		case "dblock":
			q1 := e1.Value.(*deque.Deque)
			q1.PushBack(&Elem{Type: "SEPARATE", Value: nil})
			return e1
		}
		q.PushBack(e1)
		q.PushBack(&Elem{Type: "SEPARATE", Value: nil})
		return res
	}
	if e1.Type == e2.Type {
		switch e1.Type {
		case "int", "uint", "flt", "bool", "str":
			q.PushBack(e2)
			q.PushBack(&Elem{Type: "SEPARATE", Value: nil})
			q.PushBack(e1)
			vm.Debug("Merging two values to dblock")
			return res
		case "dblock":
			vm.Debug("Merging two dblocks to dblock")
			mergeBlocksWithSeparator(e2, e1, false)
			return e2
		}
	} else {
		if e1.Type == "dblock" {
			mergeBlocksWithSeparator(e1, e2, true)
			return e1
		} else if e2.Type == "dblock" {
			mergeBlocksWithSeparator(e2, e1, false)
			return e2
		}
	}
	vm.Fatal("Merge failed. Stack uncertain")
	return nil
}
