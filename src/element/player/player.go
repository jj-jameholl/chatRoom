package player

import "net"

type Player struct {
	Name string
	Id   int
	Conn net.Conn
}
