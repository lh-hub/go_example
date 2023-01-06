package main

import "fmt"

type Log struct {
	msg string
}

type Customer struct {
	Name string
	log  *Log
}

func (l *Log) Add(s string) {
	l.msg += "\n" + s
}

func (l *Log) String() string {
	return l.msg
}

func (c *Customer) Log() *Log {
	return c.log
}

func main() {
	c := new(Customer)
	c.Name = "test"
	c.log = new(Log)
	c.log.msg = "test log"
	c.Log().Add("test2")
	fmt.Println(c.Log())

}
