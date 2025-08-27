package main

import "fmt"

func sum(nums ...int) {
	fmt.Println(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

/*
Within the function, the type of `nums` is equivalent to []int
*/

func main() {
	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	sum(nums...) // applying with a slice
}
