package main

import "fmt"

func main() {
	func1()
}

func func1() {
	fmt.Println("A1")
	defer func2()
	// func2()
	fmt.Println("A2")

}

func func2() {
	fmt.Println("B")
}
