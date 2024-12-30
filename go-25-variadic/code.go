package main

import "fmt"

func sum(values ...int) int {
	ans := 0
	for _, value := range values {
		ans += value
	}

	return ans
}

func main() {
	x := sum(1, 2, 3, 4, 5)
	fmt.Println(x)
}
