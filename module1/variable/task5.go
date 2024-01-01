package main

import "fmt"

func main(){
    var n uint16
    fmt.Scan(&n)
	n /= 10
    fmt.Println(n%10)
}