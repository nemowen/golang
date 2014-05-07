package main

import (
	"fmt"

	"log"

	"os"
)

func main() {

	logfile, err := os.OpenFile("D:/EN/client.log", os.O_RDWR|os.O_CREATE, 0)

	if err != nil {

		fmt.Printf("%s\r\n", err.Error())

		os.Exit(-1)

	}

	defer logfile.Close()

	logger := log.New(logfile, "\r\n", log.Ldate|log.Ltime|log.Llongfile)

	logger.Println("hello")

	logger.Println("oh....")

	logger.Fatal("test")

	logger.Fatal("test2")

}
