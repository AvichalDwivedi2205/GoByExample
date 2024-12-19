package main

import(
	"fmt"
	"math"
	"reflect"
)

const s = "sex"

func main(){
	fmt.Println(s)

	const n = 500000000000

	const d = 3e20/n
	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
	fmt.Println(reflect.TypeOf(s))
	fmt.Println(reflect.TypeOf(n))
	fmt.Println(reflect.TypeOf(d))
}

//const n int = 500000000 will error
// //Here's the explanation:

// Implicit Type: When n is implicitly typed, the line math.Sin(n) is effectively treated as math.Sin(float64(n)), allowing the code to compile and run.
// Explicit Type: When n is explicitly typed as int, it cannot be automatically converted to float64. This leads to a type mismatch error since math.Sin requires a float64 parameter.
