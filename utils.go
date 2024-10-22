package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRandom5x5x5Array() [5][5][5]int { // fungsi awal utk ngegenerate random state
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
	return arr
}

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

func printCube(cube [5][5][5]int) {
	for i := 0; i < 5; i++ {
		fmt.Printf("Layer %d:\n", i+1)
		for j := 0; j < 5; j++ {
			for k := 0; k < 5; k++ {
				fmt.Printf("%4d ", cube[i][j][k])
			}
			fmt.Println()
		}
		fmt.Println()
	}
}

func swapRandom(cube [5][5][5]int) [5][5][5]int { // fungsi utk ngeswap dua angka dgn posisi random
	rand.Seed(time.Now().UnixNano())

	x1, y1, z1 := rand.Intn(5), rand.Intn(5), rand.Intn(5)
	x2, y2, z2 := rand.Intn(5), rand.Intn(5), rand.Intn(5)

	cube[x1][y1][z1], cube[x2][y2][z2] = cube[x2][y2][z2], cube[x1][y1][z1]

	return cube
}

func calculateObjectiveFunction(cube [5][5][5]int) int { // ngitung objective function
	result := 0

	// iterasi bagian pertama
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

	// iterasi bagian kedua
	x := 0
	z := 0
	for y := 0; y < 5; y++ {
		temp := 0
		x = 0
		z = 0
		for x < 5 && z < 5 {
			temp += cube[x][y][z]
			x++
			z++
		}
		if temp == 315 {
			result++
		}
	}
	x = 4
	z = 0
	for y := 0; y < 5; y++ {
		temp := 0
		x = 4
		z = 0
		for x >= 0 && z < 5 {
			temp += cube[x][y][z]
			x--
			z++
		}
		if temp == 315 {
			result++
		}
	}

	y := 0
	z = 0
	for x := 0; x < 5; x++ {
		temp := 0
		y := 0
		z = 0
		for y < 5 && z < 5 {
			temp += cube[x][y][z]
			y++
			z++
		}
		if temp == 315 {
			result++
		}
	}
	y = 4
	z = 0
	for x := 0; x < 5; x++ {
		temp := 0
		y = 4
		z = 0
		for y >= 0 && z < 5 {
			temp += cube[x][y][z]
			y--
			z++
		}
		if temp == 315 {
			result++
		}
	}

	x = 0
	y = 0
	for z := 0; z < 5; z++ {
		temp := 0
		x = 0
		y = 0
		for x < 5 && y < 5 {
			temp += cube[x][y][z]
			x++
			y++
		}
		if temp == 315 {
			result++
		}
	}
	x = 4
	y = 0
	for z := 0; z < 5; z++ {
		temp := 0
		x = 4
		y = 0
		for x >= 0 && y < 5 {
			temp += cube[x][y][z]
			x--
			y++
		}
		if temp == 315 {
			result++
		}
	}

	// iterasi bagian ketiga
	x = 0
	y = 0
	z = 0
	temp := 0
	for x < 5 && y < 5 && z < 5 {
		temp += cube[x][y][z]
		x++
		y++
		z++
	}
	if temp == 315 {
		result++
	}
	x = 4
	y = 0
	z = 0
	temp = 0
	for x >= 0 && y < 5 && z < 5 {
		temp += cube[x][y][z]
		x--
		y++
		z++
	}
	if temp == 315 {
		result++
	}
	x = 4
	y = 0
	z = 4
	temp = 0
	for x >= 0 && y < 5 && z >= 0 {
		temp += cube[x][y][z]
		x--
		y++
		z--
	}
	if temp == 315 {
		result++
	}
	x = 0
	y = 0
	z = 4
	temp = 0
	for x < 5 && y < 5 && z >= 0 {
		temp += cube[x][y][z]
		x++
		y++
		z--
	}
	if temp == 315 {
		result++
	}
	return result
}
