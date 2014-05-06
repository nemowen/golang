package testinit

import (
	"fmt"
)

var Share string

func init() {
	fmt.Println("init run ")
	Share = "ok"
}
