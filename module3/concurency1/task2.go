func task2(n chan string, str string) {
	for i := 0; i <= 5; i++{
		n <- str + " "
	}
}