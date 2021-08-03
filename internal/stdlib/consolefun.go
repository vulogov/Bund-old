package stdlib

import (
	"fmt"

	vmmod "github.com/vulogov/Bund/internal/vm"
)

func ConsoleDebug(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	e1 := vmmod.DblockSprintf(e)
	if e1 != nil {
		vm.Debug(e1.Value.(string))
	}
	return nil, nil
}

func ConsoleInfo(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	e1 := vmmod.DblockSprintf(e)
	if e1 != nil {
		vm.Info(e1.Value.(string))
	}
	return nil, nil
}

func ConsoleWarning(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	e1 := vmmod.DblockSprintf(e)
	if e1 != nil {
		vm.Warning(e1.Value.(string))
	}
	return nil, nil
}

func ConsoleError(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	e1 := vmmod.DblockSprintf(e)
	if e1 != nil {
		vm.Error(e1.Value.(string))
	}
	return nil, nil
}

func ConsoleFatal(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	e1 := vmmod.DblockSprintf(e)
	if e1 != nil {
		vm.Fatal(e1.Value.(string))
	}
	return nil, nil
}

func ConsolePrint(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	e1 := vmmod.DblockSprintf(e)
	if e1 != nil {
		fmt.Printf(e1.Value.(string))
	}
	return nil, nil
}

func ConsolePrintln(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	e1 := vmmod.DblockSprintf(e)
	if e1 != nil {
		fmt.Println(e1.Value.(string))
	}
	return nil, nil
}

func InitConsoleFunctions(vm *vmmod.VM) {
	vm.Debug("[ BUND ] bund.InitConsoleFunctions() reached")
	vm.AddFunction("console/Debug", ConsoleDebug)
	vm.AddFunction("console/Info", ConsoleInfo)
	vm.AddFunction("console/Warning", ConsoleWarning)
	vm.AddFunction("console/Error", ConsoleError)
	vm.AddFunction("console/Fatal", ConsoleFatal)
	vm.AddFunction("console/Print", ConsolePrint)
	vm.AddFunction("console/Println", ConsolePrintln)
}
