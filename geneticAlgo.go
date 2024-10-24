package main

import (
	"fmt" 
	"math/rand" 
)

const (
    populationSize = 50
    mutationRate   = 0.01
    maxGenerations = 100
    targetFitness  = 109
)


func initializePopulation(size int) [][][][]int {
    population := make([][][][]int, size)
    for i := 0; i < size; i++ {
        population[i] = generateRandom5x5x5Array()
    }
    return population // Nyari populasi / kumpulan konfigurasi kubus buat dijadiin parent
}

func selectParent(population [][][][]int) [][][]int {
    tournamentSize := 5 // Cari 5 kandidat dari populasi yang udah dihasilin
    best := population[rand.Intn(len(population))] // Inisialisasi best pertama dengan random
    bestFitness := calculateObjectiveFunction(best)

    for i := 1; i < tournamentSize; i++ { // looping buat cari kandidat terbaik dari 5 kandidat
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

    // Cross oveer point 
    x := 3

    // Perform crossover
    for i := 0; i < len(parent1); i++ {
        for j := 0; j < len(parent1[i]); j++ {
            for k := 0; k < len(parent1[i][j]); k++ {
                if (i <= x) {
                    child1[i][j][k] = parent1[i][j][k]
                    child2[i][j][k] = parent2[i][j][k]
                } else {
                    child1[i][j][k] = parent2[i][j][k]
                    child2[i][j][k] = parent1[i][j][k]
                }
            }
        }
    }
    return child1, child2
}

func mutate(individual [][][]int) [][][]int {
    for i := 0; i < len(individual); i++ {
        for j := 0; j < len(individual[i]); j++ {
            for k := 0; k < len(individual[i][j]); k++ {
                if rand.Float64() < mutationRate { // mutation rate ini 0.01, jadi 1% kemungkinan buat mutasi
                    x := rand.Intn(126)
                    individual[i][j][k] = x
                }
            }
        }
    }
    return individual
}

func evolvePopulation(population [][][][]int) [][][][]int {
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


func geneticAlgorithm() {
    population := initializePopulation(populationSize)
    generation := 0

    for generation < maxGenerations {
        population = evolvePopulation(population)
        bestFitness := calculateObjectiveFunction(population[0])
        fmt.Printf("Generation: %d, Best Fitness: %d\n", generation, bestFitness)

        if bestFitness == targetFitness {
            printCube(population[0])
            fmt.Println("success")
            return
        }
        generation++
    }

    printCube(population[0])
    fmt.Println("fail")
}

