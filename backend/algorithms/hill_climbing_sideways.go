package algorithms

// Fungsi Hill Climbing Sideways Move
func HillClimbingSideways(initialState Cube, maxIter int) Solution {
	currentCube := initialState
	currentScore := currentCube.GetCurrentScore()

	res := NewSolution()
	res.AddSolutionItem(0, currentScore, currentCube.currentState)

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
			res.AddSolutionItem(i, currentScore, currentCube.currentState)

			i++

		} else {
			break
		}
	}

	return *res
}
