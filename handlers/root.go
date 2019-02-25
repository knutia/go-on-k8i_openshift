package handlers

import (
	"fmt"
	"net/http"
	"os"
)

// home is a simple HTTP handler function which writes a response.
func root(w http.ResponseWriter, _ *http.Request) {
	host := os.Getenv("HOSTNAME")

	demoName := os.Getenv("NAME")
	if demoName == "" {
		demoName = "Demo"
	}
	fmt.Fprint(w, "Hello "+demoName+" ! Your request was processed by "+host+".")
}
