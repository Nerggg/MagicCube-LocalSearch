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

func calculateObjectiveFunction(cube [5][5][5]int) int {
	result := 0

	// Iterasi bagian pertama
	for z := 0; z < 5; z++ {
		for x := 0; x < 5; x++ {
			temp := 0
			for y := 0; y < 5; y++ {
				temp += cube[x][y][z]
			}
			if temp == 315 {
				result++
			}
		}
	}

	for z := 0; z < 5; z++ {
		for y := 0; y < 5; y++ {
			temp := 0
			for x := 0; x < 5; x++ {
				temp += cube[x][y][z]
			}
			if temp == 315 {
				result++
			}
		}
	}

	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			temp := 0
			for z := 0; z < 5; z++ {
				temp += cube[x][y][z]
			}
			if temp == 315 {
				result++
			}
		}
	}

	// Iterasi bagian kedua
	for y := 0; y < 5; y++ {
		temp := 0
		for x, z := 0, 0; x < 5 && z < 5; x, z = x+1, z+1 {
			temp += cube[x][y][z]
		}
		if temp == 315 {
			result++
		}
	}

	for y := 0; y < 5; y++ {
		temp := 0
		for x, z := 4, 0; x >= 0 && z < 5; x, z = x-1, z+1 {
			temp += cube[x][y][z]
		}
		if temp == 315 {
			result++
		}
	}

	for x := 0; x < 5; x++ {
		temp := 0
		for y, z := 0, 0; y < 5 && z < 5; y, z = y+1, z+1 {
			temp += cube[x][y][z]
		}
		if temp == 315 {
			result++
		}
	}

	for x := 0; x < 5; x++ {
		temp := 0
		for y, z := 4, 0; y >= 0 && z < 5; y, z = y-1, z+1 {
			temp += cube[x][y][z]
		}
		if temp == 315 {
			result++
		}
	}

	// Iterasi bagian ketiga
	temp := 0
	for x, y, z := 0, 0, 0; x < 5 && y < 5 && z < 5; x, y, z = x+1, y+1, z+1 {
		temp += cube[x][y][z]
	}
	if temp == 315 {
		result++
	}

	temp = 0
	for x, y, z := 4, 0, 0; x >= 0 && y < 5 && z < 5; x, y, z = x-1, y+1, z+1 {
		temp += cube[x][y][z]
	}
	if temp == 315 {
		result++
	}

	temp = 0
	for x, y, z := 4, 0, 4; x >= 0 && y < 5 && z >= 0; x, y, z = x-1, y+1, z-1 {
		temp += cube[x][y][z]
	}
	if temp == 315 {
		result++
	}

	temp = 0
	for x, y, z := 0, 0, 4; x < 5 && y < 5 && z >= 0; x, y, z = x+1, y+1, z-1 {
		temp += cube[x][y][z]
	}
	if temp == 315 {
		result++
	}

	return result
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

	fmt.Println("Objective Function Result:", calculateObjectiveFunction(arr))
}
