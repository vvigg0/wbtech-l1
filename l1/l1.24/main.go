package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) Point {
	return Point{x: x, y: y}
}
func (first Point) Distance(second Point) float64 {
	return math.Sqrt(math.Pow((second.x-first.x), 2) + math.Pow((second.y-first.y), 2))
}
func main() {
	a := NewPoint(5, 3)
	b := NewPoint(1, 1)

	fmt.Printf("%.3f", a.Distance(b))
}
