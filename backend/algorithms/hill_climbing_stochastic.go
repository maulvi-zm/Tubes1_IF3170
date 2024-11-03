package algorithms

import "be/class"

func HillClimbingStochastic(maxIter int) class.Solution {
	currentCube := class.NewCube(5)
	currentCube.SetRandomStartState()
	currentScore := currentCube.GetCurrentScore()

	res := class.NewSolution()
	res.SetType("Stochastic Hill Climbing")
	res.AddSolutionItem(0, currentScore, currentCube.GetCurrentState())

	i := 1

	for i+1 != maxIter {
		randomSuccessor := currentCube.GetRandomSuccessor()
		randomSuccessorScore := randomSuccessor.GetCurrentScore()

		currentCube = randomSuccessor.CopyCube()
		if currentScore < randomSuccessorScore {
			currentScore = randomSuccessorScore
			res.AddSolutionItem(i, currentScore, currentCube.GetCurrentState())
		}
		i++
	}

	return *res
}
