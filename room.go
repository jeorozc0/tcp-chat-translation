package main

import (
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
				m.msg("Couldn't translate message from " + sender.user + ": " + msg)
				return
			}
			m.msg(sender.user + ": " + translatedMessgae)
		}
	}
}
