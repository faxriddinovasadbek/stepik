package main

import (
	"fmt"
)

func main() {
	var nums string
	
	fmt.Scan(&nums)

	max := 0

	for _, num := range nums{
		if max < int(num){
			max = int(num)
		}
	}

	fmt.Println(string(max))
}