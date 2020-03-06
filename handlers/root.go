package handlers

import (
	"fmt"
	"net/http"
	"os"
	"log"
)

// home is a simple HTTP handler function which writes a response.
func root(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
	host := os.Getenv("HOSTNAME")

	demoName := os.Getenv("NAME")
	if demoName == "" {
		demoName = "Demo v3.0"
	}
	fmt.Fprint(w, "Hello "+demoName+" ! Your request was processed by "+host+".")
}
