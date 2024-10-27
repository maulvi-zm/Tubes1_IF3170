package algorithms

import (
	"be/class"
)

func HillClimbingSteepest(initialState class.Cube) class.Solution {
	currentCube := initialState
	currentScore := currentCube.GetCurrentScore()

	res := class.NewSolution()
	res.AddSolutionItem(0, currentScore, currentCube.GetCurrentState())

	i := 1

	// Iterasi sampai mencapai kondisi optimal atau maksimum iterasi
	for {
		bestSuccessor := currentCube.GetBestSuccessor()
		bestSuccessorScore := bestSuccessor.GetCurrentScore()

		if bestSuccessorScore < currentScore {
			currentCube = *bestSuccessor.CopyCube()
			currentScore = bestSuccessorScore
			res.AddSolutionItem(i, currentScore, currentCube.GetCurrentState())
			i++
		} else {
			break
		}
	}

	return *res
}
