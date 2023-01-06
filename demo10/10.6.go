package main

import (
	"./structPack"
	"fmt"
)

func main() {
	struct1 := new(structPack.ExpStruct)
	struct1.Mi1 = 10
	struct1.Mi2 = 16.

	fmt.Printf("Mi1 = %d\n", struct1.Mi1)
	fmt.Printf("Mi2 = %f\n", struct1.Mi2)
}
