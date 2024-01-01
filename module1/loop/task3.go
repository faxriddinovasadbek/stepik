package main 

import "fmt"

func main(){
    var a, b, sum int
    fmt.Scan(&a)
    
    
    for i := 0; i < a; i++ {
        fmt.Scan(&b)
        if 10 <= b && b <= 99 && b % 8 == 0{
            sum += b
        }
    }

    fmt.Println(sum)
}