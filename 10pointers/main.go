package main

import "fmt"

func main(){
	fmt.Println("Welcome to a class on Pointers")

	// var ptr *int
	// fmt.Println("Value of the pointer is", ptr)
	
	myNumber := 23
	var ptr = &myNumber

	fmt.Println("Value of actual pointer is", ptr)
	fmt.Println("Value of actual pointer is", *ptr)
}