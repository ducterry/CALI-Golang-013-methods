package main

type Point struct {
	pointX0, pointY0 float64
}

func (p Point) IsAbove(pointY1 float64) bool {
	return p.pointY0 > pointY1
}
