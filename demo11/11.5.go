package main

import (
	"fmt"
)

type List []int

func (l List) Len() int {
	return len(l)
}

func (l *List) Append(val int) {
	*l = append(*l, val)
}

type Appender interface {
	Append(int)
}

func CountInfo(a Appender, start, end int) {
	for i := start; i <= end; i++ {
		a.Append(i)
	}
}

func LongEnough(l Lener) bool {
	return l.Len()*10 > 42
}

type Lener interface {
	Len() int
}

func main() {
	var lst List
	// CountInto(lst, 1, 10)
	if LongEnough(lst) {
		fmt.Printf("- lst is long enough\n")
	}

	plst := new(List)
	CountInfo(plst, 1, 10)
	if LongEnough(plst) {
		fmt.Printf("- lst is long enough\n")
	}

}
