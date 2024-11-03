package main

import (
	"fmt"
	"math"
	"math/rand"
)

func simulatedAnnealing(cube *[][][]int, initialTemp float64, coolingRate float64, maxIterations int) ([][][]int, int, int, []int) {
	currentState := *cube
	currentValue := calculateObjectiveFunction(currentState)

	temperature := initialTemp
	iter := 0
	stuckCount := 0
	iterOF := []int{}

	for i := 0; i < maxIterations && currentValue < 109; i++ {
		newState := copy3DArray(currentState)
		newState = swapRandom(newState)
		newValue := calculateObjectiveFunction(newState)

		// calculate probability of accepting worse solution
		delta := newValue - currentValue
		probability := math.Exp(float64(delta) / temperature)

		// accept new solution if better or worse based on probability
		if delta > 0 || rand.Float64() < probability {
			currentState = copy3DArray(newState)
			currentValue = newValue
		} else {
			stuckCount++ // Increment stuck count jika solusi tidak berubah
		}

		// cool down the temperature
		temperature *= coolingRate
		fmt.Printf("Iteration: %d, Current value: %d, Temperature: %.5f, Stuck Count: %d\n", i, currentValue, temperature, stuckCount)

		// break if temperature nearly zero or stuck count exceeds limit
		if temperature < 1e-1000 || currentValue == 109 {
			break
		}
		iter++
		iterOF = append(iterOF, currentValue)
	}

	fmt.Printf("Final value: %d\n", currentValue)
	fmt.Println("Final solution:")
	printCube(currentState)
	if currentValue == 109 {
		fmt.Println("success")
	} else {
		fmt.Println("fail")
	}

	return currentState, currentValue, stuckCount, iterOF
}
