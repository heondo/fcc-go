package main

import "fmt"

func main() {
	slice1 := []float64{
		1, 2, 3, 4, 5, 6, 7, 31413,
	}

	slice1Average := findAverage(slice1)

	first, _ := returnMult()

	fmt.Println(slice1Average)
	fmt.Println(first)
}

// func f2() (r int) {
// 	r = 1
// 	return
//   } will automatically return the named variable when the return statement is seen

func findAverage(floatSlice []float64) float64 {
	total := 0.0

	for _, val := range floatSlice {
		total += val
	}

	return total / float64(len(floatSlice))
}

func returnMult() (int, int) {
	return 2, 3
}
