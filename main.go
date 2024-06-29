package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "test handler")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/test", testHandler)

	server := &http.Server{
		Addr:    ":" + os.Getenv("APP_CONTAINER_PORT"),
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}
