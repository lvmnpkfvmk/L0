package main

// Разработать программу нахождения расстояния между двумя точками, которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

func (p *Point) String() string {
	return fmt.Sprintf("Point {%.2f, %.2f}", p.x, p.y)
}

func (p *Point) DistanceTo(other *Point) float64 {
	dx := p.x - other.x
	dy := p.y - other.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	x := NewPoint(1.0, 2.0)
	y := NewPoint(3.0, 4.0)

	distance := x.DistanceTo(y)

	fmt.Println(x, y) // Point {1.00, 2.00} Point {4.00, 6.00}
	fmt.Printf("Distance: %.2f\n", distance) // Distance: 2.83
}