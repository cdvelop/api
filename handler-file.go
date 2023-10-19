package api

import (
	"net/http"

	"github.com/cdvelop/fileserver"
	"github.com/cdvelop/model"
)

func (c config) createFile(u *model.User, o *model.Object, w http.ResponseWriter, r *http.Request) {
	// fmt.Printf("Estás en el Manejador de subida de archivos %s\n", o.Name)

	form_data, err := paramsCheckIn(true, false, true, o, w, r)
	if err != nil {
		c.error(u, w, r, err, o)
		return
	}

	data_out, err := fileserver.CreateFileInServer(r, o, form_data)
	if err != nil {
		c.error(u, w, r, err, o)
		return
	}

	c.success(w, "create", "ok", o, data_out...)
}
