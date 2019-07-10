package main

import (
	"github.com/Tedyst/GameManager/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/addserver", handlers.AddServerHandler)
	http.HandleFunc("/removeserver", handlers.RemoveServerHandler)
	http.HandleFunc("/listservers", handlers.GetServerListHandler)
	http.HandleFunc("/getserver", handlers.GetServerForPlayer)
	err := http.ListenAndServe(":3001", nil)
	if err != nil {
		log.Println("Error running server: ", err)
	}
}
