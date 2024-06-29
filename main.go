package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/takumi616/english-vocabularies-database/infrastructure"
)

func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "test handler")
}

func main() {

	port := os.Getenv("APP_CONTAINER_PORT")

	mux := infrastructure.NewRouting()

	httpServer := infrastructure.NewHttpServer(port, mux)
	httpServer.Run()
}
