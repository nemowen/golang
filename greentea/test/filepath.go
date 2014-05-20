package main

import (
	"time"
)

func main() {

	t, _ := time.Parse("20060102", "20140519")
	//println(t.Add(1 * time.Hour).String())
	println(time.Now().Sub(t.Add(2*24*time.Hour)) > 0)
}
