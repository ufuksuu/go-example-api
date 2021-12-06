package main

import (
	"bookstore/models/db"
	"bookstore/server"
)

func main() {
	db.ConnectDatabase()
	server.Init()
}
