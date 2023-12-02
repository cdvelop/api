package api

import (
	"net/http"

	"github.com/cdvelop/model"
)

// action ej: create, read, update,delete, error
func (c config) jsonResponse(p *petition, code int, message string, body_out ...map[string]string) {

	p.w.Header().Set("Content-Action", "application/json; charset=utf-8")
	p.w.Header().Set("X-Content-Action-Options", "nosniff")
	// p.w.Header().Set("Decode", p.decode)
	// ¡¡¡ antes de write header !!! de lo contrario no se guardan
	p.w.WriteHeader(code)

	// fmt.Println("p.decode:", p.decode)

	var out []byte
	var err string

	o := model.Object{ObjectName: "error"}

	if p.o != nil {
		o = *p.o
	}

	var object_name = o.ObjectName
	if p.multiple {
		out, err = c.EncodeResponses(o.Response(body_out, p.action, message))
		if err != "" {
			out = []byte(`{"Action":"error", "Message":"` + err + `"}`)
		}
	} else {
		if message == "error" {
			object_name = ""
		}
		// fmt.Println("OBJECT:", object_name)
		// fmt.Println("ANTES DE ENCODE:", body_out)

		out, err = c.EncodeMaps(body_out, object_name)
		if err != "" {
			out = []byte(err)
		}
	}

	// temp, _ := c.DecodeMaps(out, object_name)
	// fmt.Println("LO QUE SE ENVIÓ:", temp)
	//NOTIFY HERE

	p.w.Write(out)
}

func (c config) success(p *petition, action, message string, data ...map[string]string) {
	c.jsonResponse(p, http.StatusOK, message, data...)
}

// no puedes: reason
func (c config) unauthorized(p *petition, reason string) {
	c.error(p, "no puedes "+reason+" si no tiene una session en el sistema", http.StatusNetworkAuthenticationRequired)
}

// status default: StatusBadRequest
func (c config) error(p *petition, err string, status ...int) {
	// p.decode = "error"
	var code = http.StatusBadRequest
	for _, n := range status {
		code = n
	}
	p.w.Header().Set("Status", err)
	// p.w.Header().Set("Status-Text", err)

	// fmt.Println("CODIGO RESPUESTA:", code)

	p.w.WriteHeader(code)

	c.logError(p, err, code)

	// c.jsonResponse(p, code, "error", map[string]string{"error": err.Error()})
	// fmt.Fprintln(p.w, `[{"o":["error","`+err.Error()+`"]}]`)
	p.w.Write([]byte(err))
}

func (c config) logError(p *petition, err string, code int) {

	var auth_state string
	if p.u == nil {
		p.u = &model.User{Name: "unregistered"}
		auth_state = "auth:" + p.err
	}

	c.Log(p.r.Method, p.r.RemoteAddr, "user:", p.u.Name, "id:", p.u.Id, err, "code:", code, auth_state)

}
