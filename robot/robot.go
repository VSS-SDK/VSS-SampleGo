package robot

import (
	"github.com/VSS-SDK/VSS-SampleGo/pose"
	"math/rand"
)

type Robot struct {
	*pose.Pose
	SpeedX     float32
	SpeedY     float32
	SpeedAngle float32
}

func ZeroRobot() *Robot {
	return &Robot{
		pose.ZeroPose(),
		0,
		0,
		0,
	}
}

func NewRobot(x float32, y float32, angle float32, speedX float32, speedY float32, speedAngle float32) *Robot {
	return &Robot{
		pose.NewPose(x, y, angle),
		speedX,
		speedY,
		speedAngle,
	}
}

func RandRobot() *Robot {
	return &Robot{
		pose.RandPose(),
		rand.Float32() + 1,
		rand.Float32() + 1,
		rand.Float32() + 1,
	}
}
