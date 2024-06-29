package handlers

import (
	"io"
	"log"
	"net/http"
)

func AddNewVocabulary(w http.ResponseWriter, r *http.Request) {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("failed to add new vocabulary: %v", err.Error())
	}

	w.Write(reqBody)
}
