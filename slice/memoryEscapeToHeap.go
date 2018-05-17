package main

func main() {
	str := getString()
	_ = str

}

// go tool compile -m memoryEscapeToHeap.go
func getString() *string {
	var s string
	s = "nemowen"
	return &s
}
