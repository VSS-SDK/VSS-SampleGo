package command_sender

import (
	"github.com/VSS-SDK/VSS-SampleGo/base"
	"github.com/VSS-SDK/VSS-SampleGo/command"
	"github.com/VSS-SDK/VSS-SampleGo/command_mapper"
	"github.com/VSS-SDK/VSS-SampleGo/wheels_command"
	"github.com/golang/protobuf/proto"
	"github.com/pebbe/zmq4"
)

var (
	ErrSendClosedSocket  = base.NewSendClosedSocketError()
	ErrCreateSocket      = base.NewCreateSocketError()
	ErrConnectBindSocket = base.NewConnectBindSocketError()
	ErrSendBytes         = base.NewSendBytesError()
)

type Sender interface {
	Send(command *command.Command) error
}

type sender struct {
	socket *zmq4.Socket
	mapper command_mapper.Mapper
}

func (s *sender) Send(command *command.Command) error {
	if s.socket == nil {
		return ErrSendClosedSocket
	}

	globalCommands := s.mapper.FromCommand(command)

	bytes, _ := proto.Marshal(globalCommands)

	_, err := s.socket.SendBytes(bytes, 0)
	if err != nil {
		return ErrSendBytes
	}

	return nil
}

func NewSender() (Sender, error) {
	socket, err := zmq4.NewSocket(zmq4.PAIR)
	if err != nil {
		return nil, ErrCreateSocket
	}

	err = socket.Connect("tcp://localhost:5556")
	if err != nil {
		return nil, ErrConnectBindSocket
	}

	sender := &sender{
		socket,
		command_mapper.NewMapper(wheels_command.NewMapper()),
	}

	return sender, nil
}
