package api

import (
	"fmt"

	"github.com/cdvelop/model"
)

func (c config) isHandlerOk(p *petition, action_type, api_name string) error {

	// fmt.Println("TOTAL MANEJADORES", len(c.GetObjects()))
	// for _, o := range c.GetObjects() {
	// 	fmt.Printf("Estás en el Manejador de lectura de data de %s\n", o.ObjectName)
	// }

	if action_type == "crud" {
		p.action = "crud"
		p.multiple = true
		return nil
	}

	h, err := c.GetObjectByName(api_name)
	if err != nil {
		return model.Error("error isHandlerOk", err)
	}

	switch action_type {

	case "create":
		if h.BackHandler.CreateApi != nil {
			p.action = "create"
			p.o = h
			return nil
		}

	case "read":
		if h.BackHandler.ReadApi != nil {
			p.action = "read"
			p.o = h
			return nil
		}

	case "update":
		if h.BackHandler.UpdateApi != nil {
			p.action = "update"
			p.o = h
			return nil
		}

	case "delete":
		if h.BackHandler.DeleteApi != nil {
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
