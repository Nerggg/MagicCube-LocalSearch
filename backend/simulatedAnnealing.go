package main

import (
	"fmt"
	"math"
	"math/rand"
)

func simulatedAnnealing(cube *[][][]int, initialTemp float64, coolingRate float64, maxIterations int) {
	currentState := *cube
	currentValue := calculateObjectiveFunction(currentState)

	temperature := initialTemp

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
		}

		// cool down the temperature
		temperature *= coolingRate
		fmt.Printf("Iteration: %d, Current value: %d, Temperature: %.5f\n", i, currentValue, temperature)

		// break if temperature nearly zero
		if temperature < 1e-1000 {
			break
		}
	}

	fmt.Printf("Final value: %d\n", currentValue)
	fmt.Println("Final solution:")
	printCube(currentState)
	if currentValue == 109 {
		fmt.Println("success")
	} else {
		fmt.Println("fail")
	}
}
