package vm

import (
	"github.com/gammazero/deque"
	"github.com/lrita/cmap"
)

type NS struct {
	Name         string
	VM           *VM
	Stack        deque.Deque
	Fun          cmap.Cmap
	Options      cmap.Cmap
	LambdasStack deque.Deque
}

func NewNS(v *VM, name string) *NS {
	v.Debug("Creating NAMESPACE: %v", name)
	return &NS{Name: name, VM: v}
}
