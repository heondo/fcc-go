// package main

// import "fmt"

// func main() {
// 	// x := [6]float64{1, 2, 3, 4, 5, 6}

// 	// someSlice := x[:]

// 	// slice2 := make([]float64, 5)

// 	slice2 := make([]float64, 3, 4) // this is a slice of 3 an underlying array of 9 elements
// 	// so it can only get so big

// 	slice2 = append(slice2, 4, 5, 6)
// 	fmt.Println(slice2)
// 	// fmt.Println(someSlice)
// 	// fmt.Println(append(someSlice, 4, 5, 6))
// 	// why doesnt javascript put this in their language
// }

// // // this is a bit weird to me...have to use arrays when i know the exact length and dont realyl want to modify
// // // use slices when i want, seems like i shouldnt modify the length whenever possible.
