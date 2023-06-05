package main

import (
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage: ./executable <number_of_iterations>")
	}
	initialPoints := ReadSeed()
	numIters, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalln("Provide an integer argument for number of iterations")
	}
	for i, j := 0, initialPoints; i < numIters; i, j = i+1, NextStep(j) {
		j.Println()
	}
}
