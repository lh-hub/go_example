package main

import "fmt"

func Add(a, b int) {
	fmt.Println(a + b)
}

func callback(y int, f func(int, int)) {
	f(y, 2)
}

func main() {
	callback(1, Add)
}
