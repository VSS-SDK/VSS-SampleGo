package command_test

import (
	"github.com/VSS-SDK/VSS-SampleGo/command"
	"github.com/VSS-SDK/VSS-SampleGo/wheels_command"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMapper(t *testing.T) {
	mapper := command.NewMapper(wheels_command.NewMapper())

	assert.NotNil(t, mapper)
}
