package api

import (
	"fmt"
	"net/http"

	"github.com/cdvelop/cutkey"
	"github.com/cdvelop/model"
)

func (c config) delete(o *model.Object, w http.ResponseWriter, r *http.Request) {

	fmt.Printf("Estás en la página de eliminación de %s\n", o.Name)

	data, err := cutkey.Decode(r.Body, o)
	if err != nil {
		c.error(w, err, o)
		return
	}

	err = o.Delete(&data)
	if err != nil {
		c.error(w, err, o)
		return
	}

	c.success(w, "delete", "ok", o)
}
