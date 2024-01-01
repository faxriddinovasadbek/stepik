package main

import "fmt"

func main(){
	var a,in int

	count := 0

	fmt.Scan(&a)

	var arr[]int

	for i := 0; i < a; i++ {
		fmt.Scan(&in)
		arr = append(arr, in)
	}

	min := arr[0]
	
	for i := 0; i < a; i++{
		if min == arr[i]{
			count++
		}else if min > arr[i]{
			min = arr[i]
			count = 1
		}
	}

	fmt.Println(count)
}