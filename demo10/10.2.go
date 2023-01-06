package main

import (
	"fmt"
	"strings"
)

type Persion struct {
	firstName string
	lastName  string
}

func upPersion(p *Persion) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}

func main() {
	var pers1 Persion
	pers1.firstName = "Chris"
	pers1.lastName = "Woodward"
	upPersion(&pers1)
	fmt.Println(pers1)
}
