package main

import "fmt"

func main() {
	x := [6]float64{1, 2, 3, 4, 5, 6}

	someSlice := x[1:3]

	fmt.Println(someSlice)
	// why doesnt javascript put this in their language
}
