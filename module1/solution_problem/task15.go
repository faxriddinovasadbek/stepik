package main

import (
	"fmt"
	
)

func main(){
	var str , n, temp string
	fmt.Scan(&str, &n)
	for _, let := range str{
		if string(let) != n{
			temp+=string(let)
		}
	}
	fmt.Println(temp)
}