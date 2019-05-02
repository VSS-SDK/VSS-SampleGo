package command_mapper_test

import (
	"github.com/VSS-SDK/VSS-SampleGo/command_mapper"
	"github.com/VSS-SDK/VSS-SampleGo/wheels_command"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMapper(t *testing.T) {
	mapper := command_mapper.NewMapper(wheels_command.NewMapper())

	assert.NotNil(t, mapper)
}
