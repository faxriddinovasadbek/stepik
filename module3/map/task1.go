m := make(map[int]int)
  var a int
  for i := 0; i < 10; i++{
    fmt.Scan(&a)
    if m[a] == 0 {
      m[a] = work(a)
    }
    fmt.Print(m[a], " ")
  }