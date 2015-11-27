package main

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"
)

var memprofile = flag.String("memprofile", "", "write memory profile to this file")

func main() {
	//vCallToFunctionWhichAllocatesLotsOfMemory()
	flag.Parse()
	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		pprof.WriteHeapProfile(f)
	}
}
