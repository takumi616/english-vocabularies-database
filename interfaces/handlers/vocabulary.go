package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/takumi616/english-vocabularies-database/entities"
	"github.com/takumi616/english-vocabularies-database/usecases/ports"
)

type VocabularyHandler struct {
	InputPorts ports.VocabularyInputPort
}

func (vh *VocabularyHandler) AddNewVocabulary(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Title   string `json:"title"`
		Example string `json:"example"`
	}

	//Convert json http request data into go struct
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Fatalf("failed to decode request body: %v", err)
	}

	vocabulary := entities.NewVocabulary(req.Title, req.Example)
	err := vh.InputPorts.AddNewVocabulary(r.Context(), w, vocabulary)

	log.Println(err)
}
