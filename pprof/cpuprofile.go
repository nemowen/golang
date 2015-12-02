package main

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"
)

// Now run the program with this flag: 		progexec -cpuprofile=progexec.prof
// and then you can use the gopprof tool as: gopprof progexec progexec.prof

// The gopprof program is a slight variant of Google’s pprof C++ profiler; for
// more info on this tool, see http://code.google.com/p/google-perftools/.

// topN
// This shows the top N samples in the profile, e.g.: top5
// It shows the 10 most heavily used functions during the execution, an output
// like:
// Total: 3099 samples
// 626 20.2% 20.2% 626 20.2% scanblock
// 309 10.0% 30.2% 2839 91.6% main.FindLoops ...
// The 5th column is an indicator of how heavy that function is used.

var cpuprofile = flag.String("cpuprofile", "", "write cpuprofile to file")

func main() {
	flag.Parse()
	f, err := os.Create(*cpuprofile)
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
}
