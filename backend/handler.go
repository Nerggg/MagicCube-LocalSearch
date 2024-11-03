package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type SearchRequest struct {
	Cube      [][][]int `json:"cube"`
	Algorithm string    `json:"algorithm"`
}

var lastResult map[string]interface{}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == http.MethodPost {
		// Proses POST request
		var requestData SearchRequest
		err := json.NewDecoder(r.Body).Decode(&requestData)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		fmt.Println("received search request")
		fmt.Printf("Cube: %v\n", requestData.Cube)
		fmt.Printf("Algorithm: %s\n", requestData.Algorithm)

		startTime := time.Now()
		var finalState [][][]int
		var iterOF []int
		var finalValue, stuckCount int

		// Jalankan algoritma berdasarkan permintaan
		switch requestData.Algorithm {
		case "Simulated Annealing":
			finalState, finalValue, stuckCount, iterOF = simulatedAnnealing(&requestData.Cube, 10000, 0.999, 1000000)
		case "Stochastic Hill Climbing":
			// finalState, finalValue, stuckCount = stochasticHillClimbing(&requestData.Cube, 500000)
		case "Steepest Ascent Hill Climbing":
			// finalState, finalValue, stuckCount = steepestAscentHillClimbing(&requestData.Cube)
		case "Sideways Move Hill Climbing":
			// finalState, finalValue, stuckCount = sidewaysMoveHillClimbing(&requestData.Cube)
		default:
			http.Error(w, "Invalid algorithm", http.StatusBadRequest)
			return
		}

		duration := time.Since(startTime).Milliseconds()
		fmt.Printf("Duration: %d ms\n", duration)

		lastResult = map[string]interface{}{
			"finalState": finalState,
			"finalValue": finalValue,
			"stuckCount": stuckCount,
			"duration":   duration,
			"iterOF":     iterOF,
		}

		// Kirim respons JSON untuk POST
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(lastResult); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		}

	} else if r.Method == http.MethodGet {
		// Penanganan GET request untuk mengirimkan hasil terakhir
		if lastResult == nil {
			// Jika belum ada hasil dari POST request sebelumnya
			http.Error(w, "No results available. Please run a POST request first.", http.StatusNotFound)
			return
		}

		// Kirimkan hasil terakhir dalam JSON
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(lastResult); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		}
	} else {
		http.Error(w, "Invalid request method. Use POST or GET.", http.StatusMethodNotAllowed)
	}
}
