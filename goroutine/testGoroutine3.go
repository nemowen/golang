package main

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	i    int
	max  int
	text string
}

func main() {
	hello := Task{0, 3, "hello"}
	world := Task{0, 5, "world"}

	wait := new(sync.WaitGroup)

	gomethod := func(t *Task) {
		for t.i < t.max {
			time.Sleep(1 * time.Millisecond)
			fmt.Println(t.text)
			t.i++

		}
		wait.Done()
	}

	go gomethod(&hello)

	go gomethod(&world)
	wait.Add(1)

	wait.Wait()

}
