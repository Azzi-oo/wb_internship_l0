package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) Point {
	return Point{x, y}
}

func (p1 Point) Distance(p2 Point) float64 {
	dx := p2.x - p1.x
	dy := p2.y - p1.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	point1 := NewPoint(3, 4)
	point2 := NewPoint(6, 8)

	distance := point1.Distance(point2)

	fmt.Printf("Расстояние между точками: %.2f\n", distance)
}
