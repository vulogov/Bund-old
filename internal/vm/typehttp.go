package vm

import (
	"fmt"

	"github.com/levigross/grequests"
)

type HTTP struct {
	URL string
	RO  *grequests.RequestOptions
}

func HttpFactory(vm *VM) *Elem {
	return &Elem{Type: "http", Value: &HTTP{URL: "http://127.0.0.1", RO: &grequests.RequestOptions{
		InsecureSkipVerify: true,
		IsAjax:             false,
	}}}
}

func HttpToString(vm *VM, e *Elem) string {
	if e.Type == "http" {
		return fmt.Sprintf("%v", e.Value.(*HTTP).URL)
	}
	vm.Error("trying to convert a http and it is not a http: %v %T", e.Type, e.Value)
	return "<?>"
}

func HttpFromString(vm *VM, d string) *Elem {
	vm.Debug("Making HTTP object from %v", d)
	ro := &grequests.RequestOptions{
		InsecureSkipVerify: true,
		IsAjax:             false,
	}
	return &Elem{Type: "http", Value: &HTTP{URL: d, RO: ro}}
}

func HttpCompare(vm *VM, e1 *Elem, e2 *Elem) int {
	if e1.Type == "http" && e2.Type == "http" {
		r1 := e1.Value.(*HTTP).URL
		r2 := e2.Value.(*HTTP).URL
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

func HttpDup(vm *VM, e *Elem) *Elem {
	if e.Type == "http" {
		res := HttpFactory(vm)
		res.Value.(*HTTP).URL = StringDeepCopy(e.Value.(*HTTP).URL)
		return res
	}
	return nil
}

func RegisterHttp(vm *VM) {
	vm.RegisterType("http", HttpFactory, HttpToString, HttpFromString, HttpCompare, HttpDup)
}
