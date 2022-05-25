package main

import "fmt"

func main() {

	mySlice := []int{1, 2, 3, 4, 5}
	fmt.Println(mySlice)

	newSlice := []int{10, 20, 30}
	mySlice = append(mySlice, newSlice...)

	fmt.Printf("mySlice NOW contains %d Length is %d and capacity is %d\n", mySlice, len(mySlice), cap(mySlice))
}
