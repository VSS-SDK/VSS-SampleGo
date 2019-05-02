package robot_test

import (
	"github.com/VSS-SDK/VSS-SampleGo/robot"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestZeroRobot(t *testing.T) {
	robot := robot.ZeroRobot()

	assert.NotNil(t, robot)
	assert.Zero(t, robot.X)
	assert.Zero(t, robot.Y)
	assert.Zero(t, robot.Angle)
	assert.Zero(t, robot.SpeedX)
	assert.Zero(t, robot.SpeedY)
	assert.Zero(t, robot.SpeedAngle)
}

func TestRandRobot(t *testing.T) {
	robot := robot.RandRobot()

	assert.NotNil(t, robot)
	assert.NotNil(t, robot.X)
	assert.NotZero(t, robot.Y)
	assert.NotZero(t, robot.Angle)
	assert.NotZero(t, robot.SpeedX)
	assert.NotZero(t, robot.SpeedY)
	assert.NotZero(t, robot.SpeedAngle)
}

func TestNewRobot(t *testing.T) {
	rand := robot.RandRobot()
	robot := robot.NewRobot(rand.X, rand.Y, rand.Angle, rand.SpeedX, rand.SpeedY, rand.SpeedAngle)

	assert.NotNil(t, robot)
	assert.Equal(t, robot.X, rand.X)
	assert.Equal(t, robot.Y, rand.Y)
	assert.Equal(t, robot.Angle, rand.Angle)
	assert.Equal(t, robot.SpeedX, rand.SpeedX)
	assert.Equal(t, robot.SpeedY, rand.SpeedY)
	assert.Equal(t, robot.SpeedAngle, rand.SpeedAngle)
}
