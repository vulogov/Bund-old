package vm

import (
	"github.com/gammazero/deque"
	"github.com/lrita/cmap"
)

type VM struct {
	Name        string
	Mode        bool
	IsIgnore    deque.Deque
	NSStack     deque.Deque
	NS          cmap.Cmap
	Current     *deque.Deque
	CurrentNS   *NS
	CurrentElem *Elem
	RootNS      *NS
	Options     cmap.Cmap
	Opcodes     cmap.Cmap
	Types       cmap.Cmap
	Functions   cmap.Cmap
	Operators   cmap.Cmap
	Sys         cmap.Cmap
	Aliases     cmap.Cmap
}

func NewVM(name string) *VM {
	res := VM{Name: name, Current: nil, CurrentNS: nil, CurrentElem: nil, Mode: true}
	res.IsIgnore.PushBack(false)
	res.Debug("Creating BUND VM: %s", name)
	res.RegisterTypes()
	res.RegisterOpcodes()
	return &res
}

func (vm *VM) IsStack() bool {
	if vm.CurrentNS == nil {
		return false
	}
	if vm.Current == nil {
		return false
	}
	return true
}

func (vm *VM) CanGet() bool {
	if !vm.IsStack() {
		return false
	}
	if vm.Current.Len() == 0 {
		return false
	}
	return true
}
