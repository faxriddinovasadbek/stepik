package main

import "fmt"

func main(){
    var a, in int

    fmt.Scan(&a)

    var arr[]int

    for i := 1; i <= a; i++ {
        fmt.Scan(&in)
        arr = append(arr,in)
    }

    for i := 0; i <= a-1; i++ {
        if i % 2 == 0{
            fmt.Printf("%d ", arr[i])
        }
    }
}