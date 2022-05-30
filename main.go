package main

import (
	"log"
	"net/http"

	"github.com/averagebit/circle-docker/api/app"
)

func main() {
	handler := http.HandlerFunc(app.Server)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
