package api

import (
	"fmt"
	"log"
	"net/http"
)

func (c config) StartServer() {
	mux := c.SetupMuxRoutes()
	addr := "localhost:8080"
	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	// go func() {
	fmt.Printf("Servidor escuchando en http://%s/\n", addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
	// }()

}
