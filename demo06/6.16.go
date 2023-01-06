package main

import "fmt"

func Add2() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}

func Adder(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}

func main() {
	p2 := Add2()
	fmt.Println(p2(2))

	TwoAdder := Adder(2)
	fmt.Println(TwoAdder(3))
}
