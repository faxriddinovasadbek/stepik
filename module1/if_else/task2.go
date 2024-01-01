package main

import(
    "fmt"
)

func main(){
    var a uint32

    fmt.Scan(&a)

    birlik := a % 10
    onlik := (a / 10) % 10
    yuzlik := a / 100
    
    if birlik != onlik && birlik != yuzlik && onlik != yuzlik{
        fmt.Println("YES")
    }else{
        fmt.Println("NO")
    }
}