package algorithms

import (
	"be/class"
)

// Fungsi Hill Climbing Sideways Move
func HillClimbingSideways(maxIter int) class.Solution {
	currentCube := class.NewCube(5)
	currentCube.SetRandomStartState()
	currentScore := currentCube.GetCurrentScore()

	res := class.NewSolution()
	res.SetType("Hill Climbing with Sideways Move")
	res.AddSolutionItem(0, currentScore, currentCube.GetCurrentState())

	i := 1

	maxCheck := 0

	for maxCheck != maxIter {

		bestSuccessor := currentCube.GetBestSuccessor()
		bestSuccessorScore := bestSuccessor.GetCurrentScore()

		if bestSuccessorScore == currentScore {
			maxCheck++
		}

		if bestSuccessorScore >= currentScore {
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
