package algorithms

// package main

import (
	"be/class"
	"math"
	"strconv"
	"time"

	"math/rand"
)

// /////////////////////////////////// //
// Cube functions for DNA manipulation //
// /////////////////////////////////// //
func WeightedRandom(max int) int {
	rand.Seed(time.Now().UnixNano())

	// Generate a random float in the range [0, 1)
	randomFloat := rand.Float64()

	// Subtract from 1 to flip the distribution, then apply power to skew towards the higher end
	skewed := 1 - math.Pow(randomFloat, 1.5) // Adjust exponent for more or less skew (higher = more weight towards max)

	// Scale to range 0 to max
	return int(skewed * float64(max))
}

func FindNumberInCube(cube *class.Cube, startIndex int, number int) int {
	for i := startIndex; i < 125; i++ {
		if cube.GetSmallCubeValue(i) == number {
			return i
		}
	}
	return -1
}

func Crossover(HigherHalfParent, LowerHalfParent *class.Cube) *class.Cube {
	// Create new cube
	child := class.NewCube(5)

	// Get random index to split the DNA
	splitIndex := WeightedRandom(125)

	// Create map of swapped indices
	swappedIndices := make(map[int]bool)
	for i := splitIndex; i < 125; i++ {
		swappedIndices[i] = false
	}

	// Combine DNA of parents
	for i := 0; i < 125; i++ {
		if i < splitIndex {
			child.SetSmallCubeValue(i, HigherHalfParent.GetSmallCubeValue(i))
		} else {
			// Check if i has been swapped
			if swappedIndices[i] {
				continue
			}

			// Get values of the two parents
			HPValue := HigherHalfParent.GetSmallCubeValue(i) // HP = Higher Half Parent
			LPValue := LowerHalfParent.GetSmallCubeValue(i)  // LP = Lower Half Parent

			// Check if LowerParentValue in HigherParent's second half
			HPIndex := FindNumberInCube(HigherHalfParent, splitIndex, LPValue)

			// If it is, swap i with HPIndex. Else, set i to HPValue
			if HPIndex != -1 && !swappedIndices[HPIndex] {
				child.SetSmallCubeValue(i, LPValue)
				child.SetSmallCubeValue(HPIndex, HPValue)
				swappedIndices[i] = true
				swappedIndices[HPIndex] = true
			} else {
				child.SetSmallCubeValue(i, HPValue)
			}
		}
	}

	child.SetCurrentScore()
	return child
}

func Mutate(cube *class.Cube) *class.Cube {
	// Get two random indices to swap
	index1 := rand.Intn(125)
	index2 := rand.Intn(125)

	// Determine if mutation will occur
	// Mutation rate is 0.3
	rand.Seed(time.Now().UnixNano())
	mutationRate := rand.Float64()

	if mutationRate < 0.3 {
		// Swap the values at the two indices
		cube.SwitchState(index1, index2)
	}

	return cube
}

// //////////////////////////// //
// Population class and methods //
// //////////////////////////// //
type Population struct {
	Cubes         []*class.Cube
	BestCubeIndex int
	AvgScore      int
	PopulationNum int

	// Attributes for fitness function calculation
	MinScore   int
	MaxFitness int
}

func NewPopulation(populationNum int) *Population {
	return &Population{
		Cubes:         make([]*class.Cube, populationNum),
		BestCubeIndex: 0,
		AvgScore:      0,
		PopulationNum: populationNum,
		MinScore:      0,
		MaxFitness:    0,
	}
}

func (p *Population) GetCube(index int) *class.Cube {
	return p.Cubes[index]
}

func (p *Population) GetBestCube() *class.Cube {
	return p.Cubes[p.BestCubeIndex]
}

func (p *Population) GetAvgScore() int {
	return p.AvgScore
}

func (p *Population) GetPopulationNum() int {
	return p.PopulationNum
}

func (p *Population) GetMinScore() int {
	return p.MinScore
}

func (p *Population) GetMaxFitness() int {
	return p.MaxFitness
}

func (p *Population) SetBestCubeIndex(index int) {
	p.BestCubeIndex = index
}

