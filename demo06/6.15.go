package main

import "fmt"

func main() {
	for i := 0; i < 4; i++ {
		g := func(i int) { fmt.Println(i) }
		g(i)
		fmt.Println(g)
	}
}
