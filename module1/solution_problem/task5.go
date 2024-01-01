package main 

import "fmt"

func main(){
	var a, b, c int

	fmt.Scan(&a,&b,&c)

    if c < a + b && b < a + c && a < b +c{
        fmt.Println("Существует")
	}else{
		fmt.Println("Не существует")
	}
}