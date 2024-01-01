package main

import(
	"fmt"
)

func main(){

	var a,b int

	fmt.Scanln(&a)
	fmt.Scanln(&b)

	a = a * a
	b = b * b
	c := a + b

	fmt.Print(c)

}