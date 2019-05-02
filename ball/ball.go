package ball

import (
	"github.com/VSS-SDK/VSS-SampleGo/point"
	"math/rand"
)

type Ball struct {
	*point.Point
	SpeedX float32
	SpeedY float32
}

func ZeroBall() *Ball {
	return &Ball{
		point.ZeroPoint(),
		0,
		0,
	}
}

func NewBall(x float32, y float32, speedX float32, speedY float32) *Ball {
	return &Ball{
		point.NewPoint(x, y),
		speedX,
		speedY,
	}
}

func RandBall() *Ball {
	return &Ball{
		point.RandPoint(),
		rand.Float32() + 1,
		rand.Float32() + 1,
	}
}
