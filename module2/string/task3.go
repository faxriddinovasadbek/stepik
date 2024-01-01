package main

import (
	"fmt"
	"strings"
)

func main(){
	var text , let string
	fmt.Scan(&text, &let)

	fmt.Println(strings.Index(text, let))
}