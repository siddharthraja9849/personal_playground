package main

import "fmt"

type Shape interface {
	getArea() (string, float64)
}

type Square struct {
	sideLength float64
	shapeType  string
}

type Triangle struct {
	height    float64
	base      float64
	shapeType string
}

func (s Square) getArea() (string, float64) {
	return s.shapeType, s.sideLength * s.sideLength
}

func (t Triangle) getArea() (string, float64) {
	return t.shapeType, 0.5 * t.base * t.height
}

func printArea(s Shape) {
	shapeType, area := s.getArea()
	fmt.Printf("Area %s is %f\n", shapeType, area)
}

func main() {
	sq := Square{sideLength: 5}
	tr := Triangle{height: 10, base: 12.5}
	printArea(sq)
	printArea(tr)
}
