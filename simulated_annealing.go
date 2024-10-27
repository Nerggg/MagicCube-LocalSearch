package main

import (
	"fmt"
	"math"
	"math/rand"
)

func acceptanceProbability(delta int, temp float64) float64 {
	return math.Exp(float64(delta) / temp)
}

func coolingSchedule(initialTemperature float64, iteration int) float64 {
	return initialTemperature / float64(iteration+1)
}

func simulatedAnnealing(cube *[][][]int, initialTemperature float64, coolingRate float64) {
	currentCube := *cube
	currentObjectiveFunction := calculateObjectiveFunction(currentCube)
	temperature := initialTemperature

	for temperature > 1 {
		newCube := swapRandom(currentCube)
		newObjectiveFunction := calculateObjectiveFunction(newCube)

		delta := newObjectiveFunction - currentObjectiveFunction
		if delta > 0 || acceptanceProbability(delta, temperature) > rand.Float64() {
			currentCube = newCube
			currentObjectiveFunction = newObjectiveFunction
		}
		temperature *= coolingRate
	}

	*cube = currentCube
	if currentObjectiveFunction == 109 {
		printCube(currentCube)
		fmt.Println("success")
	} else {
		printCube(currentCube)
		fmt.Println("fail")
	}
}