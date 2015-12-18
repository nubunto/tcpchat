package server

import (
	"github.com/nubunto/tcpchat/types"
	"net"
)

type Connection struct {
	conn net.Conn
	types.User
}

type ConnectionList []Connection
