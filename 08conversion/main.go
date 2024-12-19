package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	fmt.Println("Welcome to our pizza app")
	fmt.Println("please rate our pizza between 1 and 5")

	reader:=bufio.NewReader(os.Stdin)

	input,_ := reader.ReadString('\n')
	fmt.Println("Thanks for rating", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	//We are using this TrimSpace because of some extra special characters that were there in the end of the program.
	//This was the error
	//strconv.ParseFloat: parsing "4\r\n": invalid syntax
	//The \r\n were removed by the TrimSpace

	if err!=nil {
		fmt.Println(err)
	}else{
		fmt.Println("Added 1 to your rating ", numRating+1)
	}

}