package api

import (
	"github.com/cdvelop/model"
)

func (c config) update(p *petition) {

	// fmt.Printf("Estás en la página de actualización del objeto %s\nData: %s\n", o.Name, u.Name)
	data, err := c.decodeStringMapData(p)
	if err != nil {
		c.error(p, model.Error("update", err))
		return
	}

	err = p.o.ValidateData(false, true, data...)
	if err != nil {
		c.error(p, err)
		return
	}

	// fmt.Println("OBJETO VALIDADO: ", o.Name)

	err = p.o.Update(p.u, data...)
	if err != nil {
		c.error(p, err)
		return
	}

	// fmt.Println("DATA DESPUÉS DE ACTUALIZAR: ", recovered_data)

	c.success(p, "update", "ok", data...)
}
