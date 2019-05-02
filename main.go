package main

import (
	"fmt"
	"github.com/VSS-SDK/VSS-SampleGo/command"
	"github.com/VSS-SDK/VSS-SampleGo/command_sender"
	"github.com/VSS-SDK/VSS-SampleGo/wheels_command"
	"time"
)

func main() {
	globalCommand := buildCommand()

	fmt.Println("command")

	commandSender, err := command_sender.NewSender()
	if err != nil {
		fmt.Println("Cannot create sender")
		return
	}

	fmt.Println("socket")

	for true {
		fmt.Println("send")
		time.Sleep(1000 * time.Millisecond)
		err := commandSender.Send(globalCommand)

		if err != nil {
			return
		}
	}
}

func buildCommand() *command.Command {
	var wheelsCommands []*wheels_command.WheelsCommand

	for i := 0; i < 3; i++ {
		wheelsCommands = append(wheelsCommands, wheels_command.NewWheelsCommand(10, -10))
	}

	return command.NewCommand(wheelsCommands)
}
