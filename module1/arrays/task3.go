package main

import "fmt"

func main(){

	// channal := make(chan int)

	var sls []int
	var fn int

	for i := 0; i < 5; i++ {
		fmt.Scan(&fn)
		sls = append(sls, fn)
	}

	max := sls[0]

	for i := 1; i < 5; i++{
		if max < sls[i]{
			max = sls[i]
		}
	}

	fmt.Println(max)
}