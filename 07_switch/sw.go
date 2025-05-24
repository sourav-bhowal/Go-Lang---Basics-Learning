package main

import (
	"fmt"
	"time"
)

func main() {
	// switch number := 1; number {
	// case 1:
	// 	fmt.Println("Number is 1")
	// case 2:
	// 	fmt.Println("Number is 2")

	// }

	// switch with time
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	// switch with time
	// switch {
	// case time.Now().Hour() < 12:
	// 	fmt.Println("It's before noon", time.Now().Hour(), "AM")
	// case time.Now().Hour() >= 12:
	// 	fmt.Println("It's after noon", time.Now().Hour(), "PM")
	// default:
	// 	fmt.Println("It's not a valid time")
	// }

	// Type switch
	// Type switch is a type of switch that is used to check the type of a variable.
	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whoAmI(true)
	whoAmI(1)
	whoAmI("Hello")
}
