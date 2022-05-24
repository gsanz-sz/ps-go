package main

import (
	"fmt"
	"strings"
)

func main() {
	author := "Gabriel Sanz"
	course := "Go Fundamentals"

	fmt.Println(converter(author, course))
}

func converter(s1, s2 string) (str1, str2 string) {
	s1 = strings.ToUpper(s1)
	s2 = strings.Title(s2)

	return s1, s2
}
