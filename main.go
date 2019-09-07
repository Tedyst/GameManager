package main

import (
	"log"
	"net/http"

	"github.com/Tedyst/GameManager/pkg/handlers"
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
