package main

import "fmt"

const (
	a = 10
	b = 20
	c = 30
)

func main() {
	const a = 10
	fmt.Println(a)
	
	fmt.Println(b, c)

	const (
		d = iota // 0
		e // 1
		f // 2
	)
	fmt.Println(d, e, f)

}
