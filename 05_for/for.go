package main

import "fmt"

func main() {
	// for loop
	fmt.Println("For loop")
	forLoop()

	// for loop normal
	fmt.Println("For loop normal")
	forLoopNormal()

	// while loop
	fmt.Println("While loop")
	whileLoop()

	// do while loop
	fmt.Println("Do while loop")
	doWhileLoop()

	// for loop with if
	fmt.Println("For loop with if")
	forLoopWithIf()

	// for loop with switch
	fmt.Println("For loop with switch")
	forLoopWithSwitch()
}

// for loop with range
func forLoop() {
	for i := range 10 {
		fmt.Println(i)
	}
}

// for loop normal
func forLoopNormal() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

// while loop
func whileLoop() {
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}
}

// do while loop
func doWhileLoop() {
	i := 0
	for {
		fmt.Println(i)
		i++
		if i >= 5 {
			break
		}
	}
}

// for loop with if
func forLoopWithIf() {
	for i := range 5 {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}
}

// for loop with switch
func forLoopWithSwitch() {
	for i := range 2 {
		switch i {
		case 0:
			fmt.Println("Zero")
		case 1:
			fmt.Println("One")
		case 2:
			fmt.Println("Two")
		
		}
	}
}