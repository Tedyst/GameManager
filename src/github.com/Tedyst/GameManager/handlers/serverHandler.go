package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/Tedyst/GameManager/globals"
	"github.com/Tedyst/GameManager/utils"
	"net/http"
)

// AddServerHandler adds a Server to ServerList
func AddServerHandler(w http.ResponseWriter, r *http.Request) {
	ip := utils.GetIP(r)
	globals.ServerList = append(globals.ServerList, globals.Server{ip, 0, ip, -1, 25565})
}

// RemoveServerHandler removes a server from ServerList
func RemoveServerHandler(w http.ResponseWriter, r *http.Request) {
	ip := utils.GetIP(r)
	for i := len(globals.ServerList) - 1; i >= 0; i-- {
		elem := &globals.ServerList[i]
		if ip.Equal(elem.IP) {
			globals.ServerList = append(globals.ServerList[:i],
				globals.ServerList[i+1:]...)
		}
	}
}

// GetServerListHandler writes a list of all servers
func GetServerListHandler(w http.ResponseWriter, r *http.Request) {
	str, err := json.Marshal(globals.ServerList)
	if err != nil {
		fmt.Fprintf(w, "Error")
		return
	}
	fmt.Fprintf(w, string(str))
}

// GetServerForPlayer yes
func GetServerForPlayer(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var t globals.PlayerJSON
	err := decoder.Decode(&t)
	if err != nil {
		fmt.Fprintf(w, "Error")
		return
	}
	var server globals.Server
	for _, elem := range globals.ServerList {
		if elem.TPS > 17 && elem.PlayerCount < 30 {
			server = elem
			break
		}
	}
	print(server.IP)
	temp := globals.PlayerList[t.UUID]
	temp.Location = server
	globals.PlayerList[t.UUID] = temp
}
