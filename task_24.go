package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float32
	y float32
}

func (p *Point) X() float32 {
	return p.x
}
func (p *Point) Y() float32 {
	return p.y
}
func (p *Point) setX(x float32) {
	p.x = x
}
func (p *Point) setY(y float32) {
	p.y = y
}

func newPoint(x, y float32) *Point {
	p := new(Point)
	p.setX(x)
	p.setY(y)

	return p
}

func (p *Point) distance(p2 *Point) float64 {
	return math.Sqrt(float64((p.X()-p2.X())*(p.X()-p2.X()) + (p.Y()-p2.Y())*(p.Y()-p2.Y())))
}

func main() {
	p1 := newPoint(0, 0)
	p2 := newPoint(1, 1)

	fmt.Println("distance = ", p1.distance(p2))
}
