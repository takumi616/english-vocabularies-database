package infrastructure

import (
	"log"
	"net/http"
)

type HttpServer struct {
	Port     string
	ServeMux *http.ServeMux
}

func (hs *HttpServer) Run() {
	server := &http.Server{
		Addr:    ":" + hs.Port,
		Handler: hs.ServeMux,
	}

	log.Fatal(server.ListenAndServe())
}
