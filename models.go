package api

import (
	"net/http"
	"time"

	"github.com/cdvelop/model"
)

type config struct {
	model.LoginUser
	model.ObjectsHandler
	model.DataConverter
	model.FileApi
	model.Logger

	bootHandlers []*model.Object

	developer_mode bool
	static_cache   string
}

type petition struct {
	action   string // action ej: create,delete,crud,update,upload
	u        *model.User
	o        *model.Object
	r        *http.Request
	w        http.ResponseWriter
	t        time.Time
	e        error
	multiple bool
	decode   string
}
