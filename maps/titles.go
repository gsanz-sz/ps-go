package main

import "fmt"

func main() {

	leagueTitles := make(map[string]int)
	leagueTitles["Sunderland"] = 6
	leagueTitles["NewCastle"] = 4

	recentHead2HeadWins := map[string]int{
		"Sunderland": 3,
		"NewCastle":  2,
	}

	fmt.Printf("League titles: %v\nRecent Head to head wins: %v\n", leagueTitles, recentHead2HeadWins)
}
