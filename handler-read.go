package api

import (
	"fmt"
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

func (c config) readFile(u *model.User, w http.ResponseWriter, r *http.Request) {
	// retorna objeto estático ej imagen.jpg
	params := make(map[string]string)

	gerUrlParams(r, params)

	fmt.Printf("Estás en la página de lectura archivo %s\n", params)

	file_path, file_area, err := c.fileApi.GetFilePathByID(params)
	if err != nil {
		errorHttp(w, err, http.StatusBadRequest)
		return
	}

	// fmt.Println("AREA ARCHIVO", file_area)
	// fmt.Println("AREA USUARIO", u.Area)

	if file_area != u.Area {
		errorHttp(w, model.Error("no autorizado para leer archivo"), http.StatusUnauthorized)
		return
	}

	http.ServeFile(w, r, file_path)
}
