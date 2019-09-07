package globals

import "net"

// Server is a struct that defines every server
type Server struct {
	IP          net.IP
	PlayerCount int16
	ProxyIP     net.IP
	TPS         float32
	Port        int16
}

// Player is every connected player
type Player struct {
	Location Server
}

// ServerList is a list of all servers connected here
var ServerList []Server

// PlayerList is a list of all players connected here
var PlayerList map[string]Player

// PlayerJSON is the JSON that is send by the server here
type PlayerJSON struct {
	Name string
	UUID string
	Type int
}
