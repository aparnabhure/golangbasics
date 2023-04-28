package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	//Format has to given like that only
	fmt.Println(presentTime.Format("01-02-2006"))
	fmt.Println(presentTime.Format("01-02-2006 Monday"))
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2023, time.March, 23, 20, 04, 05, 0, time.UTC)
	fmt.Println(createdDate)

	fmt.Println(createdDate.Format("01-02-2006 Monday"))

	//To build app for windows, it will create .exe file GOOS="windows" go build
}
