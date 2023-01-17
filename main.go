package main

import (
	"log"

	"github.com/jayanthkrishna/golang-csrf-project/db"
	"github.com/jayanthkrishna/golang-csrf-project/server"
	"github.com/jayanthkrishna/golang-csrf-project/server/middleware/myjwt"
)

var host = "localhost"
var port = "9000"

func main() {
	db.InitDB()

	jwtErr := myjwt.InitJWT()

	if jwtErr != nil {
		log.Println("error Initializing JWT")
		log.Fatal(jwtErr)
	}

	serverErr := server.StartServer(host, port)

	if serverErr != nil {
		log.Println("Error Starting Server")
		log.Fatal(serverErr)
	}
}
