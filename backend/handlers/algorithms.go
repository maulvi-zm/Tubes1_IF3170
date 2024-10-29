package handlers

import (
	"be/algorithms"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func HillClimbingSteepestHandler(w http.ResponseWriter, r *http.Request) {
	result := algorithms.HillClimbingSteepest()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func HillClimbingStochasticHandler(w http.ResponseWriter, r *http.Request) {
	result := algorithms.HillClimbingStochastic()
	fmt.Fprintf(w, "Hill Climbing (Stochastic) Result: %v", result)
}

func HillClimbingSidewayMoveHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Unable to parse form data", http.StatusBadRequest)
		return
	}

	maxSidewayMoveStr := r.FormValue("maxSidewayMove")
	maxSidewayMove, err := strconv.Atoi(maxSidewayMoveStr)
	if err != nil {
		http.Error(w, "Invalid maxSidewayMove value", http.StatusBadRequest)
		return
	}

	result := algorithms.HillClimbingSideways(maxSidewayMove)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func HillClimbingRandomRestartHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Unable to parse form data", http.StatusBadRequest)
		return
	}

	maxRestartStr := r.FormValue("maxRestart")
	maxRestart, err := strconv.Atoi(maxRestartStr)
	if err != nil {
		http.Error(w, "Invalid maxRestart value", http.StatusBadRequest)
		return
	}

	result := algorithms.HillClimbingRandomRestart(maxRestart)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func GeneticAlgorithmHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Unable to parse form data", http.StatusBadRequest)
	}

	populationNumStr := r.FormValue("populationNum")
	populationNum, err := strconv.Atoi(populationNumStr)
	if err != nil {
		http.Error(w, "Invalid populationNum value", http.StatusBadRequest)
	}

	iterationStr := r.FormValue("iteration")
	iteration, err := strconv.Atoi(iterationStr)
	if err != nil {
		http.Error(w, "Invalid iteration value", http.StatusBadRequest)
	}

	result := algorithms.GeneticAlgorithm(populationNum, iteration)
	fmt.Fprintf(w, "Genetic Algorithm Result: %v", result)
}

func SimulatedAnnealingHandler(w http.ResponseWriter, r *http.Request) {
	result := algorithms.SimulatedAnnealing()
	fmt.Fprintf(w, "Simulated Annealing Result: %v", result)
}
