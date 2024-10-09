package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRandomNumbers() []int {
	numbers := make([]int, 125)
	for i := 0; i < 125; i++ {
		numbers[i] = i + 1
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(numbers), func(i, j int) {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	})

	return numbers
}

func main() {
	var arr [5][5][5]int

	randomNumbers := generateRandomNumbers()

	index := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			for k := 0; k < 5; k++ {
				arr[i][j][k] = randomNumbers[index]
				index++
			}
		}
	}

	for i := 0; i < 5; i++ {
		fmt.Printf("Layer %d:\n", i+1)
		for j := 0; j < 5; j++ {
			for k := 0; k < 5; k++ {
				fmt.Printf("%4d ", arr[i][j][k])
			}
			fmt.Println()
		}
		fmt.Println()
	}
}
