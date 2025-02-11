package main

import "fmt"

var moviesInDB = []string{
	"Movie 12",
	"Movie 13",
	"Movie 14",
	"Movie 15",
	"Movie 16",
	"Movie 17",
	"Movie 18",
	"Movie 19",
	"Movie 20",
	"Movie 21",
	"Movie 22",
}

func main() {
	workingWithMap()
}

func workingWithMap() {
	m := make(map[string]string)
	n := map[string]string{"Matheus": "Dias"}

	m["Matheus"] = "Matheus"
	n["Matheus"] = "Matheus"

	fmt.Println(m == nil)
}

func boundsCheck(items []int) {
	_ = items[3]
	fmt.Println(items[0])
	fmt.Println(items[1])
	fmt.Println(items[2])
	fmt.Println(items[3])
}

func performaceInSlices() {
	resultsFromApi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	movies := make([]string, 0, 10)

	for _, id := range resultsFromApi {
		movie := moviesInDB[id]
		fmt.Println(len(movies), cap(movies))
		movies = append(movies, movie)
	}

	fmt.Println(len(movies), cap(movies))
	fmt.Println(movies)
}
