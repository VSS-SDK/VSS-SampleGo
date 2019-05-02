package wheels_command

import "github.com/VSS-SDK/VSS-SampleGo/protos"

type Mapper interface {
	ToWheelsCommand(robot *protos.Robot_Command) *WheelsCommand
	FromWheelsCommand(command *WheelsCommand) *protos.Robot_Command
}

type mapper struct {
}

func NewMapper() Mapper {
	return &mapper{}
}

func (m *mapper) ToWheelsCommand(robot *protos.Robot_Command) *WheelsCommand {
	return NewWheelsCommand(*robot.LeftVel, *robot.RightVel)
}

func (m *mapper) FromWheelsCommand(command *WheelsCommand) *protos.Robot_Command {
	return &protos.Robot_Command{
		LeftVel:  &command.LeftVel,
		RightVel: &command.RightVel,
	}
}
