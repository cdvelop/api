package api

import (
	"fmt"
	"net/http"

	"github.com/cdvelop/cutkey"
	"github.com/cdvelop/model"
)

func (c config) create(o *model.Object, w http.ResponseWriter, r *http.Request) {

	fmt.Printf("Estás en la página de creación del objeto %s\n", o.Name)

	data, err := cutkey.Decode(r.Body, o)
	if err != nil {
		c.error(w, err, o)
		return
	}

	out, err := o.Create(&data)
	if err != nil {
		c.error(w, err, o)
		return
	}

	c.success(w, "create", "ok", o, *out...)
}
