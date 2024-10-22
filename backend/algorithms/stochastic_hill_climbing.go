package algorithms

// Fungsi Stochastic Hill Climbing
func stochasticHillClimbing(initialState Cube, maxIter int) Solution {
	currentCube := initialState
	currentScore := currentCube.GetCurrentScore()

	res := NewSolution()
	res.AddSolutionItem(0, currentScore, currentCube.currentState)

	i := 1

	// Iterasi sampai mencapai kondisi optimal atau maksimum iterasi
	for {
		if i+1 == maxIter {
			break
		}

		randomSuccessor := currentCube.GetRandomSuccessor()
		randomSuccessorScore := randomSuccessor.GetCurrentScore()

		// Jika tetangga acak lebih baik, pindah ke tetangga tersebut
		currentCube = *randomSuccessor.CopyCube()
		currentScore = randomSuccessorScore
		res.AddSolutionItem(i, currentScore, currentCube.currentState)
		i++
	}

	return *res
}
