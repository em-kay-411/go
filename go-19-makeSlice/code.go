package main

import "fmt"

func main() {
	slc := make([]int, 0)

	slc = append(slc, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12)
	fmt.Println(slc)
	fmt.Println(len(slc))
	fmt.Println(cap(slc))
}
