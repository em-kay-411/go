package main

import "fmt"

func main() {

	for i := 0; i <= 20; i++ {
		fmt.Println(i)
	}

	count := 0

	for count < 10 {
		fmt.Println(count)
		count++
	}
}
