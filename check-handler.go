package api

import (
	"fmt"

	"github.com/cdvelop/model"
)

func (c config) isHandlerOk(action_type, api_name string) (*model.Object, error) {

	switch action_type {

	case "create":

		for _, h := range c.createHandlers {
			if h.Name == api_name {
				return h, nil
			}
		}

	case "read":
		for _, h := range c.readHandlers {
			if h.Name == api_name {
				return h, nil
			}
		}

	case "update":
		for _, h := range c.updateHandlers {
			if h.Name == api_name {
				return h, nil
			}
		}

	case "delete":
		for _, h := range c.deleteHandlers {
			if h.Name == api_name {
				return h, nil
			}
		}

	case "file":
		for _, h := range c.fileHandlers {
			if h.Name == api_name {
				return h, nil
			}
		}
	}

	// for _, f := range c.fileHandlers {
	// 	fmt.Println("FILE HANDLERS: ", f.Name)
	// }

	return nil, fmt.Errorf("no existe el controlador: %v para la acción: %v", api_name, action_type)
}
