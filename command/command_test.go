package command_test

import (
	"github.com/VSS-SDK/VSS-SampleGo/command"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestZeroCommand(t *testing.T) {
	cmd := command.ZeroCommand()

	assert.NotNil(t, cmd)
	assert.Empty(t, cmd.WheelsCommands)
}

func TestRandCommand(t *testing.T) {
	cmd := command.RandCommand()

	assert.NotNil(t, cmd)
	assert.NotEmpty(t, cmd.WheelsCommands)
}

func TestNewCommand(t *testing.T) {
	rand := command.RandCommand()
	cmd := command.NewCommand(rand.WheelsCommands)

	assert.NotNil(t, cmd)
	assert.ElementsMatch(t, cmd.WheelsCommands, rand.WheelsCommands)
}
