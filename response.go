package api

import (
	"fmt"
	"net/http"

	"github.com/cdvelop/model"
)

// action ej: create, read, update,delete, error
func (c config) jsonResponse(p *petition, code int, message string, body_out ...map[string]string) {

	p.w.Header().Set("Content-Action", "application/json; charset=utf-8")
	p.w.Header().Set("X-Content-Action-Options", "nosniff")
	p.w.Header().Set("Decode", p.decode)
	// ¡¡¡ antes de write header !!! de lo contrario no se guardan
	p.w.WriteHeader(code)

	// fmt.Println("p.decode:", p.decode)

	var out []byte
	var err error

	o := model.Object{Name: "error"}

	if p.o != nil {
		o = *p.o
	}

	var object_name = o.Name
	if p.multiple {
		out, err = c.EncodeResponses([]model.Response{o.Response(body_out, p.action, message)})
		if err != nil {
			out = []byte(`{"Action":"error", "Message":"` + err.Error() + `"}`)
		}
	} else {
		if message == "error" {

			object_name = ""
		}
		// fmt.Println("OBJECT:", object_name)
		// fmt.Println("ANTES DE ENCODE:", body_out)

		out, err = c.EncodeMaps(body_out, object_name)
		if err != nil {
			out = []byte(err.Error())
		}
	}

	// temp, _ := c.DecodeMaps(out, object_name)
	// fmt.Println("LO QUE SE ENVIÓ:", temp)
	//NOTIFY HERE

	p.w.Write(out)
}

func (c config) success(p *petition, action, message string, data ...map[string]string) {
	p.decode = p.o.Name
	c.jsonResponse(p, http.StatusOK, message, data...)
}

// no puedes: reason
func (c config) unauthorized(p *petition, reason string) {
	c.error(p, model.Error("no puedes", reason, "si no tiene una session en el sistema"), http.StatusNetworkAuthenticationRequired)
}

// status default: StatusBadRequest
func (c config) error(p *petition, err error, status ...int) {
	p.decode = "error"
	var code = http.StatusBadRequest
	for _, n := range status {
		code = n
	}

	c.logError(p, err, code)

	c.jsonResponse(p, code, "error", map[string]string{"error": err.Error()})
}

func errorHttp(p *petition, err error, code int) {
	p.w.WriteHeader(code)

	if p.multiple {
		fmt.Fprintln(p.w, `[{"o":["error","`+err.Error()+`"]}]`)
	} else {

	}
}

func (c config) logError(p *petition, err error, status ...int) {

	var code = http.StatusInternalServerError
	for _, s := range status {
		code = s
	}

	if p.u == nil {
		p.u = &model.User{Name: "unregistered"}
	}

	var auth_state string
	if p.e != nil {
		auth_state = "auth:" + p.e.Error()
	}

	c.Log(p.r.Method, p.r.RemoteAddr, "user:", p.u.Name, "id:", p.u.Id, err, "code:", code, auth_state)

}
