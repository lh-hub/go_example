package main

import "fmt"

func main() {
	str := "hello world"
	for pi, char := range str {
		fmt.Printf("%d --- %c\n", pi, char)
	}
}
