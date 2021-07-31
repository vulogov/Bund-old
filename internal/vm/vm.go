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
}

func NewVM(name string) *VM {
	res := VM{Name: name, Current: nil, CurrentNS: nil, CurrentElem: nil, Mode: true}
	res.IsIgnore.PushBack(false)
	res.Debug("Creating BUND VM: %s", name)
	res.RegisterOpcodes()
	return &res
}
