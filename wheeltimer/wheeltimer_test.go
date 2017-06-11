package wheeltimer

import (
	"fmt"
	"testing"
	"time"
)

func TestNewTimer(t *testing.T) {
	f := func(t ...string) error {
		fmt.Println("ok")
		return nil
	}
	AddTask("1111", f, nil, 1*time.h)
	StartTimer()
	time.Sleep(3 * time.Second)
	StopTimer()
}
