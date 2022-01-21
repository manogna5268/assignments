package howzatt

import "fmt"

var indPlayers = make(map[string]int)
var score int = 0

func FinalScore() {

	for _, value := range indPlayers {
		score = value + score
		fmt.Println(score)
	}

}

func IndianBatsmen() {
	indPlayers["Virat Kohli"] = 90
	indPlayers["Rohith Sharma"] = 30
	indPlayers["Hardik Pandya"] = 40
	indPlayers["Ravindra Jadeja"] = 45
	indPlayers["MS Dhoni"] = 60
	indPlayers["Mayank Agarwal"] = 0
	indPlayers["Jaspreet Bhumra"] = 10
	indPlayers["Mohammed Shami"] = 35
	indPlayers["Jayant Yadav"] = 0
	indPlayers["Harshal Patel"] = 3
	indPlayers["Bhuvaneshwar Kumar"] = 6

	fmt.Println(indPlayers)
	for _, value := range indPlayers {
		score = value + score

	}
	fmt.Println(score)

}

func SLBowlers() {

}

func SLBatsmen() {

}

func IndianBowlers() {

}

func MatchResult() {

}
