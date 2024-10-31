package main

import (
	h "be/handlers"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Endpoints for algorithms
	r.HandleFunc("/hill-climbing/steepest", h.HillClimbingSteepestHandler).Methods("POST")
	r.HandleFunc("/hill-climbing/stochastic", h.HillClimbingStochasticHandler).Methods("POST")
	r.HandleFunc("/hill-climbing/sidewaymove", h.HillClimbingSidewayMoveHandler).Methods("POST")
	r.HandleFunc("/hill-climbing/randomrestart", h.HillClimbingRandomRestartHandler).Methods("POST")
	r.HandleFunc("/genetic-algorithm", h.GeneticAlgorithmHandler).Methods("POST")
	r.HandleFunc("/simulated-annealing", h.SimulatedAnnealingHandler).Methods("POST")

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the Algorithms API!"))
	})

	// Allow CORS
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", corsHandler(r)))
}
