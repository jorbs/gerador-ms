package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"time"
)

var Version = "0.1"

func main() {
	version := flag.Bool("v", false, "prints the version")
	minimum := flag.Int("min", 1, "defines minumum value")
	maximum := flag.Int("max", 60, "defines maximum value")
	seed := flag.Int64("seed", 1, "defines a seed for numbers generation")
	n := flag.Int("n", 5, "number of bets")

	flag.Parse()

	if *version {
		fmt.Println(Version)
		os.Exit(0)
	}

	if *minimum >= *maximum {
		panic("invalid interval.")
	}

	for i := 0; i < *n; i++ {
		generateNumbers(*minimum, *maximum, *seed)
	}
}

func generateNumbers(min, max int, seed int64) {
	timeSeed := rand.NewSource(time.Now().UnixNano() + seed)
	r1 := rand.New(timeSeed)

	result := make(map[int]int)

	for i := 0; i < 6; {
		chosen := r1.Intn(max)
		_, ok := result[chosen]

		if chosen >= min && !ok {
			result[chosen] = 1
			i++
		}
	}

	var keys []int

	for k, _ := range result {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for _, v := range keys {
		fmt.Print(v, " ")
	}
	fmt.Println("")
}
