package main

import "fmt"

func main(){

	var time int

	fmt.Scan(&time)

	t := time / 3600
	s := time % 3600 / 60 

	fmt.Printf("It is %d hours %d minutes.", t,s)

	
}