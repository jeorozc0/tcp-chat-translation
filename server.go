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
			s.user()
		case CMD_JOIN:
			s.join()
		case CMD_SERVERS:
			s.servers()
		case CMD_LANG:
			s.language()
		case CMD_MSG:
			s.message()
		case CMD_QUIT:
			s.quit()

		}
	}
}
func (s *server) newClient(conn net.Conn) {
	log.Printf("new client has connected: 5s", conn.RemoteAddr().String())

	c := &client{
		conn:     conn,
		user:     "anon",
		commands: s.comamnds,
	}
	c.readInput()

}
