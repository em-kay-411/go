package main

import "fmt"

func main() {

	// Anonymous Function
	y := func() string {
		return "Hello World"
	}()
	fmt.Println(y)

	// Assigning the anonymous function to a variable
	x := func() {
		fmt.Println("Hello World")
	}
	x()
}
