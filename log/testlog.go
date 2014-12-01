package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
)

var (
	logFileName = flag.String("log", "nserver.log", "log file name")
)

func main() {
	flag.Parse()
	logfile, logerr := os.OpenFile(*logFileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if logerr != nil {
		fmt.Println("Fail to find", *logfile, "nserver start failed")
		os.Exit(1)
	}
	log.SetOutput(logfile)
	log.SetFlags(log.Ltime | log.Ldate | log.Lshortfile)
	log.Printf("server abort! cause:%v\n", "test log file")
	fmt.Println(runtime.GOOS, runtime.GOARCH)

}
