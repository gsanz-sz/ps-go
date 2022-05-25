package main

import "fmt"

func main() {
	mySlice := []string{"a", "b", "c", "d", "e"}
	fmt.Println(mySlice[3])

	sliceOfSlice := mySlice[1:3]
	fmt.Println(sliceOfSlice)
}
