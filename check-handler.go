package api

import (
	"fmt"
)

func (c config) isHandlerOk(p *petition, action_type, api_name string) error {

	if action_type == "crud" {
		p.action = "crud"
		p.multiple = true
		return nil
	}

	h, err := c.GetObjectByName(api_name)
	if err != nil {
		return err
	}

	switch action_type {

	case "create":
		if h.CreateApi != nil {
			p.action = "create"
			p.o = h
			return nil
		}

	case "read":
		if h.ReadApi != nil {
			p.action = "read"
			p.o = h
			return nil
		}

	case "update":
		if h.UpdateApi != nil {
			p.action = "update"
			p.o = h
			return nil
		}

	case "delete":
		if h.DeleteApi != nil {
			p.action = "delete"
			p.o = h
			return nil
		}

	case "upload":
		p.action = "upload"
		p.o = h
		return nil
	}

	return fmt.Errorf("no existe el controlador: %v para la acción: %v", api_name, action_type)
}
