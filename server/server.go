package server

import (
	"log"
	"net/http"

	"github.com/jayanthkrishna/golang-csrf-project/server/middleware"
)

func StartServer(host, port string) error {
	hostname := host + ":" + port

	log.Printf("Listening on : %s\n", hostname)
	handler := middleware.NewHandler()
	http.Handle("/", handler)

	return http.ListenAndServe(hostname, nil)
}
