package api

import (
	"fmt"
	"net/http"

	"github.com/cdvelop/model"
)

// action ej: create, read, update,delete, error
func (c config) jsonResponse(w http.ResponseWriter, code int, action, message string, o *model.Object, data_out ...map[string]string) {

	w.Header().Set("Content-Action", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Action-Options", "nosniff")
	w.WriteHeader(code)

	// fmt.Println("HERE 2", message, "Objeto: ", o)
	var object_name = "error"
	var module_name = "api.error"

	if o != nil {
		object_name = o.Name
		module_name = o.ModuleName()
	}

	r := model.Response{
		Action:  action,
		Data:    data_out,
		Object:  object_name,
		Module:  module_name,
		Message: message,
	}

	jsonBytes, err := c.EncodeResponses([]model.Response{r})
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

func (c config) error(w http.ResponseWriter, err error, o *model.Object) {
	c.jsonResponse(w, http.StatusBadRequest, "error", err.Error(), o)
}
