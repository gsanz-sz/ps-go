package main

import "fmt"

func main() {
	//dddLengthMins := 275
	//cawLengthMins := 30

	if dddLengthMins, cawLengthMins := 275, 30; dddLengthMins > cawLengthMins {
		fmt.Println("Docker Deep Dive is longer than Containers on AWS Wavelength.")
		if dddLengthMins > 240 {
			fmt.Println("Still longer")
		}
	} else if dddLengthMins == cawLengthMins {
		fmt.Println("Docker Deep Dive and Containers on AWS Wavelength are the same length.")
	} else {
		fmt.Println("Containers on AWS Wavelength must be longer than Docker Deep Dive.")
	}
}
