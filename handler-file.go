package api

func (c config) upload(p *petition) {
	const this = "api upload error "
	// fmt.Printf("Estás en el Manejador de subida de archivos %s\n", p.o.ObjectName)

	data_out, err := c.FileUpload(p.o.ObjectName, p.u.Area, p.r, p.w)
	if err != "" {
		c.error(p, this+err)
		return
	}
	// fmt.Println("api upload", p.o.ObjectName)

	// fmt.Println(" success SALIDA DATA UPLOAD:", data_out)

	c.success(p, "upload", "ok", data_out...)
}
