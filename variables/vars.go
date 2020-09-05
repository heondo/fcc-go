package main

import "fmt"

func main() {
	// x := 8
	// fmt.Println("My first variable, well not really but whatever", x)
	// var (
	// 	x = 10
	// 	y = 12
	// 	z = 15
	// )
	fmt.Println("Let me multiply convert C => F for you")
	var celciusInput float64
	fmt.Scanf("%f", &celciusInput)

	output := (celciusInput * (9.0 / 5.0)) + 32

	fmt.Println(output)
}

// variables are pretty straight forward, im not sure when i should use var
// use const for static variables
