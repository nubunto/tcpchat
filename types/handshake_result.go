package types

import "fmt"

/*
  A result from a Handshake.
*/

type HandshakeCode int

const (
	Ok = iota + 200
	AlreadyExists
	GenericError
)

type HandshakeResult struct {
	Code    int
	Message string
}

func (h HandshakeResult) Error() string {
	return fmt.Sprintf("[Code %d] Error: %s", h.Code, h.Message)
}
