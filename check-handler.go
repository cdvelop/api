package api

import (
	"fmt"

	"github.com/cdvelop/model"
)

func (c config) isHandlerOk(action_type, handler_name string) (*model.Object, error) {

	switch action_type {

	case "create":
		for _, h := range c.createHandlers {
			if h.Api() == handler_name {
				return h, nil
			}
		}

	case "read":
		for _, h := range c.readHandlers {
			if h.Api() == handler_name {
				return h, nil
			}
		}

	case "update":
		for _, h := range c.updateHandlers {
			if h.Api() == handler_name {
				return h, nil
			}
		}

	case "delete":
		for _, h := range c.deleteHandlers {
			if h.Api() == handler_name {
				return h, nil
			}
		}

	case "file":
		for _, h := range c.pathFileHandlers {
			if h.Api() == handler_name {
				return h, nil
			}
		}
	}

	return nil, fmt.Errorf("no existe el controlador: %v para la acción: %v", handler_name, action_type)
}
