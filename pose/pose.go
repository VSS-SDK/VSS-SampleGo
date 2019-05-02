package pose

import (
	"github.com/VSS-SDK/VSS-SampleGo/point"
	"math/rand"
)

type Pose struct {
	*point.Point
	Angle float32
}

func ZeroPose() *Pose {
	return &Pose{
		point.ZeroPoint(),
		0,
	}
}

func NewPose(x float32, y float32, angle float32) *Pose {
	return &Pose{
		point.NewPoint(x, y),
		angle,
	}
}

func RandPose() *Pose {
	return &Pose{
		point.RandPoint(),
		rand.Float32(),
	}
}
