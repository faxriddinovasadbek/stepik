package main

import(
	"fmt"
)

func IsPalindrom(str string) string {
	str1 := []rune(str)
	for i := 0; i < len(str1)/2; i++{
		if string(str1[i]) != string(str1[len(str1)-i-1]){
			return "Нет"
		}
	}
	return "Палиндром"
}

func main(){
	var pal string

	fmt.Scan(&pal)

	fmt.Println(IsPalindrom(pal))
}