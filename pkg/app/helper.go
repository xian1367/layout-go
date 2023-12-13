package app

import (
	"fmt"
)

func VarDump(expression ...interface{}) {
	fmt.Println(fmt.Sprintf("%#v", expression))
}
