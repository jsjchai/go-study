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

//goroutine
func main() {
	go printf("goroutine-1")
	go printf("goroutine-2")
	printf("main")
}
