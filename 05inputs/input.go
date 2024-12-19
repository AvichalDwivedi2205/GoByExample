package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	welcome:= "Welcome to userinput"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter the rating for our pizza")

	//comma ok syntax
	//input, err := reader.readString('\n')
	//This will be comma error syntax and will be use if there are any errors
	input, _:=reader.ReadString('\n')
	fmt.Println("Thanks for rating,", input)
	fmt.Printf("type of the rating is %T", input)
}