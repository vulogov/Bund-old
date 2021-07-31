package vm

import (
	"github.com/lrita/cmap"
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

var BundTypes cmap.Cmap
