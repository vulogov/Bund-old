package stdlib

import (
	"fmt"

	vfs "github.com/c2fo/vfs/v5"

	vmmod "github.com/vulogov/Bund/internal/vm"
)

func VfsFileElement(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	if e.Type == "str" {
		eh, err := vm.GetType("file")
		vm.OnError(err, "Error in fs/File")
		res := eh.FromString(vm, e.Value.(string))
		return res, nil
	}
	return nil, fmt.Errorf("Expecting to have a str value with URL, but have this: %v", e.Type)
}

func VfsFileExistsElement(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	if e.Type == "file" && e.Value != nil {
		f := e.Value.(vfs.File)
		vm.Put(e)
		res, err := f.Exists()
		if err != nil {
			return nil, err
		}
		return &vmmod.Elem{Type: "bool", Value: res}, nil
	}
	return nil, fmt.Errorf("Expecting to have a str value with URL, but have this: %v", e.Type)
}

func InitVFSFunctions(vm *vmmod.VM) {
	vm.Debug("[ BUND ] bund.InitVFSFunctions() reached")
	vm.AddFunction("fs/File", VfsFileElement)
	vm.AddFunction("fs/File/Exists", VfsFileExistsElement)

}
