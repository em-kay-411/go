package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func multiply(x, y int) (op1 int, op2 int, ans int) {
	op1 = x
	op2 = y
	ans = x * y
	return
}

func divide(x int, y int) (int, int) {
	return x / y, x % y
}

func main() {
	a := 5
	b := 2

	fmt.Println("Addition : ", add(a, b))
	fmt.Println("Subtraction : ", subtract(a, b))

	_, _, ans := multiply(a, b)

	fmt.Println("Multiplication : ", ans)

	quotient, remainder := divide(a, b)
	fmt.Println("Division : ", quotient, remainder)
}
