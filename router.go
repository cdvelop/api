package api

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/cdvelop/model"
	. "github.com/cdvelop/output"
)

func (c config) ServeMuxAndRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		u := c.auth.GetUser(r)

		action_type, handler_name := getMethodAndObjectFromPath(r.URL.Path)

		if u == nil { // si el usuario no es valido solo puede acceder a la pagina principal
			action_type = ""
		}

		if action_type != "" && r.Method != "GET" {
			PrintInfo(fmt.Sprintf("[%v]: [%v]: [%v]", r.Method, action_type, handler_name))
		}

		switch action_type {

		case "create":
			if r.Method != http.MethodPost {
				c.error(w, r, fmt.Errorf("método %v no permitido para crear", r.Method), nil)
				return
			}

			content_type := r.Header.Get("Content-Type")
			if strings.HasPrefix(content_type, "multipart/form-data") {
				action_type = "file"
			}

			h, err := c.isHandlerOk(action_type, handler_name)
			if err != nil {
				// fmt.Println("HERE 2 ", " action ", action_type, err)
				c.error(w, r, err, h)
				return
			}

			if action_type == "file" {
				c.createFile(u, h, w, r)

			} else {
				c.create(u, h, w, r)
			}

		case "read":
			h, err := c.isHandlerOk(action_type, handler_name)
			if err != nil {
				c.error(w, r, err, h)
				return
			}
			c.read(u, h, w, r)

		case "update":
			if r.Method != http.MethodPost {
				c.error(w, r, fmt.Errorf("método %v no permitido para actualizar", r.Method), nil)
				return
			}

			h, err := c.isHandlerOk(action_type, handler_name)
			if err != nil {
				c.error(w, r, err, h)
				return
			}

			c.update(u, h, w, r)

		case "delete":
			if r.Method != http.MethodPost {
				c.error(w, r, fmt.Errorf("método %v no permitido para eliminar", r.Method), nil)
				return
			}

			h, err := c.isHandlerOk(action_type, handler_name)
			if err != nil {
				c.error(w, r, err, h)
				return
			}

			c.delete(u, h, w, r)

		case "file":

			// fmt.Println("ROUTER API READ FILE")

			if r.Method != http.MethodGet {
				c.error(w, r, fmt.Errorf("método %v no permitido en el Manejador de archivos", r.Method), nil)
				return
			}

			h, err := c.isHandlerOk(action_type, handler_name)
			if err != nil {
				c.error(w, r, err, h)
				return
			}

			c.readFile(u, h, w, r)

		case "static":
			if r.Method != http.MethodGet {
				c.error(w, r, fmt.Errorf("método %v no permitido para archivos estáticos", r.Method), nil)
				return
			}

			c.static(w, r)

		default:
			if r.URL.Path == "/" && r.Method == http.MethodGet {

				index_content, err := os.ReadFile(filepath.Join(INDEX_FOLDER, "index.html"))
				if err != nil {
					log.Println(err)
					return
				}

				t, err := template.New("").Parse(string(index_content))
				if err != nil {
					log.Println(err)
					return
				}

				var responses []model.Response
				var data []byte

				for _, o := range c.bootHandlers {
					// PrintError("boot handler:" + o.Name)
					resp, err := o.AddBootResponse(u)
					if err != nil {
						PrintError("error boot response:", o.Name, err.Error())
					} else if len(resp) != 0 {
						responses = append(responses, resp...)
					}

				}

				data, err = c.EncodeResponses(responses)
				if err != nil {
					log.Println(err)
					return
				}

				var actions = model.BootActions{
					// JsonBootActions: "sin data x",
					JsonBootActions: string(data),
				}

				err = t.Execute(w, actions)
				if err != nil {
					logError(w, r, fmt.Errorf("error al retornar pagina %v", err))
					return
				}

				// w.Header.Set()
				// fmt.Fprint(w, "¡Hola! Esta es la página principal.")

				// w.Write()

				// http.ServeFile(w, r, INDEX_FOLDER+"/index.html")
			} else {
				logError(w, r, fmt.Errorf("error not found %v", r.URL.Path))
				http.NotFound(w, r)
			}
		}
	})

	return mux
}
