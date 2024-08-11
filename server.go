package tcpchattranslation

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
