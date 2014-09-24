package main

import (
	"fmt"
	"sync"
)

func doit() {
	fmt.Println("ok")
}

func main() {
	var once sync.Once
	once.Do(doit)
	once.Do(doit)
	once.Do(doit)
	once.Do(doit)

}
