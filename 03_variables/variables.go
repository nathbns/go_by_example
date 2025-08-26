package main

import "fmt"

func main() {
	var a = "first variable"
	fmt.Println("a = ", a)

	var b, c int = 0, 1 // you can either specify type or not.
	fmt.Println("b = ", b, " && c = ", c)
	fmt.Println("b + c = ", b+c)

	var d = true
	fmt.Println("d = ", d)

	var e int
	fmt.Println("e = ", e)

	f := "golang" // another notation to declare, without specify the type.
	fmt.Println("f = ", f)
}
