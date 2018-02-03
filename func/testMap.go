package main

import (
	"fmt"
	"sync"
)

func main() {
	var m = new(sync.Map)
	m.Store("a", 2)
	m.Store("b", 3)

	v, _ := m.Load("a")

	vint := v.(int)

	fmt.Println("aaaa---->", vint)
}
