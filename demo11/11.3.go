package main

import "fmt"

type T1 struct {
	p1    string
	p2    float32
	count float32
}

func (t T1) getValue() float32 {
	return t.p2 * t.count
}

type T2 struct {
	make  string
	p2    string
	price float32
}

func (t T2) getValue() float32 {
	return t.price
}

type valueable interface {
	getValue() float32
}

func showValue(v valueable) {
	fmt.Println(v.getValue())
}

func main() {
	var o valueable = T1{"Go", 577.2, 4}
	showValue(o)
	o = T2{"BMW", "M3", 660000}
	showValue(o)
}
