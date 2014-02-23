package main

func main() {
	c := make(chan int)
	//o := make(chan bool)

	go func() {
		for v := range c {
			print(v)
		}
		//o <- true
	}()

	for i := 0; i < 100000; i++ {
		select {
		case c <- 0:
		case c <- 1:
		}
	}

	//close(c)
	//<-o
}
