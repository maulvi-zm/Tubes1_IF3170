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
	var data struct {
		MaxStochasticMove int `json:"maxStochasticMove"`
	}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Unable to parse form data", http.StatusBadRequest)
		return
	}

	maxStochasticMove := data.MaxStochasticMove
	result := algorithms.HillClimbingStochastic(maxStochasticMove)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func HillClimbingSidewayMoveHandler(w http.ResponseWriter, r *http.Request) {
	var data struct {
		MaxSidewayMove int `json:"maxSidewayMove"`
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Unable to parse form data", http.StatusBadRequest)
		return
	}

	maxSidewayMove := data.MaxSidewayMove

	result := algorithms.HillClimbingSideways(maxSidewayMove)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func HillClimbingRandomRestartHandler(w http.ResponseWriter, r *http.Request) {
	var data struct {
		MaxRandomRestart int `json:"maxRandomRestart"`
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Unable to parse JSON data", http.StatusBadRequest)
		return
	}

	maxRestart := data.MaxRandomRestart

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

	w.Header().Set("Content-Type", "application/json")

	jsonData, err := result.ConvertToJson()
	if err != nil {
		http.Error(w, "Unable to convert solution to JSON", http.StatusInternalServerError)
		return
	}

	// print the json size in kb to check if it's too big
	fmt.Println("JSON size in KB: ", float64(len(jsonData))/1000)

	w.Write([]byte(jsonData))
}
