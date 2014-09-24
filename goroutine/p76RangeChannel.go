package main

func main() {
	c := make(chan int)
	go func() {
		for i := 0; i < 20; i++ {
			c <- i
		}
		close(c)
	}()

	// 方法一 可以用 range 迭代器从 channel 中接收数据，直到 channel close 方才终止循环。
	for v := range c {
		println(v)
	}

	// 方法二 接收方还可以用另一种⽅方式代替迭代器接收数据并判断 channel close 状态。
	// for {
	// 	if v, ok := <-c; ok {
	// 		println(v)
	// 	} else {
	// 		break
	// 	}
	// }

}
