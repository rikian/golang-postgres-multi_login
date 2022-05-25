package main

import (
	"golang/config"
	"golang/controllers/routings"

	"log"
	"net/http"
)

func main() {
	dbConnection, status := config.ConnectDB()
	if status {
		log.Println(dbConnection)
		return
	}

	log.Println(dbConnection)

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		routings.SlashRouting(res, req)
	})

	log.Println("Golang server listening on port 9091...")
	http.ListenAndServe("127.0.0.1:9091", nil)
}