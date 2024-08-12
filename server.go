package main

import (
	"log"
	"net"
)

type server struct {
	rooms    map[string]*room
	comamnds chan command
}

func newServer() *server {
	return &server{
		rooms:    make(map[string]*room),
		comamnds: make(chan command),
	}
}

func (s *server) run() {
	for cmd := range s.comamnds {
		switch cmd.id {
		case CMD_USER:
			s.user(cmd.client, cmd.args)
		case CMD_JOIN:
			s.join(cmd.client, cmd.args)
		case CMD_SERVERS:
			s.servers(cmd.client)
		case CMD_LANG:
			s.language(cmd.client, cmd.args)
		case CMD_MSG:
			s.message(cmd.client, cmd.args)
		case CMD_QUIT:
			s.quit(cmd.client)

		}
	}
}
func (s *server) newClient(conn net.Conn) {
	log.Printf("new client has connected: %s", conn.RemoteAddr().String())

	c := &client{
		conn:     conn,
		user:     "anon",
		language: "English",
		commands: s.comamnds,
	}
	c.readInput()

}
