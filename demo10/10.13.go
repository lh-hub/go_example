package main

import "fmt"

type B struct {
	thing int
}

func (b *B) change() { b.thing = 1 }

// func (b B) write() string { return fmt.Sprint(b) }
func (b B) write() string { b.thing = 3; return fmt.Sprint(b)}

func main() {
	var b1 B
	fmt.Println(b1)
	b1.change()
	fmt.Println(b1.write())
	fmt.Println(b1)

	b2 := new(B)
	b2.change()
	fmt.Println(b2.write())
	fmt.Println(b2)
}
