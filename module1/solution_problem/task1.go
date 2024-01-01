package main

import "fmt"

func main(){
	var a, count int

	fmt.Scanf("%v", &a)

	for a != 0{
		count += a % 10
		a /= 10
	}

	fmt.Println(count)

}