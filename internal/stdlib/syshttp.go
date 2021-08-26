package stdlib

import (
	"fmt"

	// "github.com/gammazero/deque"
	// "github.com/levigross/grequests"
	vmmod "github.com/vulogov/Bund/internal/vm"
)

func HttpGetElement(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	if e.Type == "str" {
	}
	return nil, fmt.Errorf("")
}

func HttpPostElement(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	if e.Type == "str" {
	}
	return nil, fmt.Errorf("Expecting an index as int for args/Slice/Len")
}

func InitHttpFunctions(vm *vmmod.VM) {
	vm.Debug("[ BUND ] bund.InitHttpFunctions() reached")
	vm.AddFunction("http/Get", HttpGetElement)
	vm.AddFunction("http/Post", HttpPostElement)
}
