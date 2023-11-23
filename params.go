package api

import (
	"net/http"
	"strings"
)

func paramsCheckIn(p *petition, its_new, its_update_or_delete bool) (params map[string]string, err string) {
	const this = "paramsCheckIn error "
	params, err = getParams(p)
	if err != "" {
		return nil, this + err
	}

	// fmt.Println("PARÁMETROS RECIBIDOS: ", params)

	err = p.o.ValidateData(its_new, its_update_or_delete, params)
	if err != "" {
		return nil, this + err
	}

	return params, ""
}

// content_type = file
func getParams(p *petition) (params map[string]string, err string) {
	const this = "getParams error "
	e := p.r.ParseForm()
	if e != nil {
		return nil, this + e.Error()
	}
	params = make(map[string]string)

	for key, values := range p.r.PostForm {
		if len(values) > 1 {
			params[key] = strings.Join(values, ",")
		} else {
			params[key] = values[0]
		}
	}

	gerUrlParams(p.r, params)

	return params, ""
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
