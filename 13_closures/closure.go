package main

import (
	"fmt"
)

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func strSeq() func() (string, string) {
	c := 'a'
	str := string(c)
	i := 0
	return func() (string, string) {
		s := string(c)
		if i > 0 {
			str += string(c)
		}

		i++
		c++
		return s, str
	}
}

func main() {
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	
	newInts := intSeq()
	fmt.Println(newInts())

	nextChar := strSeq()
	fmt.Println(nextChar())
	fmt.Println(nextChar())
}
