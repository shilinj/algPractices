package learning

import (
	"algstu/learning/lsyntax"
)

func TestDefer() {
	lsyntax.DeferRunningTime()
	lsyntax.DeferFuncArgs()
}

func TestSlice() {
	lsyntax.SliceCapAppend()
}

func TestArgs() {
	lsyntax.TestSwapInt()
	lsyntax.TestSwapPoint()
}
