package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	if 7 != 16 {
		t.Log("我就是要报一个错，你要怎么地！哼哼！！！")
		t.FailNow()
	}
}
