package main

import "fmt"

func trace(s string) string {
	fmt.Println("进入函数", s)
	return s
}

func un(s string) {
	fmt.Println("离开函数", s)
}

func a() {
	defer un(trace("A"))
	fmt.Println("in A")
}

func b() {
	defer un(trace("B"))
	fmt.Println("B")
	a()
}

func main() {
	b()
}
