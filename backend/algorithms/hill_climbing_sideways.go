package algorithms

import (
	"be/class"
)

// Fungsi Hill Climbing Sideways Move
func HillClimbingSideways(initialState class.Cube, maxIter int) class.Solution {
	currentCube := initialState
	currentScore := currentCube.GetCurrentScore()

	res := class.NewSolution()
	res.AddSolutionItem(0, currentScore, currentCube.GetCurrentState())

	i := 1

	maxCheck := 0

	// Iterasi sampai mencapai kondisi optimal atau maksimum iterasi
	for {
		if maxCheck == maxIter {
			break
		}

		bestSuccessor := currentCube.GetBestSuccessor()
		bestSuccessorScore := bestSuccessor.GetCurrentScore()

		if bestSuccessorScore == currentScore {
			maxCheck++
		}

		if bestSuccessorScore <= currentScore {
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
