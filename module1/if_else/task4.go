package main

import "fmt"

func main(){
    var a int32

    fmt.Scan(&a)

    sum1 := a % 10
    sum1 += a / 10 % 10
    sum1 += a / 100 % 10

    sum2 := a / 1000 % 10
    sum2 += a / 10000 % 10
    sum2 += a / 100000 % 10    

    if sum1 == sum2{
        fmt.Println("YES")
    }else{
        fmt.Println("NO")
    }
}