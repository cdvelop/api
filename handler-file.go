package api

func (c config) upload(p *petition) {
	// fmt.Printf("Estás en el Manejador de subida de archivos %s\n", o.Name)

	data_out, err := c.FileUpload(p.o.Name, p.u.Area, p.r, p.w)
	if err != nil {
		c.error(p, err)
		return
	}

	c.success(p, "create", "ok", data_out...)
}
