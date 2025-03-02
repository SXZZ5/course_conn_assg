package main

import (
	"assignment/api"
	"assignment/db"
)

func main() {
	server := api.Server{}
	server.Init()
	db := db.Storage{}
	db.Connect()
	db.PrintAll()
	server.Run()
}
