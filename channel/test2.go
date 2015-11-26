package main

import (
	"fmt"
)

type tfunc func(a, b int)

func dof(f tfunc) tfunc {
	return func(a, b int) {
		f(a, b)
	}
}

func aaa(a, b int) {

}
