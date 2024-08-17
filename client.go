package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type client struct {
	conn     net.Conn
	user     string
	language string
	room     *room
	commands chan<- command
}

const PROMPT = ">> "

func (c *client) readInput() {
	writer := bufio.NewWriter(c.conn)
	reader := bufio.NewReader(c.conn)
	for {
		writer.WriteString(PROMPT)
		writer.Flush()
		msg, err := reader.ReadString('\n')
		if err != nil {
			return
		}

		msg = strings.Trim(msg, "\r\n")

		args := strings.Split(msg, " ")
		cmd := strings.TrimSpace(args[0])

		if !strings.HasPrefix(cmd, "/") {
			c.commands <- command{
				id:     CMD_MSG,
				client: c,
				args:   args,
			}
		} else {
			switch cmd {
			case "/user":
				c.commands <- command{
					id:     CMD_USER,
					client: c,
					args:   args,
				}
			case "/join":
				c.commands <- command{
					id:     CMD_JOIN,
					client: c,
					args:   args,
				}
			case "/servers":
				c.commands <- command{
					id:     CMD_SERVERS,
					client: c,
				}

			case "/quit":
				c.commands <- command{
					id:     CMD_QUIT,
					client: c,
				}
			case "/lang":
				c.commands <- command{
					id:     CMD_LANG,
					client: c,
					args:   args,
				}
			default:
				c.err(fmt.Errorf("unknown command: %s", cmd))
			}
		}

	}
}

func (c *client) err(err error) {
	c.conn.Write([]byte("ERR: " + err.Error() + "\n"))
	c.conn.Write([]byte(PROMPT))
}

func (c *client) msg(msg string) {
	c.conn.Write([]byte(msg + "\n"))
	c.conn.Write([]byte(PROMPT))
}

func (c *client) brd(brd string) {
	c.conn.Write([]byte(brd + "\n"))
	c.conn.Write([]byte(PROMPT))
}
