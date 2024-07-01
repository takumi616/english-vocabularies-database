package infrastructure

import (
	"net/http"

	"github.com/takumi616/english-vocabularies-database/interfaces/handlers"
)

type Routing struct {
	VocabularyHandler *handlers.VocabularyHandler
}

func (r *Routing) Setup() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /vocabularies", r.VocabularyHandler.AddNewVocabulary)

	return mux
}
