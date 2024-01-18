package api

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"
	"time"

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
		APP_PORT = "4433"
	}

	mux := c.ServeMuxAndRoutes()
	addr := APP_DOMAIN + ":" + APP_PORT

	server := &http.Server{
		Addr:    addr,
		Handler: mux, //conexión

		// ReadTimeout es la duración máxima para leer la solicitud completa, incluido el cuerpo.
		ReadTimeout: 0, // 1 * time.Minute, 0 no hay tiempo de espera
		// WriteTimeout es la duración máxima antes de que se agoten las escrituras de la respuesta. Se restablece cada vez que se lee el encabezado de una nueva solicitud. Al igual que ReadTimeout, no permite que los controladores tomen decisiones por solicitud.
		WriteTimeout:   30 * time.Minute,
		MaxHeaderBytes: 1 << 20,
	}

	out.PrintOK("Servidor escuchando en: " + addr)

	if c.sslHandler != nil {

		server.TLSConfig = &tls.Config{
			Rand:           nil,
			GetCertificate: c.sslHandler.GetCertificate,
		}

		if err := server.ListenAndServeTLS("", ""); err != nil {
			log.Println(err)
		}

	} else {

		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}

	}

}
