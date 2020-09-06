// // / hash maps....

// package main

// import "fmt"

// func main() {
// 	// x := make(map[string]string)

// 	// x["key"] = "hello"

// 	// delete(x, "key")

// 	// fmt.Println(x)

// 	elements := map[string]map[string]string{
// 		"H": map[string]string{
// 			"name":  "Hydrogen",
// 			"state": "gas",
// 		},
// 		"He": map[string]string{
// 			"name":  "Helium",
// 			"state": "gas",
// 		},
// 		"Li": map[string]string{
// 			"name":  "Lithium",
// 			"state": "solid",
// 		},
// 		"Be": map[string]string{
// 			"name":  "Beryllium",
// 			"state": "solid",
// 		},
// 		"B": map[string]string{
// 			"name":  "Boron",
// 			"state": "solid",
// 		},
// 		"C": map[string]string{
// 			"name":  "Carbon",
// 			"state": "solid",
// 		},
// 		"N": map[string]string{
// 			"name":  "Nitrogen",
// 			"state": "gas",
// 		},
// 		"O": map[string]string{
// 			"name":  "Oxygen",
// 			"state": "gas",
// 		},
// 		"F": map[string]string{
// 			"name":  "Fluorine",
// 			"state": "gas",
// 		},
// 		"Ne": map[string]string{
// 			"name":  "Neon",
// 			"state": "gas",
// 		},
// 	}

// 	// name, ok := elements["Un"]
// 	// fmt.Println(name, ok)

// 	key := "Ne"

// 	if el, ok := elements[key]; ok {
// 		fmt.Println("This was found", el["name"], el["state"])
// 	}

// 	// fmt.Println(elements["Li"])
// }

// // when creating arrays and objects, use the type casts and curly braces to define elements
