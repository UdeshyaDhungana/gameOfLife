package main

// We work with just pointers here, okay!

import (
	"encoding/json"
	"log"
	"os"
)

func ReadSeed() *PointSet {
	file, err := os.ReadFile("seed.json")
	if err != nil {
		log.Fatalf("%v", err)
	}

	var pointsArr []*Point

	err = json.Unmarshal(file, &pointsArr)

	if err != nil {
		log.Fatalf("%v", err)
	}

	points := PointSet{map[Point]bool{}}
	for _, p := range pointsArr {
		points.insert(*p)
	}

	return &points
}
