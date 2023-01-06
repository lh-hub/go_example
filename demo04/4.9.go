package main

import "fmt"

func main() {
	var n int16 = 34
	var m int32

	m = int32(n)
	fmt.Printf("32 bit int is: %d\n", m)
	fmt.Printf("32 bit int is: %d\n", n)

	var c1 complex64 = 5 + 10i
	println(c1)

	b3 := 10 > 5
	println(b3)

}
