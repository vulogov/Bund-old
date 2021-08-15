package vm

import (
	"fmt"
	"strings"

	"github.com/peterh/liner"

	"github.com/vulogov/Bund/internal/conf"
)

func (vm *VM) DbgFun(pos int, cmd *Elem) {
	if *conf.DbgFun == "" {
		return
	}
	if !vm.IsStack() {
		vm.Error("Request debug of function for non-existing stack")
	}
	if vm.CurrentNS.CurrentLambdaName == "" {
		return
	}
	if *conf.DbgFun == vm.CurrentNS.CurrentLambdaName {
		for {
			vm.Debug("CMD: %v", cmd)
			line := liner.NewLiner()
			dcmd, err := line.Prompt(fmt.Sprintf("%v f(%v) ", pos, *conf.DbgFun))
			if err != nil {
				*conf.DbgFun = ""
			}
			dcmd = strings.TrimSpace(dcmd)
			if dcmd != "" {
				vm.Debug("Function(%v) debug command: %v", *conf.DbgFun, dcmd)
				vm.Exec(dcmd)
			} else {
				vm.Debug("Function(%v) moving to next", *conf.DbgFun)
				break
			}
		}
	}
}

func (vm *VM) DbgStack(msg string, cmd *Elem) {
	if *conf.DbgFun == vm.CurrentNS.CurrentLambdaName {
		fmt.Println(msg)
		vm.Exec("dumpstack")
	}
}
