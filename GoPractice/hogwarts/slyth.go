package hogwarts

import (
	"fmt"
)

var sumSlyth int = 0
var slScore [40]int

func SlythScores() {

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

	for i := 0; i < len(slScore); i++ {
		fmt.Println("Score of Student Id:", i+1, "of group Slytherin is ", slScore[i])
		sumSlyth = sumSlyth + slScore[i]
	}
	fmt.Println("***************")
}
func SlythAverage() {
	var average float32 = float32(sumSlyth) / 40
	fmt.Println("Average points of Slytherin is: ", average)

}
