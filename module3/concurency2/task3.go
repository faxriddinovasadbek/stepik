func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	my := make(chan int)
	go func() {
        defer close(my)
		select {
		case m := <-firstChan:
			my <- m * m
		case m := <- secondChan:
			my <- m * 3
		case <- stopChan:
			break
		}
	}()

	return my
}