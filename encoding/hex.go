package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	set := "030300000032C5FD"
	setBytes, _ := hex.DecodeString(set)
	fmt.Printf("% X", setBytes) // 03 03 00 00 00 32 C5 FD
}
