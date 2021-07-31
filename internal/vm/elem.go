package vm

import (
	"fmt"
)

type Elem struct {
	Type  string
	Value interface{}
}

const (
	Eq  = 0
	Gt  = 1
	Ls  = -1
	IDK = -2
	Ne  = 2
)

type FactoryFun func() *Elem
type ToStringFun func(v *VM, e *Elem) string
type FromStringFun func(v *VM, d string) *Elem
type CompareFun func(v *VM, e1 *Elem, e2 *Elem) int
type DupFun func(v *VM, e *Elem) *Elem

type ElemHandler struct {
	Type       string
	ToString   ToStringFun
	FromString FromStringFun
	Compare    CompareFun
	Duplicate  DupFun
}

func (vm *VM) RegisterType(t string, ts ToStringFun, fs FromStringFun, cf CompareFun, df DupFun) bool {
	if _, ok := vm.Types.Load(t); ok {
		return true
	}
	vm.Types.Store(t, &ElemHandler{Type: t, ToString: ts, FromString: fs, Compare: cf, Duplicate: df})
	vm.Debug("Register BUND datatype: %v", t)
	return true
}

func (vm *VM) GetType(t string) (*ElemHandler, error) {
	if res, ok := vm.Types.Load(t); ok {
		return res.(*ElemHandler), nil
	}
	return nil, fmt.Errorf("BUND do not have datatype: %v", t)
}
