package main

import "fmt"

func main(){
    var a,b,sum int

    fmt.Scan(&a, &b)

    for a <= b {
        sum += a
        a++
    }
    fmt.Println(sum)
}