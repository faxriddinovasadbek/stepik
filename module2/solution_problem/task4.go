package main

import (
	"fmt"
	"strconv"
)

func main(){
	var str, my_str string

	fmt.Scan(&str)

	for _, i := range str {
		my_str += strconv.Itoa(((int(i)-48)*(int(i)-48)))
	}

	fmt.Println(my_str)

}