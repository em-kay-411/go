package main

import "fmt"

func main() {
	slc := []int{1, 2, 3, 4, 5}
	fmt.Println(len(slc))
	fmt.Printf("%T\n", slc)

	for i := 0; i < len(slc); i++ {
		fmt.Println(slc[i])
	}

	for _, value := range slc {
		fmt.Println(value)
	}
}
