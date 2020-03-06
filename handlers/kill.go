package handlers

import (
	"net/http"
	"os"
	"log"
)

// home is a simple HTTP handler function which writes a response.
func kill(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
	log.Fatal("You used the "+ r.URL.String() +" url that kills the server")
	os.Exit(3)
}