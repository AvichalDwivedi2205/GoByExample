package main

import (
	"fmt"
)

func main(){
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println(a)

	arr:=[10]int {1,2,3,4,5,6,7,8,9}
	fmt.Println(arr)

	arr1:=[3]bool{true,false}
	fmt.Println(arr1)

	length := len(arr)
	fmt.Println("The length o fthe array is", length)

}