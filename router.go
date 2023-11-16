package api

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/cdvelop/model"
	out "github.com/cdvelop/output"
)

func (c config) ServeMuxAndRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		action_type, api_name := getMethodAndObjectFromPath(r.URL.Path)

		// fmt.Println("ACTION", action_type, "API NAME", api_name)

		var registered_user bool
		u, err := c.GetLoginUser(r)
		if err == nil {
			registered_user = true
		}

		// fmt.Println("registered_user", registered_user, u)

		p := &petition{
			u: u,
			o: &model.Object{Name: "error"},
			r: r,
			w: w,
			t: time.Now(),
			e: err,
		}

		if action_type != "" && r.Method != "GET" {
			out.PrintInfo(fmt.Sprintf("método:[%v]: acción:[%v]: objeto[%v]\n", r.Method, action_type, api_name))
		}

		if r.Method == "POST" {

			if !registered_user {
				c.unauthorized(p, "realizar cambios")
				return
			}

			err := c.isHandlerOk(p, action_type, api_name)
			if err != nil {
				c.error(p, err)
				return
			}

			switch p.action {
			case "create":
				c.create(p)

			case "update":
				c.update(p)

			case "delete":
				c.delete(p)

			case "upload":
				c.upload(p)

			default:
				c.error(p, model.Error("acción:", p.action, "no permitida con método", r.Method))
			}

		} else if r.Method == "GET" {

			switch action_type {

			case "read":
				if !registered_user {
					c.unauthorized(p, "obtener información")
					return
				}

				err := c.isHandlerOk(p, action_type, api_name)
				if err != nil {
					c.error(p, err)
					return
				}
				c.read(p)

			case "file":
				if !registered_user {
					c.unauthorized(p, "leer archivos")
					return
				}

				c.readFile(p)

			case "static":
				c.static(w, r)

			default:
				if r.URL.Path == "/" {

					index_content, err := os.ReadFile(filepath.Join(INDEX_FOLDER, "index.html"))
					if err != nil {
						c.error(p, err, http.StatusInternalServerError)
						return
					}

					t, err := template.New("").Parse(string(index_content))
					if err != nil {
						c.error(p, err, http.StatusInternalServerError)
						return
					}

					var responses []model.Response
					var data []byte

					for _, o := range c.bootHandlers {
						// PrintError("boot handler:" + o.Name)
						resp, err := o.AddBootResponse(p.u)
						if err != nil {
							out.PrintError("error boot response:", o.Name, err.Error())
						} else if len(resp) != 0 {
							responses = append(responses, resp...)
						}

					}

					data, err = c.EncodeResponses(responses...)
					if err != nil {
						c.error(p, err, http.StatusInternalServerError)
						return
					}

					var actions = model.BootActions{
						// JsonBootActions: "sin data x",
						JsonBootActions: string(data),
					}

					err = t.Execute(w, actions)
					if err != nil {
						c.error(p, fmt.Errorf("error al retornar pagina %v", err), http.StatusInternalServerError)
						return
					}
					// w.Header.Set()
					// fmt.Fprint(w, "¡Hola! Esta es la página principal.")
					// w.Write()
					// http.ServeFile(w, r, INDEX_FOLDER+"/index.html")
				} else {
					c.error(p, fmt.Errorf("error not found %v", r.URL.Path))
					http.NotFound(w, r)
				}
			}

		} else {
			c.error(p, model.Error("método ", r.Method, "no permitido"), http.StatusMethodNotAllowed)
		}

	})

	return mux
}
