package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// DAFTAR FUNGSI YANG ADA DISINI
// generateRandom5x5x5Array => ngegenerate random initial state, ngereturn array 3d
// printCube => ngeprint array 3d
// swapRandom => nukar 2 angka dalam posisi random dalam sebuah array 3d
// calculateObjectiveFunction => ngitung objective function, ngereturn int
// generateSuccessor => cari semua suksesor yg mungkin, ngereturn array of array 3d
// copy3DArray => nyalin array 3d, ngereturn array 3d baru yg sama
// generateMaximumSuccessor => cari satu suksesor dgn objective function tertinggi, ngereturn array 3d

func generateRandom5x5x5Array() [][][]int {
	arr := make([][][]int, 5)
	for i := range arr {
		arr[i] = make([][]int, 5)
		for j := range arr[i] {
			arr[i][j] = make([]int, 5)
		}
	}

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

func printCube(cube [][][]int) {
	for i := 0; i < len(cube); i++ {
		fmt.Printf("Layer %d:\n", i+1)
		for j := 0; j < len(cube); j++ {
			for k := 0; k < len(cube); k++ {
				fmt.Printf("%4d ", cube[i][j][k])
			}
			fmt.Println()
		}
		fmt.Println()
	}
}

func swapRandom(cube [][][]int) [][][]int { // fungsi utk ngeswap dua angka dgn posisi random
	rand.Seed(time.Now().UnixNano())

	x1, y1, z1 := rand.Intn(len(cube)), rand.Intn(len(cube)), rand.Intn(len(cube))
	x2, y2, z2 := rand.Intn(len(cube)), rand.Intn(len(cube)), rand.Intn(len(cube))

	cube[x1][y1][z1], cube[x2][y2][z2] = cube[x2][y2][z2], cube[x1][y1][z1]

	return cube
}

func calculateObjectiveFunction(cube [][][]int) int { // ngitung objective function
	result := 0

	// iterasi bagian pertama
	for z := 0; z < 5; z++ {
		for x := 0; x < 5; x++ {
			temp := 0
			for y := 0; y < 5; y++ {
				temp += cube[x][y][z]
			}
			result -= int(math.Abs(float64(315 - temp)))
		}
	}

	for z := 0; z < 5; z++ {
		for y := 0; y < 5; y++ {
			temp := 0
			for x := 0; x < 5; x++ {
				temp += cube[x][y][z]
			}
			result -= int(math.Abs(float64(315 - temp)))
		}
	}

	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			temp := 0
			for z := 0; z < 5; z++ {
				temp += cube[x][y][z]
			}
			result -= int(math.Abs(float64(315 - temp)))
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
		result -= int(math.Abs(float64(315 - temp)))
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
		result -= int(math.Abs(float64(315 - temp)))
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
		result -= int(math.Abs(float64(315 - temp)))
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
		result -= int(math.Abs(float64(315 - temp)))
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
		result -= int(math.Abs(float64(315 - temp)))
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
		result -= int(math.Abs(float64(315 - temp)))
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
	result -= int(math.Abs(float64(315 - temp)))
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
	result -= int(math.Abs(float64(315 - temp)))
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
	result -= int(math.Abs(float64(315 - temp)))
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
	result -= int(math.Abs(float64(315 - temp)))
	return result
}

func generateSuccessor(originalArr [][][]int) [][][][]int {
	store := [][]int{}
	result := [][][][]int{}

	for i1 := 0; i1 < len(originalArr); i1++ {
		for j1 := 0; j1 < len(originalArr[i1]); j1++ {
			for k1 := 0; k1 < len(originalArr[i1][j1]); k1++ {
				for i2 := 0; i2 < len(originalArr); i2++ {
					for j2 := 0; j2 < len(originalArr[i2]); j2++ {
						for k2 := 0; k2 < len(originalArr[i2][j2]); k2++ {
							if (i1 == i2 && j1 == j2 && k1 == k2) || inStore(store, i1, j1, k1, i2, j2, k2) {
								continue
							}

							arr := copy3DArray(originalArr)
							arr[i1][j1][k1], arr[i2][j2][k2] = arr[i2][j2][k2], arr[i1][j1][k1]
							result = append(result, arr)
							store = append(store, []int{i1, j1, k1, i2, j2, k2})
							store = append(store, []int{i2, j2, k2, i1, j1, k1})
						}
					}
				}
			}
		}
	}
	return result
}

func inStore(store [][]int, i1, j1, k1, i2, j2, k2 int) bool {
	for _, s := range store {
		if s[0] == i1 && s[1] == j1 && s[2] == k1 && s[3] == i2 && s[4] == j2 && s[5] == k2 {
			return true
		}
	}
	return false
}

func copy3DArray(original [][][]int) [][][]int {
	xLen := len(original)
	yLen := len(original[0])
	zLen := len(original[0][0])

	copyArray := make([][][]int, xLen)
	for i := range original {
		copyArray[i] = make([][]int, yLen)
		for j := range original[i] {
			copyArray[i][j] = make([]int, zLen)
			for k := range original[i][j] {
				copyArray[i][j][k] = original[i][j][k]
			}
		}
	}

	return copyArray
}

func generateMaximumSuccessor(originalArr [][][]int) [][][]int {
	store := [][]int{}
	result := [][][]int{}
	maxTemp := -999999

	for i1 := 0; i1 < len(originalArr); i1++ {
		for j1 := 0; j1 < len(originalArr[i1]); j1++ {
			for k1 := 0; k1 < len(originalArr[i1][j1]); k1++ {
				for i2 := 0; i2 < len(originalArr); i2++ {
					for j2 := 0; j2 < len(originalArr[i2]); j2++ {
						for k2 := 0; k2 < len(originalArr[i2][j2]); k2++ {
							if (i1 == i2 && j1 == j2 && k1 == k2) || inStore(store, i1, j1, k1, i2, j2, k2) {
								continue
							}

							arr := copy3DArray(originalArr)
							arr[i1][j1][k1], arr[i2][j2][k2] = arr[i2][j2][k2], arr[i1][j1][k1]
							store = append(store, []int{i1, j1, k1, i2, j2, k2})
							store = append(store, []int{i2, j2, k2, i1, j1, k1})

							currentTemp := calculateObjectiveFunction(arr)
							if currentTemp > maxTemp {
								result = arr
								maxTemp = currentTemp
							}
						}
					}
				}
			}
		}
	}
	return result
}
