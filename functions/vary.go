package main

import "fmt"

func main() {
	bestFinish := championshipResults(12, 5, 6, 4, 3, 3)

	fmt.Println("The best finish is", bestFinish)
}

func championshipResults(finish ...int) int {
	bestFinish := finish[0]

	for _, f := range finish {
		if f < bestFinish {
			bestFinish = f
		}
	}
	return bestFinish
}
