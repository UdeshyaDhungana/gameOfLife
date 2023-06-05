package main

// Iterates through the neighbors of current points, and applies rule
func NextStep(previousPoints *PointSet) *PointSet {
	var visitedPoints = PointSet{map[Point]bool{}}
	nextPoints := PointSet{map[Point]bool{}}
	for point := range previousPoints.getPoints() {
		neighbors := Neighbors(point)
		for _, n := range neighbors {
			isVisited := visitedPoints.contains(n)
			if isVisited {
				continue
			}
			aliveInNextStep := ApplyRule(n, previousPoints, previousPoints.contains(n))
			if aliveInNextStep {
				nextPoints.insert(n)
			}
		}
	}
	return &nextPoints
}
