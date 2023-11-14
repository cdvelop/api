package api

import (
	"net/http"
)

func (c config) create(p *petition) {

	data, err := c.decodeStringMapData(p)
	if err != nil {
		c.error(p, err, http.StatusInternalServerError)
		return
	}
	// fmt.Printf("\nEstás en creación objeto %s\n", p.o.Name)

	err = p.o.ValidateData(true, false, data...)
	if err != nil {
		c.error(p, err)
		return
	}

	err = p.o.Create(p.u, data...)
	if err != nil {
		c.error(p, err)
		return
	}

	c.success(p, "create", "ok", data...)
}
