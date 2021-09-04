package stdlib

import (
	"fmt"

	"net/url"

	"github.com/levigross/grequests"

	vmmod "github.com/vulogov/Bund/internal/vm"
)

func HttpSetProxyElement(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	if e.Type != "http" {
		if !vm.CanGet() {
			return nil, fmt.Errorf("Stack is too shallow for http/Set/Proxy")
		}
		e1 := vm.Take()
		if e1.Type != "str" {
			fmt.Errorf("http/Set/Proxy expects to have a string object: %v", e.Type)
		}
		u, err := url.Parse(e1.Value.(string))
		vm.OnError(err, "Error parsing URL")
		h := e.Value.(*vmmod.HTTP)
		h.RO.Proxies["http"] = u
		h.RO.Proxies["https"] = u
		h.RO.Proxies["ftp"] = u
		return e, nil
	}
	return nil, fmt.Errorf("http/Set/Proxy expects to have a http object: %v", e.Type)
}

func HttpSetAgentElement(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	if e.Type != "http" {
		if !vm.CanGet() {
			return nil, fmt.Errorf("Stack is too shallow for http/Set/Agent")
		}
		e1 := vm.Take()
		if e1.Type != "str" {
			fmt.Errorf("http/Set/Agent expects to have a string object: %v", e.Type)
		}
		h := e.Value.(*vmmod.HTTP)
		h.RO.UserAgent = e1.Value.(string)
		return e, nil
	}
	return nil, fmt.Errorf("http/Set/Proxy expects to have a http object: %v", e.Type)
}

func HttpSetHeadersElement(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	if e.Type != "http" {
		if !vm.CanGet() {
			return nil, fmt.Errorf("Stack is too shallow for http/Set/Headers")
		}
		e1 := vm.Take()
		if e1.Type != "dblock" {
			fmt.Errorf("http/Set/Headers expects to have a dblock object: %v", e.Type)
		}
		hdr, err := vmmod.Block2Dict(vm, e1)
		vm.OnError(err, "Error in getting data from DBLOCK")
		h := e.Value.(*vmmod.HTTP)
		h.RO.Headers = *hdr
		return e, nil
	}
	return nil, fmt.Errorf("http/Set/Headers expects to have a http object: %v", e.Type)
}

func HttpSetParamsElement(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	if e.Type != "http" {
		if !vm.CanGet() {
			return nil, fmt.Errorf("Stack is too shallow for http/Set/Params")
		}
		e1 := vm.Take()
		if e1.Type != "dblock" {
			fmt.Errorf("http/Set/Params expects to have a dblock object: %v", e.Type)
		}
		hdr, err := vmmod.Block2Dict(vm, e1)
		vm.OnError(err, "Error in getting data from DBLOCK")
		h := e.Value.(*vmmod.HTTP)
		h.RO.Params = *hdr
		return e, nil
	}
	return nil, fmt.Errorf("http/Set/Params expects to have a http object: %v", e.Type)
}

func HttpSetDataElement(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	if e.Type != "http" {
		if !vm.CanGet() {
			return nil, fmt.Errorf("Stack is too shallow for http/Set/Data")
		}
		e1 := vm.Take()
		if e1.Type != "dblock" {
			fmt.Errorf("http/Set/Data expects to have a dblock object: %v", e.Type)
		}
		hdr, err := vmmod.Block2Dict(vm, e1)
		vm.OnError(err, "Error in getting data from DBLOCK")
		h := e.Value.(*vmmod.HTTP)
		h.RO.Data = *hdr
		return e, nil
	}
	return nil, fmt.Errorf("http/Set/Data expects to have a http object: %v", e.Type)
}

func HttpGetElement(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	switch e.Type {
	case "str":
		ro := &grequests.RequestOptions{
			InsecureSkipVerify: true,
			IsAjax:             false,
		}
		resp, err := grequests.Get(e.Value.(string), ro)
		vm.OnError(err, "Error in GET(%v)", e.Value.(string))
		return &vmmod.Elem{Type: "str", Value: resp.String()}, nil
	case "http":
		h := e.Value.(*vmmod.HTTP)
		resp, err := grequests.Get(h.URL, h.RO)
		vm.OnError(err, "Error in GET(%v)", h.URL)
		return &vmmod.Elem{Type: "str", Value: resp.String()}, nil
	}
	return nil, fmt.Errorf("http/Get expects to have a string with proper URL: %v", e.Type)
}

func HttpPostElement(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	switch e.Type {
	case "str":
		ro := &grequests.RequestOptions{
			InsecureSkipVerify: true,
			IsAjax:             false,
		}
		resp, err := grequests.Post(e.Value.(string), ro)
		vm.OnError(err, "Error in GET(%v)", e.Value.(string))
		return &vmmod.Elem{Type: "str", Value: resp.String()}, nil
	case "http":
		h := e.Value.(*vmmod.HTTP)
		resp, err := grequests.Post(h.URL, h.RO)
		vm.OnError(err, "Error in GET(%v)", h.URL)
		return &vmmod.Elem{Type: "str", Value: resp.String()}, nil
	}
	return nil, fmt.Errorf("http/Get expects to have a string with proper URL: %v", e.Type)
}

func InitHttpFunctions(vm *vmmod.VM) {
	vm.Debug("[ BUND ] bund.InitHttpFunctions() reached")
	vm.AddFunction("http/Get", HttpGetElement)
	vm.AddFunction("http", HttpGetElement)
	vm.AddFunction("http/Post", HttpPostElement)
	vm.AddFunction("http/Set/Proxy", HttpSetProxyElement)
	vm.AddFunction("http/Set/Agent", HttpSetAgentElement)
	vm.AddFunction("http/Set/Headers", HttpSetHeadersElement)
	vm.AddFunction("http/Set/Params", HttpSetParamsElement)
	vm.AddFunction("http/Set/Data", HttpSetDataElement)
}
