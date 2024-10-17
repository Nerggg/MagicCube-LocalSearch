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
		for x > 5 && y < 5 {
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

func main() {
	array := [5][5][5]int{
		{
			{10, 36, 67, 98, 104},
			{39, 70, 96, 102, 8},
			{68, 99, 105, 6, 37},
			{97, 103, 9, 40, 66},
			{101, 7, 38, 69, 100},
		},
		{
			{114, 20, 46, 52, 83},
			{18, 49, 55, 81, 112},
			{47, 53, 84, 115, 16},
			{51, 82, 113, 19, 50},
			{85, 111, 17, 48, 54},
		},
		{
			{93, 124, 5, 31, 62},
			{122, 3, 34, 65, 91},
			{1, 32, 63, 94, 125},
			{35, 61, 92, 123, 4},
			{64, 95, 121, 2, 33},
		},
		{
			{72, 78, 109, 15, 41},
			{76, 107, 13, 44, 75},
			{110, 11, 42, 73, 79},
			{14, 45, 71, 77, 108},
			{43, 74, 80, 106, 12},
		},
		{
			{26, 57, 88, 119, 25},
			{60, 86, 117, 23, 29},
			{89, 120, 21, 27, 58},
			{118, 24, 30, 56, 87},
			{22, 28, 59, 90, 116},
		},
	}

	// var arr [5][5][5]int

	// randomNumbers := generateRandomNumbers()

	// index := 0
	// for i := 0; i < 5; i++ {
	// 	for j := 0; j < 5; j++ {
	// 		for k := 0; k < 5; k++ {
	// 			arr[i][j][k] = randomNumbers[index]
	// 			index++
	// 		}
	// 	}
	// }

	// for i := 0; i < 5; i++ {
	// 	fmt.Printf("Layer %d:\n", i+1)
	// 	for j := 0; j < 5; j++ {
	// 		for k := 0; k < 5; k++ {
	// 			fmt.Printf("%4d ", arr[i][j][k])
	// 		}
	// 		fmt.Println()
	// 	}
	// 	fmt.Println()
	// }

	fmt.Println(calculateObjectiveFunction(array))
}
