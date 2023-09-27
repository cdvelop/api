package api

import (
	"net/http"

	"github.com/cdvelop/cutkey"
	"github.com/cdvelop/model"
)

func (c config) update(u *model.User, o *model.Object, w http.ResponseWriter, r *http.Request) {

	data, err := cutkey.Decode(r.Body, o)
	if err != nil {
		c.error(w, r, err, o)
		return
	}

	// fmt.Printf("Estás en la página de actualización del objeto %s\nData: %s\n", o.Name, data)

	err = o.ValidateData(false, true, data...)
	if err != nil {
		c.error(w, r, err, o)
		return
	}

	// fmt.Println("OBJETO VALIDADO: ", o.Name)

	recovered_data, err := o.Update(u, data...)
	if err != nil {
		c.error(w, r, err, o)
		return
	}

	// fmt.Println("DATA DESPUÉS DE ACTUALIZAR: ", recovered_data)

	c.success(w, "update", "ok", o, recovered_data...)
}
