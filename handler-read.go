package api

import (
	"fmt"
	"net/http"

	"github.com/cdvelop/model"
)

func (c config) readone(o *model.Object, w http.ResponseWriter, r *http.Request) {

	fmt.Printf("Estás en la página de lectura de la data de %s\n", o.Name)

	params, err := paramsCheckIn(r, false, false, o)
	if err != nil {
		c.error(w, err, o)
		return
	}

	data, err := o.ReadOne(params)
	if err != nil {
		c.error(w, err, o)
		return
	}

	c.success(w, "readone", "ok", o, *data)

}

func (c config) readall(o *model.Object, w http.ResponseWriter, r *http.Request) {

	fmt.Printf("Estás en la página de lectura de toda la data de %s\n", o.Name)

	params, err := paramsCheckIn(r, false, false, o)
	if err != nil {
		c.error(w, err, o)
		return
	}

	data, err := o.ReadAll(params)
	if err != nil {
		c.error(w, err, o)
		return
	}

	c.success(w, "readall", "ok", o, *data...)
}
