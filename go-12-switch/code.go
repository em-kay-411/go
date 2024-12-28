package main

import "fmt"

func main() {
	x := 3

	switch x {
	case 3:
		fmt.Println("The number is 3")
		fallthrough

	case 4:
		fmt.Println("The number is 4")

	case 5:
		fmt.Println("The number is 5")

	case 6:
		fmt.Println("The number is 6")

	case 7:
		fmt.Println("The number is 7")

	case 8:
		fmt.Println("The number is 8")
	default:
		fmt.Println("This is a number")

	}
}
