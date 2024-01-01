package main

import (
	"fmt"
	"unicode"
)

func IsPro(parol string) string {
	ru_str := []rune(parol)
	
	if len(ru_str) < 5{
		return "Wrong password"
	}

	for _, i := range ru_str {
		if !(unicode.Is(unicode.Latin,i) || unicode.IsDigit(i)){
			return "Wrong password"
		}
	}

	return "Ok"
}

func main(){
	var str string
	fmt.Scan(&str)
	fmt.Println(IsPro(str))
}