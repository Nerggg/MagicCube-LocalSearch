package main

import "fmt"

func steepestAscentHillClimbing(cube *[][][]int) {
	currentObjectiveFunction := calculateObjectiveFunction(*cube)
	for currentObjectiveFunction < 109 {
		tempCube := generateMaximumSuccessor(*cube)
		tempObjectiveFunction := calculateObjectiveFunction(tempCube)
		if tempObjectiveFunction > currentObjectiveFunction {
			*cube = tempCube
			currentObjectiveFunction = tempObjectiveFunction
		} else {
			printCube(*cube)
			fmt.Println("fail")
			break
		}
	}
	if currentObjectiveFunction == 109 {
		printCube(*cube)
		fmt.Print("success")
	}
}

func stochasticHillClimbing(cube *[][][]int, NMax int) {
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

func stochasticHillClimbingHelper(cube *[][][]int, NMax int, currentObjectiveFunction int, proceed *bool) {
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
