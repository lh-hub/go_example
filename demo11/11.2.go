package main

import "fmt"

type Shaper interface {
	Area() float32
}

type T1 struct {
	side float32
}

func (t *T1) Area() float32 {
	return t.side * t.side
}

type T2 struct {
	length, width float32
}

func (t T2) Area() float32 {
	return t.length * t.width
}

func main() {
	t1 := &T1{5}
	t2 := T2{5, 3}

	shapes := []Shaper{t1, t2}
	for n, _ := range shapes {
		fmt.Println(shapes[n])
		fmt.Println(shapes[n].Area())
	}
}
