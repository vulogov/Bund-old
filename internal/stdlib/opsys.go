package stdlib

import (
	"fmt"

	vmmod "github.com/vulogov/Bund/internal/vm"
)

func SwapOperator(vm *vmmod.VM, e1 *vmmod.Elem, e2 *vmmod.Elem) (*vmmod.Elem, error) {
	vm.Put(e1)
	return e2, nil
}

func UseOperator(vm *vmmod.VM, e1 *vmmod.Elem, e2 *vmmod.Elem) (*vmmod.Elem, error) {
	if e1.Type == "str" && e2.Type == "str" {
		vm.Debug("Locating %v in %v", e2.Value.(string), e1.Value.(string))
		if !vm.HasUserFunctionInNS(e2.Value.(string), e1.Value.(string)) {
			return nil, fmt.Errorf("Function %v at %v not found", e2.Value.(string), e1.Value.(string))
		}
		n := vm.AsNS(e1.Value.(string))
		if n == nil {
			return nil, fmt.Errorf("namespace %v not found", e1.Value.(string))
		}
		c := n.GetLambda(e2.Value.(string))
		nsname := vm.CurrentNS.Name
		res := vm.SetUserFunctionInNS(nsname, e2.Value.(string), c)
		if !res {
			return nil, fmt.Errorf("Error registering function %v %v", nsname, e2.Value.(string))
		}
		return &vmmod.Elem{Type: "CALL", Value: e2.Value.(string)}, nil
	}
	return nil, fmt.Errorf("Expecting strings in stack for -> operator")
}

func InitOpSys(vm *vmmod.VM) {
	vm.Debug("[ BUND ] bund.InitOpSys() reached")
	vm.AddOperator("<=>", SwapOperator)
	vm.AddOperator("->", UseOperator)
	vm.AddOperator("→", UseOperator)
}
