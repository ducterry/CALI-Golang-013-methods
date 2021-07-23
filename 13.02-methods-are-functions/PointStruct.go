package main

type Point struct {
	pointX0, pointY0 float64
}

func IsAbove(myPoint Point, pointY1 float64) bool {
	return myPoint.pointY0 > pointY1
}
