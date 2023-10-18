package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cdvelop/model"
)

// action ej: create, read, update,delete, error
func (c config) jsonResponse(w http.ResponseWriter, code int, action, message string, obj_in *model.Object, data_out ...map[string]string) {

	w.Header().Set("Content-Action", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Action-Options", "nosniff")
	w.WriteHeader(code)

	o := model.Object{Name: "error"}

	if obj_in != nil {
		o = *obj_in
	}

	jsonBytes, err := c.EncodeResponses([]model.Response{o.Response(data_out, action, message)})
	if err != nil {
		fmt.Fprintln(w, `{"Action":"error", "Message":"`+err.Error()+`"}`)
		return
	}

	//NOTIFY HERE

	w.Write(jsonBytes)
}

func (c config) success(w http.ResponseWriter, action, message string, o *model.Object, data ...map[string]string) {
	c.jsonResponse(w, http.StatusOK, action, message, o, data...)
}

func (c config) error(u *model.User, w http.ResponseWriter, r *http.Request, err error, o *model.Object) {

	logError(u, r, err)

	c.jsonResponse(w, http.StatusBadRequest, "error", err.Error(), o)
}

func logError(u *model.User, r *http.Request, err error) {

	if u == nil {
		u = &model.User{Name: "unregistered"}
	}

	log.Printf("%v %v user:%v id:%v %v", r.Method, r.RemoteAddr, u.Name, u.Id, err)
}
