package api

import (
	"github.com/cdvelop/cutkey"
	"github.com/cdvelop/model"
)

type config struct {
	*cutkey.Cut

	createHandlers []*model.Object
	readHandlers   []*model.Object
	updateHandlers []*model.Object
	deleteHandlers []*model.Object

	pathFileHandlers []*model.Object

	developer_mode bool
	static_cache   string
}
