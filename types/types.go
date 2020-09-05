package main

import "fmt"

func main() {
	fmt.Println(len("hello this is i love the len function"))
	fmt.Println("Something with a \n new line"[10]) // returns the byte representation of the value
	fmt.Println("hello"[0])                         // returns the byte representation of the value
	fmt.Println("hhllo"[1])                         // returns the byte representation of the value
}

// integer types uint8, uint16, uint32, uint64, int8, int16, int32 and int64
// float types float32 and float64, complex64, complex128
