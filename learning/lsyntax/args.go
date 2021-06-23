package lsyntax

import "fmt"

func swapInt(x, y int) {
	x, y = y, x
	fmt.Printf("swap: x=%v, y=%v\n", x, y)
}

func TestSwapInt() {
	x := 5
	y := 10
	swapInt(x, y)
	fmt.Printf("x=%v, y=%v\n", x, y)
}

func swapPoint(a, b *int) {
	*a, *b = *b, *a
}

func TestSwapPoint() {
	a := 1
	b := 2
	swapPoint(&a, &b)
	fmt.Printf("a=%v, b=%v\n", a, b)
}
