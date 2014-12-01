package main

func main() {
	c1, c2 := make(chan int), make(chan string)

	go func() {
		for {
			select {
			case v, ok := <-c1:
				if ok {
					println("c1=", v)
				}
			case v, ok := <-c2:
				if ok {
					println("c2=", v)
				}

			}
		}

	}()

	c1 <- 1
	c2 <- "you "
	c1 <- 32
	c2 <- "is "
	c2 <- "died"
	close(c2)
	close(c1)
	//o <- true
	println("ok")

}
