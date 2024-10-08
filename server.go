package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"strings"

	lang "jeorozco.com/go/tcp-chat-translation/language"
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
		language: "english",
		commands: s.comamnds,
	}
	c.readInput()

}

func (s *server) user(c *client, args []string) {
	c.user = args[1]
	c.brd(fmt.Sprintf("Your new username is %s", c.user))
}
func (s *server) join(c *client, args []string) {
	if len(args) < 2 {
		c.brd("room name is required. usage: /join ROOM_NAME")
		return
	}
	roomName := args[1]
	r, ok := s.rooms[roomName]
	if !ok {
		r = &room{
			name:    roomName,
			members: make(map[net.Addr]*client),
		}
		s.rooms[roomName] = r
	}

	r.members[c.conn.RemoteAddr()] = c

	s.quitRoom(c)
	c.room = r

	r.broadcast(c, fmt.Sprintf("%s has joined the room", c.user))
	c.brd(fmt.Sprintf("Welcome to %s", r.name))
}
func (s *server) servers(c *client) {
	var rooms []string
	for name := range s.rooms {
		rooms = append(rooms, name)
	}
	c.brd(fmt.Sprintf("Availble rooms: %s", strings.Join(rooms, ",")))
}
func (s *server) language(c *client, args []string) {
	if len(args) < 2 || args[1] == "" {
		languages := (strings.Join(lang.GetLanguages(), ", "))
		c.brd(fmt.Sprintf("Supported languages are: %s", languages))
		return
	}
	ok := lang.IsValidLanguage(args[1])
	if !ok {
		c.brd(fmt.Sprintf("%s is not a supported language.", args[1]))
		return
	}
	c.language = args[1]
	c.brd(fmt.Sprintf("Language set to %s", c.language))
}

func (s *server) message(c *client, args []string) {
	if c.room == nil {
		c.err(errors.New("You must be in a room to send messages"))
		return
	}
	c.room.message(c, strings.Join(args[0:], " "))
}

func (s *server) quit(c *client) {
	log.Printf("User has disconnected: %s", c.conn.RemoteAddr().String())

	s.quitRoom(c)

	c.brd("See you soon!")
	c.conn.Close()
}

func (s *server) quitRoom(c *client) {
	if c.room != nil {
		oldRoom := s.rooms[c.room.name]
		delete(s.rooms[c.room.name].members, c.conn.RemoteAddr())
		oldRoom.broadcast(c, fmt.Sprintf("%s has left the room", c.user))
	}
}
