package main

import "fmt"

func main() {
	slc := []int{1, 2, 3, 4, 5}

	// slice = slc[initial index : final index]

	fmt.Println(slc[1:3])
	fmt.Println(slc[2:])
	fmt.Println(slc[:3])
}
