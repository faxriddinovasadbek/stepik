func main() {
    var x, y int

	fmt.Scan(&x, &y)
    res, err := divide(x, y)

	if err != nil{
        fmt.Println("ошибка")
	}else{
        fmt.Println(res)
	}
}