package api

import (
	"fmt"
	"net/http"

	"github.com/cdvelop/cutkey"
	"github.com/cdvelop/model"
)

func (c config) update(o *model.Object, w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Estás en la página de actualización del objeto %s\n", o.Name)

	data, err := cutkey.Decode(r.Body, o)
	if err != nil {
		c.error(w, err, o)
		return
	}

	err = o.ValidateData(false, true, data...)
	if err != nil {
		c.error(w, err, o)
		return
	}

	recovered_data, err := o.Update(data...)
	if err != nil {
		c.error(w, err, o)
		return
	}

	c.success(w, "update", "ok", o, recovered_data...)
}
