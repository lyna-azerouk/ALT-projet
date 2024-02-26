package main

import (
	"log"
	router2 "serveur/server/router"

	_ "github.com/lib/pq"
)

func main() {
	router := router2.SetUpRouter()
	err := router.Run(":8080")
	if err == nil {
		log.Fatal("Error while starting the server: " + err.Error())
		return
	}

}
