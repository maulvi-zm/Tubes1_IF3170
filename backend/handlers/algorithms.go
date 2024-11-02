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
	var data struct {
		Iteration     int `json:"iteration"`
		PopulationNum int `json:"populationNum"`
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Unable to parse JSON data", http.StatusBadRequest)
		return
	}

	iteration := data.Iteration
	populationNum := data.PopulationNum
	fmt.Println("Population Num: ", populationNum)
	fmt.Println("Iteration: ", iteration)

	result := algorithms.GeneticAlgorithm(populationNum, iteration)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
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
