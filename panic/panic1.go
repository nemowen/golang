package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Starting the program.")
	panic("A severe error occurred : stopping the program.")
	fmt.Println("Ending the program")
}
