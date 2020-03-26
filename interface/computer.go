package main

import "fmt"

type Computer interface {
	getInfo()
}
type Mac struct {
}

func (mac Mac) getInfo() {
	fmt.Println("Apple computer")
}

type Lenovo struct {
}

func (lenovo Lenovo) getInfo() {
	fmt.Println("Lenovo computer")
}

func main() {
	var computer Computer

	computer = new(Mac)
	computer.getInfo()

	computer = new(Lenovo)
	computer.getInfo()
}
