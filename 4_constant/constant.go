package main

import (
	"fmt"
	"math"
)

const s string = "constant"

// Constant expressions perform arithmetic with arbitrary precision.

func main() {
	fmt.Println(s)

	const n = 500000000

	const e = 3e20 / n
	fmt.Println("e = ", e)
	fmt.Println("e = ", int64(e))

	fmt.Println(math.Sin(n))
	/*
		A number can be given a type by using it in a context that requires one,
		such as a variable assignment or function call.
		For example, here math.Sin expects a float64
	*/
}
