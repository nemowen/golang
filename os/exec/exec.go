package main

import (
	"os"
	"os/exec"
)

func main() {
	run1()
}

func run1() {
	com1 := exec.Command("ls")
	p, _ := com1.StdoutPipe()
	com2 := exec.Command("wc", "-l")
	com2.Stdin = p
	com2.Stderr = os.Stderr
	com2.Stdout = os.Stdout
	com1.Start()
	com2.Run()
	com2.Wait()
}

func run2(ok int, i string) {
	com := exec.Command("/bin/sh", "-c", `ls|wc -l`)
	com.Stdout = os.Stdout
	com.Stderr = os.Stderr
	com.Start()
	com.Run()
	com.Wait()
}
