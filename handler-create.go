package api

import (
	"net/http"

	"github.com/cdvelop/cutkey"
	"github.com/cdvelop/model"
)

func (c config) create(u *model.User, o *model.Object, w http.ResponseWriter, r *http.Request) {

	// fmt.Printf("Estás en el Manejador de creación del objeto %s\n", o.Name)

	data, err := cutkey.Decode(r.Body, o)
	if err != nil {
		c.error(u, w, r, err, o)
		return
	}

	err = o.ValidateData(true, false, data...)
	if err != nil {
		c.error(u, w, r, err, o)
		return
	}

	err = o.Create(u, data...)
	if err != nil {
		c.error(u, w, r, err, o)
		return
	}

	c.success(w, "create", "ok", o, data...)
}
