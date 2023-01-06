package main

import "fmt"

func main() {
	var arr1 [5]int

	for i := 0; i < len(arr1); i++ {
		arr1[i] = i * 2
	}

	for i := 0; i < len(arr1); i++ {
		fmt.Printf("%d is %d\n", i, arr1[i])
	}

	for i := range arr1 {
		fmt.Println(i)
	}
}
