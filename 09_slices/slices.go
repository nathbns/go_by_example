package main 

import (
	"fmt"
	//"slices"
)

// if you already work with dynamic arrays, this should be good enough for you. You can skip.

func main() {
	var s []string 
	fmt.Println("uninit:", s, s == nil, len(s) == 0) // uninitialize slice is equal to nil and len == 0

	s = make([]string, 3) // use the builtin `make` to create slice with non-zero length
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("append:", s)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("c:", c)

	l := s[2:5]
	fmt.Println("l:", l)

	twoD := make([][]int, 3)
	for i := range 3 {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := range innerLen {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}
