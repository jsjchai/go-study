package main

import "fmt"

// 这种因式分解关键字的写法一般用于声明全局变量
var (
	id   int
	name string
)

func main() {
	var hello string = "hello"
	fmt.Println(hello)

	var a int = 1
	var b int = 1
	fmt.Println(a + b)

	var f = 0.1
	fmt.Println(f)

	var flag = true
	fmt.Println(flag)

	tmp := "abcdefg"
	fmt.Println(tmp)

	id = 10000
	name = "Tom"

	fmt.Print("id=", id, " name=", name)

}
