package api

import (
	"fmt"
	"net/http"
	"strings"
)

func (c config) ServeMuxAndRoutes() *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		action_type, handler_name := getMethodAndObjectFromPath(r.URL.Path)

		// fmt.Printf("action_type: [%s] handler_name: [%s] method[%s]\n", action_type, handler_name, r.Method)

		switch action_type {

		case "create":
			if r.Method != http.MethodPost {
				c.error(w, fmt.Errorf("método %v no permitido para crear", r.Method), nil)
				return
			}

			content_type := r.Header.Get("Content-Type")
			if strings.HasPrefix(content_type, "multipart/form-data") {
				action_type = "file"
			}

			h, err := c.isHandlerOk(action_type, handler_name)
			if err != nil {
				// fmt.Println("HERE 2 ", " action ", action_type, err)

				c.error(w, err, h)
				return
			}

			if action_type == "file" {
				c.createFile(h, w, r)

			} else {
				c.create(h, w, r)
			}

		case "read":
			h, err := c.isHandlerOk(action_type, handler_name)
			if err != nil {
				c.error(w, err, h)
				return
			}
			c.read(h, w, r)

		case "update":
			if r.Method != http.MethodPost {
				c.error(w, fmt.Errorf("método %v no permitido para actualizar", r.Method), nil)
				return
			}

			h, err := c.isHandlerOk(action_type, handler_name)
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

			h, err := c.isHandlerOk(action_type, handler_name)
			if err != nil {
				c.error(w, err, h)
				return
			}

			c.delete(h, w, r)

		case "file":

			// fmt.Println("ROUTER API READ FILE")

			if r.Method != http.MethodGet {
				c.error(w, fmt.Errorf("método %v no permitido en el Manejador de archivos", r.Method), nil)
				return
			}

			h, err := c.isHandlerOk(action_type, handler_name)
			if err != nil {
				c.error(w, err, h)
				return
			}

			c.readFile(h, w, r)

		case "static":
			if r.Method != http.MethodGet {
				c.error(w, fmt.Errorf("método %v no permitido para archivos estáticos", r.Method), nil)
				return
			}

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
