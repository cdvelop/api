package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/cdvelop/model"
)

func paramsCheckIn(its_new, its_update_or_delete, content_is_file bool, o *model.Object, w http.ResponseWriter, r *http.Request) (map[string]string, error) {

	params, err := getParams(o, content_is_file, w, r)
	if err != nil {
		return nil, err
	}

	// fmt.Println("PARÁMETROS RECIBIDOS: ", params)

	err = o.ValidateData(its_new, its_update_or_delete, params)
	if err != nil {
		return nil, err
	}

	return params, nil
}

// content_type = file
func getParams(o *model.Object, content_is_file bool, w http.ResponseWriter, r *http.Request) (map[string]string, error) {

	if content_is_file {

		r.Body = http.MaxBytesReader(w, r.Body, o.ConfigFile().MaximumFileSize) // 220 KB

		err := r.ParseMultipartForm(o.ConfigFile().MaximumFileSize) // Specify the maximum memory allowed for parsing (e.g., 10 MB)
		if err != nil {
			if strings.Contains(err.Error(), "multipart") {
				return nil, fmt.Errorf("CreateFile error ParseMultipartForm %v", err)
			} else {
				return nil, fmt.Errorf("error tamaño de archivo excedido máximo admitido: %v kb", o.ConfigFile().MaximumFileSize)
			}
		}

	} else {
		err := r.ParseForm()
		if err != nil {
			return nil, err
		}
	}

	params := make(map[string]string)
	for key, values := range r.PostForm {
		if len(values) > 1 {
			params[key] = strings.Join(values, ",")
		} else {
			params[key] = values[0]
		}
	}

	for key, values := range r.URL.Query() {
		if _, exists := params[key]; !exists {
			if len(values) > 1 {
				params[key] = strings.Join(values, ",")
			} else {
				params[key] = values[0]
			}
		}
	}

	return params, nil
}
