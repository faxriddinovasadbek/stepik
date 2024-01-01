package main

import (
	"fmt"
	"strings"
)

func main(){
	var str string
	var my_str string

	fmt.Scan(&str)

	for i := 0; i < len(str); i++ {
		if strings.Count(str, string(str[i])) == 1{
			my_str+=string(str[i])
		}
	}
	fmt.Println(my_str)
}