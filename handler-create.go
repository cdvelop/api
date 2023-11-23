package api

import (
	"net/http"
)

func (c config) create(p *petition) {
	const this = "api create error "
	data, err := c.decodeStringMapData(p)
	if err != "" {
		c.error(p, this+err, http.StatusInternalServerError)
		return
	}
	// fmt.Printf("\nEstás en creación objeto %s\n", p.o.ObjectName)

	err = p.o.ValidateData(true, false, data...)
	if err != "" {
		c.error(p, this+err)
		return
	}

	err = p.o.BackHandler.Create(p.u, data...)
	if err != "" {
		c.error(p, this+err)
		return
	}

	c.success(p, "create", "ok", data...)
}
