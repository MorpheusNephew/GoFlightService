package server

import (
	"fmt"
	"net/http"
)

// Initialize initializes the web server and sets up the /status endpoint
func Initialize() {
	http.HandleFunc("/status", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintln(w, "We're up and running!!! #ChargeOn")
	})

	port := ":3000"
	fmt.Println("Running on localhost port", port)
	http.ListenAndServe(port, nil)
}
