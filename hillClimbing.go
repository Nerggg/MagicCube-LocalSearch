package main

import "fmt"

func stochasticHillClimbing(cube *[5][5][5]int, NMax int) {
	currentObjectiveFunction := calculateObjectiveFunction(*cube)
	proceed := true
	for currentObjectiveFunction < 109 {
		stochasticHillClimbingHelper(cube, NMax, currentObjectiveFunction, &proceed)
		if !proceed {
			printCube(*cube)
			fmt.Println("fail")
			break
		}
		currentObjectiveFunction = calculateObjectiveFunction(*cube)
		// fmt.Println(currentObjectiveFunction)
	}

	if currentObjectiveFunction == 109 {
		printCube(*cube)
		fmt.Println("success")
	}
}

func stochasticHillClimbingHelper(cube *[5][5][5]int, NMax int, currentObjectiveFunction int, proceed *bool) {
	i := 0
	for i < NMax {
		tempCube := swapRandom(*cube)
		if calculateObjectiveFunction(tempCube) > currentObjectiveFunction {
			*cube = tempCube
			break
		}
		i++
	}
	if i == NMax {
		*proceed = false
	}
}
