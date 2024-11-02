package algorithms

import (
	"be/class"
)

func HillClimbingRandomRestart(maxRestart int) class.Solution {
	currentCube := class.NewCube(5)
	currentCube.SetRandomStartState()
	currentScore := currentCube.GetCurrentScore()
	res := class.NewSolution()
	res.AddSolutionItem(0, currentScore, currentCube.GetCurrentState())

	i := 1

	maxIter := 0
	
	for (maxIter != maxRestart){
		bestSuccessor := currentCube.GetBestSuccessor()
		bestSuccessorScore := bestSuccessor.GetCurrentScore()

		if bestSuccessorScore <= currentScore {
			currentCube.SetRandomStartState()
			currentScore = currentCube.GetCurrentScore()
			res.AddSolutionItem(i, currentScore, currentCube.GetCurrentState())
			maxIter++
		} else {
			currentCube = bestSuccessor.CopyCube()
			currentScore = bestSuccessorScore
			res.AddSolutionItem(i, currentScore, currentCube.GetCurrentState())
		}

		i++
	}

	return *res

}
