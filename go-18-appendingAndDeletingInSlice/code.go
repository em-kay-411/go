package main

import "fmt"

func main() {
	slc := []int{1, 2, 3, 4}

	// Appending to the slice
	slc = append(slc, 5, 6, 7, 8)
	fmt.Println(slc)

	// Deleting from a slice
	slc = append(slc[:2], slc[3:]...)
	fmt.Println(slc)
}
