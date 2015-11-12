package main

import (
	"fmt"
	"math"
)

type Point3 struct {
	x, y, z float64
}

func (p Point3) Abs() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y + p.z*p.z)
}

func main() {
	p := &Point3{3, 4, 5}
	fmt.Printf("%p result : %f p:%s\n", p, p.Abs(), p)

}
