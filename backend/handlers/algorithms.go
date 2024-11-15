package handlers

import (
	"be/algorithms"
	"encoding/json"
	"net/http"
	// "strconv"
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

	result := algorithms.GeneticAlgorithm(populationNum, iteration)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func SimulatedAnnealingHandler(w http.ResponseWriter, r *http.Request) {
	result := algorithms.SimulatedAnnealing()

	// jsonData, err := result.ConvertToJson()
	// if err != nil {
	// 	http.Error(w, "Unable to convert solution to JSON", http.StatusInternalServerError)
	// 	return
	// }

	// print the json size in kb to check if it's too big
	// fmt.Println("JSON size in KB: ", float64(len(jsonData))/1000)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
	// w.Write([]byte(jsonData))
}
