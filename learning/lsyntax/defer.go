package lsyntax

import "fmt"

// defer func()在函数return之后调用
func DeferRunningTime() {
	var i = 5
	defer func() {
		i += 1
		fmt.Println("defer i =", i)
	}()
	fmt.Println("func i =", i)
}

// defer func()的参数, 值传递
func DeferFuncArgs() {
	i := 0
	defer fmt.Println("defer i = ", i)
	i++
	fmt.Println("func i =", i)
}
