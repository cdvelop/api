package api

import (
	"fmt"
	"log"
	"net/http"

	out "github.com/cdvelop/output"
)

func (c config) StartServer() {
	mux := c.ServeMuxAndRoutes()
	addr := "localhost:8080"
	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	out.PrintOK(fmt.Sprintf("Servidor escuchando en http://%v/", addr))
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
