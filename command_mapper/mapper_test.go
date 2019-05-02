package command_mapper_test

import (
	"github.com/VSS-SDK/VSS-SampleGo/command"
	"github.com/VSS-SDK/VSS-SampleGo/command_mapper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMapper(t *testing.T) {
	mapper := command_mapper.NewMapper()
	assert.NotNil(t, mapper)
}

func TestMapper_FromCommand(t *testing.T) {
	mapper := command_mapper.NewMapper()

	command := command.RandCommand()
	globalCommand := mapper.FromCommand(command)

	assert.NotNil(t, globalCommand)
	assert.NotEmpty(t, globalCommand.RobotCommands)

	for index, robotCommand := range globalCommand.RobotCommands {
		assert.NotNil(t, robotCommand)
		assert.Equal(t, robotCommand.LeftVel, &command.WheelsCommands[index].LeftVel)
		assert.Equal(t, robotCommand.RightVel, &command.WheelsCommands[index].RightVel)
	}
}
