package main

import "fmt"

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	// 缓冲区大小为2
	ch := make(chan int, 2)

	ch <- 1
	ch <- 2
	//ch <- 100

	fmt.Println(<-ch)
	fmt.Println(<-ch)

	fmt.Println("------------")

	c := make(chan int, 10)

	//cap(c) 获取通道容量
	go fibonacci(cap(c), c)

	for i := range c {
		fmt.Println(i)
	}
}
