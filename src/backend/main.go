package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// arr := [][][]int{ // ini array yg udah perfect magic cube
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
	// 		{121, 7, 108, 20, 59},
	// 		{29, 28, 122, 125, 11},
	// 		{51, 15, 41, 124, 84},
	// 		{78, 54, 99, 24, 60},
	// 		{36, 110, 46, 22, 101},
	// 	},
	// }

	// var arr [][][]int = generateRandom5x5x5Array()
	// printCube(arr)
	// simulatedAnnealing(&arr, 10000, 0.999, 1000000)
	// http.HandleFunc("/calculate")

	// stochasticHillClimbing(&arr, 500000)
	// steepestAscentHillClimbing(&arr)
	// sidewaysMoveHillClimbing(&arr)
	fmt.Println("Server is running on port 8080")
	r := mux.NewRouter()

	// Tentukan route yang Anda gunakan, misalnya:
	r.HandleFunc("/search", searchHandler).Methods("GET", "POST")

	// Konfigurasi CORS
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"http://localhost:3000"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})

	// Gunakan CORS middleware
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(r)))
}
