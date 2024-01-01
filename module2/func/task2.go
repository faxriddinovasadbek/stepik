func minimumFromFour() int {
    min := 1000
    var num int
    for i := 0; i < 4; i++{
        fmt.Scan(&num)
        if min > num{
           min = num
        }
    }
    return min
}