package handlers

import (
	"io"
	"log"
	"net/http"

	"github.com/takumi616/english-vocabularies-database/usecases/inputPorts"
)

type VocabularyHandler struct {
	InputPorts inputPorts.VocabularyInputPort
}

func (vh *VocabularyHandler) AddNewVocabulary(w http.ResponseWriter, r *http.Request) {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("failed to add new vocabulary: %v", err.Error())
	}

	reqTitle := "title"
	reqExample := "example"

	addedId, _ := vh.InputPorts.AddNewVocabulary(reqTitle, reqExample)

	w.Write(reqBody)
	log.Println(addedId)
}
