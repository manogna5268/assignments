package hogwarts

import (
	"fmt"
)

var sumHuffle int = 0
var huffScore [40]int //using array data stucture to store IDs and scores of students in Hufflepuff group to minimize the number of variables and lines of code

func HuffleScores() {

	//using for loop to populate the array and print student Id and score
	for i := 0; i < len(huffScore); i++ {
		huffScore[i] = (i + 1) * 4

		fmt.Println("Score of Student Id:", i+1, "of group Hufflepuff is ", huffScore[i])
		sumHuffle = sumHuffle + huffScore[i]
	}
	fmt.Println("***************")

}

func HuffleTotal() {
	fmt.Println("Total points of group Hufflepuff is: ", sumHuffle)
	fmt.Println("***************")
}
func HuffleQuiddichReward() {
	sumHuffle = 0

	//using for loop to calculate new total and print student ID and scores
	for i := 0; i < len(huffScore); i++ {
		huffScore[i] += 1
		fmt.Println("Score of Student Id:", i+1, "of group Hufflepuff is ", huffScore[i])
		sumHuffle = sumHuffle + huffScore[i]
	}
	fmt.Println("***************")
}
func HuffleAverage() {

	//using float32 function to convert sumHuffle variable to a floating point in an attempt to get accurate results including decimals
	var average float32 = float32(sumHuffle) / 40
	fmt.Println("Average points of Hufflepuff is: ", average)
	fmt.Println()
	fmt.Println()

}
