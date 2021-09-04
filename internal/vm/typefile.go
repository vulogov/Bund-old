package vm

import (
	"fmt"

	vfs "github.com/c2fo/vfs/v5"
	"github.com/c2fo/vfs/v5/vfssimple"
)

func FileFactory(vm *VM) *Elem {
	return &Elem{Type: "file", Value: nil}
}

func FileToString(vm *VM, e *Elem) string {
	if e.Type == "file" {
		f := e.Value.(vfs.File)
		return fmt.Sprintf("f'%v'", f.URI())
	}
	vm.Error("trying to convert a file and it is not a file: %v %T", e.Type, e.Value)
	return "<?>"
}

func FileFromString(vm *VM, d string) *Elem {
	f := FileFactory(vm)
	_f, err := vfssimple.NewFile(d)
	vm.OnError(err, "Error in VFS layer")
	f.Value = _f
	return f
}

func FileCompare(vm *VM, e1 *Elem, e2 *Elem) int {
	return IDK
}

func FileDup(vm *VM, e *Elem) *Elem {
	if e.Type == "file" {
		f := e.Value.(vfs.File)
		res := FileFromString(vm, f.URI())
		return res
	}
	return nil
}

func RegisterFile(vm *VM) {
	vm.RegisterType("file", FileFactory, FileToString, FileFromString, FileCompare, FileDup)
}
