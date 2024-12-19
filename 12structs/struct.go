package main

import "fmt"

func main(){
	fmt.Println("Structs In Golang")
	// no inheritance in golang; No super or parent

	avichal := User{"Avichal", "Avichal@go.dev", true, 16}
	fmt.Println(avichal)
	fmt.Printf("Avichal details are: %v\n", avichal)
 
}

type User struct {
	Name string
	Email string
	Status bool
	Age int
}
