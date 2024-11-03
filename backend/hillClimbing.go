package main

import "fmt"

func steepestAscentHillClimbing(cube *[][][]int) ([][][]int, int, int, []int) {
	iterOF := []int{}
	currentState := *cube
	currentValue := calculateObjectiveFunction(currentState)
	for currentValue < 0 {
		newState := copy3DArray(*cube)
		newState = generateMaximumSuccessor(newState)
		newValue := calculateObjectiveFunction(newState)
		if newValue <= currentValue {
			printCube(*cube)
			fmt.Println("fail")
			break
		}
		*cube = newState
		currentValue = newValue
		fmt.Println(currentValue)
		iterOF = append(iterOF, currentValue)
	}
	if currentValue == 0 {
		printCube(*cube)
		fmt.Println("success")
	}

	return *cube, currentValue, 0, iterOF
}

func sidewaysMoveHillClimbing(cube *[][][]int) ([][][]int, int, int, []int) {
	iterOF := []int{}
	currentState := *cube
	currentValue := calculateObjectiveFunction(currentState)
	for currentValue < 0 {
		sameCounter := 0
		newState := copy3DArray(*cube)
		newState = generateMaximumSuccessor(newState)
		newValue := calculateObjectiveFunction(newState)
		if newValue == currentValue {
			sameCounter++
		}
		if newValue > currentValue {
			sameCounter = 0
		}
		if newValue < currentValue || sameCounter == 1 {
			printCube(*cube)
			fmt.Println("fail")
			break
		}

		// fmt.Println("newvalue ", newValue)
		// fmt.Println("currentvalue ", currentValue)
		// fmt.Println("same counternya ", sameCounter)

		*cube = newState
		currentValue = newValue
		fmt.Println(currentValue)
		iterOF = append(iterOF, currentValue)
	}
	if currentValue == 0 {
		printCube(*cube)
		fmt.Println("success")
	}

	return *cube, currentValue, 0, iterOF
}

func stochasticHillClimbing(cube *[][][]int, NMax int) ([][][]int, int, int, []int) {
	iterOF := []int{}
	currentState := *cube
	currentValue := calculateObjectiveFunction(currentState)
	proceed := true
	for currentValue < 0 && proceed {
		var i int
		for i = 0; i < NMax; i++ {
			if i%100000 == 0 && i != 0 {
				fmt.Println("Iterasi ke ", i)
			}
			newState := copy3DArray(currentState)
			newState = swapRandom(currentState)
			newValue := calculateObjectiveFunction(newState)
			if newValue-currentValue > 0 {
				currentState = copy3DArray(newState)
				currentValue = newValue
				break
			}
		}
		fmt.Println(currentValue)
		iterOF = append(iterOF, currentValue)

		if i == NMax {
			proceed = false
			printCube(*cube)
			fmt.Println("fail")
		}
	}

	if currentValue == 0 {
		printCube(*cube)
		fmt.Println("success")
	}

	return currentState, currentValue, 0, iterOF
}
