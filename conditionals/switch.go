package main

import "fmt"

func main() {

	switch "Kubernetes" {
	case "Docker":
		fmt.Println("Docker is a containerization engine.")
	case "Kubernetes":
		fmt.Println("Kubernetes is a containerization engine.")
		fallthrough
	case "Swarm":
		fmt.Println("Swarm is a containerization engine.")
	default:
		fmt.Println("Unknown.")
	}
}
