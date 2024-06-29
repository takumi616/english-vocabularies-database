package infrastructure

import (
	"net/http"

	"github.com/takumi616/english-vocabularies-database/interfaces/handlers"
)

func NewRouting() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /vocabularies", handlers.AddNewVocabulary)

	return mux
}
