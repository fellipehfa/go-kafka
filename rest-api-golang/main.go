package main

import (
	"rest-api/database"
	"rest-api/server"
)

func main() {
	database.StartDB()

	server := server.NewServer()

	server.Run()
}
