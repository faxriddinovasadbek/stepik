package main

import "fmt"

func main() {
    var input, max, count int

    for {
        fmt.Scan(&input)

        if input == 0 {
            break
        }
        if input > max {
            max = input
            count = 1
        } else if input == max {
            count++
        }
    }

    fmt.Println(count)
}