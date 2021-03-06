package hogwarts

import (
	"fmt"
)

var sumGryff int = 0
var gScore [40]int //using array data stucture to store IDs and scores of students in Gryffindor group to minimize the number of variables and lines of code

func GryffScores() {

	//using for loop to populate the array and print student Id and score
	for i := 0; i < len(gScore); i++ {
		gScore[i] = (i + 1) * 5

		fmt.Println("Score of Student Id:", i+1, "of group Gryffindor is ", gScore[i])
		sumGryff = sumGryff + gScore[i]
	}
	fmt.Println("***************")

}

func GryffTotal() {
	fmt.Println("Total points of group Gryffindor is: ", sumGryff)
}
func GryffRewards() {
	sumGryff = 0
	gScore[8] = gScore[8] + 4
	gScore[26] = gScore[26] + 4
	gScore[14] = gScore[14] + 4

	//using for loop to calculate new total and print student ID and scores
	for i := 0; i < len(gScore); i++ {
		fmt.Println("Score of Student Id:", i+1, "of group Gryffindor is ", gScore[i])
		sumGryff = sumGryff + gScore[i]
	}
	fmt.Println("***************")
}
func GryffAverage() {

	//using float32 function to convert sumGryff variable to a floating point in an attempt to get accurate results including decimals
	var average float32 = float32(sumGryff) / 40
	fmt.Println("Average points of Gryffindor is: ", average)

}
