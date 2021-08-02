package vm

import (
	"unsafe"

	"github.com/cstockton/go-conv"
)

func StringDeepCopy(s string) string {
	b := make([]byte, len(s))
	copy(b, s)
	return *(*string)(unsafe.Pointer(&b))
}

func StringFactory(vm *VM) *Elem {
	return &Elem{Type: "str", Value: ""}
}

func StringToString(vm *VM, e *Elem) string {
	if e.Type == "str" {
		switch v := e.Value.(type) {
		case string:
			return string(v)
		}
	}
	vm.Error("trying to convert a String and it is not a String: %v %T", e.Type, e.Value)
	return "<?>"
}

func StringFromString(vm *VM, d string) *Elem {
	res, err := conv.String(d)
	if err != nil {
		return nil
	}
	return &Elem{Type: "str", Value: res}
}

func StringCompare(vm *VM, e1 *Elem, e2 *Elem) int {
	if e1.Type == "str" && e2.Type == "str" {
		r1 := e1.Value.(string)
		r2 := e2.Value.(string)
		if r1 == r2 {
			return Eq
		} else if r1 > r2 {
			return Gt
		} else {
			return Ls
		}
	}
	return IDK
}

func StringDup(vm *VM, e *Elem) *Elem {
	if e.Type != "str" {
		return nil
	}
	return &Elem{Type: "str", Value: StringDeepCopy(e.Value.(string))}
}

func RegisterString(vm *VM) {
	vm.RegisterType("str", StringFactory, StringToString, StringFromString, StringCompare, StringDup)
}
