package hogwarts

import (
	"fmt"
)

var sumRaven int = 0
var ravScore [40]int

func RavScores() {

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

	for i := 0; i < len(ravScore); i++ {
		ravScore[i] -= 1
		fmt.Println("Score of Student Id:", i+1, "of group Ravenclaw is ", ravScore[i])
		sumRaven = sumRaven + ravScore[i]
	}
	fmt.Println("***************")
}
func RavAverage() {
	var average float32 = float32(sumRaven) / 40
	fmt.Println("Average points of Ravenclaw is: ", average)

}
