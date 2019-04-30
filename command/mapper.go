package command

import "github.com/VSS-SDK/VSS-SampleGo/protos"

type Mapper interface {
	ToCommand(command protos.Global_Commands) Command
	FromCommand(command Command) *protos.Global_Commands
}

type mapper struct {
}

func NewMapper() Mapper {
	return &mapper{}
}

func (m *mapper) ToCommand(command protos.Global_Commands) Command {
	var wheelsCommands []WheelsCommand

	for _, c := range command.RobotCommands {
		wheelsCommands = append(wheelsCommands, m.toWheelsCommand(c))
	}

	return NewCommand(wheelsCommands)
}

func (m *mapper) FromCommand(command Command) *protos.Global_Commands {
	var robots []*protos.Robot_Command

	for _, c := range command.wheelsCommands {
		robots = append(robots, m.fromWheelsCommand(c))
	}

	return &protos.Global_Commands{
		RobotCommands: robots,
	}
}

func (m *mapper) toWheelsCommand(robot *protos.Robot_Command) WheelsCommand {
	return NewWheelsCommand(*robot.LeftVel, *robot.RightVel)
}

func (m *mapper) fromWheelsCommand(command WheelsCommand) *protos.Robot_Command {
	return &protos.Robot_Command{
		LeftVel:  &command.leftVel,
		RightVel: &command.rightVel,
	}
}
