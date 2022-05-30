package app

import (
	"fmt"
	"net/http"
)

func Server(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
