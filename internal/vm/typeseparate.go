package vm

func SepFactory(vm *VM) *Elem {
	return &Elem{Type: "SEPARATE", Value: nil}
}

func SepToString(vm *VM, e *Elem) string {
	if e.Type == "SEPARATE" {
		return "|"
	}
	vm.Error("trying to convert a SEPARATE and it is not an SEPARATE: %v %T", e.Type, e.Value)
	return "<?>"
}

func SepFromString(vm *VM, d string) *Elem {
	return &Elem{Type: "SEPARATE", Value: nil}
}

func SepCompare(vm *VM, e1 *Elem, e2 *Elem) int {
	if e1.Type == "SEPARATE" && e2.Type == "SEPARATE" {
		return Eq
	}
	return IDK
}

func SepDup(vm *VM, e *Elem) *Elem {
	return &Elem{Type: "SEPARATE", Value: nil}
}

func RegisterSep(vm *VM) {
	vm.RegisterType("SEPARATE", SepFactory, SepToString, SepFromString, SepCompare, SepDup)
}
