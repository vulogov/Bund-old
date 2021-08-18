package stdlib

import (
	"fmt"
	"math/big"

	"github.com/gammazero/deque"

	"github.com/vulogov/Bund/internal/conf"
	vmmod "github.com/vulogov/Bund/internal/vm"
	// "github.com/goml/gobrain"
)

func PctElement(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	if e.Type == "dblock" {
		eh, err := vm.GetType("PCT")
		vm.OnError(err, "Error in ml/Perceptron")
		p := eh.Factory(vm)
		pct := p.Value.(*vmmod.PERCEPTRON)
		q := e.Value.(*deque.Deque)
		_in := 0
		_out := 0
		_mod := true
		_tdata := new([]float64)
		_tval := new([]float64)
		for i := 0; i < q.Len(); i++ {
			v := q.At(i).(*vmmod.Elem)
			dat := 0.0
			switch v.Type {
			case "int":
				dat = float64(v.Value.(*big.Int).Int64())
			case "flt":
				dat = v.Value.(float64)
			case "SEPARATE":
				if _mod {
					_mod = false
					pct.D += 1
				} else {
					_t := make([][]float64, 2)
					_t[0] = *_tdata
					_t[1] = *_tval
					pct.T = append(pct.T, _t)
					_tdata = new([]float64)
					_tval = new([]float64)
					_mod = true
					if pct.In == 0 {
						pct.In = _in
						pct.Out = _out
					} else {
						if pct.In != _in || pct.Out != _out {
							return nil, fmt.Errorf("Training data is incorrect arity: %v %v, %v %v", pct.In, _in, pct.Out, _out)
						}
					}
					_in = 0
					_out = 0
				}
				continue
			default:
				continue
			}
			if _mod {
				*_tdata = append(*_tdata, dat)
				_in += 1
			} else {
				*_tval = append(*_tval, dat)
				_out += 1
			}
		}
		pct.Hidden = pct.In * 2
		vm.Debug("Initializing PERCEPTRON IN: %v H: %v OUT: %v", pct.In, pct.Hidden, pct.Out)
		pct.P.Init(pct.In, pct.Hidden, pct.Out)
		pct.P.Train(pct.T, pct.Epoch, pct.Rate, pct.MFact, pct.Lerr)
		if *conf.Debug {
			vm.Debug("Testing newly created PERCEPTRON")
			pct.P.Test(pct.T)
		}
		return p, nil
	}
	return nil, fmt.Errorf("Expected (* ) for ml/Perceptron I have: %v", e.Type)
}

func PctUpdElement(vm *vmmod.VM, e *vmmod.Elem) (*vmmod.Elem, error) {
	if e.Type == "PCT" {
		if !vm.CanGet() {
			return nil, fmt.Errorf("Stack is too shallow for ml/Perceptron/Update ")
		}
		e1 := vm.Take()
		if e1.Type != "dblock" {
			return nil, fmt.Errorf("Expected (* ) for ml/Perceptron/Update I have: %v", e.Type)
		}
		vm.Debug("Pushing PERCEPTRON back to stack")
		vm.Put(e)
		eh, err := vm.GetType("dblock")
		vm.OnError(err, "Error in ml/Perceptron/Update")
		uvec := make([]float64, 0)
		q := e1.Value.(*deque.Deque)
		for i := 0; i < q.Len(); i++ {
			v := q.At(i).(*vmmod.Elem)
			switch v.Type {
			case "int":
				uvec = append(uvec, float64(v.Value.(*big.Int).Int64()))
			case "flt":
				uvec = append(uvec, v.Value.(float64))
			}
		}
		pct := e.Value.(*vmmod.PERCEPTRON)
		if len(uvec) != pct.In {
			return nil, fmt.Errorf("Arity of input data for ml/Perceptron/Update not matched I have: %v %v", pct.In, len(uvec))
		}
		_res := pct.P.Update(uvec)
		res := eh.Factory(vm)
		q1 := res.Value.(*deque.Deque)
		for j := 0; j < len(_res); j++ {
			q1.PushBack(&vmmod.Elem{Type: "flt", Value: _res[j]})
		}
		return res, nil
	}
	return nil, fmt.Errorf("Expected PERCEPTRON for ml/Perceptron/Update I have: %v", e.Type)
}

func InitPctFunctions(vm *vmmod.VM) {
	vm.Debug("[ BUND ] bund.InitPctFunctions() reached")
	vm.AddFunction("ml/Perceptron", PctElement)
	vm.AddFunction("ml/Perceptron/Update", PctUpdElement)
}
