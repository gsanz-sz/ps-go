package main

import (
	"fmt"
)

func main() {
	name := "Sanz"
	course := "Studying go"

	fmt.Println("Name and course are set to", name, "and", course)
	updateCourse(course)

	fmt.Println("Name and course now are set to", name, "and", course)
}

func updateCourse(course string) string {
	course = "Studying go with GoLang"
	fmt.Println("Course is set to", course)
	return course
}
