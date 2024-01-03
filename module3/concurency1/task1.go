func task(n chan int, num int) {
	n <- num + 1
}