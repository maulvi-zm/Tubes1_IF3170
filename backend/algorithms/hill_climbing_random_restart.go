package algorithms

import (
	"be/class"
	"time"
)

func HillClimbingRandomRestart(maxRestart int) class.Solution {
	currentCube := class.NewCube(5)
	currentCube.SetRandomStartState()
	currentScore := currentCube.GetCurrentScore()
	res := class.NewSolution()
	res.SetType("Random Restart Hill Climbing")
	res.AddSolutionItem(0, currentScore, currentCube.GetCurrentState())

	i := 1
	timeStart := time.Now()

	maxIter := 0
	iterasi := 1
	for maxIter <= maxRestart {
		bestSuccessor := currentCube.GetBestSuccessor()
		bestSuccessorScore := bestSuccessor.GetCurrentScore()

		if bestSuccessorScore <= currentScore {
			maxIter++
			if (maxIter <= maxRestart){
				currentCube.SetRandomStartState()
				currentScore = currentCube.GetCurrentScore()
				res.AddSolutionItem(i, currentScore, currentCube.GetCurrentState())
			}
			res.AddIterasiRestart(iterasi,maxIter-1)
			iterasi = 0	
		} else {
			currentCube = bestSuccessor.CopyCube()
			currentScore = bestSuccessorScore
			if (maxIter <= maxRestart){
				res.AddSolutionItem(i, currentScore, currentCube.GetCurrentState())
			}
		}
		iterasi++
		i++
	}
	elapsedTime := time.Since(timeStart).Milliseconds()
	res.AddTotalRestart(maxIter-1)
	res.AddMaxRestart(maxRestart)
	res.AddElapsedTime(float64(elapsedTime))
	res.AddLastScore(currentScore)


	return *res

}
