package algorithms

import (
	"be/class"
	"time"
)

func HillClimbingStochastic(maxIter int) class.Solution {
	currentCube := class.NewCube(5)
	currentCube.SetRandomStartState()
	currentScore := currentCube.GetCurrentScore()

	timeStart := time.Now()

	res := class.NewSolution()
	res.SetType("Stochastic Hill Climbing")
	res.AddSolutionItem(0, currentScore, currentCube.GetCurrentState())

	i := 1

	for i <= maxIter {
		randomSuccessor := currentCube.GetRandomSuccessor()
		randomSuccessorScore := randomSuccessor.GetCurrentScore()

		currentCube = randomSuccessor.CopyCube()
		if currentScore < randomSuccessorScore {
			currentScore = randomSuccessorScore
		}
		res.AddSolutionItem(i, currentScore, currentCube.GetCurrentState())
		i++
	}

	// Add additional info
	elapsedTime := time.Since(timeStart).Milliseconds()
	res.AddElapsedTime(float64(elapsedTime))
	res.AddLastScore(currentScore)
	res.AddAdditionalInfo("Iteration count until algorithm halts", float64(i-1))

	return *res
}
