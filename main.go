package main

import (
	"bookstore/models"
	"bookstore/server"
)

func main() {
	models.ConnectDatabase()
	server.Init()
}
