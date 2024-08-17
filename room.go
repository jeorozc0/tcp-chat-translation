package main

import (
	"fmt"
	"net"

	lang "jeorozco.com/go/tcp-chat-translation/language"
)

type room struct {
	name    string
	members map[net.Addr]*client
}

func (r *room) broadcast(sender *client, brd string) {
	for addr, m := range r.members {
		if addr != sender.conn.RemoteAddr() {
			m.brd(brd)
		}
	}
}

func (r *room) message(sender *client, msg string) {
	for addr, m := range r.members {
		if addr != sender.conn.RemoteAddr() {
			translatedMessgae, err := lang.TranslateMsg(msg, sender.language, m.language)
			if err != nil {
				m.msg(fmt.Sprintf("Couldn't translate message from: %v\n ", err))
				return
			}
			m.msg(sender.user + ": " + translatedMessgae)
		}
	}
}
