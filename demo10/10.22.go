package main

import (
	"fmt"
)

type Twoints struct {
	a int
	b int
}

func (tn *Twoints) String() string {
	return "lh"
}

func main() {
	two1 := new(Twoints)
	two1.a = 12
	two1.b = 10
	fmt.Printf("two1 is: %v\n", two1)
	fmt.Println("two1 is:", two1)
	fmt.Printf("two1 is: %T\n", two1)
	fmt.Printf("two1 is: %#v\n", two1)
}
