package main

import "fmt"

func first() {
	fmt.Println("This is function first")
}

func second() {
	fmt.Println("This is function second")
}

func main() {
	defer first()
	second()
}
