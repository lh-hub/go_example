package main

import (
	"fmt"
	"time"
)

type myTime struct {
	time.Time
}

func (t myTime) first3Chats() string {
	return t.Time.String()[0:10]
}

func main() {
	m := myTime{time.Now()}
	fmt.Println(m)
	fmt.Println(m.String())
	fmt.Println(m.first3Chats())
}
