func test(x1 *int, x2 *int) {
    tepm := *x1
    *x1 = *x2
    *x2 = tepm
    
    fmt.Println(*x1, *x2)
}