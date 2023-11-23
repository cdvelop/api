package api

import "net/http"

func (c config) delete(p *petition) {
	const this = "api delete error "
	// fmt.Printf("Estás en la página de eliminación de %s\n", o.ObjectName)

	// Leer el cuerpo de la solicitud en un slice de bytes
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

	// fmt.Println("data recibida para eliminar:", data)

	err = p.o.BackHandler.Delete(p.u, data...)
	if err != "" {
		c.error(p, this+err)
		return
	}

	c.success(p, "delete", "ok", data...)
}
