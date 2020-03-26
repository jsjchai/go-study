package main

import (
	"fmt"
	"time"
)

func printf(s string) {
	for i := 1; i <= 10; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(fmt.Sprintf("%s %d", s, i))
	}
}

//通道
func sum(array []int, res chan int) {
	sum := 0
	for _, v := range array {
		sum += v
	}
	res <- sum // 把 sum 发送到通道 c
}

//goroutine
func main() {
	go printf("goroutine-1")
	go printf("goroutine-2")
	printf("main")

	array := []int{1, 2, 8, -9, 4, 7}
	res := make(chan int)
	go sum(array, res)
	fmt.Println(res)

	s := <-res
	fmt.Println(s)

	fmt.Println(array[:len(array)/2])
	fmt.Println(array[len(array)/2:])
	fmt.Println(array[len(array)/2])
	go sum(array[:len(array)/2], res)
	go sum(array[len(array)/2:], res)

	// 从通道 c 中接收
	x, y := <-res, <-res

	fmt.Println(x, y, x+y)

}
