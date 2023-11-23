package api

import "net/http"

func (c config) update(p *petition) {
	const this = "api update error "
	// fmt.Printf("Estás en la página de actualización del objeto %s\nData: %s\n", o.ObjectName, u.Name)
	data, err := c.decodeStringMapData(p)
	if err != "" {
		c.error(p, this+err, http.StatusInternalServerError)
		return
	}

	err = p.o.ValidateData(false, true, data...)
	if err != "" {
		c.error(p, this+err)
		return
	}

	// fmt.Println("OBJETO VALIDADO: ", o.ObjectName)

	err = p.o.BackHandler.Update(p.u, data...)
	if err != "" {
		c.error(p, this+err)
		return
	}

	// fmt.Println("DATA DESPUÉS DE ACTUALIZAR: ", recovered_data)

	c.success(p, "update", "ok", data...)
}
