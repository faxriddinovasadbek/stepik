
package main

import "fmt"

func main(){
	var a,in,count int

	fmt.Scan(&a)

	var arr[]int

	for i := 0; i < a; i++{
		fmt.Scan(&in)
		arr = append(arr, in)
	}

	
	for i := 0; i < a; i++{
		if arr[i] == 0{
			count++
		}
	}

	fmt.Println(count)
}