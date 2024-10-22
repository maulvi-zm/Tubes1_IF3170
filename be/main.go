package main

import (
	"be/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()

    // Endpoints for algorithms
    r.HandleFunc("/hill-climbing/steepest", handlers.HillClimbingSteepestHandler).Methods("GET")
    r.HandleFunc("/hill-climbing/stochastic", handlers.HillClimbingStochasticHandler).Methods("GET")
    r.HandleFunc("/genetic-algorithm", handlers.GeneticAlgorithmHandler).Methods("GET")
    r.HandleFunc("/simulated-annealing", handlers.SimulatedAnnealingHandler).Methods("GET")

    r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Welcome to the Algorithms API!"))
    })

    log.Println("Server running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
