package main

import "fmt"

func my_funk(num int) int{
	var a , b , c int = 1,1,1

	for a <= num{
		if a == num	{
			return c
		}
		c++
		a , b = b, a+b
	}

	return -1
}

func main() {
	var n int
	fmt.Scan(&n)
    fmt.Println(my_funk(n))
}

