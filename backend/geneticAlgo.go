package main

import (
	"fmt"
	"math/rand"
)

const (
	populationSize = 100
	maxGenerations = 1000
	targetFitness  = 0
)

func initializePopulation(cube *[][][]int, size int) [][][][]int {
	population := make([][][][]int, size)
	population[0] = *cube
	for i := 1; i < size; i++ {
		population[i] = generateRandom5x5x5Array()
	}
	return population // Nyari populasi / kumpulan awal konfigurasi kubus
}

func selectParent(population [][][][]int) [][][]int {
	tournamentSize := 20                           // Cari 20 kandidat dari populasi yang udah dihasilin
	best := population[rand.Intn(len(population))] // Inisialisasi best pertama dengan random
	bestFitness := calculateObjectiveFunction(best)

	for i := 0; i < tournamentSize; i++ { // looping buat cari kandidat terbaik dari 20 kandidat
		individual := population[rand.Intn(len(population))]
		individualFitness := calculateObjectiveFunction(individual)
		if individualFitness > bestFitness {
			best = individual // kalo fitnessnya lebih bagus, update best
			bestFitness = individualFitness
		}
	}
	return best // ngebalikin kandidat terbaik yang nantinya bakal jadi parent
}

// function memperanakkan parent yang udah dipilih
func crossover(parent1, parent2 [][][]int) ([][][]int, [][][]int) {
	// Bikin child baru dari parent
	child1 := copy3DArray(parent1)
	child2 := copy3DArray(parent2)

	// Cross over point
	x := 3

	// lakuin crossover
	for i := 0; i < len(parent1); i++ {
		for j := 0; j < len(parent1[i]); j++ {
			for k := 0; k < len(parent1[i][j]); k++ {
				if i <= x { // kalo dia blom lewatin cross over point, gaada perubahan
					child1[i][j][k] = parent1[i][j][k]
					child2[i][j][k] = parent2[i][j][k]
				} else {
					child1[i][j][k] = parent2[i][j][k]
					child2[i][j][k] = parent1[i][j][k]
				}
			}
		}
	}
	return child1, child2 // dapetin 2 child
}

func mutate(individual [][][]int) [][][]int {
	usedNumbers := make(map[int]bool)
	duplicates := make([][3]int, 0) // buat nyimpen posisi duplicate

	//sekali pass: cari duplicate
	for i := 0; i < len(individual); i++ {
		for j := 0; j < len(individual[i]); j++ {
			for k := 0; k < len(individual[i][j]); k++ {
				num := individual[i][j][k]
				if usedNumbers[num] {
					duplicates = append(duplicates, [3]int{i, j, k})
				} else {
					usedNumbers[num] = true
				}
			}
		}
	}

	// pass kedua : ganti duplicate dengan angka random
	for _, pos := range duplicates {
		newNum := rand.Intn(125) + 1
		for usedNumbers[newNum] {
			newNum = rand.Intn(125) + 1
		}
		individual[pos[0]][pos[1]][pos[2]] = newNum
		usedNumbers[newNum] = true
	}

	// kemungkinan 7% untuk swap 2 baris
	if rand.Float64() < 0.07 {
		// Select two random positions
		i1, j1, k1 := rand.Intn(len(individual)), rand.Intn(len(individual[0])), rand.Intn(len(individual[0][0]))
		i2, j2, k2 := rand.Intn(len(individual)), rand.Intn(len(individual[0])), rand.Intn(len(individual[0][0]))

		// swap baris
		individual[i1][j1][k1], individual[i2][j2][k2] = individual[i2][j2][k2], individual[i1][j1][k1]
	}

	// 3% kemungkinan untuk swap 2 kolom
	if rand.Float64() < 0.03 {
		// Select two random columns
		col1 := rand.Intn(len(individual[0]))
		col2 := rand.Intn(len(individual[0]))

		//swap kolom
		for i := 0; i < len(individual); i++ {
			for j := 0; j < len(individual[i]); j++ {
				individual[i][j][col1], individual[i][j][col2] = individual[i][j][col2], individual[i][j][col1]
			}
		}
	}

	return individual
}

func evolvePopulation(population [][][][]int) [][][][]int { // Ini fungsi buat manggil dan dapetin populasi baru
	newPopulation := make([][][][]int, 0, len(population))

	for len(newPopulation) < len(population) {
		parent1 := selectParent(population)
		parent2 := selectParent(population)
		child1, child2 := crossover(parent1, parent2)
		newPopulation = append(newPopulation, mutate(child1))
		if len(newPopulation) < len(population) {
			newPopulation = append(newPopulation, mutate(child2))
		}
	}
	return newPopulation
}

func geneticAlgorithm(cube *[][][]int) ([][][]int, int, int, []int) {
	population := initializePopulation(cube, populationSize)
	generation := 0
	stuckCount := 0
	iterOF := []int{}
	var bestState [][][]int
	var bestFitness int

	for generation < maxGenerations {
		population = evolvePopulation(population) // Ini populasinya berubah terus setiap iterasi jadi makin bagus

		// Find the best individual in the current population
		for _, individual := range population {
			fitness := calculateObjectiveFunction(individual)
			if fitness > bestFitness {
				bestFitness = fitness
				bestState = individual
			}
		}

		if bestFitness == targetFitness {
			printCube(bestState)
			fmt.Println("success")
			return bestState, bestFitness, stuckCount, iterOF
		}
		generation++
		iterOF = append(iterOF, bestFitness)
	}

	fmt.Printf("Best fitness: %d\n", bestFitness)
	printCube(bestState)
	return bestState, bestFitness, stuckCount, iterOF
}
