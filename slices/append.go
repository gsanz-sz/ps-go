package main

import "fmt"

func main() {
	mySlice := make([]int, 1, 4)
	fmt.Printf("Length is %d with a cap of %d\n", len(mySlice), cap(mySlice))

	for i := 1; i < 17; i++ {
		mySlice = append(mySlice, i)
		fmt.Printf("Length is %d but cap of %d\n", len(mySlice), cap(mySlice))
	}
}
