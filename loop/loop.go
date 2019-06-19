package main

import "fmt"

func main() {
	sum()

	gotoLoop()
}

func gotoLoop() {
	var i = 1
	var sum = 0

LOOP:
	for i <= 100 {

		if i == 10 {
			i = 100
			goto LOOP
		}

		sum += i
		i++
	}

	fmt.Println("goto loop sum=", sum)
}

func sum() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}

	fmt.Println("1+2+...+100=", sum)
}
