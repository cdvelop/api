package api

import (
	"fmt"
	"net/http"

	"github.com/cdvelop/model"
)

func (c config) read(o *model.Object, w http.ResponseWriter, r *http.Request) {

	fmt.Printf("Estás en la página de lectura de data de %s\n", o.Name)

	params, err := paramsCheckIn(r, false, false, o)
	if err != nil {
		c.error(w, err, o)
		return
	}

	data, err := o.Read(params)
	if err != nil {
		c.error(w, err, o)
		return
	}

	c.success(w, "read", "ok", o, data...)
}
