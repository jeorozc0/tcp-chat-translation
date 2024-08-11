package tcpchattranslation

import "net"

type client struct {
	conn     net.Addr
	user     string
	room     *room
	commands chan<- command
}
