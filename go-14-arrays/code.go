package main

import "fmt"

func main() {

	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("%T\n", arr)

	arr1 := [...]string{"hello", "world", "!"}
	fmt.Printf("%T\n", arr1)
	fmt.Println(len(arr1))
}
