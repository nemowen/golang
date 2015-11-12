package main

import (
	"fmt"
)

type List []int

func (l List) Len() int        { return len(l) }
func (l *List) Append(val int) { *l = append(*l, val) }

func main() {
	// A bare value
	var lst List
	lst.Append(1)
	fmt.Printf("%v (len: %d)\n", lst, lst.Len())

	// A Point value
	plst := new(List)
	plst.Append(3)
	fmt.Printf("%v (len: %d)\n", plst, plst.Len())
}
