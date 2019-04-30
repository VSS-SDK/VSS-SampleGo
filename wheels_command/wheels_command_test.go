package wheels_command_test

import (
	"github.com/VSS-SDK/VSS-SampleGo/wheels_command"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestZeroWheelsCommand(t *testing.T) {
	wheelsCommand := wheels_command.ZeroWheelsCommand()

	assert.NotNil(t, wheelsCommand)
	assert.Zero(t, wheelsCommand.LeftVel)
	assert.Zero(t, wheelsCommand.RightVel)
}

func TestRandWheelsCommand(t *testing.T) {
	wheelsCommand := wheels_command.RandWheelsCommand()

	assert.NotNil(t, wheelsCommand)
	assert.NotZero(t, wheelsCommand.LeftVel)
	assert.NotZero(t, wheelsCommand.RightVel)
}

func TestNewWheelsCommand(t *testing.T) {
	rand := wheels_command.RandWheelsCommand()
	wheelsCommand := wheels_command.NewWheelsCommand(rand.LeftVel, rand.RightVel)

	assert.NotNil(t, wheelsCommand)
	assert.Equal(t, wheelsCommand.LeftVel, rand.LeftVel)
	assert.Equal(t, wheelsCommand.RightVel, rand.RightVel)
}
