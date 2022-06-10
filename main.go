package main

import (
	"log"
	"net/http"

	"github.com/averagebit/k8s-demo-app/app"
)

func main() {
	handler := http.HandlerFunc(app.Server)
	log.Fatal(http.ListenAndServe(":3000", handler))
}
