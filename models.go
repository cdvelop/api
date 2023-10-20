package api

import (
	"net/http"

	"github.com/cdvelop/cutkey"
	"github.com/cdvelop/model"
)

type config struct {
	*cutkey.Cut
	bootHandlers []*model.Object

	createHandlers []*model.Object
	readHandlers   []*model.Object
	updateHandlers []*model.Object
	deleteHandlers []*model.Object

	fileHandlers []*model.Object
	fileApi      model.FileApi

	developer_mode bool
	static_cache   string

	auth authAdapter
}

type authAdapter interface {
	GetUser(r *http.Request) *model.User
}
