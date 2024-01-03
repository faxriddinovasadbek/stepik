package main

import (
	"fmt"
	"time"
)

func main() {
	var timee string

	fmt.Scan(&timee)

	t, err := time.Parse(time.RFC3339, timee)

	if err != nil{
		return 
	}

	fmt.Println(t.Format(time.UnixDate))
	
}
