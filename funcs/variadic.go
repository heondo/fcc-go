package main

import "fmt"

func add(a ...int) int {
	total := 0
	for _, val := range a {
		total += val
	}
	return total
}

func main() {
	xs := []int{1, 2, 3, 5, 6}

	fmt.Println(add(xs...))
}
