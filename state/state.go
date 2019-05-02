package state

import (
	"github.com/VSS-SDK/VSS-SampleGo/ball"
	"github.com/VSS-SDK/VSS-SampleGo/robot"
	"math/rand"
)

type State struct {
	YellowRobots []*robot.Robot
	BlueRobots   []*robot.Robot
	Ball         *ball.Ball
}

func ZeroState() *State {
	var yellowRobots []*robot.Robot
	var blueRobots []*robot.Robot

	return &State{
		yellowRobots,
		blueRobots,
		ball.ZeroBall(),
	}
}

func RandState() *State {
	var yellowRobots []*robot.Robot
	var blueRobots []*robot.Robot

	qtdYellow := rand.Intn(10) + 1
	for i := 0; i < qtdYellow; i++ {
		yellowRobots = append(yellowRobots, robot.RandRobot())
	}

	qtdBlue := rand.Intn(10) + 1
	for i := 0; i < qtdBlue; i++ {
		blueRobots = append(blueRobots, robot.RandRobot())
	}

	return &State{
		yellowRobots,
		blueRobots,
		ball.RandBall(),
	}
}

func NewState(yellowRobots []*robot.Robot, blueRobots []*robot.Robot, ball *ball.Ball) *State {
	return &State{
		yellowRobots,
		blueRobots,
		ball,
	}
}
