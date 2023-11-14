package api

import (
	"bytes"
	"io"
	"net/http"

	"github.com/cdvelop/model"
)

func (c config) readBody(p *petition) ([]byte, error) {
	// Cerrar el cuerpo de la solicitud original al final de la función
	defer p.r.Body.Close()

	body, err := io.ReadAll(p.r.Body)
	if err != nil {
		c.error(p, model.Error("error al leer el cuerpo de la solicitud"), http.StatusInternalServerError)
		return nil, err
	}

	p.r.Body = io.NopCloser(bytes.NewBuffer(body))

	return body, nil
}

func (c config) decodeStringMapData(p *petition) ([]map[string]string, error) {

	body, err := c.readBody(p)
	if err != nil {
		return nil, err
	}

	data, err := c.DecodeMaps(body, p.o.Name)
	if err != nil {
		return nil, err
	}

	return data, nil

}
