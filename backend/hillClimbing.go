package main

import "fmt"

func steepestAscentHillClimbing(cube *[][][]int) {
	currentObjectiveFunction := calculateObjectiveFunction(*cube)
	for currentObjectiveFunction < 109 {
		tempCube := generateMaximumSuccessor(*cube)
		tempObjectiveFunction := calculateObjectiveFunction(tempCube)
		if tempObjectiveFunction <= currentObjectiveFunction {
			printCube(*cube)
			fmt.Println("fail")
			break
		}
		*cube = tempCube
		currentObjectiveFunction = tempObjectiveFunction
		fmt.Println(currentObjectiveFunction)
	}
	if currentObjectiveFunction == 109 {
		printCube(*cube)
		fmt.Println("success")
	}
}

func sidewaysMoveHillClimbing(cube *[][][]int) {
	currentObjectiveFunction := calculateObjectiveFunction(*cube)
	for currentObjectiveFunction < 109 {
		tempCube := generateMaximumSuccessor(*cube)
		tempObjectiveFunction := calculateObjectiveFunction(tempCube)
		if tempObjectiveFunction < currentObjectiveFunction {
			printCube(*cube)
			fmt.Println("fail")
			break
		}
		*cube = tempCube
		currentObjectiveFunction = tempObjectiveFunction
		fmt.Println(currentObjectiveFunction)
	}
	if currentObjectiveFunction == 109 {
		printCube(*cube)
		fmt.Println("success")
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
		fmt.Println(currentObjectiveFunction)
	}

	if currentObjectiveFunction == 109 {
		printCube(*cube)
		fmt.Println("success")
	}
}

func stochasticHillClimbingHelper(cube *[][][]int, NMax int, currentObjectiveFunction int, proceed *bool) {
	i := 0
	for i < NMax {
		if i%100000 == 0 && i != 0 {
			fmt.Println("Iterasi ke ", i)
		}
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
