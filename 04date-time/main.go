package main

import (
	"fmt"
	"time"
)


func main(){
	fmt.Println("Welcome to date and time")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 Monday"))

	createDate := time.Date(2020, time.May, 22, 23, 23, 0, 0, time.UTC)
	fmt.Println(createDate.Format("01-02-2006 Monday"))
}

//We can build a standalone program in linux and windows in go very easily.