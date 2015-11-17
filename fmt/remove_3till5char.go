package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	var input, output *os.File
	var err error
	input, err = os.Open("C:/Users/nemowen/temp/input.txt")
	if err != nil {
		fmt.Errorf("open input.txt file:", err)
	}
	defer input.Close()

	output, err = os.OpenFile("C:/Users/nemowen/temp/output.txt", os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Errorf("open output.txt file:", err)
	}
	defer output.Close()

	inputReader := bufio.NewReader(input)
	outputWrite := bufio.NewWriter(output)

	for {
		inputStr, _, e := inputReader.ReadLine()
		//fmt.Println(inputStr)
		if e == io.EOF {
			fmt.Println("EOF")
			outputWrite.Flush()
			break
		}

		outputString := string([]byte(inputStr)[2:5]) + "\r\n"
		_, e = outputWrite.WriteString(outputString)
		if err != nil {
			fmt.Println(err)
			break
		}
	}
	fmt.Println("Conversion done")
}
