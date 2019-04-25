package command

import "math/rand"

type WheelsCommand struct {
	leftVel  float32
	rightVel float32
}

func ZeroWheelsCommand() WheelsCommand {
	return WheelsCommand{
		0.0,
		0.0,
	}
}

func NewWheelsCommand(leftVel float32, rightVel float32) WheelsCommand {
	return WheelsCommand{
		leftVel,
		rightVel,
	}
}

func RandWheelsCommand() WheelsCommand {
	return WheelsCommand{
		rand.Float32(),
		rand.Float32(),
	}
}
