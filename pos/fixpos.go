package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"strings"
)

func main() {
	paths, _ := os.Getwd()
	tos := path.Join(paths, "/workspace/.metadata/.plugins/org.eclipse.core.runtime/.settings/com.comresource.eshoes.pos.client.prefs")
	// open this file
	file, err := os.Open(tos)
	defer file.Close()

	checkError(err)

	buf := bufio.NewReader(file)

	var result string

	for {

		line, _, err := buf.ReadLine()

		if err != nil {
			// Unexpected error
			if err != io.EOF {
				checkError(err)
			}

			// Reached end of file, if nothing to read then break,
			// otherwise handle the last line.
			if len(line) == 0 {
				break
			}
		}
		result += string(line) + "\n"
	}
	if !strings.Contains(result, "cr.update") {
		result += "cr.update = http://192.168.2.19:2000/posUpdate\n"
		wr, err := os.Create(tos)
		defer wr.Close()
		checkError(err)
		wr.WriteString(result)

	}
	fmt.Println("已经修复成功!!!!")
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Println(os.Stderr, "Fatal error: ", err.Error())
		os.Exit(1)
	}
}
