package handlers

import (
	"fmt"
	"net/http"
	"os"
)

// home is a simple HTTP handler function which writes a response.
func root(w http.ResponseWriter, _ *http.Request) {
	host := os.Getenv("HOSTNAME")
	fmt.Fprint(w, "Hello! Your request was processed by "+host+".")
}
