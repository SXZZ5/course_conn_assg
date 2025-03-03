package main

import (
	"assignment/api"
	"assignment/db"
)

func main() {
	storage := db.Storage{}
	storage.Connect()
	server := api.Server{}
	server.RoutesInit(&storage)
	
	storage.PrintAll()
	server.Run()
}
