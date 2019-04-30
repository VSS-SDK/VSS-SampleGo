package command

import (
	"github.com/VSS-SDK/VSS-SampleGo/wheels_command"
	"math/rand"
)

type Command struct {
	WheelsCommands []wheels_command.WheelsCommand
}

func ZeroCommand() Command {
	var wheelsCommands []wheels_command.WheelsCommand

	return Command{
		wheelsCommands,
	}
}

func NewCommand(wheelsCommands []wheels_command.WheelsCommand) Command {
	return Command{
		wheelsCommands,
	}
}

func RandCommand() Command {
	var wheelsCommands []wheels_command.WheelsCommand

	qtd := rand.Intn(10) + 1

	for i := 0; i < qtd; i++ {
		wheelsCommands = append(wheelsCommands, wheels_command.RandWheelsCommand())
	}

	return Command{
		wheelsCommands,
	}
}
