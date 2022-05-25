package main

import (
	"fmt"
	"os"
)

func main() {

	_, err := os.Open("./test1.txt")

	if err != nil { // if err is not nil, then there is an error
		fmt.Println("This is the code error:", err)
	}
}
