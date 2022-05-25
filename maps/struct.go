package main

import "fmt"

func main() {

	type courseMeta struct {
		author string
		level  string
		rating float64
	}

	gettingStartedWithK8s := courseMeta{
		author: "Gabriel Sanz",
		level:  "Beginner",
		rating: 4.5,
	}

	fmt.Println(gettingStartedWithK8s)

	fmt.Println("Author of Getting Started with Kubernetes is", gettingStartedWithK8s.author)
	gettingStartedWithK8s.rating = 10
	fmt.Println("Rating of Getting Started with Kubernetes is", gettingStartedWithK8s.rating)

}
