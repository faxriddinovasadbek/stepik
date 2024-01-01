package main 

import "fmt"

func main(){

    for {
        var inp int
        fmt.Scan(&inp)
        if inp > 100{
            break
        }else if inp < 10{
            continue
        }else{
            fmt.Println(inp)
        }
    }

}