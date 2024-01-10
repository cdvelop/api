package api

import (
	"log"
	"net/http"
	"os"

	out "github.com/cdvelop/output"
)

// Getenv("APP_DOMAIN") default: "localhost"
// Getenv("APP_PORT") default: "8080"
func (c config) StartServer() {

	APP_DOMAIN := os.Getenv("APP_DOMAIN")
	if APP_DOMAIN == "" {
		APP_DOMAIN = "localhost"
	}

	APP_PORT := os.Getenv("APP_PORT")
	if APP_PORT == "" {
		APP_PORT = "8080"
	}

	mux := c.ServeMuxAndRoutes()
	addr := APP_DOMAIN + ":" + APP_PORT
	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	out.PrintOK("Servidor escuchando en: " + addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
