package conf

import (
	"fmt"
)

var Major = 1
var Min = 1
var Rel = 0
var BVersion = fmt.Sprintf("%v.%v(rel %v)", Major, Min, Rel)
var EVersion = fmt.Sprintf("%v.%v", Major, Min)
