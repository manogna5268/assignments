package hogwarts

import (
	"fmt"
)

var sumHuffle int = 0
var huffScore [40]int

func HuffleScores() {

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

	for i := 0; i < len(huffScore); i++ {
		huffScore[i] += 1
		fmt.Println("Score of Student Id:", i+1, "of group Hufflepuff is ", huffScore[i])
		sumHuffle = sumHuffle + huffScore[i]
	}
	fmt.Println("***************")
}
func HuffleAverage() {
	var average float32 = float32(sumHuffle) / 40
	fmt.Println("Average points of Hufflepuff is: ", average)

}
