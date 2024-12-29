package main

import "fmt"

func main() {

	arr := [5]int{1, 2, 3, 4, 5}

	for i := 0; i < len(arr); i++ {
		fmt.Printf("Index, %v - %v \n", i, arr[i])
	}

	for index, value := range arr {
		fmt.Printf("Index, %v - %v \n", index, value)
	}
}
