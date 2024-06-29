package infrastructure

import (
	"log"
	"net/http"
)

type HttpServer struct {
	Port     string
	ServeMux *http.ServeMux
}

func NewHttpServer(port string, mux *http.ServeMux) *HttpServer {
	return &HttpServer{
		Port:     port,
		ServeMux: mux,
	}
}

func (hs *HttpServer) Run() {
	server := &http.Server{
		Addr:    ":" + hs.Port,
		Handler: hs.ServeMux,
	}

	log.Fatal(server.ListenAndServe())
}
