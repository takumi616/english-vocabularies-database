package main

import (
	"context"

	"github.com/takumi616/english-vocabularies-database/infrastructure"
	"github.com/takumi616/english-vocabularies-database/interfaces/gateways"
	"github.com/takumi616/english-vocabularies-database/interfaces/handlers"
	"github.com/takumi616/english-vocabularies-database/interfaces/presenters"
	"github.com/takumi616/english-vocabularies-database/usecases/interactors"
)

func main() {
	ctx := context.Background()

	config, _ := infrastructure.NewConfig()

	db, _ := infrastructure.InitDatabase(ctx, config)

	//Initialize gateway of interfaces layer
	vocabularyGateway := gateways.NewVocabularyGateway(db)
	//Initialize presenter of interfaces layer
	vocabularyPresenter := presenters.NewVocabularyPresenter()

	//Initialize interactor of usecases layer
	vocabularyInteractor := interactors.NewVocabularyInteractor(vocabularyGateway, vocabularyPresenter)

	//Initialize handler of interfaces layer
	vocabularyHandler := handlers.NewVocabularyHandler(vocabularyInteractor)

	routing := infrastructure.NewRouting(vocabularyHandler)
	mux := routing.Setup()

	httpServer := infrastructure.NewHttpServer(config.AppPort, mux)
	httpServer.Run()
}
