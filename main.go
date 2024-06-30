package main

import (
	"context"

	"github.com/takumi616/english-vocabularies-database/infrastructure"
	"github.com/takumi616/english-vocabularies-database/infrastructure/postgres"
	"github.com/takumi616/english-vocabularies-database/interfaces/gateways"
	"github.com/takumi616/english-vocabularies-database/interfaces/handlers"
	"github.com/takumi616/english-vocabularies-database/usecases/interactors"
)

func main() {
	ctx := context.Background()

	config, _ := infrastructure.NewConfig()

	//postgres := &postgres.Postgres{DbHandle: "dummy db"}
	postgres, _ := postgres.Init(ctx, config)

	gateway := &gateways.VocabularyGateway{Persistence: postgres}

	vocabularyInteractor := &interactors.VocabularyInteractor{Repo: gateway}

	vocabularyHandler := &handlers.VocabularyHandler{InputPorts: vocabularyInteractor}

	routing := infrastructure.NewRouting(vocabularyHandler)
	mux := routing.Setup()

	httpServer := infrastructure.NewHttpServer(config.AppPort, mux)
	httpServer.Run()
}
