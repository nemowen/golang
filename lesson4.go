package main

import (
	"fmt"
	"testing."
)

var b testing.TB = (*TB)(nil)

func main() {
	testTB(new(TB))
	fmt.Println(b)
}

func testTB(t testing.TB) {
	t.Fatal("aaa")
}

type TB struct {
	testing.TB
}

func (p *TB) Failed() bool { return true }
func (p *TB) Fatal(args ...interface{}) {
	fmt.Println("TB.Fatal disabled!")
}
