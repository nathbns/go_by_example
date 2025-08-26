package main

import "fmt"

func main() {
	i := 1
	for i <= 3 { // classic loop 1
		fmt.Println("i = ", i)
		i++
	}

	for j := 0; j < 3; j++ { // classic loop 2
		fmt.Println("j = ", j)
	}

	for i := range 3 { // more modern loop 1
		fmt.Println("i = ", i)
	}

	i = 0
	for {
		if i == 2 {
			break
		}
		fmt.Println("loop")
		i++
	}
	/*
		for without a condition will loop repeatedly
		until you break out of the loop or return from
		the enclosing function.
	*/
}
