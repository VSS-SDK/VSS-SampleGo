package main

import (
	"fmt"
	"github.com/VSS-SDK/VSS-SampleGo/command"
	"time"
)

func main() {
	globalCommand := command.RandCommand()

	fmt.Println("command")

	commandSender, err := command.NewSender()
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
