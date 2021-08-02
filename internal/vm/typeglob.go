package vm

import (
	"fmt"

	"github.com/cstockton/go-conv"
	glob "github.com/ganbarodigital/go_glob"
	"github.com/pieterclaerhout/go-log"
)

func GlobFactory(vm *VM) *Elem {
	return &Elem{Type: "glob", Value: "*"}
}

func GlobToString(vm *VM, e *Elem) string {
	if e.Type == "glob" {
		switch v := e.Value.(type) {
		case string:
			return fmt.Sprintf("g'%v'", string(v))
		}
	}
	log.Errorf("trying to convert a Glob and it is not a Glob: %v %T", e.Type, e.Value)
	return "<?>"
}

func GlobFromString(vm *VM, d string) *Elem {
	res, err := conv.String(d)
	if err != nil {
		return nil
	}
	return &Elem{Type: "glob", Value: res}
}

func GlobCompare(vm *VM, e1 *Elem, e2 *Elem) int {
	if e1.Type == "glob" && e2.Type == "str" {
		r1 := e1.Value.(string)
		r2 := e2.Value.(string)
		g := glob.NewGlob(r1)
		if res, err := g.Match(r2); err == nil && res == true {
			return Eq
		}
		return Ne
	}
	return IDK
}

func GlobDup(vm *VM, e *Elem) *Elem {
	if e.Type != "glob" {
		return nil
	}
	return &Elem{Type: "glob", Value: StringDeepCopy(e.Value.(string))}
}

func RegisterGlob(vm *VM) {
	vm.RegisterType("glob", GlobFactory, GlobToString, GlobFromString, GlobCompare, GlobDup)
}
