package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := make([]int, 5)
	copy(b, a)
	a[0] = 0

	fmt.Println(a)
	fmt.Println(b)
}
