package command

import (
	"github.com/VSS-SDK/VSS-SampleGo/protos"
	"github.com/VSS-SDK/VSS-SampleGo/wheels_command"
)

type Mapper interface {
	ToCommand(command protos.Global_Commands) Command
	FromCommand(command Command) *protos.Global_Commands
}

type mapper struct {
	wheels_command_mapper wheels_command.Mapper
}

func NewMapper(wheels_command_mapper wheels_command.Mapper) Mapper {
	return &mapper{
		wheels_command_mapper,
	}
}

func (m *mapper) ToCommand(command protos.Global_Commands) Command {
	var wheelsCommands []wheels_command.WheelsCommand

	for _, c := range command.RobotCommands {
		wheelsCommands = append(wheelsCommands, m.wheels_command_mapper.ToWheelsCommand(c))
	}

	return NewCommand(wheelsCommands)
}

func (m *mapper) FromCommand(command Command) *protos.Global_Commands {
	var robots []*protos.Robot_Command

	for _, c := range command.WheelsCommands {
		robots = append(robots, m.wheels_command_mapper.FromWheelsCommand(c))
	}

	return &protos.Global_Commands{
		RobotCommands: robots,
	}
}
