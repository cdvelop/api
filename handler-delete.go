package api

import (
	"github.com/cdvelop/model"
)

func (c config) delete(p *petition) {

	// fmt.Printf("Estás en la página de eliminación de %s\n", o.Name)

	// Leer el cuerpo de la solicitud en un slice de bytes
	data, err := c.decodeStringMapData(p)
	if err != nil {
		c.error(p, model.Error("delete", err))
		return
	}

	err = p.o.ValidateData(false, true, data...)
	if err != nil {
		c.error(p, err)
		return
	}

	// fmt.Println("data recibida para eliminar:", data)

	recovered_data, err := p.o.Delete(p.u, data...)
	if err != nil {
		c.error(p, err)
		return
	}

	c.success(p, "delete", "ok", recovered_data...)
}
