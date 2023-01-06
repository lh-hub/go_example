package main

import (
	"fmt"
	"math"
)

type Square struct {
	side float32
}

type T1 struct {
	ra float32
}

type I1 interface {
	Area() float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (t T1) Area() float32 {
	return t.ra * t.ra * math.Pi
}

func main() {
	var a I1
	sq1 := new(Square)
	sq1.side = 5

	a = sq1
	if t, ok := a.(*Square); ok {
		fmt.Printf("The type of areaIntf is: %T\n", t)
	}
    if u, ok := a.(*T1); ok {
        fmt.Printf("The type of areaIntf is: %T\n", u)
    } else {
        fmt.Println("areaIntf does not contain a variable of type Circle")
    }

}
