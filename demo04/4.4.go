package main

import (
	"fmt"
)

const c = "C"

var v int = 5

type T struct{}

func init() {
	// init
}
func main() {
	var a int
	a = 2
	fmt.Print(a)
	Func1()
}

func (t T) Method1() {
	// ..
}

func Func1() {
	// ..
}
