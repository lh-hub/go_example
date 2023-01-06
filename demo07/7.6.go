package main

import "fmt"

func main() {
	arrar := [3]float64{7.0, 8.5, 9.1}
	x := Sum(&arrar)
	fmt.Println(x)
}

func Sum(a *[3]float64) (sum float64) {
	for _, v := range *a {
		sum += v
	}
	return
}
