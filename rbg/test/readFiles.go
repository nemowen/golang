package main

import (
	//"io/ioutil"
	"log"
	//"bufio"
	"os"
	//"strings"
)

func main() {

	// files, _ := ioutil.ReadDir("d:/EN")

	// for _, f := range files {
	// 	fmt.Println(f.Name(), f.ModTime().Unix())
	// }

	// f, _ := os.Create("d:/EN/note.ini")
	// defer f.Close()

	f, e := os.Create("D:/EN/flag.ini")
	if e != nil {
		log.Println("打开文件失败：", "D:/EN/flag.ini")
		return
	}
	defer f.Close()
	// 将Flag标识位置为END
	f.WriteString("OK")

}
