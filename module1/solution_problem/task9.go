package main

import "fmt"

func iteration(a int) int {
  sum := 0
  for a % 10 > 0 {
    sum += a % 10
    a /= 10
  }
  return sum
}

func main() {
  var a int
  fmt.Scan(&a)
    
  for a > 10 {
    a = iteration(a)
  }
  fmt.Println(a)
}