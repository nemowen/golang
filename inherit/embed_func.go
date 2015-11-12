package main

import (
	"fmt"
)

type Base struct {
	id int64
}

type Log struct {
	Base
	msg string
}

type Customer struct {
	Name string
	Age  string
	Log
}

func (l *Log) Add(s string) {
	l.msg += "\n" + s
}

func (l *Log) String() string {
	return l.msg
}

func (c *Customer) String() string {
	return c.Name + "\nLog:" + fmt.Sprintln(c.Log)
}

func main() {
	c := &Customer{"Barak Obama", "g", Log{Base{2}, "1 - yes we can!"}}
	c.Add("2 - oooooooooooooooo")
	fmt.Println(c)
}
