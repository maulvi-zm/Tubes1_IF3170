package algorithms

import "strconv"

func GeneticAlgorithm(populationNum int, iteration int) string {
	return "Genetic Algorithm completed. Population: " + strconv.Itoa(populationNum) + ", Iteration: " + strconv.Itoa(iteration)
}

