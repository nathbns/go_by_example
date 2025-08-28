package main

import(
	"fmt"
	"time"
)

func f(s string) {
	for i := range(3) {
		fmt.Println(s, ":", i)
	}
}

func main() {
	f("direct")
	
	// go routine with our `f` function
	go f("goroutine")

	// go routine with an anonymous function
	go func(msg string) {
		fmt.Println(msg)
	}("anonymous_goroutine")

	time.Sleep(time.Second)
	fmt.Println("done.")
}
