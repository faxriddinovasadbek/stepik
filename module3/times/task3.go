package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	str1, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	str := strings.Split(str1, ",")

	tt1, tt2 := strings.TrimSpace(str[0]), strings.TrimSpace(str[1])


	t1, _ := time.Parse("02.01.2006 15:04:05", tt1)
	t2, _ := time.Parse("02.01.2006 15:04:05", tt2)

	if t1.After(t2){
		fmt.Println(t1.Sub(t2))
	}else{
		fmt.Println(t2.Sub(t1))
	}

}
