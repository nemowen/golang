package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	out, err := os.OpenFile("C:/Users/nemowen/w.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Fscanf(os.Stdout, "open file error:%s", err)
		os.Exit(1)
	}
	defer out.Close()

	w := bufio.NewWriter(out)
	w.WriteString("here is the simple\r\n")
	w.Flush()

}
