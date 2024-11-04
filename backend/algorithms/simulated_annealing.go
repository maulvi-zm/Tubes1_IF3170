package algorithms

import (
	"be/class"
	"fmt"
	"math"
	"time"
)

const (
	T0                    = 100
	alpha                 = 0.9
	threshold             = 0.5
	TEMPERATURE_THRESHOLD = 7
)

func getCurrentTemperature(k int) float64 {
	floatK := float64(k)
	return T0 / (1 + alpha*math.Log(1+floatK))
}

func getProbability(deltaE int, temperature float64) float64 {
	floatDeltaE := float64(deltaE)
	return math.Exp(floatDeltaE / temperature)
}

func isAcceptance(probability float64) bool {
	return probability > threshold
}

func SimulatedAnnealing() class.Solution {
	currentCube := class.NewCube(5)
	currentCube.SetRandomStartState()
	currentScore := currentCube.GetCurrentScore()

	timeStart := time.Now()

	fmt.Println("Initial Score: ", currentScore)
	res := class.NewSolution()
	res.SetType("Simulated Annealing")
	res.AddSolutionItem(0, currentScore, currentCube.GetCurrentState())

	i := 1

	stuck := 0

	for {
		successor := currentCube.GetRandomSuccessor()
		successorScore := successor.GetCurrentScore()

		scoreDifference := successorScore - currentScore

		currentTemperature := getCurrentTemperature(i)

		if currentTemperature < TEMPERATURE_THRESHOLD || currentScore == 0 {
			break
		}

		probability := getProbability(scoreDifference, currentTemperature)
		if scoreDifference > 0 {
			currentCube = successor.CopyCube()
			currentScore = successorScore

		} else {
			stuck++

			if isAcceptance(probability) {
				currentCube = successor.CopyCube()
				currentScore = successorScore
			}
		}

		if i%10000 == 0 {
			res.AddSolutionItem(i, currentScore, currentCube.GetCurrentState(), probability)
			fmt.Println("Iteration: ", i, " Score: ", currentScore)
		}

		i++
	}

	if i%10000 != 0 {
		res.AddSolutionItem(i, currentScore, currentCube.GetCurrentState())
	}

	elapsedTime := time.Since(timeStart).Milliseconds()
	res.AddElapsedTime(float64(elapsedTime))
	res.AddLastScore(currentScore)
	res.AddAdditionalInfo("Stuck in Local Optimum Frequence", float64(stuck))

	return *res
}
