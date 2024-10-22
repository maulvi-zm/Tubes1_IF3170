package algorithms

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

		if bestSuccessorScore < currentScore {
			currentCube = *bestSuccessor.CopyCube()
			currentScore = bestSuccessorScore
			res.AddSolutionItem(i, currentScore, currentCube.currentState)
			i++
		} else {
			break
		}
	}

	return *res
}
