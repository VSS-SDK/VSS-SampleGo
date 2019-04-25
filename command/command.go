package command

type Command struct {
	wheelsCommands []WheelsCommand
}

func ZeroCommand() Command {
	var wheelsCommands []WheelsCommand

	wheelsCommands = append(wheelsCommands, ZeroWheelsCommand())
	wheelsCommands = append(wheelsCommands, ZeroWheelsCommand())
	wheelsCommands = append(wheelsCommands, ZeroWheelsCommand())

	return Command{
		wheelsCommands,
	}
}

func NewCommand(wheelsCommands []WheelsCommand) Command {
	return Command{
		wheelsCommands,
	}
}

func RandCommand() Command {
	var wheelsCommands []WheelsCommand

	wheelsCommands = append(wheelsCommands, RandWheelsCommand())
	wheelsCommands = append(wheelsCommands, RandWheelsCommand())
	wheelsCommands = append(wheelsCommands, RandWheelsCommand())

	return Command{
		wheelsCommands,
	}
}
