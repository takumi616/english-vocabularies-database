package main

import (
	"github.com/takumi616/english-vocabularies-database/infrastructure"
	"github.com/takumi616/english-vocabularies-database/interfaces/gateways"
	"github.com/takumi616/english-vocabularies-database/interfaces/handlers"
	"github.com/takumi616/english-vocabularies-database/interfaces/presenters"
	"github.com/takumi616/english-vocabularies-database/usecases/interactors"
)

func main() {
	//ctx := context.Background()

	//Infrastructure layer
	config, _ := infrastructure.NewConfig()
	postgres := &infrastructure.Postgres{PgConfig: config.PgConfig}

	//Interfaces layer
	//Initialize gateway
	vocabularyGateway := &gateways.VocabularyGateway{Postgres: postgres}
	//Initialize presenter
	vocabularyPresenter := &presenters.VocabularyPresenter{}

	//Usecases layer
	//Initialize interactor with gateway(implements repository) and presenter(implements outputport)
	vocabularyInteractor := &interactors.VocabularyInteractor{Repo: vocabularyGateway, OutputPort: vocabularyPresenter}

	//Initialize handler with interactor(implements inputport)
	vocabularyHandler := &handlers.VocabularyHandler{InputPorts: vocabularyInteractor}

	//Infrastructure layer
	//Initialize routing
	routing := &infrastructure.Routing{VocabularyHandler: vocabularyHandler}
	mux := routing.Setup()
	//Initialize http server and run it
	httpServer := &infrastructure.HttpServer{Port: config.AppPort, ServeMux: mux}
	httpServer.Run()
}
