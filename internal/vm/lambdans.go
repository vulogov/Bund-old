package vm

import (
	"github.com/gammazero/deque"
)

func (ns *NS) HasLambda(name string) bool {
	if _, ok := ns.Fun.Load(name); ok {
		return true
	}
	return false
}

func (ns *NS) GetLambda(name string) *deque.Deque {
	var res *deque.Deque
	if _res, ok := ns.Fun.Load(name); ok {
		ns.VM.Debug("Returning LAMBDA from %v: %v", ns.Name, name)
		res = _res.(*deque.Deque)
	} else {
		ns.VM.Debug("Creating LAMBDA in %v: %v", ns.Name, name)
		res = new(deque.Deque)
		ns.Fun.Store(name, res)
	}
	return res
}

func (ns *NS) InLambda(name string) bool {
	if _, ok := ns.Fun.Load(name); ok {
		ns.LambdasStack.PushBack(name)
		ns.VM.Debug("We are going in Lambda(%v)", name)
		return true
	}
	ns.VM.Error("Attempt to go in Lambda(%v) failed", name)
	return false
}

func (ns *NS) NameOfCurrentLambda() string {
	if ns.LambdasStack.Len() < 1 {
		ns.VM.Error("Attempt to get Lambda function name on empty Lambdas stack")
		return ""
	}
	return ns.LambdasStack.Back().(string)
}

func (ns *NS) CurrentLambda() *deque.Deque {
	var res *deque.Deque
	if ns.LambdasStack.Len() < 1 {
		ns.VM.Error("Attempt to select Lambda function on empty Lambdas stack")
		return nil
	}
	name := ns.LambdasStack.Back().(string)
	if _res, ok := ns.Fun.Load(name); ok {
		ns.VM.Debug("Returning LAMBDA from %v: %v", ns.Name, name)
		res = _res.(*deque.Deque)
		return res
	}
	ns.VM.Error("Something is seriously wrong, current lambda %v not found", name)
	return nil
}

func (ns *NS) CloseLambda() bool {
	if ns.LambdasStack.Len() < 1 {
		ns.VM.Error("Attempt to close Lambda fuction on empty Lambdas stack")
		return false
	}
	ln := ns.LambdasStack.PopBack().(string)
	ns.VM.Debug("Closing lambda %v. Stack size: %v", ln, ns.LambdasStack.Len())
	return true
}
