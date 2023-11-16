package api

import "fmt"

func (c config) upload(p *petition) {
	// fmt.Printf("Estás en el Manejador de subida de archivos %s\n", p.o.Name)

	data_out, err := c.FileUpload(p.o.Name, p.u.Area, p.r, p.w)
	if err != nil {

		c.error(p, err)
		return
	}
	fmt.Println(p.o.Name)

	// fmt.Println(" success SALIDA DATA UPLOAD:", data_out)

	c.success(p, "upload", "ok", data_out...)
}
