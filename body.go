package api

import (
	"bytes"
	"io"
	"net/http"
)

func (c config) readBody(p *petition) (out []byte, err string) {
	// Cerrar el cuerpo de la solicitud original al final de la función
	defer p.r.Body.Close()

	body, e := io.ReadAll(p.r.Body)
	if e != nil {
		c.error(p, "error al leer el cuerpo de la solicitud", http.StatusInternalServerError)
		return nil, err
	}

	p.r.Body = io.NopCloser(bytes.NewBuffer(body))

	return body, ""
}

func (c config) decodeStringMapData(p *petition) (data []map[string]string, err string) {

	body, err := c.readBody(p)
	if err != "" {
		return nil, "decodeStringMapData " + err
	}

	return c.DecodeMaps(body, p.o.ObjectName)

}
