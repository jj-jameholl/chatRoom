package base

import (
	"net"
)

type Server interface {
	Handle(method string, params []string, conn net.Conn)
	Name() string
}
