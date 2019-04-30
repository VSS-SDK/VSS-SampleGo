package wheels_command_test

import (
	"github.com/VSS-SDK/VSS-SampleGo/wheels_command"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMapper(t *testing.T) {
	wheelsCommandMapper := wheels_command.NewMapper()

	assert.NotNil(t, wheelsCommandMapper)
}

func TestMapper_FromWheelsCommand(t *testing.T) {
	randWheelsCommand := wheels_command.RandWheelsCommand()
	wheelsCommandMapper := wheels_command.NewMapper()

	robotCommand := wheelsCommandMapper.FromWheelsCommand(randWheelsCommand)

	assert.NotNil(t, robotCommand)
	assert.Equal(t, randWheelsCommand.RightVel, *robotCommand.RightVel)
	assert.Equal(t, randWheelsCommand.LeftVel, *robotCommand.LeftVel)
}
