package main

import "fmt"

func main() {
	courses := make([]string, 5, 10)
	courses[0] = "Go Fundamentals"
	courses[1] = "Go Fundamentals Part 2"
	courses[2] = "Go Fundamentals Part 3"

	fmt.Printf("Length of slice is %d and capacity is %d", len(courses), cap(courses))

	fmt.Println(courses)

	for _, i := range courses {
		fmt.Println(i)
	}
}
