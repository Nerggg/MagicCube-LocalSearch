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

func main() {
	// array := [5][5][5]int{ // ini array yg udah perfect magic cube
	// 	{
	// 		{25, 16, 80, 104, 90},
	// 		{115, 98, 4, 1, 97},
	// 		{42, 111, 85, 2, 75},
	// 		{66, 72, 27, 102, 48},
	// 		{67, 18, 119, 106, 5},
	// 	},
	// 	{
	// 		{91, 77, 71, 6, 70},
	// 		{52, 64, 117, 69, 13},
	// 		{30, 118, 21, 123, 23},
	// 		{26, 39, 92, 44, 114},
	// 		{116, 17, 14, 73, 95},
	// 	},
	// 	{
	// 		{47, 61, 45, 76, 86},
	// 		{107, 43, 38, 33, 94},
	// 		{89, 68, 63, 58, 37},
	// 		{32, 93, 88, 83, 19},
	// 		{40, 50, 81, 65, 79},
	// 	},
	// 	{
	// 		{31, 53, 112, 109, 10},
	// 		{12, 82, 34, 87, 100},
	// 		{103, 3, 105, 8, 96},
	// 		{113, 57, 9, 62, 74},
	// 		{56, 120, 55, 49, 35},
	// 	},
	// 	{
	// 		{121, 108, 7, 20, 59},
	// 		{29, 28, 122, 125, 11},
	// 		{51, 15, 41, 124, 84},
	// 		{78, 54, 99, 24, 60},
	// 		{36, 110, 46, 22, 101},
	// 	},
	// }

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

	fmt.Println(calculateObjectiveFunction(arr))
}
