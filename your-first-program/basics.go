package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, world...")
	os.Exit(404)
}

// FMT library has package fmt
// that is because it is a library and not a program
// programs need main sto run in the terminal or after build
// we import "fmt" because package fmt exists
