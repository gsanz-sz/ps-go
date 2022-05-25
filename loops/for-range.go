package main

import "fmt"

func main() {

	coursesInProgr := []string{"go", "python", "ruby"}

	coursesCompleted := []string{"go", "C#", "python"}

	for _, i := range coursesInProgr {
		fmt.Println(i)
		for _, j := range coursesCompleted {
			if i == j {
				fmt.Println("Hey we found a clash", i, "is in both lists")
			}
		}
	}
}
