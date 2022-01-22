package hogwarts

import (
	"fmt"
)

var sumSlyth int = 0
var slScore [40]int //using array data stucture to store IDs and scores of students in Slytherin group to minimize the number of variables and lines of code

func SlythScores() {

	//using for loop to populate the array and print student Id and score
	for i := 0; i < len(slScore); i++ {
		slScore[i] = (i + 1) * 2

		fmt.Println("Score of Student Id:", i+1, "of group Slytherin is ", slScore[i])
		sumSlyth = sumSlyth + slScore[i]
	}
	fmt.Println("***************")

}

func SlythTotal() {
	fmt.Println("Total points of group Slytherin is: ", sumSlyth)
}
func SlythPenalty() {
	sumSlyth = 0
	slScore[9] -= 4
	slScore[24] -= 4
	slScore[28] -= 4

	//using for loop to calculate new total and print student ID and scores
	for i := 0; i < len(slScore); i++ {
		fmt.Println("Score of Student Id:", i+1, "of group Slytherin is ", slScore[i])
		sumSlyth = sumSlyth + slScore[i]
	}
	fmt.Println("***************")
}
func SlythAverage() {

	//using float32 function to convert sumSlyth variable to a floating point in an attempt to get accurate results including decimals
	var average float32 = float32(sumSlyth) / 40
	fmt.Println("Average points of Slytherin is: ", average)

}
