package main

import (
	"fmt"
	"os"
	"strconv"
)

type Celsius float32

func (c Celsius) String() string {
	return strconv.FormatFloat(float64(c), 'f', 1, 32) + "℃"
}

type Day int

var zh_CN_Day = []string{"星期一", "星期二", "星期三", "星期四", "星期五", "星期六", "星期日"}

func (d Day) String() string {
	return zh_CN_Day[d]
}

func prints(args ...interface{}) {
	for i, v := range args {
		if i > 0 {
			os.Stdout.WriteString(" ")
		}
		switch t := v.(type) {
		case string:
			os.Stdout.WriteString(t)
		case fmt.Stringer:
			os.Stdout.WriteString(t.String())
		case int:
			os.Stdout.WriteString(strconv.Itoa(t))
		default:
			os.Stdout.WriteString("???")
		}
	}
	os.Stdout.WriteString("\n")
}

func main() {
	prints("nemowen", 28, Day(7), Celsius(36.5222))
}
