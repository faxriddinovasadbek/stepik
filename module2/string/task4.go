package main

import "fmt"

func main(){
	var str string
	
	fmt.Scan(&str)

	for i := 0; i < len(str); i++ {
		if i % 2 == 1 {
			fmt.Print(string(str[i]))
		}
	}
}