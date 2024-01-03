func calculator(arguments <-chan int, done <-chan struct{}) <-chan int {

	my := make(chan int)

	var summa int
	go func() {
        defer close(my)
		for {
			select {
			case m := <-arguments:
				summa += m
			case <-done:
				my <- summa
				return
			}

		}
	}()

	return my
}