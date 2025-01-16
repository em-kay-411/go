package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func subtract(a int, b int) int {
	return a - b
}

func calculate(a int, b int, f func(int, int) int) int {
	return f(a, b)
}

func returningAFunction() func() {
	return func() {
		fmt.Println("Hello world")
	}
}

func main() {
	a := 31
	b := 10

	fmt.Println(calculate(a, b, add))
	fmt.Println(calculate(a, b, subtract))

	x := returningAFunction()
	x()
}
