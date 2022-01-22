package howzatt

import "fmt"

var indScore int = 0      //using integer data type as we want to score the value in integer format
var srilankaScore int = 0 //using integer data type as we want to score the value in integer format

func IndianBatsmen() {

	fmt.Println("INDIA VS SRI LANKA CRICKET MATCH")
	fmt.Println()

	indPlayers := make(map[string]int) //using maps data structure to store Indian Batsmen names and their scores, this way we can prevent mismatching of scores to players

	//populating values into map
	indPlayers["Virat Kohli"] = 0
	indPlayers["KL Rahul"] = 55
	indPlayers["Shikhar Dhawan"] = 29
	indPlayers["Rishabh Pant"] = 85
	indPlayers["Shreyas Iyer"] = 11
	indPlayers["Venkatesh Iyer"] = 22
	indPlayers["Jaspreet Bhumra"] = 0
	indPlayers["Shardul Thakur"] = 40
	indPlayers["Ravichandran Ashwin"] = 25
	indPlayers["Yuzvendra Chahal"] = 0
	indPlayers["Bhuvaneshwar Kumar"] = 0

	//using for range loop to calculate the sum of scores of indian batsmen
	for _, value := range indPlayers {

		indScore = value + indScore

	}
	fmt.Println("Final India score is <", indScore, "/6>")
	fmt.Println()
	fmt.Println("INDIAN BATSMEN NAMES AND THEIR RESPECTIVE SCORES ARE:")

	//using for range loop to print names of Indian Batsmen and their score
	for key, value := range indPlayers {
		fmt.Println(key, "--->", value)
	}
	fmt.Println()

}

func SLBowlers() {
	sriLankaBowlers := make(map[string]int) //using maps data structure to store Sri Lankan Bowlers names and their scores, this way we can prevent mismatching of scores to players

	//populating values into map
	sriLankaBowlers["Lungi Ngidi"] = 35
	sriLankaBowlers["Sisanda Magala"] = 64
	sriLankaBowlers["Aiden Markram"] = 34
	sriLankaBowlers["Keshav Maharaj"] = 52
	sriLankaBowlers["Andile Phehlukwayo"] = 44
	sriLankaBowlers["Tabraiz Shamsi"] = 57

	fmt.Println("RUNS CONCEDED BY SRI LAKNKAN BOWLERS ARE:")
	for key, value := range sriLankaBowlers {
		fmt.Println(key, "--->", value)
	}
	fmt.Println()

}

func SLBatsmen() {
	sriLankanBatsmen := make(map[string]int)
	sriLankanBatsmen["Janneman Malan"] = 91
	sriLankanBatsmen["Quinton de Kock"] = 78
	sriLankanBatsmen["Temba Bavuma"] = 35
	sriLankanBatsmen["Aiden Markram"] = 37
	sriLankanBatsmen["Rassie van der Dussen"] = 37
	sriLankanBatsmen["David Miller"] = 0
	sriLankanBatsmen["Andile Phehlukwayo"] = 0
	sriLankanBatsmen["Sisanda Magala"] = 0
	sriLankanBatsmen["Keshav Maharaj"] = 0
	sriLankanBatsmen["Tabraiz Shamsi"] = 0
	sriLankanBatsmen["Lungi Nigdi"] = 0

	fmt.Println("SRI LANKAN BATSMEN NAMES AND THEIR RESPECTIVE SCORES ARE:")

	//using for range loop to calculate the sum of scores of indian batsmen and to print batsmen names and their scores
	for key, value := range sriLankanBatsmen {
		fmt.Println(key, "--->", value)

		srilankaScore = value + srilankaScore

	}
	fmt.Println()
}

func IndianBowlers() {
	indianBowlers := make(map[string]int)
	indianBowlers["Jasprit Bumrah"] = 37
	indianBowlers["Bhuvaneshwar Kumar"] = 67
	indianBowlers["Ravichandran Ashwin"] = 68
	indianBowlers["Yuzvendra Chahal"] = 47
	indianBowlers["Shardul Thakur"] = 35
	indianBowlers["Venkatesh Iyer"] = 28
	indianBowlers["Shreyas Iyer"] = 1

	fmt.Println("RUNS CONCEDED BY INDIAN BOWLERS ARE:")
	for key, value := range indianBowlers {
		fmt.Println(key, "--->", value)
	}
	fmt.Println()
}

func MatchResult() {

	//using if else condition to determine which country won the match based on their scores
	if srilankaScore > indScore {

		fmt.Println("Score of India is", indScore)
		fmt.Println("Score of Sri Lanka is", srilankaScore)
		fmt.Println("Srilanka won the match")

	} else {
		fmt.Println("Score of India is", indScore)
		fmt.Println("Score of Sri Lanka is", srilankaScore)
		fmt.Println("India won the match")
	}

}
