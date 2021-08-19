package stdlib

import (
	"strings"

	"github.com/gammazero/deque"

	vmmod "github.com/vulogov/Bund/internal/vm"
)

func StringTrimElement(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	vm.Debug("str/Strip: %v", e.Type)
	if e.Type == "str" {
		res := strings.TrimSpace(e.Value.(string))
		return &vmmod.Elem{Type: "str", Value: res}, nil
	}
	return e, nil
}

func StringLinesElement(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	vm.Debug("str/Lines: %v", e.Type)
	if e.Type == "str" {
		_res := strings.Split(e.Value.(string), "\n")
		eh, err := vm.GetType("dblock")
		vm.OnError(err, "Error in str/Lines")
		res := eh.Factory(vm)
		q := res.Value.(*deque.Deque)
		for i := 0; i < len(_res); i++ {
			q.PushBack(&vmmod.Elem{Type: "str", Value: strings.TrimSpace(_res[i])})
		}
		return res, nil
	}
	return e, nil
}

func InitStringFunctions(vm *vmmod.VM) {
	vm.Debug("[ BUND ] bund.InitStringFunctions() reached")
	vm.AddFunction("str/Strip", StringTrimElement)
	vm.AddFunction("str/Lines", StringLinesElement)
}
