package main

import "fmt"

func main() {
	testMap := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
		"D": 4,
		"E": 5,
		"F": 6,
		"G": 7,
		"H": 8,
		"I": 9,
	}

	for mapKey, mapValue := range testMap {
		fmt.Printf("Key: %s, Value: %d\n", mapKey, mapValue)
	}

	fmt.Println(testMap["C"])

	delete(testMap, "C")
	fmt.Println(testMap)
}
