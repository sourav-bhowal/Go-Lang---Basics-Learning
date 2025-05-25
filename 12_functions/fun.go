package main

// func add(a, b int) int { // func function name, (parameters) and return type
// 	return a + b
// }

// func getName()(string, string, bool) { // multiple return values
// 	return "John Doe", "Hello", true
// }

// Func inside a function can be used to create closures or callbacks
func getGreeting() func(string) string {
	return func(name string) string {
		return "Hello, " + name
	}
}

func process(myfun func(a int) int) {
	myfun(10)
}

func main() {
	// result := add(3, 4)
	// println("The sum is:", result)

	// name, greeting, state := getName()
	// println(greeting, name, state)

	greet := getGreeting()

	println(greet("Alice")) // Output: Hello, Alice

	process(func(a int) int {
		println("Processing number:", a)
		return a * 2 // Example processing, returning double the number
	}) // Passing an anonymous function as an argument

	println("Function executed successfully") // Indicating the end of the program
}
