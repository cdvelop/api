package api

import (
	"net/http"
	"strings"
)

func paramsCheckIn(p *petition, its_new, its_update_or_delete bool) (map[string]string, error) {

	params, err := getParams(p)
	if err != nil {
		return nil, err
	}

	// fmt.Println("PARÁMETROS RECIBIDOS: ", params)

	err = p.o.ValidateData(its_new, its_update_or_delete, params)
	if err != nil {
		return nil, err
	}

	return params, nil
}

// content_type = file
func getParams(p *petition) (map[string]string, error) {

	err := p.r.ParseForm()
	if err != nil {
		return nil, err
	}

	params := make(map[string]string)
	for key, values := range p.r.PostForm {
		if len(values) > 1 {
			params[key] = strings.Join(values, ",")
		} else {
			params[key] = values[0]
		}
	}

	gerUrlParams(p.r, params)

	return params, nil
}

func gerUrlParams(r *http.Request, params_out map[string]string) {

	for key, values := range r.URL.Query() {
		if _, exists := params_out[key]; !exists {
			if len(values) > 1 {
				params_out[key] = strings.Join(values, ",")
			} else {
				params_out[key] = values[0]
			}
		}
	}

}
