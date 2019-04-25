package command

import (
	"../base"
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
	Send(command Command) error
}

type sender struct {
	socket *zmq4.Socket
	mapper Mapper
}

func (s *sender) Send(command Command) error {
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
		NewMapper(),
	}

	return sender, nil
}
