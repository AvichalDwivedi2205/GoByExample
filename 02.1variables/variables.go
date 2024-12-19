package main

import (
	"fmt"
	"reflect"
)


func main(){
	var a = "SexyBoi69"
	fmt.Println(a)

	var b,c,d = 1,2,3
	fmt.Println(b,c,d)
	fmt.Println(reflect.TypeOf(b))
	//Variables must be used somewhere in golandg or it will show error
	 
	var e = true
	fmt.Println(e)

	//Shortcut for declaring a variable
	f:="Sexyboi69"
	fmt.Println(f)
	fmt.Println(reflect.TypeOf(f))

	g:=2
	fmt.Println(reflect.TypeOf(g))

	f="hh"
	fmt.Println(f)
	//f=8 npt ok as string

	h:=true
	fmt.Println(h)

}