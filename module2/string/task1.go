package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main(){

	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	text_far := []rune(text)

	// fmt.Println(text_far)

	if unicode.IsUpper(text_far[0]) && string(text_far[len(text_far)-1]) == "."{
		fmt.Println("Right")
	}else{
		fmt.Println("Wrong")
	}

}