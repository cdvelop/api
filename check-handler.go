package api

func (c config) isHandlerOk(p *petition, action_type, api_name string) (err string) {
	const this = "api isHandlerOk error "
	// fmt.Println("TOTAL MANEJADORES", len(c.GetAllObjects()))
	// for _, o := range c.GetAllObjects() {
	// 	fmt.Printf("Estás en el Manejador de lectura de data de %s\n", o.ObjectName)
	// }

	if action_type == "crud" {
		p.action = "crud"
		p.multiple = true
		return ""
	}

	h, err := c.GetObjectByName(api_name)
	if err != "" {
		return this + err
	}

	switch action_type {

	case "create":
		if h.BackHandler.CreateApi != nil {
			p.action = "create"
			p.o = h
			return ""
		}

	case "read":
		if h.BackHandler.ReadApi != nil {
			p.action = "read"
			p.o = h
			return ""
		}

	case "update":
		if h.BackHandler.UpdateApi != nil {
			p.action = "update"
			p.o = h
			return ""
		}

	case "delete":
		if h.BackHandler.DeleteApi != nil {
			p.action = "delete"
			p.o = h
			return ""
		}

	case "upload":
		p.action = "upload"
		p.o = h
		return ""
	}

	return this + "no existe el controlador: " + api_name + " para la acción: " + action_type
}
