package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/takumi616/english-vocabularies-database/usecases/inputPorts"
)

type VocabularyHandler struct {
	InputPorts inputPorts.VocabularyInputPort
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

	addedId, _ := vh.InputPorts.AddNewVocabulary(req.Title, req.Example)

	log.Println(addedId)
}
