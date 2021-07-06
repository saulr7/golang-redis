package main

import (
	"go-redis/db"
	"go-redis/server"
)

func main() {

	db.ConnectRedis()

	server.Serve()

}
