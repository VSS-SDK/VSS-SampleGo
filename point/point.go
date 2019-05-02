package point

import "math/rand"

type Point struct {
	X float32
	Y float32
}

func ZeroPoint() *Point {
	return &Point{
		0,
		0,
	}
}

func NewPoint(x float32, y float32) *Point {
	return &Point{
		x,
		y,
	}
}

func RandPoint() *Point {
	return &Point{
		rand.Float32() + 1,
		rand.Float32() + 1,
	}
}
