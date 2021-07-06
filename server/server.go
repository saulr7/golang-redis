package server

import (
	"go-redis/handlers"
)

func Serve() {

	e := handlers.Routes()

	e.Logger.Fatal(e.Start(":1323"))

}
