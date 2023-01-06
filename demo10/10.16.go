package main

import (
	"./person"
	"fmt"
)

func main() {
	p := new(person.Person)
	p.SetFirstName("lh")
	fmt.Println(p.FirstName())
}
