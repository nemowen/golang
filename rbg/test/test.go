package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// set tty input buffer to 1 char and disable echo.
	// HACK: exec command twice to make sure it will work on both mac and linux.
	exec.Command("/bin/stty", "-f", "/dev/tty", "-icanon", "min", "1", "-echo").Run()
	exec.Command("/bin/stty", "-F", "/dev/tty", "-icanon", "min", "1", "-echo").Run()

	// read input directly from tty.
	f, e := os.Open("/dev/tty")

	if e != nil {
		panic("cannot open /dev/tty. error: " + e.Error())
	}

	// a deadloop reading tty to demo concept.
	var buf [1]byte

	for {
		n, e := f.Read(buf[:])

		if e != nil || n == 0 {
			break
		}

		fmt.Printf("Got input: 0x%x\n", buf[0])
	}
}
