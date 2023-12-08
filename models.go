package api

import (
	"net/http"
	"time"

	"github.com/cdvelop/model"
)

type config struct {
	model.AuthBackendAdapter
	model.ObjectsHandler
	model.ModuleHandler
	model.DataConverter
	model.FileApi
	model.Logger

	production_mode bool
	static_cache    string
}

type petition struct {
	action          string // action ej: create,delete,crud,update,upload
	object_response string // nombre del objeto respuesta
	u               *model.User
	o               *model.Object
	r               *http.Request
	w               http.ResponseWriter
	t               time.Time
	err             string
	multiple        bool
}
