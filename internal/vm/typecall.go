package vm

import (
	"fmt"

	"github.com/cstockton/go-conv"
)

func CallFactory(vm *VM) *Elem {
	return &Elem{Type: "CALL", Value: ""}
}

func CallToString(vm *VM, e *Elem) string {
	if e.Type == "CALL" {
		switch v := e.Value.(type) {
		case string:
			return fmt.Sprintf("[%v]", string(v))
		}
	}
	vm.Error("trying to convert a CALL and it is not a CALL: %v %T", e.Type, e.Value)
	return "<?>"
}

func CallFromString(vm *VM, d string) *Elem {
	res, err := conv.String(d)
	if err != nil {
		return nil
	}
	return &Elem{Type: "CALL", Value: res}
}

func CallCompare(vm *VM, e1 *Elem, e2 *Elem) int {
	return IDK
}

func CallDup(vm *VM, e *Elem) *Elem {
	if e.Type != "CALL" {
		return nil
	}
	return &Elem{Type: "CALL", Value: StringDeepCopy(e.Value.(string))}
}

func RegisterCALL(vm *VM) {
	vm.RegisterType("CALL", CallFactory, CallToString, CallFromString, CallCompare, CallDup)
}
