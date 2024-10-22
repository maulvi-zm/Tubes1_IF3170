package handlers

import (
	"be/algorithms"
	"fmt"
	"net/http"
)

func HillClimbingSteepestHandler(w http.ResponseWriter, r *http.Request) {
    result := algorithms.HillClimbingSteepest()
    fmt.Fprintf(w, "Hill Climbing (Steepest Accent) Result: %v", result)
}

func HillClimbingStochasticHandler(w http.ResponseWriter, r *http.Request) {
    result := algorithms.HillClimbingStochastic()
    fmt.Fprintf(w, "Hill Climbing (Stochastic) Result: %v", result)
}

func GeneticAlgorithmHandler(w http.ResponseWriter, r *http.Request) {
    result := algorithms.GeneticAlgorithm()
    fmt.Fprintf(w, "Genetic Algorithm Result: %v", result)
}

func SimulatedAnnealingHandler(w http.ResponseWriter, r *http.Request) {
    result := algorithms.SimulatedAnnealing()
    fmt.Fprintf(w, "Simulated Annealing Result: %v", result)
}
