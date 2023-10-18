package api

import (
	"net/http"

	"github.com/cdvelop/model"
)

func (c config) read(u *model.User, o *model.Object, w http.ResponseWriter, r *http.Request) {

	// fmt.Printf("Estás en el Manejador de lectura de data de %s\n", o.Name)

	params, err := paramsCheckIn(false, false, false, o, w, r)
	if err != nil {
		c.error(u, w, r, err, o)
		return
	}

	data, err := o.Read(u, params)
	if err != nil {
		c.error(u, w, r, err, o)
		return
	}

	c.success(w, "read", "ok", o, data...)
}

func (c config) readFile(u *model.User, o *model.Object, w http.ResponseWriter, r *http.Request) {
	// retorna objeto estático ej imagen.jpg
	// fmt.Printf("Estás en la página de lectura del archivo %s\n", o.Name)

	params, err := paramsCheckIn(false, false, false, o, w, r)
	if err != nil {
		c.error(u, w, r, err, o)
		return
	}

	file_path, err := o.FilePath(u, params)
	if err != nil {
		c.error(u, w, r, err, o)
		return
	}

	http.ServeFile(w, r, file_path)
}
