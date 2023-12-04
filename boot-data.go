package api

import (
	"github.com/cdvelop/model"
)

func (c config) BackendLoadBootData(u *model.User) (out, err string) {
	const this = "BackendLoadBootData error :"

	if u == nil {
		return "", this + "no se puede entregar información si no te has registrado"
	}

	var responses []model.Response
	for _, o := range c.GetObjects() {

		// fmt.Println("BackHandler.BootResponse", o.ObjectName)
		// fmt.Println("Estado Back:", o.BackHandler.BootResponse)

		if o.BackHandler.BootResponse != nil {
			resp, err := o.BackHandler.AddBootResponse(u)
			if err != "" {
				return "", this + o.ObjectName + " " + err
			} else if len(resp) != 0 {
				responses = append(responses, resp...)
			}
		}
	}

	boot_data_byte, err := c.EncodeResponses(responses...)
	if err != "" {
		return "", this + err
	}

	return string(boot_data_byte), ""
}
