package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type SearchRequest struct {
	Cube               [][][]int `json:"cube"`
	Algorithm          string    `json:"algorithm"`
	PopulationSize     int       `json:"populationSize,omitempty"`     // Optional field for Genetic Algorithm
	MaxGenerations     int       `json:"maxGenerations,omitempty"`     // Optional field for Genetic Algorithm
	Temperature        float64   `json:"temperature,omitempty"`        // Optional field for Simulated Annealing
	CoolingRate        float64   `json:"coolingRate,omitempty"`        // Optional field for Simulated Annealing
	MaxIterations      int       `json:"maxIterations,omitempty"`      // Optional field for Simulated Annealing
	MaxStateGeneration int       `json:"maxStateGeneration,omitempty"` // Optional field for Simulated Annealing
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
		if requestData.Algorithm == "Genetic Algorithm" {
			fmt.Printf("Population Size: %d\n", requestData.PopulationSize)
			fmt.Printf("Max Generations: %d\n", requestData.MaxGenerations)
		}

		if requestData.Algorithm == "Simulated Annealing" {
			fmt.Printf("Temperature: %f\n", requestData.Temperature)
			fmt.Printf("Cooling Rate: %f\n", requestData.CoolingRate)
			fmt.Printf("Max Iterations: %d\n", requestData.MaxIterations)
		}

		if requestData.Algorithm == "Stochastic Hill Climbing" {
			fmt.Printf("Max State Generation: %f\n", requestData.Temperature)
		}

		startTime := time.Now()
		var finalState [][][]int
		var iterOF []int
		var finalValue, stuckCount int

		// Jalankan algoritma berdasarkan permintaan
		switch requestData.Algorithm {
		case "Simulated Annealing":
			finalState, finalValue, stuckCount, iterOF = simulatedAnnealing(&requestData.Cube, requestData.Temperature, requestData.CoolingRate, requestData.MaxIterations)
			duration := time.Since(startTime).Milliseconds()
			fmt.Printf("Duration: %d ms\n", duration)
			lastResult = map[string]interface{}{
				"finalState":    finalState,
				"finalValue":    finalValue,
				"stuckCount":    stuckCount,
				"duration":      duration,
				"iterOF":        iterOF,
				"temperature":   requestData.Temperature,
				"coolingRate":   requestData.CoolingRate,
				"maxIterations": requestData.MaxIterations,
			}

		case "Stochastic Hill Climbing":
			finalState, finalValue, stuckCount, iterOF = stochasticHillClimbing(&requestData.Cube, requestData.MaxStateGeneration)
			duration := time.Since(startTime).Milliseconds()
			fmt.Printf("Duration: %d ms\n", duration)
			lastResult = map[string]interface{}{
				"finalState": finalState,
				"finalValue": finalValue,
				"stuckCount": stuckCount,
				"duration":   duration,
				"iterOF":     iterOF,
			}

		case "Random Restart Hill Climbing":
			finalState, finalValue, stuckCount, iterOF = steepestAscentHillClimbing(&requestData.Cube)
			duration := time.Since(startTime).Milliseconds()
			fmt.Printf("Duration: %d ms\n", duration)
			lastResult = map[string]interface{}{
				"finalState": finalState,
				"finalValue": finalValue,
				"stuckCount": stuckCount,
				"duration":   duration,
				"iterOF":     iterOF,
			}

		case "Steepest Ascent Hill Climbing":
			finalState, finalValue, stuckCount, iterOF = steepestAscentHillClimbing(&requestData.Cube)
			duration := time.Since(startTime).Milliseconds()
			fmt.Printf("Duration: %d ms\n", duration)
			lastResult = map[string]interface{}{
				"finalState": finalState,
				"finalValue": finalValue,
				"stuckCount": stuckCount,
				"duration":   duration,
				"iterOF":     iterOF,
			}

		case "Sideways Move Hill Climbing":
			finalState, finalValue, stuckCount, iterOF = sidewaysMoveHillClimbing(&requestData.Cube)
			duration := time.Since(startTime).Milliseconds()
			fmt.Printf("Duration: %d ms\n", duration)
			lastResult = map[string]interface{}{
				"finalState": finalState,
				"finalValue": finalValue,
				"stuckCount": stuckCount,
				"duration":   duration,
				"iterOF":     iterOF,
			}

		case "Genetic Algorithm":
			finalState, finalValue, iterOF = geneticAlgorithm(&requestData.Cube, requestData.PopulationSize, requestData.MaxGenerations)
			duration := time.Since(startTime).Milliseconds()
			fmt.Printf("Duration: %d ms\n", duration)
			lastResult = map[string]interface{}{
				"finalState":     finalState,
				"finalValue":     finalValue,
				"duration":       duration,
				"iterOF":         iterOF,
				"generation":     requestData.MaxGenerations,
				"populationSize": requestData.PopulationSize,
			}

		default:
			http.Error(w, "Invalid algorithm", http.StatusBadRequest)
			return
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
