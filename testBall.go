package main

import "fmt"
import "ball"
import "os"
import "encoding/json"

func main() {

	ba := ball.Ball{}
	ba.Red = [6]int{1, 21, 36, 4, 6, 9}
	ba.Blue = 11
	ba.NO = "20140110"
	ba.OpenDate = "20140123"

	file, _ := os.OpenFile("d:/ball.txt", os.O_RDWR|os.O_CREATE, 0644)

	//close the file
	defer func() {
		file.Close()
	}()

	s, _ := json.Marshal(ba)

	ret, _ := file.Write(s)

	fmt.Println(ret)
}
