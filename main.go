package main

import (
	"log"

	"github.com/vSivarajah/distributed-kafka-go/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
