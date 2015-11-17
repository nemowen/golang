package main

import (
	"os"
	"fmt"
	"io"
	"bufio"
)

func cat(r *bufio.Reader) {
	for(
		buf,err:=r.ReadBytes('\n')
		if err==io.EOF{
			break
		}
		fmt.Fprintf(os.Stdout,"%s",buf)
	)
	return
}
