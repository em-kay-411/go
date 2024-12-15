package main

import "fmt"

func main() {
	a := 6

	/*
		Relational Operators
			 ==
			 >=
			 <=
			 <
			 >
		Logical Operators
			&&
			||
			!
	*/

	if a > 5 && a < 10 {
		fmt.Println("a is greater than 10")
	} else if a >= 10 && a <= 20 {
		fmt.Println("a is 10")
	} else {
		fmt.Println("a is less than 10")
	}
}
