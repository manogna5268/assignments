package hogwartss

import "fmt"

type Group struct {
	Student    map[int]int
	TotalScore int
	Average    float32
}

func HogScores() {

	var gryffindor Group
	gryffindor.TotalScore = 0
	gryffindor.Average = 0.0
	gryffindor.Student = make(map[int]int)
	for i := 1; i <= 40; i++ {
		gryffindor.Student[i] = i * 2
		gryffindor.TotalScore = gryffindor.TotalScore + gryffindor.Student[i]
		fmt.Println("score of StudentID", i, "is ", gryffindor.Student[i])

	}
	fmt.Println("Total Score of gryffindor is", gryffindor.TotalScore)
	gryffindor.Average = float32(gryffindor.TotalScore) / float32(len(gryffindor.Student))
	fmt.Println("Average of gryffindor is ", gryffindor.Average)

	var slytherin Group
	slytherin.TotalScore = 0
	slytherin.Average = 0.0
	slytherin.Student = make(map[int]int)
	for i := 1; i <= 40; i++ {
		slytherin.Student[i] = i * 3
		slytherin.TotalScore = slytherin.TotalScore + slytherin.Student[i]
		fmt.Println("score of StudentID", i, "is ", slytherin.Student[i])

	}
	fmt.Println("Total Score of slytherin is", slytherin.TotalScore)
	slytherin.Average = float32(slytherin.TotalScore) / float32(len(slytherin.Student))
	fmt.Println("Average of slytherin is ", slytherin.Average)

}
