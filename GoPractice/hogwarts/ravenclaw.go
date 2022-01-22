package hogwarts

import (
	"fmt"
)

var sumRaven int = 0
var ravScore [40]int //using array data stucture to store IDs and scores of students in Ravenclaw group to minimize the number of variables and lines of code

func RavScores() {

	//using for loop to populate the array and print student Id and score
	for i := 0; i < len(ravScore); i++ {
		ravScore[i] = (i + 1) * 3

		fmt.Println("Score of Student Id:", i+1, "of group Ravenclaw is ", ravScore[i])
		sumRaven = sumRaven + ravScore[i]
	}
	fmt.Println("***************")

}

func RavTotal() {
	fmt.Println("Total points of group Ravenclaw is: ", sumRaven)
}
func RavQuiddichPenalty() {
	sumRaven = 0

	//using for loop to calculate new total and print student ID and scores
	for i := 0; i < len(ravScore); i++ {
		ravScore[i] -= 1
		fmt.Println("Score of Student Id:", i+1, "of group Ravenclaw is ", ravScore[i])
		sumRaven = sumRaven + ravScore[i]
	}
	fmt.Println("***************")
}
func RavAverage() {

	//using float32 function to convert sumRaven variable to a floating point in an attempt to get accurate results including decimals
	var average float32 = float32(sumRaven) / 40
	fmt.Println("Average points of Ravenclaw is: ", average)

}
