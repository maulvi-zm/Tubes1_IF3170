package algorithms

import (
	"be/class"
)

func HillClimbingSteepest() class.Solution {
	currentCube := class.NewCube(5)
	currentCube.SetRandomStartState()
	currentScore := currentCube.GetCurrentScore()

	res := class.NewSolution()
	res.SetType("Steepest Ascent Hill Climbing")
	res.AddSolutionItem(0, currentScore, currentCube.GetCurrentState())

	i := 1

	for {
		bestSuccessor := currentCube.GetBestSuccessor()
		bestSuccessorScore := bestSuccessor.GetCurrentScore()

		if bestSuccessorScore > currentScore {
			currentCube = bestSuccessor.CopyCube()
			currentScore = bestSuccessorScore
			res.AddSolutionItem(i, currentScore, currentCube.GetCurrentState())
			i++
		} else {
			break
		}
	}

	return *res
}
