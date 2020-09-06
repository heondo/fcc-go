package main

import "fmt"

func main() {
	// i := 0
	// for i < 10 {
	// 	fmt.Println(i)
	// 	i = i + 1
	// }

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }
	// for i := 1; i <= 100; i++ {
	// 	if i%3 == 0 {
	// 		fmt.Println(i)
	// 	}
	// }

	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz", i)
		} else if i%5 == 0 {
			fmt.Println("Buzz", i)
		}
	}
}
