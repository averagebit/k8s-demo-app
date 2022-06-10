package app

import (
	"fmt"
	"net/http"
	"os"
)

func Server(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	if name == "" {
		fmt.Fprintf(w, "Hello from Go!")
		return
	}
	fmt.Fprintf(w, "Hello from %s!", os.Getenv("NAME"))
}
