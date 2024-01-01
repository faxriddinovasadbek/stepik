package main

import(
    "fmt"

)

func main(){
    var a int32

    fmt.Scan(&a)

    for a >= 10 {
        a /= 10
    }
    fmt.Println(a)
}