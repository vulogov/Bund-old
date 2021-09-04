package vm

import (
	"fmt"

	gabs "github.com/Jeffail/gabs/v2"
)

func JsonFactory(vm *VM) *Elem {
	return &Elem{Type: "json", Value: gabs.New()}
}

func JsonToString(vm *VM, e *Elem) string {
	if e.Type == "json" {
		j := e.Value.(*gabs.Container)
		return fmt.Sprintf("%v", j.String())
	}
	vm.Error("trying to convert a json and it is not a json: %v %T", e.Type, e.Value)
	return "<?>"
}

func JsonFromString(vm *VM, d string) *Elem {
	vm.Debug("Making JSON object from %v", d)
	j, err := gabs.ParseJSON([]byte(d))
	vm.OnError(err, "Error converting JSON from string")
	return &Elem{Type: "json", Value: j}
}

func JsonCompare(vm *VM, e1 *Elem, e2 *Elem) int {
	return IDK
}

func JsonDup(vm *VM, e *Elem) *Elem {
	if e.Type == "json" {
		return &Elem{Type: "json", Value: JsonFromString(vm, JsonToString(vm, e))}
	}
	return nil
}

func RegisterJson(vm *VM) {
	vm.RegisterType("json", JsonFactory, JsonToString, JsonFromString, JsonCompare, JsonDup)
}
