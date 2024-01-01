package main

import "fmt"

func main(){

    var arr[]int

    var in int

    for i := 0; i <= 4; i++ {
        fmt.Scan(&in)
        arr = append(arr,in)
    }

    fmt.Println(arr[4])

}