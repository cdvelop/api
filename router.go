package api

import (
	"fmt"
	"net/http"
)

func (c config) SetupMuxRoutes() *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		action_type, handler_name := getMethodAndObjectFromPath(r.URL.Path)

		// fmt.Println("action_type ", action_type, " handler_name ", handler_name)

		switch action_type {

		case "create":
			if r.Method != http.MethodPost {
				c.error(w, fmt.Errorf("método %v no permitido para crear", r.Method), nil)
				return
			}

			h, err := c.isHandlerOk("create", handler_name)
			if err != nil {
				c.error(w, err, h)
				return
			}

			c.create(h, w, r)

		case "readone":

			h, err := c.isHandlerOk("readone", handler_name)
			if err != nil {
				c.error(w, err, h)
				return
			}
			c.readone(h, w, r)

		case "readall":

			h, err := c.isHandlerOk("readall", handler_name)
			if err != nil {
				c.error(w, err, h)
				return
			}
			c.readall(h, w, r)

		case "file":

			h, err := c.isHandlerOk("file", handler_name)
			if err != nil {
				c.error(w, err, h)
				return
			}
			c.readFile(h, w, r)

		case "update":
			if r.Method != http.MethodPost {
				c.error(w, fmt.Errorf("método %v no permitido para actualizar", r.Method), nil)
				return
			}

			h, err := c.isHandlerOk("update", handler_name)
			if err != nil {
				c.error(w, err, h)
				return
			}

			c.update(h, w, r)

		case "delete":
			if r.Method != http.MethodPost {
				c.error(w, fmt.Errorf("método %v no permitido para eliminar", r.Method), nil)
				return
			}

			h, err := c.isHandlerOk("delete", handler_name)
			if err != nil {
				c.error(w, err, h)
				return
			}

			c.delete(h, w, r)

		case "static":
			c.static(w, r)

		default:
			if r.URL.Path == "/" && r.Method == http.MethodGet {

				// fmt.Fprint(w, "¡Hola! Esta es la página principal.")

				http.ServeFile(w, r, INDEX_FOLDER+"/index.html")

			} else {
				http.NotFound(w, r)
			}
		}
	})

	return mux
}
