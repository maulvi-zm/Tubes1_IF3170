package algorithms

// package main

import (
	"be/class"
	"fmt"
	"math"
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
	skewed := 1 - math.Pow(randomFloat, 3.0) // Adjust exponent for more or less skew (higher = more weight towards max)

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
			if HPIndex != -1 {
				child.SetSmallCubeValue(i, LPValue)
				child.SetSmallCubeValue(HPIndex, HPValue)
				swappedIndices[i] = true
				swappedIndices[HPIndex] = true
			} else {
				child.SetSmallCubeValue(i, HPValue)
			}
		}
	}

	return child
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

// func (p *Population) BreedPopulation() *Population {
// 	newPopulation := NewPopulation(p.GetPopulationNum())

// 	for i := 0; i < p.GetPopulationNum(); i++ {
// 		parent1 := p.GetWeightedRandomCube()
// 		parent2 := p.GetWeightedRandomCube()

// 		child := parent1.Crossover(parent2)
// 		child.Mutate()

// 		newPopulation.Cubes[i] = child
// 	}

// 	return newPopulation
// }

func GeneticAlgorithm(populationNum int, iteration int) class.Solution {
	HigherHalfParent := class.NewCube(5)
	LowerHalfParent := class.NewCube(5)

	HigherHalfParent.SetRandomStartState()
	LowerHalfParent.SetRandomStartState()

	// Solution
	res := class.NewSolution()
	res.SetType("Genetic Algorithm")
	res.AddSolutionItem(0, 0, HigherHalfParent.GetCurrentState())

	// Print parents
	fmt.Println("Higher Half Parent:")
	for i := 0; i < 27; i++ {
		fmt.Print(HigherHalfParent.GetSmallCubeValue(i), " ")
	}

	fmt.Println("\nLower Half Parent:")
	for i := 0; i < 27; i++ {
		fmt.Print(LowerHalfParent.GetSmallCubeValue(i), " ")
	}

	// child := Crossover(HigherHalfParent, LowerHalfParent)
	// child.PrintCube()

	return *res
}
