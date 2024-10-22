package algorithms

import (
	"fmt"
)

// Fungsi Hill Climbing Steepest Ascent
func HillClimbingSteepest(initialState Cube) Solution {
	currentCube := initialState
	currentScore := currentCube.GetCurrentScore()

	res := NewSolution()
	res.AddSolutionItem(0, currentScore, currentCube.currentState)

	i := 1

	// Iterasi sampai mencapai kondisi optimal atau maksimum iterasi
	for {
		bestSuccessor := currentCube.GetBestSuccessor()
		bestSuccessorScore := bestSuccessor.GetCurrentScore()

		// If the best successor is better than the current cube, move to the new state
		if bestSuccessorScore < currentScore {
			currentCube = *bestSuccessor.CopyCube()
			currentScore = bestSuccessorScore
			res.AddSolutionItem(i, currentScore, currentCube.currentState)

			i++

			// Output the current score (optional, for tracking progress)
			fmt.Println("Current score:", currentScore)
		} else {
			// If no better successor is found, stop the algorithm
			break
		}
	}

	return *res
}
