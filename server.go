package api

import (
	"fmt"
	"log"
	"net/http"

	. "github.com/cdvelop/gotools"
)

func (c config) StartServer() {
	mux := c.ServeMuxAndRoutes()
	addr := "localhost:8080"
	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	Print("ok", fmt.Sprintf("Servidor escuchando en http://%s/", addr))
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
