package main

import "fmt"

//常量用作枚举：
const (
	Monday  = 1
	Tuesday = 2
)

func main() {

	const SIZE int = 10

	fmt.Println(SIZE)

	//可修改的常量
	const (
		a = iota
		b = iota
		c = iota
	)
	fmt.Println(a, b, c)

	const (
		x = iota
		y
		z
	)
	fmt.Println(x, y, z)
}
