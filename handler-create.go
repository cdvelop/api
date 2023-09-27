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
		c.error(w, r, err, o)
		return
	}

	err = o.ValidateData(true, false, data...)
	if err != nil {
		c.error(w, r, err, o)
		return
	}

	err = o.Create(u, data...)
	if err != nil {
		c.error(w, r, err, o)
		return
	}

	c.success(w, "create", "ok", o, data...)
}

func (c config) createFile(u *model.User, o *model.Object, w http.ResponseWriter, r *http.Request) {
	// retorna objeto estático ej imagen.jpg
	// fmt.Printf("Estás en el Manejador de subida de archivos %s\n", o.Name)

	params, err := paramsCheckIn(true, false, true, o, w, r)
	if err != nil {
		c.error(w, r, err, o)
		return
	}

	// fmt.Println("OBJETO: ", o)

	data, err := o.CreateFile(u, r, params)
	if err != nil {
		c.error(w, r, err, o)
		return
	}

	c.success(w, "create", "ok", o, data...)
}
