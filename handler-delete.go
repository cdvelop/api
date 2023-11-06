package api

import (
	"net/http"

	"github.com/cdvelop/cutkey"
	"github.com/cdvelop/model"
)

func (c config) delete(u *model.User, o *model.Object, w http.ResponseWriter, r *http.Request) {

	// fmt.Printf("Estás en la página de eliminación de %s\n", o.Name)

	data, err := cutkey.Decode(r.Body, o)
	if err != nil {
		c.error(u, w, r, err, o)
		return
	}

	err = o.ValidateData(false, true, data...)
	if err != nil {
		c.error(u, w, r, err, o)
		return
	}

	// fmt.Println("data recibida para eliminar:", data)

	recovered_data, err := o.Delete(u, data...)
	if err != nil {
		c.error(u, w, r, err, o)
		return
	}

	c.success(w, "delete", "ok", o, recovered_data...)
}
