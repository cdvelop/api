package api

import (
	"fmt"
	"net/http"

	"github.com/cdvelop/model"
)

func (c config) readFile(o *model.Object, w http.ResponseWriter, r *http.Request) {
	// retorna objeto estático ej imagen.jpg
	fmt.Printf("Estás en la página de lectura del archivo %s\n", o.Name)

	params, err := paramsCheckIn(r, false, false, o)
	if err != nil {
		c.error(w, err, o)
		return
	}

	file_path, err := o.FilePath(params)
	if err != nil {
		c.error(w, err, o)
		return
	}

	http.ServeFile(w, r, file_path)
}
