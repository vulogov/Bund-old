package vm

import (
	"fmt"

	"github.com/goml/gobrain"

	"github.com/vulogov/Bund/internal/conf"
)

type PERCEPTRON struct {
	T      [][][]float64
	In     int
	Out    int
	Hidden int
	D      int
	Epoch  int
	Rate   float64
	MFact  float64
	Lerr   bool
	P      *gobrain.FeedForward
}

func PctFactory(vm *VM) *Elem {
	res := &PERCEPTRON{P: &gobrain.FeedForward{}}
	res.Epoch = 1000
	res.Rate = 0.6
	res.MFact = 0.4
	res.In = 0
	res.Out = 0
	res.Hidden = 0
	if *conf.Debug {
		res.Lerr = true
	} else {
		res.Lerr = false
	}
	return &Elem{Type: "PCT", Value: res}
}

func PctToString(vm *VM, e *Elem) string {
	if e.Type == "PCT" {
		pct := e.Value.(*PERCEPTRON)
		return fmt.Sprintf("PERCEPTRON[IN %v, H: %v, OUT: %v, D: %v]", pct.In, pct.Hidden, pct.Out, pct.D)
	}
	vm.Error("trying to convert a PCT and it is not a PCT: %v %T", e.Type, e.Value)
	return "<?>"
}

func PctFromString(vm *VM, d string) *Elem {
	return PctFactory(vm)
}

func PctCompare(vm *VM, e1 *Elem, e2 *Elem) int {
	return IDK
}

func PctDup(vm *VM, e *Elem) *Elem {
	if e.Type == "PCT" {
		res := PctFactory(vm)
		pct := res.Value.(*PERCEPTRON)
		pcte := e.Value.(*PERCEPTRON)
		pct.T = pcte.T
		pct.In = pcte.In
		pct.Out = pcte.Out
		pct.Hidden = pcte.Hidden
		pct.Lerr = pcte.Lerr
		pct.P.Init(pct.In, pct.Hidden, pct.Out)
		pct.P.Train(pct.T, pct.Epoch, pct.Rate, pct.MFact, pct.Lerr)
		if *conf.Debug {
			pct.P.Test(pct.T)
		}
		return res
	}
	return nil
}

func RegisterPerceptron(vm *VM) {
	vm.RegisterType("PCT", PctFactory, PctToString, PctFromString, PctCompare, PctDup)
}
