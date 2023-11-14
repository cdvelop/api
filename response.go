package api

import (
	"fmt"
	"net/http"

	"github.com/cdvelop/model"
	out "github.com/cdvelop/output"
)

// action ej: create, read, update,delete, error
func (c config) jsonResponse(p *petition, code int, action, message string, body_out ...map[string]string) {

	p.w.Header().Set("Content-Action", "application/json; charset=utf-8")
	p.w.Header().Set("X-Content-Action-Options", "nosniff")
	p.w.WriteHeader(code)

	var out []byte
	var err error

	o := model.Object{Name: "error"}

	if p.o != nil {
		o = *p.o
	}

	if p.multiple {
		out, err = c.EncodeResponses([]model.Response{o.Response(body_out, action, message)})
		if err != nil {
			out = []byte(`{"Action":"error", "Message":"` + err.Error() + `"}`)
		}
	} else {

		// fmt.Println("ANTES DE ENCODE:", body_out)

		out, err = c.EncodeMaps(body_out, o.Name)
		if err != nil {
			out = []byte(err.Error())
		}
	}

	// temp, _ := c.DecodeMaps(out, o.Name)
	// fmt.Println("LO QUE SE ENVIÓ:", temp)
	//NOTIFY HERE

	p.w.Write(out)
}

func (c config) success(p *petition, action, message string, data ...map[string]string) {
	c.jsonResponse(p, http.StatusOK, action, message, data...)
}

// no puedes: reason
func (c config) unauthorized(p *petition, reason string) {
	c.error(p, model.Error("no puedes", reason, "si no tiene una session en el sistema"), http.StatusNetworkAuthenticationRequired)
}

// status default: StatusBadRequest
func (c config) error(p *petition, err error, status ...int) {

	var code = http.StatusBadRequest
	for _, n := range status {
		code = n
	}
	c.logError(p, err, code)

	c.jsonResponse(p, code, "error", err.Error())
}

func errorHttp(w http.ResponseWriter, err error, code int) {
	w.WriteHeader(code)
	fmt.Fprintln(w, `[{"o":["error","`+err.Error()+`"]}]`)
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

	message_log := fmt.Sprint(p.r.Method, p.r.RemoteAddr, "user:", p.u.Name, "id:", p.u.Id, err, "code:", code, auth_state)

	if c.Logger != nil {
		c.Log(message_log)
	} else {
		out.PrintWarning(message_log)
	}

}
