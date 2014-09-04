package main

import "os"
import "strings"

func main() {
	hostname, _ := os.Hostname()
	println(hostname)

}
