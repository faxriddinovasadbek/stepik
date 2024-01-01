package main

import "fmt"

func main(){
    var a int
    fmt.Scan(&a)
    c := a * 2 
    c += 100
    fmt.Println(c)
}