func (p *Population) SetAvgScore(score int) {
	p.AvgScore = score
}

func (p *Population) SetMinScore(score int) {
	p.MinScore = score
}

func (p *Population) SetMaxFitness(fitness int) {
	p.MaxFitness = fitness
}

// Function to get a random cube from the population, with weighted probability based on the cube's fitness
func (p *Population) GetWeightedRandomCube() *class.Cube {
	rand.Seed(time.Now().UnixNano())
	randomFitness := rand.Intn(p.GetMaxFitness())

	// Find random cube based on the random fitness
	currentTotalFitness := 0
	for i := 0; i < p.GetPopulationNum(); i++ {
		fitness := p.GetCube(i).GetCurrentScore() - p.GetMinScore()
		currentTotalFitness += fitness

		if currentTotalFitness >= randomFitness {
			return p.GetCube(i)
		}
	}

	return p.GetCube(p.GetPopulationNum() - 1)
}

func (p *Population) GeneratePopulation() {
	SumScore := 0

	for i := 0; i < p.PopulationNum; i++ {
		p.Cubes[i] = class.NewCube(5)
		p.Cubes[i].SetRandomStartState()

		// Sum scores to get average
		Score := p.Cubes[i].GetCurrentScore()
		SumScore += Score

		// Set best cube of population
		if Score > p.GetBestCube().GetCurrentScore() {
			p.SetBestCubeIndex(i)
		}

		// Set min score of population
		if Score < p.GetMinScore() {
			p.SetMinScore(Score)
		}
	}

	// Set average score of population
	p.SetAvgScore(SumScore / p.GetPopulationNum())

	// Set max fitness of population
	maxFitness := p.GetBestCube().GetCurrentScore() - p.GetMinScore()
	p.SetMaxFitness(maxFitness)
}

func (p *Population) BreedPopulation() *Population {
	newPopulation := NewPopulation(p.GetPopulationNum())
	SumScore := 0

	for i := 0; i < p.GetPopulationNum(); i++ {
		parent1 := p.GetWeightedRandomCube()
		parent2 := p.GetWeightedRandomCube()

		child := Crossover(parent1, parent2)
		child = Mutate(child)

		newPopulation.Cubes[i] = child

		// Sum scores to get average
		Score := child.GetCurrentScore()
		SumScore += Score

		// Set best cube of population
		if Score > newPopulation.GetBestCube().GetCurrentScore() {
			newPopulation.SetBestCubeIndex(i)
		}

		// Set min score of population
		if Score < newPopulation.GetMinScore() {
			newPopulation.SetMinScore(Score)
		}
	}

	// Set average score of population
	newPopulation.SetAvgScore(SumScore / newPopulation.GetPopulationNum())

	// Set max fitness of population
	maxFitness := newPopulation.GetBestCube().GetCurrentScore() - newPopulation.GetMinScore()
	newPopulation.SetMaxFitness(maxFitness)

	return newPopulation
}

func GeneticAlgorithm(populationNum int, iteration int) class.Solution {
	population := NewPopulation(populationNum)
	population.GeneratePopulation()

	timeStart := time.Now()

	res := class.NewSolution()
	res.SetType("Genetic Algorithm")
	res.AddSolutionItem(0, population.GetBestCube().GetCurrentScore(), population.GetBestCube().GetCurrentState())
	res.AddAdditionalInfo(strconv.Itoa(0), float64(population.GetAvgScore()))

	for i := 1; i <= iteration; i++ {
		newPopulation := population.BreedPopulation()

		population = newPopulation
		res.AddSolutionItem(i, population.GetBestCube().GetCurrentScore(), population.GetBestCube().GetCurrentState())
		res.AddAdditionalInfo(strconv.Itoa(i), float64(population.GetAvgScore()))
	}

	// Add additional info
	elapsedTime := time.Since(timeStart).Milliseconds()
	res.AddElapsedTime(float64(elapsedTime))
	res.AddLastScore(population.GetBestCube().GetCurrentScore())
	res.AddAdditionalInfo("Iteration count", float64(iteration))
	res.AddAdditionalInfo("Population count", float64(populationNum))

	return *res
}
