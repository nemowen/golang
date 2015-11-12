package main

import (
	"fmt"
)

type B struct {
}

func (B) Magic() { fmt.Println("base magic") }

func (b B) MoreMagic() {
	b.Magic()
	b.Magic()
}

type Vo struct {
	B
}

func (Vo) Magic() {
	fmt.Println("Vo magic")
}

func main() {
	v := new(Vo)
	v.Magic()
	v.MoreMagic()
}
