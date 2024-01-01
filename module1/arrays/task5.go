package main 

import "fmt"

func main(){
	var n , in, count int

	fmt.Scan(&n)

	var arr[]int

	for i := 0; i < n; i++{
		fmt.Scan(&in)
		arr = append(arr, in)
	}

	for i := 0; i < n; i++{
		if arr[i] > 0 {
			count++
		}
	}

	fmt.Println(count)
}