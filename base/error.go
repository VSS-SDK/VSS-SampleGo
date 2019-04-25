package base

import "errors"

type Error error

func NewSendClosedSocketError() Error {
	return errors.New("Cannot send message, closed socket")
}

func NewCreateSocketError() Error {
	return errors.New("Cannot create socket")
}

func NewConnectBindSocketError() Error {
	return errors.New("Cannot connect/bind socket")
}

func NewSendBytesError() Error {
	return errors.New("Cannot send bytes")
}
