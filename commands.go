package main

type commandID int

const (
	CMD_USER commandID = iota
	CMD_JOIN
	CMD_SERVERS
	CMD_LANG
	CMD_MSG
	CMD_QUIT
)

type command struct {
	id     commandID
	client *client
	args   []string
}
