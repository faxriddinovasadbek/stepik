package main

import (
	"fmt"
	"strings"
)

func main(){
	var num1, num2 string

	fmt.Scan(&num1, &num2)

	for _, num := range num1{
		if 0 != strings.Count(num2, string(num)){
			fmt.Print(string(num), " ")
		}	
	}
}