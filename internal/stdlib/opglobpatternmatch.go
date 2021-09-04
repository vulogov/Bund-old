package stdlib

import (
	"fmt"

	glob "github.com/ganbarodigital/go_glob"

	"github.com/vulogov/Bund/internal/vm"
)

func GlobPatternMatchOperator(v *vm.VM, e1 *vm.Elem, e2 *vm.Elem) (*vm.Elem, error) {
	return EqOperator(v, e1, e2)
}

func matchshortprefix(v *vm.VM, e1 *vm.Elem, e2 *vm.Elem) bool {
	g := glob.NewGlob(e1.Value.(string))
	_, res, err := g.MatchShortestPrefix(e2.Value.(string))
	if err != nil {
		v.OnError(err, "Error matching short prefix")
	}
	return res
}

func matchshortestsuffix(v *vm.VM, e1 *vm.Elem, e2 *vm.Elem) bool {
	g := glob.NewGlob(e1.Value.(string))
	_, res, err := g.MatchShortestSuffix(e2.Value.(string))
	if err != nil {
		v.OnError(err, "Error matching shortest suffix")
	}
	return res
}

func matchlongestprefix(v *vm.VM, e1 *vm.Elem, e2 *vm.Elem) bool {
	g := glob.NewGlob(e1.Value.(string))
	_, res, err := g.MatchLongestPrefix(e2.Value.(string))
	if err != nil {
		v.OnError(err, "Error matching shortest suffix")
	}
	return res
}

func matchlongestsuffix(v *vm.VM, e1 *vm.Elem, e2 *vm.Elem) bool {
	g := glob.NewGlob(e1.Value.(string))
	_, res, err := g.MatchLongestSuffix(e2.Value.(string))
	if err != nil {
		v.OnError(err, "Error matching shortest suffix")
	}
	return res
}

func rungpo(v *vm.VM, e1 *vm.Elem, e2 *vm.Elem, gpofun func(*vm.VM, *vm.Elem, *vm.Elem) bool) (*vm.Elem, error) {
	if e1.Type != "glob" {
		return nil, fmt.Errorf("Pattern for matching not found")
	}
	if e2.Type != "str" {
		return nil, fmt.Errorf("Pattern matching is done only with strings")
	}
	res := gpofun(v, e1, e2)
	if res == true {
		return &vm.Elem{Type: "bool", Value: true}, nil
	} else {
		return &vm.Elem{Type: "bool", Value: false}, nil
	}
}

func GlobPatternMatchPrefixOperator(v *vm.VM, e1 *vm.Elem, e2 *vm.Elem) (*vm.Elem, error) {
	return rungpo(v, e1, e2, matchshortprefix)
}

func GlobPatternMatchSuffixOperator(v *vm.VM, e1 *vm.Elem, e2 *vm.Elem) (*vm.Elem, error) {
	return rungpo(v, e1, e2, matchshortestsuffix)
}

func GlobPatternMatchLongestPrefixOperator(v *vm.VM, e1 *vm.Elem, e2 *vm.Elem) (*vm.Elem, error) {
	return rungpo(v, e1, e2, matchlongestprefix)
}

func GlobPatternMatchLongestSuffixOperator(v *vm.VM, e1 *vm.Elem, e2 *vm.Elem) (*vm.Elem, error) {
	return rungpo(v, e1, e2, matchlongestsuffix)
}

func InitGPMOperators(vm *vm.VM) {
	vm.Debug("[ BUND ] bund.InitGPMOperators() reached")
	vm.AddOperator("===", GlobPatternMatchOperator)
	vm.AddOperator("<===", GlobPatternMatchPrefixOperator)
	vm.AddOperator("===>", GlobPatternMatchSuffixOperator)
	vm.AddOperator("=<==", GlobPatternMatchLongestPrefixOperator)
	vm.AddOperator("==>=", GlobPatternMatchLongestSuffixOperator)
}
