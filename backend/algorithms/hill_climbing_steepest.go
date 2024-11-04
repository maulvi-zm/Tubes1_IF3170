package algorithms

import (
	"be/class"
	"time"
)

func HillClimbingSteepest() class.Solution {
	currentCube := class.NewCube(5)
	currentCube.SetRandomStartState()
	currentScore := currentCube.GetCurrentScore()

	timeStart := time.Now()

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

	// Add additional info
	elapsedTime := time.Since(timeStart).Milliseconds()
	res.AddElapsedTime(float64(elapsedTime))
	res.AddLastScore(currentScore)
	res.AddAdditionalInfo("Iteration count until algorithm halts", float64(i-1))

	return *res
}
