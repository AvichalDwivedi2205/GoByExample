package main

import "fmt"

func main(){

	i:=1
	for i<=3{
		fmt.Println(i)
		i+=1
	}

	for j:=0; j<3;j++{
		fmt.Println(j)
	}

	for {
		fmt.Println("ok bro")
		break;
	}

	for i:= range 'a'{
		fmt.Println(i)
	}

	

}