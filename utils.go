package main

type Point struct {
	X, Y int
}

func Neighbors(p Point) []Point {
	neighbors := make([]Point, 0)
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			neighbors = append(neighbors, Point{X: p.X + i, Y: p.Y + j})
		}
	}
	return neighbors
}

// game of life's actual logic
func ApplyRule(p Point, previous *PointSet, isAlive bool) bool {
	neighbors := Neighbors(p)
	aliveNeighbors := 0
	for _, n := range neighbors {
		if previous.contains(n) {
			aliveNeighbors++
		}
	}
	if !isAlive {
		return aliveNeighbors == 3
	}

	return aliveNeighbors == 2 || aliveNeighbors == 3
}
