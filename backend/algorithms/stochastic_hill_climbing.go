package algorithms

import (
	"be/class"
)

// Fungsi Stochastic Hill Climbing
func stochasticHillClimbing(initialState class.Cube, maxIter int) class.Solution {
	currentCube := initialState
	currentScore := currentCube.GetCurrentScore()

	res := class.NewSolution()
	res.AddSolutionItem(0, currentScore, currentCube.GetCurrentState())

	i := 1

	// Iterasi sampai mencapai kondisi optimal atau maksimum iterasi
	for i+1 != maxIter {

		randomSuccessor := currentCube.GetRandomSuccessor()
		randomSuccessorScore := randomSuccessor.GetCurrentScore()

		// Jika tetangga acak lebih baik, pindah ke tetangga tersebut
		currentCube = *randomSuccessor.CopyCube()
		currentScore = randomSuccessorScore
		res.AddSolutionItem(i, currentScore, currentCube.GetCurrentState())
		i++
	}

	return *res
}
