package main

import (
	"fmt"
	"time"
)

func main() {

	// basic switch
	i := 2
	fmt.Print("write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday: // use comma for multiple args in `case`
		fmt.Println("it's weekend")
	default:
		fmt.Println("it's weekday")
	}

	// use switch like an `if|else` statement
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatami := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("it's a bool")
		case int:
			fmt.Println("it's an integer")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	whatami(true)
	whatami(1)
	whatami(1.0)
	whatami("string")
}
