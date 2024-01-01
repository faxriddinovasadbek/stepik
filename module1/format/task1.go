package main

import (
  "fmt"
  "math"
)

func main() {
  var num float64
  fmt.Scan(&num)

  result := processNumber(num)

  fmt.Println(result)
}

func processNumber(num float64) string {
  if num <= 0 {
    return fmt.Sprintf("число %.2f не подходит", num)
  }

  if num > 10000 {
    return fmt.Sprintf("%e", num)
  }

  squared := math.Pow(num, 2)
  rounded := math.Round(squared*1e4) / 1e4

  return fmt.Sprintf("%.4f", rounded)
}