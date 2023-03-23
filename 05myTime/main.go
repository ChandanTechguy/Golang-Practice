package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()
	fmt.Println(presentTime)
	
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))  // time should be according that format only

	createDate := time.Date(2023 , time.March , 23 ,12,44,0,0, time.Local)
	fmt.Println(createDate)

}