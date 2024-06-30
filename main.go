package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/takumi616/english-vocabularies-database/infrastructure"
	"github.com/takumi616/english-vocabularies-database/infrastructure/postgres"
	"github.com/takumi616/english-vocabularies-database/interfaces/gateways"
	"github.com/takumi616/english-vocabularies-database/interfaces/handlers"
	"github.com/takumi616/english-vocabularies-database/usecases/interactors"
)

func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "test handler")
}

func main() {

	port := os.Getenv("APP_CONTAINER_PORT")

	postgres := &postgres.Postgres{DbHandle: "dummy db"}

	gateway := &gateways.VocabularyGateway{Persistence: postgres}

	vocabularyInteractor := &interactors.VocabularyInteractor{Repo: gateway}

	vocabularyHandler := &handlers.VocabularyHandler{InputPorts: vocabularyInteractor}

	routing := infrastructure.NewRouting(vocabularyHandler)
	mux := routing.Setup()

	httpServer := infrastructure.NewHttpServer(port, mux)
	httpServer.Run()
}
