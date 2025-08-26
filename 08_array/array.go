package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("emp:", a)
	
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	// declare inline arrays
	b := [5]int{1,2,3,4,5}
	fmt.Println("dcl:", b)

	b = [...]int{1,2,3,4,5} // [...] compiler count the len(`array`)
	fmt.Println("dcl:", b)

	b = [...]int{100, 3: 400, 500} // specify index, elm in between zero-valued
	fmt.Println("idx:", b)

    var twoD [2][3]int // two dim array
    for i := range 2 {
        for j := range 3 {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)

	// init 2d array
	twoD = [2][3]int {
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d:", twoD)
}
