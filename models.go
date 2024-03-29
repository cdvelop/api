package api

import (
	"crypto/tls"
	"net/http"
	"time"

	"github.com/cdvelop/model"
)

type config struct {
	model.SessionBackendAdapter
	model.ObjectsHandlerAdapter
	model.DataConverter
	model.FileApi
	model.Logger
	model.BackendBootDataUser

	production_mode bool
	static_cache    string

	sslHandler
}

type petition struct {
	action   string // action ej: create,delete,crud,update,upload
	u        *model.User
	o        *model.Object
	r        *http.Request
	w        http.ResponseWriter
	t        time.Time
	err      string
	multiple bool
}

type sslHandler interface {
	GetCertificate(hello *tls.ClientHelloInfo) (*tls.Certificate, error)
}
