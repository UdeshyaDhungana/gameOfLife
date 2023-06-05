package main

import "fmt"

type PointSet struct {
	Points map[Point]bool
}

func (s *PointSet) getPoints() map[Point]bool {
	return s.Points
}

func (s *PointSet) insert(p Point) {
	s.Points[p] = true
}

func (s *PointSet) contains(p Point) bool {
	_, ok := s.Points[p]
	return ok
}

func (s *PointSet) Println() {
	fmt.Printf("[")
	for p := range s.Points {
		fmt.Printf("(%d, %d)", p.X, p.Y)
	}
	fmt.Printf("]\n")
}
