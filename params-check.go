package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/cdvelop/model"
)

func getParams(r *http.Request) (*map[string]string, error) {
	err := r.ParseForm()
	if err != nil {
		return nil, err
	}

	params := make(map[string]string)
	for key := range r.Form {
		params[key] = r.Form.Get(key)
	}

	for key, values := range r.URL.Query() {
		if len(values) > 1 {
			params[key] = strings.Join(values, ",")
		} else {
			params[key] = values[0]
		}
	}

	return &params, nil
}

func paramsCheckIn(r *http.Request, its_new, its_update_or_delete bool, o *model.Object) (*map[string]string, error) {

	params, err := getParams(r)
	if err != nil {
		return nil, err
	}

	fmt.Println("PARÁMETROS RECIBIDOS: ", params)

	err = o.ValidateData(its_new, its_update_or_delete, params)
	if err != nil {
		return nil, err
	}

	return params, nil
}
